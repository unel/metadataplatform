package fs

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"time"

	"github.com/unel/metadataplatform/store"
)

const (
	dirEntities  = "entities"
	dirRelations = "relations"
	dirJobs      = "jobs"
)

// Config задаёт параметры FS-стора.
type Config struct {
	Basedir string
}

// Store — FS-реализация всех трёх store-интерфейсов.
// Данные хранятся как JSON-файлы в поддиректориях basedir.
// Debug-only: не предназначена для продакшн-использования.
// TODO: add RWMutex for concurrent access
type Store struct {
	basedir string
	log     *slog.Logger
}

// New инициализирует Store: создаёт поддиректории и возвращает готовый стор.
func New(cfg Config, log *slog.Logger) (*Store, error) {
	if log == nil {
		log = slog.Default()
	}
	for _, sub := range []string{dirEntities, dirRelations, dirJobs} {
		if err := os.MkdirAll(filepath.Join(cfg.Basedir, sub), 0o755); err != nil {
			log.Error("store/fs init failed",
				slog.String("basedir", cfg.Basedir),
				slog.String("error", err.Error()),
			)
			return nil, err
		}
	}
	log.Info("store/fs initialized", slog.String("basedir", cfg.Basedir))
	return &Store{basedir: cfg.Basedir, log: log}, nil
}

// dir возвращает абсолютный путь к поддиректории.
func (s *Store) dir(sub string) string {
	return filepath.Join(s.basedir, sub)
}

// Entities возвращает EntityStore. *Store напрямую удовлетворяет интерфейсу — обёртка не нужна.
func (s *Store) Entities() store.EntityStore { return s }

// Upsert сохраняет entity. Timestamps выставляются стором; входящие игнорируются.
func (s *Store) Upsert(ctx context.Context, e store.Entity) error {
	if err := validateID(e.ID); err != nil {
		return err
	}
	dir := s.dir(dirEntities)
	path, err := recordPath(dir, e.ID)
	if err != nil {
		return err
	}
	return upsertRecord(dir, path, &e,
		func(createdAt, updatedAt time.Time) {
			e.CreatedAt = createdAt
			e.UpdatedAt = updatedAt
		},
		entityCreatedAt,
	)
}

// Get возвращает entity по id.
func (s *Store) Get(ctx context.Context, id string) (store.Entity, error) {
	if err := validateID(id); err != nil {
		return store.Entity{}, err
	}
	dir := s.dir(dirEntities)
	path, err := recordPath(dir, id)
	if err != nil {
		return store.Entity{}, err
	}
	var e store.Entity
	if err := readRecord(path, &e); err != nil {
		return store.Entity{}, err
	}
	return e, nil
}

// Delete удаляет entity. Не идемпотентный.
func (s *Store) Delete(ctx context.Context, id string) error {
	if err := validateID(id); err != nil {
		return err
	}
	dir := s.dir(dirEntities)
	path, err := recordPath(dir, id)
	if err != nil {
		return err
	}
	return removeRecord(path)
}

// List возвращает все entity. Порядок не гарантирован.
func (s *Store) List(ctx context.Context) ([]store.Entity, error) {
	var results []store.Entity
	err := listDir(s.dir(dirEntities), func(name string, data []byte) error {
		var e store.Entity
		if err := json.Unmarshal(data, &e); err != nil {
			return fmt.Errorf("%w: %s", store.ErrReadRecord, strings.TrimSuffix(name, ".json"))
		}
		results = append(results, e)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if results == nil {
		results = []store.Entity{}
	}
	return results, nil
}
