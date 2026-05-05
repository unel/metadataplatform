package fs

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/unel/metadataplatform/store"
)

// Jobs возвращает JobStore поверх этого стора.
func (s *Store) Jobs() store.JobStore { return &jobStore{s} }

type jobStore struct{ s *Store }

func (js *jobStore) Upsert(ctx context.Context, j store.Job) error {
	if err := validateID(j.ID); err != nil {
		return err
	}
	dir := js.s.dir(dirJobs)
	path, err := recordPath(dir, j.ID)
	if err != nil {
		return err
	}
	return upsertRecord(dir, path, &j,
		func(createdAt, updatedAt time.Time) {
			j.CreatedAt = createdAt
			j.UpdatedAt = updatedAt
		},
		jobCreatedAt,
	)
}

func (js *jobStore) Get(ctx context.Context, id string) (store.Job, error) {
	if err := validateID(id); err != nil {
		return store.Job{}, err
	}
	dir := js.s.dir(dirJobs)
	path, err := recordPath(dir, id)
	if err != nil {
		return store.Job{}, err
	}
	var j store.Job
	if err := readRecord(path, &j); err != nil {
		return store.Job{}, err
	}
	return j, nil
}

func (js *jobStore) Delete(ctx context.Context, id string) error {
	if err := validateID(id); err != nil {
		return err
	}
	dir := js.s.dir(dirJobs)
	path, err := recordPath(dir, id)
	if err != nil {
		return err
	}
	return removeRecord(path)
}

func (js *jobStore) List(ctx context.Context) ([]store.Job, error) {
	var results []store.Job
	err := listDir(js.s.dir(dirJobs), func(name string, data []byte) error {
		var j store.Job
		if err := json.Unmarshal(data, &j); err != nil {
			return fmt.Errorf("%w: %s", store.ErrReadRecord, strings.TrimSuffix(name, ".json"))
		}
		results = append(results, j)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if results == nil {
		results = []store.Job{}
	}
	return results, nil
}
