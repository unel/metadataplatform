package integration_test

// Вспомогательные функции для тестов JSONL-роутера.

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store/fs"
	"github.com/unel/metadataplatform/store/router"
)

type routerResponse struct {
	OK        bool            `json:"ok"`
	Error     string          `json:"error,omitempty"`
	ErrorCode string          `json:"errorCode,omitempty"`
	ID        string          `json:"id,omitempty"`
	Data      json.RawMessage `json:"data,omitempty"`
}

// newTestRouter создаёт роутер с реальным FS-стором на net.Pipe().
// logBuf — опциональный буфер для перехвата лог-вывода; nil → slog.Default().
func newTestRouter(t *testing.T, logBuf *bytes.Buffer) (clientConn net.Conn) {
	t.Helper()

	s, err := fs.New(fs.Config{Basedir: t.TempDir()}, slog.Default())
	require.NoError(t, err)

	var logger *slog.Logger
	if logBuf != nil {
		logger = slog.New(slog.NewTextHandler(logBuf, &slog.HandlerOptions{Level: slog.LevelDebug}))
	} else {
		logger = slog.Default()
	}

	r := router.New(s.Entities(), s.Relations(), s.Jobs(), logger)

	serverConn, clientConn := net.Pipe()
	t.Cleanup(func() {
		serverConn.Close()
		clientConn.Close()
	})

	go r.Handle(context.Background(), serverConn)
	return clientConn
}

func sendLine(t *testing.T, conn net.Conn, line string) {
	t.Helper()
	_, err := fmt.Fprintf(conn, "%s\n", line)
	require.NoError(t, err)
}

func readResponse(t *testing.T, conn net.Conn) routerResponse {
	t.Helper()
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	sc := bufio.NewScanner(conn)
	require.True(t, sc.Scan(), "ожидаем ответ от роутера")
	var resp routerResponse
	require.NoError(t, json.Unmarshal(sc.Bytes(), &resp))
	return resp
}

// isConnClosed проверяет что ошибка означает закрытое соединение (не таймаут).
func isConnClosed(err error) bool {
	if err == nil {
		return false
	}
	if err == io.ErrUnexpectedEOF {
		return true
	}
	msg := err.Error()
	return msg == "io: read/write on closed pipe" ||
		bytes.Contains([]byte(msg), []byte("closed"))
}
