package fs_test

// Вспомогательные функции: FS-стор, builders для моделей.

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
	"github.com/unel/metadataplatform/store/fs"
)

// newFSStore создаёт fs.Store во временной директории.
func newFSStore(t *testing.T) (*fs.Store, string) {
	t.Helper()
	basedir := t.TempDir()
	s, err := fs.New(fs.Config{Basedir: basedir}, nil)
	require.NoError(t, err, "fs.New must not fail in test setup")
	return s, basedir
}

// statDir — обёртка для os.Stat.
func statDir(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// rawJSON конвертирует строку в json.RawMessage.
func rawJSON(s string) json.RawMessage {
	return json.RawMessage(s)
}

// defaultEntity возвращает Entity с минимально заполненными полями.
func defaultEntity(id string) store.Entity {
	return store.Entity{
		ID:   id,
		Type: "file",
		Meta: rawJSON(`{}`),
	}
}

// defaultRelation возвращает Relation с минимально заполненными полями.
func defaultRelation(id, fromID, toID string) store.Relation {
	return store.Relation{
		ID:     id,
		FromID: fromID,
		ToID:   toID,
		Type:   "derivedFrom",
		Value:  rawJSON(`{}`),
		Meta:   rawJSON(`{}`),
	}
}

// defaultJob возвращает Job с минимально заполненными полями.
func defaultJob(id, entityID string) store.Job {
	return store.Job{
		ID:       id,
		EntityID: entityID,
		Kind:     "hash",
		Worker:   "hash.sha256",
		Status:   "pending",
		Payload:  rawJSON(`{}`),
	}
}
