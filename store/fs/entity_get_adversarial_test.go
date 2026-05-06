package fs_test

// AC-ENTITY-GET-02, AC-ENTITY-GET-04
// Get — казалось бы просто. Но "файл есть, прочитать нельзя" — уже интереснее.

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

// AC-ENTITY-GET-03: пустой id → ошибка, не ErrNotFound и не паника.
// Важно: пустой id — это MISSING_ID, не "файл не найден".
func TestEntity_Get_EmptyID_ReturnsError(t *testing.T) {
	es := newEntityStore(t)

	_, err := es.Get(context.Background(), "")
	require.Error(t, err, "пустой id должен вернуть ошибку")
	assert.False(t, errors.Is(err, store.ErrNotFound),
		"пустой id — это MISSING_ID, не NOT_FOUND")
}

// AC-ENTITY-GET-02: несуществующая запись → ErrNotFound, не просто "какая-то ошибка".
func TestEntity_Get_NonExistentID_ReturnsErrNotFound(t *testing.T) {
	es := newEntityStore(t)

	_, err := es.Get(context.Background(), "ent-999")
	require.Error(t, err)
	assert.True(t, errors.Is(err, store.ErrNotFound),
		"ожидаем store.ErrNotFound, получили: %v", err)
}

// AC-ENTITY-GET-04: файл существует, но права 000 → READ_ERROR, не ErrNotFound.
// Важно: os.IsNotExist(err) == false здесь — это другая ветка кода.
func TestEntity_Get_FileExistsButUnreadable_ReturnsReadError(t *testing.T) {
	basedir := t.TempDir()
	s, err := fs.New(fs.Config{Basedir: basedir}, testLogger(t))
	require.NoError(t, err)
	es := s.Entities()
	ctx := context.Background()

	require.NoError(t, es.Upsert(ctx, store.Entity{ID: "ent-perm", Type: "service"}))

	filePath := filepath.Join(basedir, "entities", "ent-perm.json")
	require.NoError(t, os.Chmod(filePath, 0o000))
	t.Cleanup(func() { os.Chmod(filePath, 0o644) })

	_, getErr := es.Get(ctx, "ent-perm")
	require.Error(t, getErr)

	// это НЕ должно быть ErrNotFound — файл существует, просто нечитаем
	assert.False(t, errors.Is(getErr, store.ErrNotFound),
		"ошибка чтения не должна мимикрировать под ErrNotFound")
}
