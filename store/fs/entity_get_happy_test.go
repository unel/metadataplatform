package fs_test

// AC-ENTITY-GET-01
// AC-ENTITY-GET-03: см. entity_get_adversarial_test.go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AC-ENTITY-GET-01: получение существующей entity возвращает корректные данные.
func TestEntity_Get_ExistingRecord_ReturnsCorrectData(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	ctx := context.Background()

	entity := defaultEntity("entity-get-01")
	entity.Subtype = "image"
	entity.Name = "photo.jpg"
	entity.Meta = rawJSON(`{"path":"/photos/photo.jpg"}`)
	require.NoError(t, s.Upsert(ctx, entity))

	// Act
	got, err := s.Get(ctx, entity.ID)

	// Assert
	require.NoError(t, err, "Get of existing entity must succeed")
	assert.Equal(t, entity.ID, got.ID)
	assert.Equal(t, entity.Type, got.Type)
	assert.Equal(t, entity.Subtype, got.Subtype)
	assert.Equal(t, entity.Name, got.Name)
	assert.JSONEq(t, string(entity.Meta), string(got.Meta))
}
