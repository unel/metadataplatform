package router

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"

	"github.com/unel/metadataplatform/store"
)

// Router маппит входящие JSONL-команды на вызовы store-интерфейсов.
type Router struct {
	entities  store.EntityStore
	relations store.RelationStore
	jobs      store.JobStore
	log       *slog.Logger
}

// New создаёт роутер с переданными сторами.
// Сторы передаются через интерфейсы — роутер не зависит от конкретных реализаций.
func New(entities store.EntityStore, relations store.RelationStore, jobs store.JobStore, log *slog.Logger) *Router {
	if log == nil {
		log = slog.Default()
	}
	return &Router{
		entities:  entities,
		relations: relations,
		jobs:      jobs,
		log:       log,
	}
}

// envelope — первый pass decode входящего запроса.
type envelope struct {
	Op   string          `json:"op"`
	Type string          `json:"type"`
	ID   string          `json:"id"`
	Data json.RawMessage `json:"data"`
}

// response — формат всех ответов роутера.
type response struct {
	OK        bool   `json:"ok"`
	Error     string `json:"error,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
}

type responseWithID struct {
	OK        bool   `json:"ok"`
	ID        string `json:"id,omitempty"`
	Error     string `json:"error,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
}

type responseWithData struct {
	OK        bool   `json:"ok"`
	Data      any    `json:"data,omitempty"`
	Error     string `json:"error,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
}

// Handle обрабатывает JSONL-соединение: читает команды и пишет ответы до EOF или ошибки.
func (r *Router) Handle(ctx context.Context, conn net.Conn) {
	enc := json.NewEncoder(conn)
	dec := json.NewDecoder(conn)

	for {
		var env envelope
		if err := dec.Decode(&env); err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
				return
			}
			// parse error — поток испорчён, закрываем соединение
			r.log.Error("parse_error",
				slog.String("op", "unknown"),
				slog.String("parse_error", err.Error()),
			)
			conn.Close()
			return
		}

		resp := r.dispatch(ctx, env)
		_ = enc.Encode(resp)
	}
}

// dispatch обрабатывает один запрос и возвращает ответ для сериализации.
func (r *Router) dispatch(ctx context.Context, env envelope) any {
	switch env.Op {
	case "upsert":
		return r.handleUpsert(ctx, env)
	case "get":
		return r.handleGet(ctx, env)
	case "delete":
		return r.handleDelete(ctx, env)
	case "list":
		return r.handleList(ctx, env)
	default:
		r.log.Error("unknown op",
			slog.String("op", env.Op),
			slog.String("type", env.Type),
		)
		return response{OK: false, ErrorCode: "UNKNOWN_OP", Error: fmt.Sprintf("unknown op: %s", env.Op)}
	}
}

// mapStoreError преобразует ошибки стора в response с соответствующим errorCode.
func mapStoreError(err error) response {
	switch {
	case errors.Is(err, store.ErrNotFound):
		return response{OK: false, ErrorCode: "NOT_FOUND", Error: "not found"}
	case errors.Is(err, store.ErrMissingID):
		return response{OK: false, ErrorCode: "MISSING_ID", Error: "id is required"}
	case errors.Is(err, store.ErrReadExisting):
		return response{OK: false, ErrorCode: "READ_ERROR", Error: "failed to read existing record"}
	case errors.Is(err, store.ErrReadRecord):
		return response{OK: false, ErrorCode: "READ_ERROR", Error: "failed to read record"}
	case errors.Is(err, store.ErrWriteRecord):
		return response{OK: false, ErrorCode: "WRITE_ERROR", Error: "failed to write record"}
	case errors.Is(err, store.ErrDeleteRecord):
		return response{OK: false, ErrorCode: "DELETE_ERROR", Error: "failed to delete record"}
	case errors.Is(err, store.ErrListRecords):
		return response{OK: false, ErrorCode: "LIST_ERROR", Error: "failed to list records"}
	default:
		// TODO: в debug-режиме возвращать детали ошибки; сейчас скрываем детали реализации.
		return response{OK: false, ErrorCode: "INTERNAL_ERROR", Error: "internal error"}
	}
}
