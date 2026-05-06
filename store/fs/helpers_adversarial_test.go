package fs_test

// Вспомогательные функции для adversarial тестов.
// testLogger — slog в stderr только WARN+, чтобы не засорять вывод.
// Если тест хочет проверить лог — он создаёт свой bytes.Buffer и logger напрямую.

import (
	"log/slog"
	"os"
	"testing"
)

// testLogger возвращает slog.Logger с уровнем WARN — тихий по умолчанию.
func testLogger(t *testing.T) *slog.Logger {
	t.Helper()
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelWarn}))
}
