package fs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/unel/metadataplatform/store"
)

// atomicWrite записывает объект атомарно: temp-файл в той же директории + rename.
// WARNING-4: tmp закрывается явно перед os.Rename; defer только для cleanup при ошибке.
func atomicWrite(dir, targetPath string, obj any) error {
	tmp, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return fmt.Errorf("%w: %w", store.ErrWriteRecord, err)
	}

	// cleanup при ошибке — только если tmp ещё не переименован
	tmpName := tmp.Name()
	var writeErr error
	defer func() {
		if writeErr != nil {
			os.Remove(tmpName)
		}
	}()

	enc := json.NewEncoder(tmp)
	if writeErr = enc.Encode(obj); writeErr != nil {
		tmp.Close()
		return fmt.Errorf("%w: %w", store.ErrWriteRecord, writeErr)
	}
	if writeErr = tmp.Sync(); writeErr != nil {
		tmp.Close()
		return fmt.Errorf("%w: %w", store.ErrWriteRecord, writeErr)
	}
	// Закрываем явно перед rename — дескриптор не должен жить после переименования.
	if writeErr = tmp.Close(); writeErr != nil {
		return fmt.Errorf("%w: %w", store.ErrWriteRecord, writeErr)
	}
	if writeErr = os.Rename(tmpName, targetPath); writeErr != nil {
		return fmt.Errorf("%w: %w", store.ErrWriteRecord, writeErr)
	}
	return nil
}

func readRecord(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return store.ErrNotFound
		}
		return fmt.Errorf("%w: %w", store.ErrReadRecord, err)
	}
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("%w: %w", store.ErrReadRecord, err)
	}
	return nil
}

func removeRecord(path string) error {
	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return store.ErrNotFound
		}
		return fmt.Errorf("%w: %w", store.ErrDeleteRecord, err)
	}
	return nil
}

func listDir(dir string, fn func(name string, data []byte) error) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("%w: %w", store.ErrListRecords, err)
	}
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasSuffix(name, ".json") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return fmt.Errorf("%w: %w", store.ErrListRecords, err)
		}
		if err := fn(name, data); err != nil {
			return err
		}
	}
	return nil
}
