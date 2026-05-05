package happy_test

// AC-ROUTER-14, AC-ROUTER-15
// AC-ROUTER-16: см. adversarial/router_error_test.go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// AC-ROUTER-14: после успешного upsert роутер логирует DEBUG op=upsert type=entity id=<id>.
func TestRouter_Upsert_Entity_LogsDebugWithOpTypeID(t *testing.T) {
	r, buf, _ := newRouterWithSlogBuf(t)
	conn := routerConn(t, r)

	entityID := "log-upsert-entity-01"
	req := map[string]any{
		"op":   "upsert",
		"type": "entity",
		"data": map[string]any{
			"id":   entityID,
			"type": "file",
			"meta": map[string]any{},
		},
	}

	resp := sendRequest(t, conn, req)

	require.Equal(t, true, resp["ok"])
	logOutput := buf.String()
	assert.Contains(t, logOutput, "DEBUG", "router must log DEBUG entry for successful upsert entity")
	assert.Contains(t, logOutput, "op=upsert", "DEBUG log must contain op=upsert")
	assert.Contains(t, logOutput, "type=entity", "DEBUG log must contain type=entity")
	assert.Contains(t, logOutput, entityID, "DEBUG log must contain entity id")
}

// AC-ROUTER-14: после успешного get роутер логирует DEBUG op=get type=entity id=<id>.
func TestRouter_Get_Entity_LogsDebugWithOpTypeID(t *testing.T) {
	r, buf, s := newRouterWithSlogBuf(t)
	ctx := context.Background()
	entityID := "log-get-entity-01"
	require.NoError(t, s.Entities().Upsert(ctx, defaultEntity(entityID)))

	conn := routerConn(t, r)
	req := map[string]any{
		"op":   "get",
		"type": "entity",
		"id":   entityID,
	}

	resp := sendRequest(t, conn, req)

	require.Equal(t, true, resp["ok"])
	logOutput := buf.String()
	assert.Contains(t, logOutput, "DEBUG", "router must log DEBUG entry for successful get entity")
	assert.Contains(t, logOutput, "op=get", "DEBUG log must contain op=get")
	assert.Contains(t, logOutput, "type=entity", "DEBUG log must contain type=entity")
	assert.Contains(t, logOutput, entityID, "DEBUG log must contain entity id")
}

// AC-ROUTER-15: после успешного list роутер логирует DEBUG op=list type=entity count=N.
func TestRouter_List_Entities_LogsDebugWithCount(t *testing.T) {
	r, buf, s := newRouterWithSlogBuf(t)
	ctx := context.Background()
	require.NoError(t, s.Entities().Upsert(ctx, defaultEntity("log-list-e-01")))
	require.NoError(t, s.Entities().Upsert(ctx, defaultEntity("log-list-e-02")))

	conn := routerConn(t, r)
	req := map[string]any{
		"op":   "list",
		"type": "entity",
	}

	resp := sendRequest(t, conn, req)

	require.Equal(t, true, resp["ok"])
	logOutput := buf.String()
	assert.Contains(t, logOutput, "DEBUG", "router must log DEBUG entry for successful list entity")
	assert.Contains(t, logOutput, "op=list", "DEBUG log must contain op=list")
	assert.Contains(t, logOutput, "type=entity", "DEBUG log must contain type=entity")
	assert.Contains(t, logOutput, "count=2", "DEBUG log must contain count=2")
}
