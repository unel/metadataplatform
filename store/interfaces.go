package store

import "context"

// EntityStore — CRUD для Entity.
type EntityStore interface {
	Upsert(ctx context.Context, e Entity) error
	Get(ctx context.Context, id string) (Entity, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]Entity, error)
}

// RelationStore — CRUD для Relation.
type RelationStore interface {
	Upsert(ctx context.Context, r Relation) error
	Get(ctx context.Context, id string) (Relation, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]Relation, error)
}

// JobStore — CRUD для Job.
type JobStore interface {
	Upsert(ctx context.Context, j Job) error
	Get(ctx context.Context, id string) (Job, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]Job, error)
}
