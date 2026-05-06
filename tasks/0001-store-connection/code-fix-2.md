# Code Fix 2 — store/connection

## Замечания из code-review-1.md

### CR-1 (критичная): Race на WaitGroup при shutdown

**Файл:** `cmd/store/store.go`

Исправлено архитектурным изменением: `acceptLoop` теперь управляет собственным локальным `sync.WaitGroup` для conn-горутин. Внешний `wg` (в `RunStore`) считает только acceptLoop-горутины. `acceptLoop` перед возвратом вызывает `wg.Wait()` через `defer wg.Wait()` — ждёт все свои conn-горутины. Таким образом когда внешний `wg.Wait()` возвращает управление, все conn-горутины уже завершены, и race исключён структурно.

---

### CR-2 (важная): `RunStoreWithExitCode` принимает `chan struct{}` вместо `<-chan struct{}`

**Файл:** `cmd/store/store.go`, строка 75

Сигнатура изменена с `stopCh chan struct{}` на `stopCh <-chan struct{}`.

---

### CR-3 (важная): `applyDefaults` вызывается дважды

**Файл:** `cmd/store/store.go`

Единственная точка вызова — `RunStore`. Из `RunStoreWithExitCode` вызов убран. `loadConfig` также вызывает `applyDefaults` внутри себя — это остаётся как есть (loadConfig используется независимо от RunStore).

---

### CR-4 (важная): `acceptLoop` принимает `*sync.WaitGroup` и делает скрытый side effect

**Файл:** `cmd/store/store.go`

`acceptLoop` больше не принимает `*sync.WaitGroup`. Локальный `wg` объявлен внутри `acceptLoop`, `defer wg.Wait()` гарантирует ожидание conn-горутин перед возвратом. Внешний `wg` не мутируется из `acceptLoop`.

---

### CR-5 (важная): Неограниченное число параллельных соединений

**Файлы:** `cmd/store/config.go`, `cmd/store/store.go`

Добавлено поле `MaxConnections int` в `StoreConfig` (yaml: `max_connections`). Дефолт `1024` в `applyDefaults`. В `acceptLoop` передаётся `sem chan struct{}` (buffered channel размером `MaxConnections`). Перед запуском conn-горутины делается неблокирующий `select` на `sem <- struct{}{}`: если слот есть — занимаем, если нет — соединение закрывается с WARN-логом. После завершения conn-горутины слот освобождается через `defer func() { <-sem }()`.

---

### CR-6 (минорная): `waitForSocket` объявлена но не используется

**Файл:** `cmd/store/testhelpers_test.go`

Функция `waitForSocket` удалена.

---

### CR-7 (минорная): `cancel()` вызывается дважды в select

**Файл:** `cmd/store/store.go`

Вызов `cancel()` вынесен за блок `select` — вызывается один раз после его завершения независимо от ветки.

---

### CR-8 (минорная): `StoreProcess.logs []string` не используется

**Файл:** `cmd/store/testhelpers_test.go`

Поле `logs []string` и `mu sync.Mutex` удалены из `StoreProcess`. Инициализация `logs: nil` убрана из `newStoreProcess`.

---

## Затронутые файлы

- `cmd/store/store.go` — CR-1, CR-2, CR-3, CR-4, CR-5, CR-7
- `cmd/store/config.go` — CR-5
- `cmd/store/testhelpers_test.go` — CR-6, CR-8
