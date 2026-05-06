package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// StoreProcess — обёртка вокруг запущенного store-процесса.
// Нет, это не мок. Это настоящий процесс. Мы честные люди.
type StoreProcess struct {
	t       *testing.T
	sockDir string
	socks   []string

	doneCh chan struct{}

	// cancel — функция остановки store; зависит от реализации
	// (в тестах подменяем на embedded runner)
	cancel func()
	runner *embeddedRunner
}

// embeddedRunner запускает store in-process через RunStore.
// Так мы не зависим от наличия бинаря — юнит-тесты без компиляции.
type embeddedRunner struct {
	cfg    Config
	logBuf *logBuffer
	stopCh chan struct{}
	doneCh chan struct{}
}

// logBuffer — thread-safe буфер логов для перехвата вывода store.
type logBuffer struct {
	mu   sync.Mutex
	msgs []logEntry
}

type logEntry struct {
	level string
	msg   string
	attrs map[string]string
}

func (b *logBuffer) add(level, msg string, attrs map[string]string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.msgs = append(b.msgs, logEntry{level: level, msg: msg, attrs: attrs})
}

func (b *logBuffer) all() []logEntry {
	b.mu.Lock()
	defer b.mu.Unlock()
	out := make([]logEntry, len(b.msgs))
	copy(out, b.msgs)
	return out
}

func (b *logBuffer) hasLevel(level string) bool {
	for _, e := range b.all() {
		if strings.EqualFold(e.level, level) {
			return true
		}
	}
	return false
}

func (b *logBuffer) hasMsg(substr string) bool {
	for _, e := range b.all() {
		if strings.Contains(e.msg, substr) {
			return true
		}
	}
	return false
}

func (b *logBuffer) hasLevelAndMsg(level, substr string) bool {
	for _, e := range b.all() {
		if strings.EqualFold(e.level, level) && strings.Contains(e.msg, substr) {
			return true
		}
	}
	return false
}

// waitForMsg ждёт появления записи в логах с таймаутом.
func (b *logBuffer) waitForMsg(substr string, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if b.hasMsg(substr) {
			return true
		}
		time.Sleep(5 * time.Millisecond)
	}
	return false
}

// waitForLevelAndMsg ждёт появления записи нужного уровня с таймаутом.
func (b *logBuffer) waitForLevelAndMsg(level, substr string, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if b.hasLevelAndMsg(level, substr) {
			return true
		}
		time.Sleep(5 * time.Millisecond)
	}
	return false
}

// connIDs возвращает все conn_id из лога.
func (b *logBuffer) connIDs() []string {
	seen := map[string]struct{}{}
	var ids []string
	for _, e := range b.all() {
		if id, ok := e.attrs["conn_id"]; ok {
			if _, dup := seen[id]; !dup {
				seen[id] = struct{}{}
				ids = append(ids, id)
			}
		}
	}
	return ids
}

// waitForConnID ждёт появления хотя бы одного conn_id в attrs логов.
// conn_id живёт в structured fields (attrs), а не в msg —
// используй этот метод вместо waitForMsg("conn_id").
func (b *logBuffer) waitForConnID(timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if len(b.connIDs()) > 0 {
			return true
		}
		time.Sleep(5 * time.Millisecond)
	}
	return false
}

// waitForNConnIDs ждёт появления не менее n различных conn_id в attrs логов.
func (b *logBuffer) waitForNConnIDs(n int, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if len(b.connIDs()) >= n {
			return true
		}
		time.Sleep(5 * time.Millisecond)
	}
	return false
}

// --- Хелперы для создания процесса store ---

// newStoreProcess создаёт и запускает store с заданными сокетами во временной директории.
// Возвращает StoreProcess; cleanup регистрируется через t.Cleanup.
func newStoreProcess(t *testing.T, sockNames []string, extraCfg ...func(*Config)) *StoreProcess {
	t.Helper()
	dir := t.TempDir()

	var paths []SocketConfig
	for _, name := range sockNames {
		paths = append(paths, SocketConfig{
			Path: filepath.Join(dir, name),
			Ops:  []string{"query", "upsert", "delete"},
		})
	}

	cfg := Config{
		Store: StoreConfig{
			DBUrl:        "postgres://localhost/test_unused",
			MaxLineBytes: 1048576,
			Sockets:      paths,
		},
	}

	for _, apply := range extraCfg {
		apply(&cfg)
	}

	logs := &logBuffer{}
	sp := &StoreProcess{
		t:       t,
		sockDir: dir,
		doneCh:  make(chan struct{}),
	}

	runner := &embeddedRunner{
		cfg:    cfg,
		logBuf: logs,
		stopCh: make(chan struct{}),
		doneCh: make(chan struct{}),
	}
	sp.runner = runner

	go func() {
		defer close(runner.doneCh)
		RunStore(cfg, logs, runner.stopCh)
	}()

	// ждём пока store залогирует "listening on socket <path>" для каждого сокета —
	// это гарантирует что listenUnix выполнен и сокет принимает подключения.
	for _, sc := range cfg.Store.Sockets {
		sp.socks = append(sp.socks, sc.Path)
		if !logs.waitForMsg(sc.Path, 2*time.Second) {
			t.Fatalf("socket did not start listening: %s", sc.Path)
		}
	}

	// Stop идемпотентен — безопасно вызывать из Cleanup даже если тест уже вызвал Stop().
	t.Cleanup(func() { sp.Stop() })

	return sp
}

