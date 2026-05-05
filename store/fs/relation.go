package fs

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/unel/metadataplatform/store"
)

// Relations возвращает RelationStore поверх этого стора.
func (s *Store) Relations() store.RelationStore { return &relationStore{s} }

type relationStore struct{ s *Store }

func (rs *relationStore) Upsert(ctx context.Context, r store.Relation) error {
	if err := validateID(r.ID); err != nil {
		return err
	}
	dir := rs.s.dir(dirRelations)
	path, err := recordPath(dir, r.ID)
	if err != nil {
		return err
	}
	return upsertRecord(dir, path, &r,
		func(createdAt, updatedAt time.Time) {
			r.CreatedAt = createdAt
			r.UpdatedAt = updatedAt
		},
		relationCreatedAt,
	)
}

func (rs *relationStore) Get(ctx context.Context, id string) (store.Relation, error) {
	if err := validateID(id); err != nil {
		return store.Relation{}, err
	}
	dir := rs.s.dir(dirRelations)
	path, err := recordPath(dir, id)
	if err != nil {
		return store.Relation{}, err
	}
	var r store.Relation
	if err := readRecord(path, &r); err != nil {
		return store.Relation{}, err
	}
	return r, nil
}

func (rs *relationStore) Delete(ctx context.Context, id string) error {
	if err := validateID(id); err != nil {
		return err
	}
	dir := rs.s.dir(dirRelations)
	path, err := recordPath(dir, id)
	if err != nil {
		return err
	}
	return removeRecord(path)
}

func (rs *relationStore) List(ctx context.Context) ([]store.Relation, error) {
	var results []store.Relation
	err := listDir(rs.s.dir(dirRelations), func(name string, data []byte) error {
		var r store.Relation
		if err := json.Unmarshal(data, &r); err != nil {
			return fmt.Errorf("%w: %s", store.ErrReadRecord, strings.TrimSuffix(name, ".json"))
		}
		results = append(results, r)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if results == nil {
		results = []store.Relation{}
	}
	return results, nil
}
