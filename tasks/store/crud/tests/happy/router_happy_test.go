package happy_test

// AC-ROUTER-01, AC-ROUTER-02, AC-ROUTER-03, AC-ROUTER-04
// AC-ROUTER-INIT-01

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
	"github.com/unel/metadataplatform/store/fs"
	"github.com/unel/metadataplatform/store/router"
)

// AC-ROUTER-INIT-01: роутер принимает интерфейсы, не конкретные типы.
// Если это компилируется — интерфейсы соблюдены.
func TestRouter_New_AcceptsStoreInterfaces(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	var entities store.EntityStore = s.Entities()
	var relations store.RelationStore = s.Relations()
	var jobs store.JobStore = s.Jobs()

	// Act
	r := router.New(entities, relations, jobs, nil)

	// Assert
	assert.NotNil(t, r, "router.New must return non-nil *Router")
}

// AC-ROUTER-01: upsert entity через роутер — успех, ответ содержит ok=true и id.
func TestRouter_Upsert_Entity_ReturnsOkWithID(t *testing.T) {
	// Arrange
	r, _, _ := newRouterWithSlogBuf(t)
	conn := routerConn(t, r)

	req := map[string]any{
		"op":   "upsert",
		"type": "entity",
		"data": map[string]any{
			"id":   "router-entity-01",
			"type": "file",
			"meta": map[string]any{},
		},
	}

	// Act
	resp := sendRequest(t, conn, req)

	// Assert
	assert.Equal(t, true, resp["ok"], "upsert entity must return ok=true")
	assert.Equal(t, "router-entity-01", resp["id"], "upsert entity must return id from data")
}

// AC-ROUTER-02: get entity через роутер — успех, ответ содержит data с полями entity.
func TestRouter_Get_ExistingEntity_ReturnsOkWithData(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	r := router.New(s.Entities(), s.Relations(), s.Jobs(), nil)
	ctx := context.Background()
	entity := defaultEntity("router-get-entity-01")
	require.NoError(t, s.Entities().Upsert(ctx, entity))

	conn := routerConn(t, r)
	req := map[string]any{
		"op":   "get",
		"type": "entity",
		"id":   entity.ID,
	}

	// Act
	resp := sendRequest(t, conn, req)

	// Assert
	assert.Equal(t, true, resp["ok"])
	data, ok := resp["data"].(map[string]any)
	require.True(t, ok, "response must contain 'data' object")
	assert.Equal(t, entity.ID, data["id"])
	assert.Equal(t, entity.Type, data["type"])
}

// AC-ROUTER-03: delete entity через роутер — успех, ответ содержит ok=true.
func TestRouter_Delete_ExistingEntity_ReturnsOk(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	r := router.New(s.Entities(), s.Relations(), s.Jobs(), nil)
	ctx := context.Background()
	entity := defaultEntity("router-delete-entity-01")
	require.NoError(t, s.Entities().Upsert(ctx, entity))

	conn := routerConn(t, r)
	req := map[string]any{
		"op":   "delete",
		"type": "entity",
		"id":   entity.ID,
	}

	// Act
	resp := sendRequest(t, conn, req)

	// Assert
	assert.Equal(t, true, resp["ok"])
}

// AC-ROUTER-04: list entity через роутер — успех, ответ содержит data массив.
func TestRouter_List_Entities_ReturnsOkWithDataArray(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	r := router.New(s.Entities(), s.Relations(), s.Jobs(), nil)
	ctx := context.Background()
	require.NoError(t, s.Entities().Upsert(ctx, defaultEntity("router-list-e-01")))
	require.NoError(t, s.Entities().Upsert(ctx, defaultEntity("router-list-e-02")))

	conn := routerConn(t, r)
	req := map[string]any{
		"op":   "list",
		"type": "entity",
	}

	// Act
	resp := sendRequest(t, conn, req)

	// Assert
	assert.Equal(t, true, resp["ok"])
	data, ok := resp["data"].([]any)
	require.True(t, ok, "list response must contain 'data' array")
	assert.Len(t, data, 2)
}

// Compile-time: *fs.Store реализует EntityStore напрямую.
// Relations() и Jobs() возвращают соответствующие интерфейсы — проверяется в TestRouter_New_AcceptsStoreInterfaces.
var _ store.EntityStore = (*fs.Store)(nil)
