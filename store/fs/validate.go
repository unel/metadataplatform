package fs

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/unel/metadataplatform/store"
)

func validateID(id string) error {
	if id == "" {
		return store.ErrMissingID
	}
	return nil
}

// recordPath строит путь к файлу записи и проверяет что он не выходит за basedir.
// CRITICAL-1: защита от path traversal через id.
func recordPath(basedir, id string) (string, error) {
	p := filepath.Join(basedir, id+".json")
	if !strings.HasPrefix(filepath.Clean(p)+string(filepath.Separator), filepath.Clean(basedir)+string(filepath.Separator)) {
		return "", fmt.Errorf("%w: invalid id", store.ErrMissingID)
	}
	return p, nil
}
