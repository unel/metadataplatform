package happy_test

// AC-FS-INIT-01, AC-FS-INIT-03

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store/fs"
)

// AC-FS-INIT-01: успешная инициализация создаёт Store и три поддиректории.
func TestFSInit_NewWithNonexistentBasedir_CreatesStoreAndSubdirs(t *testing.T) {
	// Arrange
	basedir := t.TempDir() + "/store"
	cfg := fs.Config{Basedir: basedir}

	// Act
	store, err := fs.New(cfg, nil)

	// Assert
	require.NoError(t, err, "fs.New must succeed when basedir can be created")
	require.NotNil(t, store, "fs.New must return non-nil *Store on success")

	for _, sub := range []string{"entities", "relations", "jobs"} {
		info, statErr := statDir(basedir + "/" + sub)
		assert.NoErrorf(t, statErr,
			"subdirectory %q must exist after successful fs.New", sub)
		if info != nil {
			assert.Truef(t, info.IsDir(),
				"%q must be a directory", sub)
		}
	}
}

// AC-FS-INIT-03: успешная инициализация пишет INFO-лог с текстом "store/fs initialized" и полем basedir.
func TestFSInit_SuccessfulInit_LogsInfoWithBasedir(t *testing.T) {
	// Arrange
	basedir := t.TempDir() + "/logtest"
	cfg := fs.Config{Basedir: basedir}

	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, &slog.HandlerOptions{Level: slog.LevelDebug}))

	// Act
	_, err := fs.New(cfg, logger)

	// Assert
	require.NoError(t, err)
	logOutput := buf.String()
	assert.Contains(t, logOutput, "INFO",
		"log must contain INFO level entry")
	assert.Contains(t, logOutput, "store/fs initialized",
		"INFO log entry must contain message 'store/fs initialized'")
	assert.Contains(t, logOutput, basedir,
		"INFO log entry must contain basedir path as field value")
}
