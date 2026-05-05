package router

import (
	"context"
	"fmt"
	"log/slog"
)

// handleDelete обрабатывает op=delete.
func (r *Router) handleDelete(ctx context.Context, env envelope) any {
	id := env.ID

	switch env.Type {
	case "entity":
		if err := r.entities.Delete(ctx, id); err != nil {
			r.log.Error("delete error",
				slog.String("op", "delete"),
				slog.String("type", "entity"),
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("delete ok",
			slog.String("op", "delete"),
			slog.String("type", "entity"),
			slog.String("id", id),
		)
		return response{OK: true}

	case "relation":
		if err := r.relations.Delete(ctx, id); err != nil {
			r.log.Error("delete error",
				slog.String("op", "delete"),
				slog.String("type", "relation"),
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("delete ok",
			slog.String("op", "delete"),
			slog.String("type", "relation"),
			slog.String("id", id),
		)
		return response{OK: true}

	case "job":
		if err := r.jobs.Delete(ctx, id); err != nil {
			r.log.Error("delete error",
				slog.String("op", "delete"),
				slog.String("type", "job"),
				slog.String("id", id),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("delete ok",
			slog.String("op", "delete"),
			slog.String("type", "job"),
			slog.String("id", id),
		)
		return response{OK: true}

	default:
		return response{OK: false, ErrorCode: "UNKNOWN_TYPE", Error: fmt.Sprintf("unknown type: %s", env.Type)}
	}
}
