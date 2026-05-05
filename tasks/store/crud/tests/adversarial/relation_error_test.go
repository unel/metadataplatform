package adversarial_test

// AC-RELATION-UPSERT-03, AC-RELATION-GET-02, AC-RELATION-DELETE-02
// Relations — те же паттерны что Entity, но другой путь. Проверяем что путь правильный.

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
	"github.com/unel/metadataplatform/store/fs"
)

func newRelationStore(t *testing.T) store.RelationStore {
	t.Helper()
	s, err := fs.New(fs.Config{Basedir: t.TempDir()}, testLogger(t))
	require.NoError(t, err)
	return s.Relations()
}

// AC-RELATION-UPSERT-03: пустой id → ошибка.
// FS не проверяет FK — но MISSING_ID всё равно должен быть.
func TestRelation_Upsert_EmptyID_ReturnsError(t *testing.T) {
	rs := newRelationStore(t)

	err := rs.Upsert(context.Background(), store.Relation{
		ID:     "",
		FromID: "ent-a",
		ToID:   "ent-b",
		Type:   "uses",
	})
	assert.Error(t, err, "пустой id должен вернуть ошибку")
}

// AC-RELATION-GET-02: несуществующая relation → ErrNotFound.
func TestRelation_Get_NonExistentID_ReturnsErrNotFound(t *testing.T) {
	rs := newRelationStore(t)

	_, err := rs.Get(context.Background(), "rel-999")
	require.Error(t, err)
	assert.True(t, errors.Is(err, store.ErrNotFound),
		"ожидаем store.ErrNotFound, получили: %v", err)
}

// AC-RELATION-DELETE-02: delete несуществующей relation → ErrNotFound (не идемпотентный).
func TestRelation_Delete_NonExistentID_ReturnsErrNotFound(t *testing.T) {
	rs := newRelationStore(t)

	err := rs.Delete(context.Background(), "rel-999")
	require.Error(t, err)
	assert.True(t, errors.Is(err, store.ErrNotFound),
		"delete несуществующей relation должен вернуть ErrNotFound, получили: %v", err)
}

// FS не проверяет from_id/to_id — это explicit по спеке.
// Тест документирует это поведение чтобы никто не "починил" его случайно.
func TestRelation_Upsert_NonExistentFromToIDs_Succeeds(t *testing.T) {
	rs := newRelationStore(t)

	err := rs.Upsert(context.Background(), store.Relation{
		ID:     "rel-fk",
		FromID: "ghost-entity-1",
		ToID:   "ghost-entity-2",
		Type:   "uses",
	})
	assert.NoError(t, err,
		"FS-реализация намеренно не проверяет FK — это задокументированное ограничение §7")
}
