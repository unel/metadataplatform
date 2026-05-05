package happy_test

// AC-ENTITY-UPSERT-01, AC-ENTITY-UPSERT-02

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AC-ENTITY-UPSERT-01: создание новой записи — входящие timestamps игнорируются,
// created_at и updated_at выставляются стором.
func TestEntity_Upsert_NewRecord_SetsTimestamps(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	ctx := context.Background()
	before := time.Now().UTC()

	entity := defaultEntity("entity-upsert-01")
	entity.CreatedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) // должно быть проигнорировано
	entity.UpdatedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) // должно быть проигнорировано

	// Act
	err := s.Upsert(ctx, entity)
	after := time.Now().UTC()

	// Assert
	require.NoError(t, err, "Upsert of new entity must succeed")

	got, getErr := s.Get(ctx, entity.ID)
	require.NoError(t, getErr)

	assert.True(t,
		!got.CreatedAt.Before(before) && !got.CreatedAt.After(after),
		"created_at must be set by store, not taken from input: got %v", got.CreatedAt,
	)
	assert.True(t,
		!got.UpdatedAt.Before(before) && !got.UpdatedAt.After(after),
		"updated_at must be set by store, not taken from input: got %v", got.UpdatedAt,
	)
	assert.Equal(t, entity.ID, got.ID)
	assert.Equal(t, entity.Type, got.Type)
}

// AC-ENTITY-UPSERT-01: создание новой записи — данные сохраняются корректно.
func TestEntity_Upsert_NewRecord_PersistsAllFields(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	ctx := context.Background()

	entity := defaultEntity("entity-fields-01")
	entity.Subtype = "video"
	entity.Name = "test.mp4"
	entity.Description = "test video file"
	entity.Meta = rawJSON(`{"path":"/data/test.mp4","size":1024}`)

	// Act
	err := s.Upsert(ctx, entity)

	// Assert
	require.NoError(t, err)
	got, getErr := s.Get(ctx, entity.ID)
	require.NoError(t, getErr)

	assert.Equal(t, entity.ID, got.ID)
	assert.Equal(t, entity.Type, got.Type)
	assert.Equal(t, entity.Subtype, got.Subtype)
	assert.Equal(t, entity.Name, got.Name)
	assert.Equal(t, entity.Description, got.Description)
	assert.JSONEq(t, string(entity.Meta), string(got.Meta))
}

// AC-ENTITY-UPSERT-02: обновление существующей записи — created_at сохраняется,
// updated_at обновляется.
func TestEntity_Upsert_ExistingRecord_PreservesCreatedAtAndUpdatesUpdatedAt(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	ctx := context.Background()

	entity := defaultEntity("entity-upsert-02")
	require.NoError(t, s.Upsert(ctx, entity))

	original, err := s.Get(ctx, entity.ID)
	require.NoError(t, err)

	// небольшая пауза чтобы updated_at гарантированно отличался
	time.Sleep(2 * time.Millisecond)

	// Act: повторный upsert с изменёнными данными
	entity.Name = "updated name"
	err = s.Upsert(ctx, entity)

	// Assert
	require.NoError(t, err, "Upsert of existing entity must succeed")

	updated, getErr := s.Get(ctx, entity.ID)
	require.NoError(t, getErr)

	assert.True(t,
		updated.CreatedAt.Equal(original.CreatedAt),
		"created_at must be preserved on update: original=%v got=%v",
		original.CreatedAt, updated.CreatedAt,
	)
	assert.True(t,
		!updated.UpdatedAt.Before(original.UpdatedAt),
		"updated_at must be >= previous updated_at after update: original=%v got=%v",
		original.UpdatedAt, updated.UpdatedAt,
	)
	assert.Equal(t, "updated name", updated.Name, "updated field must be persisted")
}
