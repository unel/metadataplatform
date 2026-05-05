package adversarial_test

// AC-ENTITY-LIST-03, AC-ENTITY-LIST-04
// List — то место где "partial success" особенно соблазнителен. Спека говорит: нет.

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
)

// AC-ENTITY-LIST-03: директория entities недоступна для чтения → (nil, err).
func TestEntity_List_UnreadableDirectory_ReturnsError(t *testing.T) {
	basedir := t.TempDir()
	es := newEntityStoreAt(t, basedir)
	ctx := context.Background()

	// создаём пару записей чтобы было что терять
	require.NoError(t, es.Upsert(ctx, store.Entity{ID: "ent-a", Type: "service"}))
	require.NoError(t, es.Upsert(ctx, store.Entity{ID: "ent-b", Type: "service"}))

	entitiesDir := filepath.Join(basedir, "entities")
	require.NoError(t, os.Chmod(entitiesDir, 0o333)) // write/exec, но не read
	t.Cleanup(func() { os.Chmod(entitiesDir, 0o755) })

	result, err := es.List(ctx)
	assert.Error(t, err, "ожидаем ошибку при нечитаемой директории")
	assert.Nil(t, result, "при ошибке list должен вернуть nil, не частичный результат")
}

// AC-ENTITY-LIST-04: один файл невалидный JSON → (nil, err), НЕ частичный результат.
// Это проверяет fail-fast семантику — половина данных хуже чем ничего.
func TestEntity_List_PartialDecodeFailure_ReturnsNilNotPartial(t *testing.T) {
	basedir := t.TempDir()
	es := newEntityStoreAt(t, basedir)
	ctx := context.Background()

	// один валидный файл
	require.NoError(t, es.Upsert(ctx, store.Entity{ID: "a", Type: "service"}))

	// один файл с мусором — имитируем повреждение
	corruptPath := filepath.Join(basedir, "entities", "b.json")
	require.NoError(t, os.WriteFile(corruptPath, []byte(`{broken json`), 0o644))

	result, err := es.List(ctx)
	require.Error(t, err, "ожидаем ошибку при невалидном JSON файле")
	assert.Nil(t, result, "при decode failure должен вернуть nil, не частичный результат")

	// ошибка должна содержать id сломанного файла ("b")
	assert.Contains(t, err.Error(), "b",
		"ошибка должна идентифицировать проблемный файл по id")
}

// AC-ENTITY-LIST-05: List игнорирует не-.json файлы в директории.
// temp-файлы с префиксом .tmp- или другие файлы без .json не должны попадать в результат.
// Если кто-то "оптимизирует" фильтрацию — этот тест поймает.
func TestEntity_List_IgnoresNonJsonFiles(t *testing.T) {
	basedir := t.TempDir()
	es := newEntityStoreAt(t, basedir)
	ctx := context.Background()

	// один валидный entity
	require.NoError(t, es.Upsert(ctx, store.Entity{ID: "only-valid", Type: "service"}))

	// посторонние файлы в директории
	entitiesDir := filepath.Join(basedir, "entities")
	require.NoError(t, os.WriteFile(filepath.Join(entitiesDir, ".tmp-abc123"), []byte(`{"id":"intruder"}`), 0o644))
	require.NoError(t, os.WriteFile(filepath.Join(entitiesDir, "README"), []byte("not json"), 0o644))

	result, err := es.List(ctx)
	require.NoError(t, err, "посторонние файлы не должны вызывать ошибку")
	require.Len(t, result, 1, "только один валидный entity должен быть возвращён")
	assert.Equal(t, "only-valid", result[0].ID)
}

// дополнительно: список возвращает nil (не пустой срез) при ошибке
func TestEntity_List_OnError_ReturnsNilSliceNotEmpty(t *testing.T) {
	basedir := t.TempDir()
	es := newEntityStoreAt(t, basedir)
	ctx := context.Background()

	corruptPath := filepath.Join(basedir, "entities", "z.json")
	require.NoError(t, os.WriteFile(corruptPath, []byte(`not-json-at-all`), 0o644))

	result, err := es.List(ctx)
	require.Error(t, err)

	// json.RawMessage{} != nil, поэтому проверяем именно nil
	var nilSlice []store.Entity
	assert.Equal(t, nilSlice, result,
		"при ошибке List возвращает nil-срез, не []Entity{}")

	// убеждаемся что это не просто interface-nil
	if result != nil {
		j, _ := json.Marshal(result)
		t.Errorf("ожидали nil, получили: %s", string(j))
	}
}
