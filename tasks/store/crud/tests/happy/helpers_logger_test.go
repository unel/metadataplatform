package happy_test

// Вспомогательная функция: роутер с slog-логгером на bytes.Buffer.

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/unel/metadataplatform/store/fs"
	"github.com/unel/metadataplatform/store/router"
)

// newRouterWithSlogBuf создаёт роутер с захватывающим slog-логгером и FS-стором.
func newRouterWithSlogBuf(t *testing.T) (*router.Router, *bytes.Buffer, *fs.Store) {
	t.Helper()
	s, _ := newFSStore(t)
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))
	r := router.New(s.Entities(), s.Relations(), s.Jobs(), logger)
	return r, &buf, s
}
