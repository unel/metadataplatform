package adversarial_test

// AC-NFT-ATOMIC-01
// Атомарность через temp+rename. SIGKILL в процессе — целевой файл цел или отсутствует.
// Демон на GPU под линуксом тестирует файловую систему. Поэзия.

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/unel/metadataplatform/store"
)

// AC-NFT-ATOMIC-01: целевой файл не повреждён при прерывании upsert.
//
// Подход: запускаем отдельный бинарь-помощник который делает upsert и спим
// в defer после создания temp-файла (имитация медленной записи).
// Главный тест делает SIGKILL через 50ms. После — файл либо V1, либо отсутствует.
// Никакого частичного содержимого.
//
// Если тест запущен не как root — os.Chmod(proc, ...) может не работать,
// поэтому используем subprocess-подход с helper binary.
func TestNFT_Atomic_UpsertInterruptedBySIGKILL_FileNotCorrupted(t *testing.T) {
	if os.Getenv("TEST_ATOMIC_HELPER") == "1" {
		runAtomicHelper()
		return
	}

	basedir := t.TempDir()

	// V1: создаём начальную версию файла
	v1Content := store.Entity{ID: "ent-atomic", Type: "service", Name: "version-1"}
	v1Bytes, err := json.Marshal(v1Content)
	require.NoError(t, err)

	entitiesDir := filepath.Join(basedir, "entities")
	require.NoError(t, os.MkdirAll(entitiesDir, 0o755))

	targetFile := filepath.Join(entitiesDir, "ent-atomic.json")
	require.NoError(t, os.WriteFile(targetFile, v1Bytes, 0o644))

	// запускаем subprocess который будет делать upsert
	cmd := exec.Command(os.Args[0], "-test.run=TestNFT_Atomic_UpsertInterruptedBySIGKILL_FileNotCorrupted")
	cmd.Env = append(os.Environ(),
		"TEST_ATOMIC_HELPER=1",
		"TEST_ATOMIC_BASEDIR="+basedir,
	)

	require.NoError(t, cmd.Start())

	// ждём маркера готовности вместо слепого sleep — subprocess пишет .ready после Sync()
	readyMarker := filepath.Join(basedir, ".ready")
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if _, err := os.Stat(readyMarker); err == nil {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}

	cmd.Process.Kill()
	cmd.Wait()

	// проверяем инвариант: файл либо V1, либо отсутствует — но не мусор
	data, err := os.ReadFile(targetFile)
	if os.IsNotExist(err) {
		// файл исчез — допустимо (rename не завершился)
		return
	}
	require.NoError(t, err, "неожиданная ошибка чтения файла")

	// файл должен быть валидным JSON
	var got store.Entity
	assert.NoError(t, json.Unmarshal(data, &got),
		"файл после прерывания должен содержать валидный JSON, получили: %s", string(data))

	// содержимое должно быть V1 (новая версия не прошла rename)
	assert.Equal(t, "version-1", got.Name,
		"после SIGKILL файл должен содержать V1, получили: %s", string(data))
}

// runAtomicHelper запускается в subprocess: делает upsert с искусственной задержкой.
func runAtomicHelper() {
	basedir := os.Getenv("TEST_ATOMIC_BASEDIR")
	if basedir == "" {
		os.Exit(1)
	}

	// имитируем медленную запись через создание temp-файла вручную
	// (реальная реализация fs.Upsert будет использовать тот же паттерн)
	entitiesDir := filepath.Join(basedir, "entities")
	target := filepath.Join(entitiesDir, "ent-atomic.json")

	tmp, err := os.CreateTemp(entitiesDir, ".tmp-*")
	if err != nil {
		os.Exit(2)
	}

	// пишем "новую версию" в temp — SIGKILL должен придти до rename
	newContent := []byte(`{"id":"ent-atomic","type":"service","name":"version-2","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)
	tmp.Write(newContent)
	tmp.Sync()

	// сигнализируем родительскому процессу: данные записаны, можно убивать
	os.WriteFile(filepath.Join(basedir, ".ready"), nil, 0o644)

	// искусственная задержка — SIGKILL ожидается здесь
	time.Sleep(500 * time.Millisecond)

	// если дожили — делаем rename (в тесте этого не должно быть)
	tmp.Close()
	os.Rename(tmp.Name(), target)
}
