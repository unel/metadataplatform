package adversarial_test

// AC-JOB-UPSERT-02, AC-JOB-GET-02, AC-JOB-DELETE-02
// Job — третий тип, те же failure modes. Проверяем что реализация не забыла про jobs/.

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
	"github.com/unel/metadataplatform/store/fs"
)

func newJobStore(t *testing.T) store.JobStore {
	t.Helper()
	s, err := fs.New(fs.Config{Basedir: t.TempDir()}, testLogger(t))
	require.NoError(t, err)
	return s.Jobs()
}

// AC-JOB-UPSERT-02: пустой id → ошибка.
func TestJob_Upsert_EmptyID_ReturnsError(t *testing.T) {
	js := newJobStore(t)

	err := js.Upsert(context.Background(), store.Job{
		ID:     "",
		Kind:   "index",
		Worker: "w1",
		Status: "pending",
	})
	assert.Error(t, err, "пустой id должен вернуть ошибку")
}

// AC-JOB-GET-02: несуществующий job → ErrNotFound.
func TestJob_Get_NonExistentID_ReturnsErrNotFound(t *testing.T) {
	js := newJobStore(t)

	_, err := js.Get(context.Background(), "job-999")
	require.Error(t, err)
	assert.True(t, errors.Is(err, store.ErrNotFound),
		"ожидаем store.ErrNotFound, получили: %v", err)
}

// AC-JOB-DELETE-02: delete несуществующего job → ErrNotFound (не идемпотентный).
func TestJob_Delete_NonExistentID_ReturnsErrNotFound(t *testing.T) {
	js := newJobStore(t)

	err := js.Delete(context.Background(), "job-999")
	require.Error(t, err)
	assert.True(t, errors.Is(err, store.ErrNotFound),
		"delete несуществующего job должен вернуть ErrNotFound, получили: %v", err)
}

// Проверяем что job-директория называется jobs/, а не job/ или что-то ещё.
// Косвенно: создаём job и убеждаемся что он находится через Get.
func TestJob_Upsert_Then_Get_PathIsCorrect(t *testing.T) {
	js := newJobStore(t)
	ctx := context.Background()

	require.NoError(t, js.Upsert(ctx, store.Job{
		ID:     "job-path-check",
		Kind:   "hash",
		Worker: "hash.sha256",
		Status: "pending",
	}))

	got, err := js.Get(ctx, "job-path-check")
	require.NoError(t, err)
	assert.Equal(t, "job-path-check", got.ID)
}
