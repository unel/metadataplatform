package fs_test

// AC-JOB-UPSERT-01, AC-JOB-GET-01, AC-JOB-DELETE-01
// AC-JOB-LIST-01, AC-JOB-LIST-02

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
)

// AC-JOB-UPSERT-01: создание новой job — timestamps выставляются стором,
// входящие игнорируются; все поля сохраняются корректно.
func TestJob_Upsert_NewRecord_SetsTimestampsAndPersistsFields(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	js := s.Jobs()
	ctx := context.Background()
	before := time.Now().UTC()

	job := defaultJob("job-upsert-01", "entity-abc")
	job.CreatedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	job.UpdatedAt = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	job.Payload = rawJSON(`{"source":"/data/file.mp4"}`)

	// Act
	err := js.Upsert(ctx, job)
	after := time.Now().UTC()

	// Assert
	require.NoError(t, err)
	got, getErr := js.Get(ctx, job.ID)
	require.NoError(t, getErr)

	assert.True(t,
		!got.CreatedAt.Before(before) && !got.CreatedAt.After(after),
		"created_at must be set by store: got %v", got.CreatedAt,
	)
	assert.True(t,
		!got.UpdatedAt.Before(before) && !got.UpdatedAt.After(after),
		"updated_at must be set by store: got %v", got.UpdatedAt,
	)
	assert.Equal(t, job.ID, got.ID)
	assert.Equal(t, job.EntityID, got.EntityID)
	assert.Equal(t, job.Kind, got.Kind)
	assert.Equal(t, job.Worker, got.Worker)
	assert.Equal(t, job.Status, got.Status)
	assert.JSONEq(t, string(job.Payload), string(got.Payload))
}

// AC-JOB-GET-01: получение существующей job возвращает корректные данные.
func TestJob_Get_ExistingRecord_ReturnsCorrectData(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	js := s.Jobs()
	ctx := context.Background()

	job := defaultJob("job-get-01", "entity-xyz")
	job.Status = "running"
	job.Progress = rawJSON(`{"done":10,"total":100}`)
	require.NoError(t, js.Upsert(ctx, job))

	// Act
	got, err := js.Get(ctx, job.ID)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, job.ID, got.ID)
	assert.Equal(t, job.Status, got.Status)
	assert.JSONEq(t, string(job.Progress), string(got.Progress))
}

// AC-JOB-DELETE-01: удаление существующей job — запись пропадает.
func TestJob_Delete_ExistingRecord_RecordNoLongerFound(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	js := s.Jobs()
	ctx := context.Background()

	job := defaultJob("job-delete-01", "entity-del")
	require.NoError(t, js.Upsert(ctx, job))

	// Act
	err := js.Delete(ctx, job.ID)

	// Assert
	require.NoError(t, err)
	_, getErr := js.Get(ctx, job.ID)
	assert.True(t,
		errors.Is(getErr, store.ErrNotFound),
		"after Delete, Get must return ErrNotFound: got %v", getErr,
	)
}

// AC-JOB-LIST-01: list возвращает все сохранённые jobs.
func TestJob_List_MultipleRecords_ReturnsAll(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	js := s.Jobs()
	ctx := context.Background()

	ids := []string{"job-list-a", "job-list-b", "job-list-c"}
	for _, id := range ids {
		require.NoError(t, js.Upsert(ctx, defaultJob(id, "entity-e")))
	}

	// Act
	result, err := js.List(ctx)

	// Assert
	require.NoError(t, err)
	require.Len(t, result, len(ids))

	gotIDs := make([]string, len(result))
	for i, j := range result {
		gotIDs[i] = j.ID
	}
	assert.ElementsMatch(t, ids, gotIDs)
}

// AC-JOB-LIST-02: list на пустом сторе возвращает пустой срез.
func TestJob_List_EmptyStore_ReturnsEmptySlice(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	js := s.Jobs()
	ctx := context.Background()

	// Act
	result, err := js.List(ctx)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 0)
}
