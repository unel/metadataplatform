package happy_test

// AC-ENTITY-DELETE-01

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
)

// AC-ENTITY-DELETE-01: удаление существующей entity — запись пропадает из стора.
func TestEntity_Delete_ExistingRecord_RecordNoLongerFound(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	ctx := context.Background()

	entity := defaultEntity("entity-delete-01")
	require.NoError(t, s.Upsert(ctx, entity))

	// предусловие: запись существует
	_, preErr := s.Get(ctx, entity.ID)
	require.NoError(t, preErr, "pre-condition: entity must exist before delete")

	// Act
	err := s.Delete(ctx, entity.ID)

	// Assert
	require.NoError(t, err, "Delete of existing entity must succeed")

	_, getErr := s.Get(ctx, entity.ID)
	assert.True(t,
		errors.Is(getErr, store.ErrNotFound),
		"after Delete, Get must return ErrNotFound: got %v", getErr,
	)
}
