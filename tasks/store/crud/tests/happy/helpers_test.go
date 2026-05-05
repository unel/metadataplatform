package happy_test

// Вспомогательные функции: FS-стор, builders для моделей.

import (
	"context"
	"encoding/json"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
	"github.com/unel/metadataplatform/store/fs"
	"github.com/unel/metadataplatform/store/router"
)

// newFSStore создаёт fs.Store во временной директории.
func newFSStore(t *testing.T) (*fs.Store, string) {
	t.Helper()
	basedir := t.TempDir()
	s, err := fs.New(fs.Config{Basedir: basedir}, nil)
	require.NoError(t, err, "fs.New must not fail in test setup")
	return s, basedir
}

// statDir — обёртка для os.Stat.
func statDir(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// rawJSON конвертирует строку в json.RawMessage.
func rawJSON(s string) json.RawMessage {
	return json.RawMessage(s)
}

// routerConn возвращает clientConn через net.Pipe с запущенным Handle.
func routerConn(t *testing.T, r *router.Router) net.Conn {
	t.Helper()
	server, client := net.Pipe()
	go func() {
		r.Handle(context.Background(), server)
		server.Close()
	}()
	t.Cleanup(func() {
		client.Close()
		server.Close()
	})
	return client
}

// sendRequest отправляет JSON-запрос и читает один JSON-ответ.
func sendRequest(t *testing.T, conn net.Conn, req any) map[string]any {
	t.Helper()
	data, err := json.Marshal(req)
	require.NoError(t, err)
	_, err = conn.Write(append(data, '\n'))
	require.NoError(t, err)

	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	defer conn.SetReadDeadline(time.Time{})

	dec := json.NewDecoder(conn)
	var resp map[string]any
	require.NoError(t, dec.Decode(&resp))
	return resp
}

// defaultEntity возвращает Entity с минимально заполненными полями.
func defaultEntity(id string) store.Entity {
	return store.Entity{
		ID:   id,
		Type: "file",
		Meta: rawJSON(`{}`),
	}
}

// defaultRelation возвращает Relation с минимально заполненными полями.
func defaultRelation(id, fromID, toID string) store.Relation {
	return store.Relation{
		ID:     id,
		FromID: fromID,
		ToID:   toID,
		Type:   "derivedFrom",
		Value:  rawJSON(`{}`),
		Meta:   rawJSON(`{}`),
	}
}

// defaultJob возвращает Job с минимально заполненными полями.
func defaultJob(id, entityID string) store.Job {
	return store.Job{
		ID:       id,
		EntityID: entityID,
		Kind:     "hash",
		Worker:   "hash.sha256",
		Status:   "pending",
		Payload:  rawJSON(`{}`),
	}
}
