package fs_test

// AC-ENTITY-DELETE-02—04
// Delete не идемпотентный. Это редкий дизайн-выбор, который легко сломать "починкой".

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
	"github.com/unel/metadataplatform/store/fs"
)

// AC-ENTITY-DELETE-02: delete несуществующей записи → ErrNotFound.
// Идемпотентный delete вернул бы nil — здесь должна быть ошибка.
func TestEntity_Delete_NonExistentID_ReturnsErrNotFound(t *testing.T) {
	es := newEntityStore(t)

	err := es.Delete(context.Background(), "ent-999")
	require.Error(t, err)
	assert.True(t, errors.Is(err, store.ErrNotFound),
		"delete несуществующей записи должен вернуть ErrNotFound, получили: %v", err)
}

// AC-ENTITY-DELETE-03: пустой id → ошибка (не паника, не ErrNotFound).
func TestEntity_Delete_EmptyID_ReturnsError(t *testing.T) {
	es := newEntityStore(t)

	err := es.Delete(context.Background(), "")
	assert.Error(t, err, "пустой id должен вернуть ошибку")
	assert.False(t, errors.Is(err, store.ErrNotFound),
		"пустой id — это MISSING_ID, не NOT_FOUND")
}

// AC-ENTITY-DELETE-04: файл существует, но директория без прав на удаление.
// Симулируем через sticky bit + owner mismatch (или просто chmod o-w на родителя).
func TestEntity_Delete_FileExistsButRemoveFails_ReturnsDeleteError(t *testing.T) {
	basedir := t.TempDir()
	es := newEntityStoreAt(t, basedir)
	ctx := context.Background()

	require.NoError(t, es.Upsert(ctx, store.Entity{ID: "ent-del", Type: "service"}))

	entitiesDir := filepath.Join(basedir, "entities")

	// убираем write из директории — os.Remove файла потребует write на директорию
	require.NoError(t, os.Chmod(entitiesDir, 0o555))
	t.Cleanup(func() { os.Chmod(entitiesDir, 0o755) })

	err := es.Delete(ctx, "ent-del")
	require.Error(t, err, "ожидаем ошибку при невозможности удалить файл")

	// не должно быть ErrNotFound — файл есть, просто нельзя удалить
	assert.False(t, errors.Is(err, store.ErrNotFound),
		"ошибка удаления не должна маскироваться под ErrNotFound")
}

func newEntityStore(t *testing.T) store.EntityStore {
	t.Helper()
	basedir := t.TempDir()
	s, err := fs.New(fs.Config{Basedir: basedir}, testLogger(t))
	require.NoError(t, err)
	return s.Entities()
}

func newEntityStoreAt(t *testing.T, basedir string) store.EntityStore {
	t.Helper()
	s, err := fs.New(fs.Config{Basedir: basedir}, testLogger(t))
	require.NoError(t, err)
	return s.Entities()
}
