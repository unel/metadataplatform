package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"syscall"
)

// handleMessage — обработчик одной строки. Сейчас эхо.
// Следующая фича подменит на реальный роутер.
func handleMessage(line string) string {
	return line
}

// isClientClose возвращает true если ошибка записи вызвана тем что клиент
// закрыл соединение со своей стороны (broken pipe = clean FIN от клиента).
func isClientClose(err error) bool {
	return errors.Is(err, syscall.EPIPE)
}

// isForcedClose возвращает true если ошибка записи вызвана тем что мы сами
// закрыли соединение (conn.Close из watcher-горутины при ctx.Done).
func isForcedClose(err error) bool {
	return errors.Is(err, net.ErrClosed)
}

func handleConn(ctx context.Context, conn net.Conn, connID uint64, maxLineBytes int, log logger) {
	defer conn.Close()

	rd := bufio.NewReaderSize(conn, maxLineBytes+1)

	connClosed := make(chan struct{})
	defer close(connClosed)
	go func() {
		select {
		case <-ctx.Done():
			conn.Close()
		case <-connClosed:
		}
	}()

	attrs := map[string]string{"conn_id": fmt.Sprintf("%d", connID)}

	var lineBytes []byte

	for {
		// ReadLine выбран намеренно: isPrefix позволяет накапливать чанки и обнаруживать
		// превышение maxLineBytes без лишних аллокаций — Scanner этого не даёт.
		chunk, isPrefix, err := rd.ReadLine()

		// накапливаем байты строки; при превышении лимита закрываем немедленно,
		// не блокируясь на дальнейшем чтении
		if len(chunk) > 0 {
			lineBytes = append(lineBytes, chunk...)
			if len(lineBytes) > maxLineBytes {
				log.add("INFO", "line exceeds max_line_bytes limit, closing connection", attrs)
				return
			}
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				if len(lineBytes) > 0 {
					// partial line при clean EOF — нормальное закрытие клиентом
					log.add("DEBUG", "connection closed (EOF, partial line)", attrs)
				} else {
					log.add("DEBUG", "connection closed (EOF)", attrs)
				}
				return
			}
			if errors.Is(err, io.ErrUnexpectedEOF) {
				log.add("DEBUG", "connection closed (unexpected EOF, partial line)", attrs)
				return
			}
			if ctx.Err() != nil {
				return
			}
			log.add("INFO", fmt.Sprintf("read error: %v", err), attrs)
			return
		}

		if isPrefix {
			// строка продолжается — читаем следующий чанк
			continue
		}

		// полная строка собрана
		line := string(lineBytes)
		lineBytes = lineBytes[:0]

		if strings.TrimSpace(line) == "" {
			continue
		}

		log.add("DEBUG", line, attrs)

		response := handleMessage(line)
		if _, werr := fmt.Fprintf(conn, "%s\n", response); werr != nil {
			switch {
			case isClientClose(werr):
				// клиент закрыл соединение со своей стороны пока мы отвечали
				log.add("DEBUG", "connection closed (EOF)", attrs)
			case isForcedClose(werr):
				// мы принудительно закрыли соединение по таймауту shutdown
				log.add("INFO", "connection forcibly closed", attrs)
			default:
				log.add("INFO", fmt.Sprintf("write error: %v", werr), attrs)
			}
			return
		}
	}
}
