package router

import (
	"context"
	"fmt"
	"log/slog"
)

// handleGet обрабатывает op=get.
func (r *Router) handleGet(ctx context.Context, env envelope) any {
	id := env.ID

	switch env.Type {
	case "entity":
		e, err := r.entities.Get(ctx, id)
		if err != nil {
			r.log.Error("get error",
				slog.String("op", "get"),
				slog.String("type", "entity"),
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("get ok",
			slog.String("op", "get"),
			slog.String("type", "entity"),
			slog.String("id", id),
		)
		return responseWithData{OK: true, Data: e}

	case "relation":
		rel, err := r.relations.Get(ctx, id)
		if err != nil {
			r.log.Error("get error",
				slog.String("op", "get"),
				slog.String("type", "relation"),
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("get ok",
			slog.String("op", "get"),
			slog.String("type", "relation"),
			slog.String("id", id),
		)
		return responseWithData{OK: true, Data: rel}

	case "job":
		j, err := r.jobs.Get(ctx, id)
		if err != nil {
			r.log.Error("get error",
				slog.String("op", "get"),
				slog.String("type", "job"),
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("get ok",
			slog.String("op", "get"),
			slog.String("type", "job"),
			slog.String("id", id),
		)
		return responseWithData{OK: true, Data: j}

	default:
		return response{OK: false, ErrorCode: "UNKNOWN_TYPE", Error: fmt.Sprintf("unknown type: %s", env.Type)}
	}
}