// Logs возвращает буфер логов.
func (sp *StoreProcess) Logs() *logBuffer {
	return sp.runner.logBuf
}

// SockPath возвращает путь к первому сокету (или по индексу).
func (sp *StoreProcess) SockPath(idx ...int) string {
	i := 0
	if len(idx) > 0 {
		i = idx[0]
	}
	return sp.socks[i]
}

// Stop останавливает store и ждёт завершения.
func (sp *StoreProcess) Stop() {
	select {
	case <-sp.runner.stopCh:
		// уже остановлен
	default:
		close(sp.runner.stopCh)
	}
	select {
	case <-sp.runner.doneCh:
	case <-time.After(6 * time.Second):
		sp.t.Log("store did not stop in time")
	}
}

// --- Сетевые хелперы ---

// dialStore подключается к Unix сокету. Умирает если не удалось.
// После успешного dial ждёт небольшую паузу чтобы acceptLoop успел залогировать
// "new connection" — без этого тест рискует захватить baseline логов до появления
// записи о подключении.
func dialStore(t *testing.T, sockPath string) net.Conn {
	t.Helper()
	conn, err := net.DialTimeout("unix", sockPath, time.Second)
	if err != nil {
		t.Fatalf("dial %s: %v", sockPath, err)
	}
	t.Cleanup(func() { conn.Close() })
	time.Sleep(25 * time.Millisecond)
	return conn
}

// sendLine отправляет строку с \n в соединение.
func sendLine(t *testing.T, conn net.Conn, line string) {
	t.Helper()
	_, err := fmt.Fprintf(conn, "%s\n", line)
	if err != nil {
		t.Fatalf("sendLine: %v", err)
	}
}

// readLine читает одну строку из соединения с таймаутом.
func readLine(t *testing.T, conn net.Conn, timeout time.Duration) string {
	t.Helper()
	conn.SetReadDeadline(time.Now().Add(timeout))
	defer conn.SetReadDeadline(time.Time{})
	sc := bufio.NewScanner(conn)
	if !sc.Scan() {
		if err := sc.Err(); err != nil {
			t.Fatalf("readLine: %v", err)
		}
		t.Fatal("readLine: EOF before reading a line")
	}
	return sc.Text()
}

// expectEOF читает из соединения и ожидает EOF (соединение закрыто сервером).
func expectEOF(t *testing.T, conn net.Conn, timeout time.Duration) {
	t.Helper()
	conn.SetReadDeadline(time.Now().Add(timeout))
	defer conn.SetReadDeadline(time.Time{})
	buf := make([]byte, 1)
	_, err := conn.Read(buf)
	if err == nil {
		t.Fatal("expectEOF: got data instead of EOF")
	}
	// net.ErrClosed, io.EOF — всё подходит; главное что соединение закрыто
}

// sockGone проверяет что сокет-файл удалён.
func sockGone(path string, timeout time.Duration) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			return true
		}
		time.Sleep(5 * time.Millisecond)
	}
	return false
}

// makeStaleSockFile создаёт файл-заглушку по пути сокета (имитирует stale socket).
func makeStaleSockFile(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, []byte("stale"), 0o600); err != nil {
		t.Fatalf("makeStaleSockFile: %v", err)
	}
}

// newTestScanner создаёт bufio.Scanner с увеличенным буфером для чтения строк в тестах.
// Буфер 4 МБ нужен для тестов с большими строками (EDGE-02 и подобные).
func newTestScanner(r interface{ Read([]byte) (int, error) }) *bufio.Scanner {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 4*1024*1024), 4*1024*1024)
	return sc
}
