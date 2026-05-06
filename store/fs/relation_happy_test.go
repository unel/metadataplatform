package fs_test

// AC-RELATION-UPSERT-01, AC-RELATION-UPSERT-02
// AC-RELATION-GET-01, AC-RELATION-DELETE-01
// AC-RELATION-LIST-01, AC-RELATION-LIST-02

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
)

// AC-RELATION-UPSERT-01: создание новой relation — timestamps выставляются стором.
func TestRelation_Upsert_NewRecord_SetsTimestamps(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	rs := s.Relations()
	ctx := context.Background()
	before := time.Now().UTC()

	rel := defaultRelation("rel-upsert-01", "from-entity-01", "to-entity-01")
	rel.CreatedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	rel.UpdatedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	// Act
	err := rs.Upsert(ctx, rel)
	after := time.Now().UTC()

	// Assert
	require.NoError(t, err)
	got, getErr := rs.Get(ctx, rel.ID)
	require.NoError(t, getErr)

	assert.True(t,
		!got.CreatedAt.Before(before) && !got.CreatedAt.After(after),
		"created_at must be set by store: got %v", got.CreatedAt,
	)
	assert.True(t,
		!got.UpdatedAt.Before(before) && !got.UpdatedAt.After(after),
		"updated_at must be set by store: got %v", got.UpdatedAt,
	)
}

// AC-RELATION-UPSERT-02: FS не проверяет существование from_id / to_id.
func TestRelation_Upsert_NonexistentFromToIDs_Succeeds(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	rs := s.Relations()
	ctx := context.Background()

	rel := defaultRelation("rel-upsert-02", "nonexistent-from", "nonexistent-to")

	// Act
	err := rs.Upsert(ctx, rel)

	// Assert
	assert.NoError(t, err,
		"FS store must not check FK constraints: upsert with nonexistent from_id/to_id must succeed",
	)
}

// AC-RELATION-GET-01: получение существующей relation возвращает корректные данные.
func TestRelation_Get_ExistingRecord_ReturnsCorrectData(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	rs := s.Relations()
	ctx := context.Background()

	rel := defaultRelation("rel-get-01", "from-a", "to-b")
	rel.Subtype = "preview"
	rel.Value = rawJSON(`{"generator":"preview-worker:v1","format":"jpeg"}`)
	require.NoError(t, rs.Upsert(ctx, rel))

	// Act
	got, err := rs.Get(ctx, rel.ID)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, rel.ID, got.ID)
	assert.Equal(t, rel.FromID, got.FromID)
	assert.Equal(t, rel.ToID, got.ToID)
	assert.Equal(t, rel.Type, got.Type)
	assert.Equal(t, rel.Subtype, got.Subtype)
	assert.JSONEq(t, string(rel.Value), string(got.Value))
}

// AC-RELATION-DELETE-01: удаление существующей relation — запись пропадает.
func TestRelation_Delete_ExistingRecord_RecordNoLongerFound(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	rs := s.Relations()
	ctx := context.Background()

	rel := defaultRelation("rel-delete-01", "from-x", "to-y")
	require.NoError(t, rs.Upsert(ctx, rel))

	// Act
	err := rs.Delete(ctx, rel.ID)

	// Assert
	require.NoError(t, err)
	_, getErr := rs.Get(ctx, rel.ID)
	assert.True(t,
		errors.Is(getErr, store.ErrNotFound),
		"after Delete, Get must return ErrNotFound: got %v", getErr,
	)
}

// AC-RELATION-LIST-01: list возвращает все сохранённые relations.
func TestRelation_List_MultipleRecords_ReturnsAll(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	rs := s.Relations()
	ctx := context.Background()

	ids := []string{"rel-list-a", "rel-list-b"}
	for _, id := range ids {
		require.NoError(t, rs.Upsert(ctx, defaultRelation(id, "from-e", "to-e")))
	}

	// Act
	result, err := rs.List(ctx)

	// Assert
	require.NoError(t, err)
	require.Len(t, result, len(ids))

	gotIDs := make([]string, len(result))
	for i, r := range result {
		gotIDs[i] = r.ID
	}
	assert.ElementsMatch(t, ids, gotIDs)
}

// AC-RELATION-LIST-02: list на пустом сторе возвращает пустой срез.
func TestRelation_List_EmptyStore_ReturnsEmptySlice(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	rs := s.Relations()
	ctx := context.Background()

	// Act
	result, err := rs.List(ctx)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}
