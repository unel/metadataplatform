package router

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/unel/metadataplatform/store"
)

// handleUpsert обрабатывает op=upsert.
func (r *Router) handleUpsert(ctx context.Context, env envelope) any {
	// data обязательна для upsert
	if env.Data == nil || string(env.Data) == "null" {
		return response{OK: false, ErrorCode: "INVALID_REQUEST", Error: "data is required for upsert"}
	}

	switch env.Type {
	case "entity":
		var e store.Entity
		if err := json.Unmarshal(env.Data, &e); err != nil {
			return response{OK: false, ErrorCode: "INVALID_REQUEST", Error: "invalid data"}
		}
		if err := r.entities.Upsert(ctx, e); err != nil {
			r.log.Error("upsert error",
				slog.String("op", "upsert"),
				slog.String("type", "entity"),
				slog.String("id", e.ID),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("upsert ok",
			slog.String("op", "upsert"),
			slog.String("type", "entity"),
			slog.String("id", e.ID),
		)
		return responseWithID{OK: true, ID: e.ID}

	case "relation":
		var rel store.Relation
		if err := json.Unmarshal(env.Data, &rel); err != nil {
			return response{OK: false, ErrorCode: "INVALID_REQUEST", Error: "invalid data"}
		}
		if err := r.relations.Upsert(ctx, rel); err != nil {
			r.log.Error("upsert error",
				slog.String("op", "upsert"),
				slog.String("type", "relation"),
				slog.String("id", rel.ID),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("upsert ok",
			slog.String("op", "upsert"),
			slog.String("type", "relation"),
			slog.String("id", rel.ID),
		)
		return responseWithID{OK: true, ID: rel.ID}

	case "job":
		var j store.Job
		if err := json.Unmarshal(env.Data, &j); err != nil {
			return response{OK: false, ErrorCode: "INVALID_REQUEST", Error: "invalid data"}
		}
		if err := r.jobs.Upsert(ctx, j); err != nil {
			r.log.Error("upsert error",
				slog.String("op", "upsert"),
				slog.String("type", "job"),
				slog.String("id", j.ID),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("upsert ok",
			slog.String("op", "upsert"),
			slog.String("type", "job"),
			slog.String("id", j.ID),
		)
		return responseWithID{OK: true, ID: j.ID}

	default:
		return response{OK: false, ErrorCode: "UNKNOWN_TYPE", Error: fmt.Sprintf("unknown type: %s", env.Type)}
	}
}
