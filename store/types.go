package store

import (
	"encoding/json"
	"errors"
	"time"
)

// Sentinel errors для операций стора.
var (
	ErrNotFound    = errors.New("not found")
	ErrMissingID   = errors.New("id is required")
	ErrReadRecord  = errors.New("failed to read record")
	ErrWriteRecord = errors.New("failed to write record")
	ErrDeleteRecord = errors.New("failed to delete record")
	ErrListRecords  = errors.New("failed to list records")
	ErrReadExisting = errors.New("failed to read existing record")
)

// Entity — любой объект в системе.
type Entity struct {
	ID          string          `json:"id"`
	Type        string          `json:"type"`
	Subtype     string          `json:"subtype,omitempty"`
	Name        string          `json:"name,omitempty"`
	Description string          `json:"description,omitempty"`
	Meta        json.RawMessage `json:"meta"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// Relation — направленная связь между двумя entity.
type Relation struct {
	ID        string          `json:"id"`
	FromID    string          `json:"from_id"`
	ToID      string          `json:"to_id"`
	Type      string          `json:"type"`
	Subtype   string          `json:"subtype,omitempty"`
	Value     json.RawMessage `json:"value"`
	Meta      json.RawMessage `json:"meta"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// Job — задача обработки.
type Job struct {
	ID         string          `json:"id"`
	EntityID   string          `json:"entity_id,omitempty"`
	RelationID string          `json:"relation_id,omitempty"`
	Kind       string          `json:"kind"`
	Worker     string          `json:"worker"`
	Status     string          `json:"status"`
	Progress   json.RawMessage `json:"progress,omitempty"`
	Error      string          `json:"error,omitempty"`
	Payload    json.RawMessage `json:"payload"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
}
