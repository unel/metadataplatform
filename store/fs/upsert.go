package fs

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/unel/metadataplatform/store"
)

// upsertRecord реализует общую логику upsert:
// читает существующую запись (сохраняет CreatedAt), выставляет timestamps, записывает атомарно.
// setTimestamps вызывается с (createdAt, updatedAt).
func upsertRecord(dir, path string, obj any, setTimestamps func(createdAt, updatedAt time.Time), readCreatedAt func([]byte) (time.Time, error)) error {
	existing, err := os.ReadFile(path)
	if err == nil {
		createdAt, err2 := readCreatedAt(existing)
		if err2 != nil {
			return fmt.Errorf("%w: %w", store.ErrReadExisting, err2)
		}
		setTimestamps(createdAt, time.Now().UTC())
	} else if os.IsNotExist(err) {
		now := time.Now().UTC()
		setTimestamps(now, now)
	} else {
		return fmt.Errorf("%w: %w", store.ErrReadExisting, err)
	}

	return atomicWrite(dir, path, obj)
}

func entityCreatedAt(data []byte) (time.Time, error) {
	var v store.Entity
	if err := json.Unmarshal(data, &v); err != nil {
		return time.Time{}, err
	}
	return v.CreatedAt, nil
}

func relationCreatedAt(data []byte) (time.Time, error) {
	var v store.Relation
	if err := json.Unmarshal(data, &v); err != nil {
		return time.Time{}, err
	}
	return v.CreatedAt, nil
}

func jobCreatedAt(data []byte) (time.Time, error) {
	var v store.Job
	if err := json.Unmarshal(data, &v); err != nil {
		return time.Time{}, err
	}
	return v.CreatedAt, nil
}
