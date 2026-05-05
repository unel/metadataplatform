package router

import (
	"context"
	"fmt"
	"log/slog"
)

// handleList обрабатывает op=list.
func (r *Router) handleList(ctx context.Context, env envelope) any {
	switch env.Type {
	case "entity":
		items, err := r.entities.List(ctx)
		if err != nil {
			r.log.Error("list error",
				slog.String("op", "list"),
				slog.String("type", "entity"),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("list ok",
			slog.String("op", "list"),
			slog.String("type", "entity"),
			slog.Int("count", len(items)),
		)
		return responseWithData{OK: true, Data: items}

	case "relation":
		items, err := r.relations.List(ctx)
		if err != nil {
			r.log.Error("list error",
				slog.String("op", "list"),
				slog.String("type", "relation"),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("list ok",
			slog.String("op", "list"),
			slog.String("type", "relation"),
			slog.Int("count", len(items)),
		)
		return responseWithData{OK: true, Data: items}

	case "job":
		items, err := r.jobs.List(ctx)
		if err != nil {
			r.log.Error("list error",
				slog.String("op", "list"),
				slog.String("type", "job"),
				slog.String("error", err.Error()),
			)
			return mapStoreError(err)
		}
		r.log.Debug("list ok",
			slog.String("op", "list"),
			slog.String("type", "job"),
			slog.Int("count", len(items)),
		)
		return responseWithData{OK: true, Data: items}

	default:
		return response{OK: false, ErrorCode: "UNKNOWN_TYPE", Error: fmt.Sprintf("unknown type: %s", env.Type)}
	}
}
