package fs_test

// AC-ENTITY-LIST-01, AC-ENTITY-LIST-02, AC-ENTITY-LIST-05

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AC-ENTITY-LIST-01: list возвращает все сохранённые entity.
func TestEntity_List_MultipleRecords_ReturnsAll(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	ctx := context.Background()

	entities := []string{"list-entity-a", "list-entity-b", "list-entity-c"}
	for _, id := range entities {
		require.NoError(t, s.Upsert(ctx, defaultEntity(id)))
	}

	// Act
	result, err := s.List(ctx)

	// Assert
	require.NoError(t, err, "List must succeed")
	require.Len(t, result, len(entities),
		"List must return exactly %d entities", len(entities))

	ids := make([]string, len(result))
	for i, e := range result {
		ids[i] = e.ID
	}
	assert.ElementsMatch(t, entities, ids,
		"List must return all inserted entity IDs")
}

// AC-ENTITY-LIST-02: list на пустом сторе возвращает пустой срез, не ошибку.
func TestEntity_List_EmptyStore_ReturnsEmptySlice(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	ctx := context.Background()

	// Act
	result, err := s.List(ctx)

	// Assert
	require.NoError(t, err, "List on empty store must not return error")
	assert.NotNil(t, result, "List must return non-nil slice even when empty")
	assert.Len(t, result, 0, "List on empty store must return zero elements")
}

// AC-ENTITY-LIST-05: list игнорирует .tmp файлы в директории entities.
func TestEntity_List_TmpFilesPresent_IgnoresTmpFiles(t *testing.T) {
	// Arrange
	s, basedir := newFSStore(t)
	ctx := context.Background()

	require.NoError(t, s.Upsert(ctx, defaultEntity("real-entity-01")))

	// создаём .tmp файл вручную — он должен быть проигнорирован
	tmpPath := filepath.Join(basedir, "entities", ".tmp-stale12345")
	require.NoError(t, os.WriteFile(tmpPath, []byte(`{"id":"ghost"}`), 0o600))

	// Act
	result, err := s.List(ctx)

	// Assert
	require.NoError(t, err, "List must succeed even with .tmp files present")
	require.Len(t, result, 1, "List must return only real entity, not tmp file")
	assert.Equal(t, "real-entity-01", result[0].ID)
}
