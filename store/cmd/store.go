package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// logger — интерфейс логирования. *logBuffer из тестов реализует его.
type logger interface {
	add(level, msg string, attrs map[string]string)
}

func RunStore(cfg Config, log logger, stopCh <-chan struct{}) bool {
	applyDefaults(&cfg)

	var connCounter uint64

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	listeners := make([]net.Listener, 0, len(cfg.Store.Sockets))
	for _, sc := range cfg.Store.Sockets {
		ln, err := listenUnix(sc.Path)
		if err != nil {
			log.add("ERROR", fmt.Sprintf("failed to create socket %s: %v", sc.Path, err), nil)
			cancel()
			for _, l := range listeners {
				l.Close()
			}
			return false
		}
		listeners = append(listeners, ln)
		log.add("INFO", fmt.Sprintf("listening on socket %s", sc.Path), nil)
	}

	// semaphore ограничивает суммарное число параллельных соединений по всем сокетам.
	sem := make(chan struct{}, cfg.Store.MaxConnections)

	for i, ln := range listeners {
		sockPath := cfg.Store.Sockets[i].Path
		wg.Add(1)
		go func(ln net.Listener, path string) {
			defer wg.Done()
			acceptLoop(ctx, ln, path, &connCounter, cfg.Store.MaxLineBytes, log, sem)
		}(ln, sockPath)
	}

	<-stopCh
	log.add("INFO", "shutting down", nil)

	for i, ln := range listeners {
		if err := ln.Close(); err != nil {
			log.add("WARN", fmt.Sprintf("failed to close listener %s: %v", cfg.Store.Sockets[i].Path, err), nil)
		}
		if err := os.Remove(cfg.Store.Sockets[i].Path); err != nil && !os.IsNotExist(err) {
			log.add("WARN", fmt.Sprintf("failed to remove socket %s: %v", cfg.Store.Sockets[i].Path, err), nil)
		}
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		cancel()
	case <-time.After(5 * time.Second):
		cancel() // форсируем закрытие активных соединений
		<-done   // ждём завершения горутин после форсированного закрытия
	}
	return true
}

func RunStoreWithExitCode(cfg Config, log logger, stopCh <-chan struct{}) int {
	if cfg.configPath != "" {
		loaded, err := loadConfig(cfg.configPath)
		if err != nil {
			log.add("ERROR", fmt.Sprintf("config error: %v", err), nil)
			return 1
		}
		cfg = loaded
	}

	if len(cfg.Store.Sockets) == 0 {
		log.add("ERROR", "no sockets configured: store.sockets must not be empty", nil)
		return 1
	}

	if !RunStore(cfg, log, stopCh) {
		return 1
	}
	return 0
}

func listenUnix(path string) (net.Listener, error) {
	os.Remove(path)
	return net.Listen("unix", path)
}

// acceptLoop принимает соединения и запускает горутину для каждого.
// Перед возвратом ждёт завершения всех запущенных conn-горутин.
// sem ограничивает число одновременно активных соединений.
func acceptLoop(ctx context.Context, ln net.Listener, sockPath string, counter *uint64, maxLineBytes int, log logger, sem chan struct{}) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}

		select {
		case sem <- struct{}{}:
			// слот получен — продолжаем
		default:
			// лимит достигнут — отклоняем соединение
			log.add("WARN", fmt.Sprintf("connection limit reached on %s, dropping connection", sockPath), nil)
			conn.Close()
			continue
		}

		connID := atomic.AddUint64(counter, 1)
		attrs := map[string]string{"conn_id": fmt.Sprintf("%d", connID)}
		log.add("INFO", fmt.Sprintf("new connection on %s", sockPath), attrs)

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			handleConn(ctx, conn, connID, maxLineBytes, log)
		}()
	}
}
