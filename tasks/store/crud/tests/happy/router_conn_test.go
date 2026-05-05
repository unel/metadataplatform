package happy_test

// AC-ROUTER-11, AC-ROUTER-12, AC-ROUTER-13
// Тесты поведения роутера на уровне соединения.

import (
	"context"
	"encoding/json"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store/router"
)

// AC-ROUTER-11: поле id в envelope при upsert игнорируется — id берётся из data.
func TestRouter_Upsert_EnvelopeIDIgnored_TakesIDFromData(t *testing.T) {
	// Arrange
	r, _, s := newRouterWithSlogBuf(t)
	conn := routerConn(t, r)

	req := map[string]any{
		"op":   "upsert",
		"type": "entity",
		"id":   "envelope-id-should-be-ignored",
		"data": map[string]any{
			"id":   "data-id-is-canonical",
			"type": "file",
			"meta": map[string]any{},
		},
	}

	// Act
	resp := sendRequest(t, conn, req)

	// Assert
	require.Equal(t, true, resp["ok"])
	assert.Equal(t, "data-id-is-canonical", resp["id"],
		"id in response must come from data, not from envelope")

	ctx := context.Background()
	_, err := s.Get(ctx, "data-id-is-canonical")
	assert.NoError(t, err, "entity with data id must be findable in store")
}

// AC-ROUTER-12: несколько запросов в одном соединении — все обрабатываются.
func TestRouter_MultipleRequestsOnSingleConnection_AllProcessed(t *testing.T) {
	// Arrange
	r, _, _ := newRouterWithSlogBuf(t)
	conn := routerConn(t, r)

	ids := []string{"multi-req-01", "multi-req-02", "multi-req-03"}

	for _, id := range ids {
		req := map[string]any{
			"op":   "upsert",
			"type": "entity",
			"data": map[string]any{
				"id":   id,
				"type": "file",
				"meta": map[string]any{},
			},
		}

		// Act
		resp := sendRequest(t, conn, req)

		// Assert
		assert.Equalf(t, true, resp["ok"], "request id=%s must return ok=true", id)
		assert.Equalf(t, id, resp["id"], "request id=%s must return correct id", id)
	}
}

// AC-ROUTER-13: клиент закрывает соединение — роутер получает EOF и выходит штатно.
func TestRouter_ClientClosesConnection_RouterExitsGracefully(t *testing.T) {
	// Arrange
	s, _ := newFSStore(t)
	r := router.New(s.Entities(), s.Relations(), s.Jobs(), nil)

	srv, cli := net.Pipe()
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		r.Handle(context.Background(), srv)
		srv.Close()
	}()
	t.Cleanup(func() {
		srv.Close()
		cli.Close()
	})

	// отправляем один запрос — убеждаемся что соединение работает
	req := map[string]any{"op": "list", "type": "entity"}
	data, _ := json.Marshal(req)
	_, writeErr := cli.Write(append(data, '\n'))
	require.NoError(t, writeErr)

	cli.SetReadDeadline(time.Now().Add(time.Second))
	dec := json.NewDecoder(cli)
	var resp map[string]any
	require.NoError(t, dec.Decode(&resp))
	cli.SetReadDeadline(time.Time{})

	// Act: клиент закрывает соединение
	cli.Close()

	// Assert: роутер завершается без паники и зависания
	select {
	case <-doneCh:
		// роутер вышел штатно — OK
	case <-time.After(2 * time.Second):
		t.Fatal("router did not exit after client closed connection (possible goroutine leak)")
	}
}
