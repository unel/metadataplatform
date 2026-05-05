package adversarial_test

// AC-FS-INIT-02, AC-FS-INIT-04
// "Это никогда не случится в проде" — а вот и случится, когда монтирование потеряется.

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store/fs"
)

// AC-FS-INIT-02: конструктор возвращает (nil, err) если basedir нельзя создать.
func TestFSInit_UncreatableBasedir_ReturnsError(t *testing.T) {
	parent := t.TempDir()

	// запрещаем запись в родительский каталог — os.MkdirAll не сможет создать subdir
	require.NoError(t, os.Chmod(parent, 0o555))
	t.Cleanup(func() { os.Chmod(parent, 0o755) })

	basedir := filepath.Join(parent, "locked", "nested")

	store, err := fs.New(fs.Config{Basedir: basedir}, slog.Default())
	assert.Nil(t, store, "ожидаем nil стор при ошибке init")
	assert.Error(t, err, "ожидаем ошибку при недоступном basedir")
}

// AC-FS-INIT-04: при ошибке init логируется ERROR с полями basedir и error.
func TestFSInit_UncreatableBasedir_LogsErrorWithFields(t *testing.T) {
	parent := t.TempDir()
	require.NoError(t, os.Chmod(parent, 0o555))
	t.Cleanup(func() { os.Chmod(parent, 0o755) })

	basedir := filepath.Join(parent, "locked")

	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	_, _ = fs.New(fs.Config{Basedir: basedir}, logger)

	logOutput := buf.String()
	assert.Contains(t, logOutput, "ERROR", "ожидаем уровень ERROR в логе")
	assert.Contains(t, logOutput, "store/fs init failed", "ожидаем сообщение store/fs init failed")
	assert.True(t, strings.Contains(logOutput, basedir),
		"ожидаем basedir=%q в логе, получили: %s", basedir, logOutput)
	assert.Contains(t, logOutput, "error", "ожидаем поле error в логе")
}

// Бонус: успешная инициализация логирует INFO, не ERROR.
func TestFSInit_ValidBasedir_LogsInfoAndNoError(t *testing.T) {
	basedir := filepath.Join(t.TempDir(), "store")

	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	store, err := fs.New(fs.Config{Basedir: basedir}, logger)
	require.NoError(t, err)
	require.NotNil(t, store)

	logOutput := buf.String()
	assert.NotContains(t, logOutput, "ERROR", "успешный init не должен логировать ERROR")
	assert.Contains(t, logOutput, "store/fs initialized", "ожидаем INFO store/fs initialized")
}
