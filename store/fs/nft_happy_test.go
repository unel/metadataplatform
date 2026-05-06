package fs_test

// AC-NFT-ATOMIC-02: uuid.NewV7() используется для генерации ID на стороне клиента.
// AC-ROUTER-INIT-01: роутер принимает интерфейсы (дополнительная проверка через compile-time).

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AC-NFT-ATOMIC-02: uuid.NewV7() генерирует валидный UUID версии 7.
// Версия 7 кодирует timestamp в старших битах — IDs монотонно возрастают.
func TestUUID_NewV7_GeneratesVersionSevenUUID(t *testing.T) {
	// Act
	id, err := uuid.NewV7()

	// Assert
	require.NoError(t, err, "uuid.NewV7 must not return error")
	assert.Equal(t, uuid.Version(7), id.Version(),
		"generated UUID must be version 7, not %d", id.Version())
}

// AC-NFT-ATOMIC-02: UUID v7 монотонно возрастает — позволяет cursor-based пагинацию.
func TestUUID_NewV7_SequentialIDs_AreMonotonicallyIncreasing(t *testing.T) {
	// Arrange: генерируем несколько ID с небольшой паузой
	const count = 5
	ids := make([]uuid.UUID, count)
	for i := range ids {
		id, err := uuid.NewV7()
		require.NoError(t, err)
		ids[i] = id
		if i < count-1 {
			time.Sleep(time.Millisecond)
		}
	}

	// Assert: каждый следующий ID >= предыдущего (UUIDv7 лексикографически монотонен)
	for i := 1; i < count; i++ {
		prev := ids[i-1].String()
		curr := ids[i].String()
		assert.True(t, curr >= prev,
			"UUIDv7[%d]=%s must be >= UUIDv7[%d]=%s", i, curr, i-1, prev)
	}
}

// AC-NFT-ATOMIC-02: строковое представление UUID v7 совместимо с форматом поля id в Entity.
func TestUUID_NewV7_StringRepresentation_CompatibleWithEntityID(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	ctx := context.Background()

	id, err := uuid.NewV7()
	require.NoError(t, err)
	idStr := id.String()

	entity := defaultEntity(idStr)

	// Act: upsert с UUID v7 как id
	upsertErr := s.Upsert(ctx, entity)
	require.NoError(t, upsertErr)

	// Assert: entity доступна по тому же ID
	got, getErr := s.Get(ctx, idStr)
	require.NoError(t, getErr)
	assert.Equal(t, idStr, got.ID,
		"entity stored with UUIDv7 id must be retrievable by the same id")
}
