package fs_test

// AC-ENTITY-UPSERT-03—06
// Upsert — самая богатая на failure modes операция. Проверяем каждый.

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
)

// AC-ENTITY-UPSERT-03: пустой id — возвращает ошибку, не паникует.
func TestEntity_Upsert_EmptyID_ReturnsError(t *testing.T) {
	es := newEntityStore(t)
	err := es.Upsert(context.Background(), store.Entity{ID: "", Type: "service"})
	assert.Error(t, err, "пустой id должен вернуть ошибку")
}

// AC-ENTITY-UPSERT-04: файл существует, но недоступен для чтения (права 000).
// Upsert должен читать существующую запись чтобы сохранить created_at.
func TestEntity_Upsert_ExistingFileUnreadable_ReturnsReadError(t *testing.T) {
	basedir := t.TempDir()
	es := newEntityStoreAt(t, basedir)
	ctx := context.Background()

	// создаём запись
	require.NoError(t, es.Upsert(ctx, store.Entity{ID: "ent-perm", Type: "service"}))

	// блокируем чтение
	filePath := filepath.Join(basedir, "entities", "ent-perm.json")
	require.NoError(t, os.Chmod(filePath, 0o000))
	t.Cleanup(func() { os.Chmod(filePath, 0o644) })

	err := es.Upsert(ctx, store.Entity{ID: "ent-perm", Type: "service", Name: "updated"})
	assert.Error(t, err, "ожидаем ошибку при невозможности прочитать существующий файл")
}

// AC-ENTITY-UPSERT-05: директория entities недоступна для записи.
func TestEntity_Upsert_DirectoryUnwritable_ReturnsWriteError(t *testing.T) {
	basedir := t.TempDir()
	es := newEntityStoreAt(t, basedir)

	entitiesDir := filepath.Join(basedir, "entities")
	require.NoError(t, os.Chmod(entitiesDir, 0o555))
	t.Cleanup(func() { os.Chmod(entitiesDir, 0o755) })

	err := es.Upsert(context.Background(), store.Entity{ID: "ent-new", Type: "service"})
	assert.Error(t, err, "ожидаем ошибку при недоступной для записи директории")
}

// AC-ENTITY-UPSERT-06: temp-файл не остаётся при ошибке записи.
// Если реализация использует defer cleanup — temp не должен торчать.
func TestEntity_Upsert_WriteError_NoTempFileLeft(t *testing.T) {
	basedir := t.TempDir()
	es := newEntityStoreAt(t, basedir)

	entitiesDir := filepath.Join(basedir, "entities")
	require.NoError(t, os.Chmod(entitiesDir, 0o555))
	t.Cleanup(func() { os.Chmod(entitiesDir, 0o755) })

	_ = es.Upsert(context.Background(), store.Entity{ID: "ent-fail", Type: "service"})

	entries, err := os.ReadDir(entitiesDir)
	if err != nil {
		t.Skip("entitiesDir unreadable after chmod — running as root; skipping tmp-file check")
	}

	for _, e := range entries {
		assert.False(t, strings.HasPrefix(e.Name(), ".tmp-"),
			"temp-файл %q не должен оставаться после ошибки", e.Name())
	}
}
