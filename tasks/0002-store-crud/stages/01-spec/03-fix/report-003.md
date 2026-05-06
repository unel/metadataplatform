---
process: 01-spec/03-fix
run: 3
date: 2026-04-28T20:27:48Z
created: 2026-04-28T20:27:48Z
see-also: tasks/0002-store-crud/stages/01-spec/03-fix/report-002.md
status: done
agent: Танк
checklist: все пункты закрыты
---

## Spec (обновлённая)

### ft.md — Функциональные требования: store/crud

---
feature: store/crud
version: 1.3.0
created: 2026-04-28T15:02:00Z
updated: 2026-04-28T20:27:48Z
---

#### Контекст

`store/crud` — CRUD-слой поверх Unix socket инфраструктуры (`store/connection`). Реализует операции `upsert`, `get`, `delete`, `list` для трёх типов: `entity`, `relation`, `job`.

Компоненты:
- **Go-интерфейсы** — `EntityStore`, `RelationStore`, `JobStore`
- **FS-реализация** — JSON-файлы на диск, для отладки
- **JSONL-роутер** — маппит входящие JSONL-команды на вызовы интерфейсов

Структура пакетов:

```
store/
  store.go          — интерфейсы + ErrNotFound + модели
  router/
    router.go       — JSONL-роутер
  fs/
    fs.go           — FS-реализация
```

---

#### 1. Go-интерфейсы

##### 1.1 Сигнатуры

```go
type EntityStore interface {
    Upsert(ctx context.Context, e Entity) error
    Get(ctx context.Context, id string) (Entity, error)
    Delete(ctx context.Context, id string) error
    List(ctx context.Context) ([]Entity, error)
}

type RelationStore interface {
    Upsert(ctx context.Context, r Relation) error
    Get(ctx context.Context, id string) (Relation, error)
    Delete(ctx context.Context, id string) error
    List(ctx context.Context) ([]Relation, error)
}

type JobStore interface {
    Upsert(ctx context.Context, j Job) error
    Get(ctx context.Context, id string) (Job, error)
    Delete(ctx context.Context, id string) error
    List(ctx context.Context) ([]Job, error)
}
```

Интерфейсы определяются на стороне потребителя (пакет `store`), не в пакете реализации.

##### 1.2 Sentinel error

```go
var ErrNotFound = errors.New("not found")
```

Объявляется на уровне пакета `store`. Caller проверяет через `errors.Is(err, store.ErrNotFound)`.

---

#### 2. Модели данных

##### 2.1 Entity

```go
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
```

##### 2.2 Relation

```go
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
    // search_tsv намеренно отсутствует в Go-struct.
    // В PostgreSQL это tsvector-колонка, управляемая триггером.
    // FS-реализация её не ведёт. Поле не попадает в прикладной код ни одной реализации.
}
```

##### 2.3 Job

```go
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
```

---

#### 3. Операции

##### 3.1 Upsert

**Входные данные:** объект типа (Entity / Relation / Job) с заполненным `id`.

**Поведение:**
1. Если `id` пустой — вернуть ошибку клиента. Стор не генерирует id.
2. Если запись с `id` не существует — создать; `created_at` выставить в `time.Now().UTC()`.
3. Если запись с `id` существует — обновить; `created_at` сохраняется из существующей записи.
4. `updated_at` — всегда `time.Now().UTC()`, даже если данные не изменились.
5. Значения `created_at` и `updated_at` из входящего запроса игнорируются.

**Результат успеха:** `{"ok":true,"id":"<uuid>"}`

**Failure modes:**

| Условие | errorCode | error |
|---|---|---|
| `id` отсутствует или пустая строка | `MISSING_ID` | `"id is required"` |
| Ошибка записи на диск (FS) | `WRITE_ERROR` | `"failed to write record"` |

---

##### 3.2 Get

**Входные данные:** поле `id` строкой.

**Поведение:**
1. Если `id` пустой — вернуть ошибку клиента.
2. Найти запись по `id`.
3. Если не найдена — вернуть `ErrNotFound`.

**Результат успеха:** `{"ok":true,"data":{<объект>}}`

**Failure modes:**

| Условие | errorCode | error |
|---|---|---|
| `id` отсутствует или пустой | `MISSING_ID` | `"id is required"` |
| Запись не найдена | `NOT_FOUND` | `"not found"` |
| Ошибка чтения с диска (FS) | `READ_ERROR` | `"failed to read record"` |

---

##### 3.3 Delete

**Входные данные:** поле `id` строкой.

**Поведение:**
1. Если `id` пустой — вернуть ошибку клиента.
2. Если запись не существует — вернуть ошибку. Delete **не идемпотентный**.
3. Удалить запись.

**Результат успеха:** `{"ok":true}`

**Failure modes:**

| Условие | errorCode | error |
|---|---|---|
| `id` отсутствует или пустой | `MISSING_ID` | `"id is required"` |
| Запись не найдена | `NOT_FOUND` | `"not found"` |
| Ошибка удаления с диска (FS) | `DELETE_ERROR` | `"failed to delete record"` |

---

##### 3.4 List

**Входные данные:** тип (`entity` / `relation` / `job`). Фильтры не принимаются — это scope `store/query`.

**Поведение:**
1. Вернуть все записи данного типа.
2. Порядок элементов в результате **не гарантирован**. Тесты используют `assert.ElementsMatch`, не `assert.Equal`.
3. `ctx` пробрасывается в каждую операцию чтения внутри List — в том числе в текущей последовательной реализации. Это обязательное требование: без ctx-пробрасывания переход к параллельной реализации потребует изменения сигнатуры.
4. При ошибке чтения или decode **любого** файла — немедленно прекратить обработку остальных и вернуть ошибку. Частичный результат не возвращается.
5. При будущей параллельной реализации — при ошибке в одной горутине отменить все остальные через переданный `ctx`.

**Результат успеха:** `{"ok":true,"data":[<объекты>]}`

Пустой список — валидный результат: `{"ok":true,"data":[]}`

**Failure modes:**

| Условие | errorCode | error |
|---|---|---|
| Ошибка чтения директории (FS) | `LIST_ERROR` | `"failed to list records"` |
| Ошибка чтения или decode одного файла | `LIST_ERROR` | `"failed to read record <id>"` |

---

#### 4. FS-реализация

##### 4.1 Структура директорий

```
{basedir}/
  entities/{uuid}.json
  relations/{uuid}.json
  jobs/{uuid}.json
```

`basedir` задаётся в конфиге. Субдиректории создаются при инициализации если не существуют (`os.MkdirAll`).

Путь к файлу записи формируется через `filepath.Join(dir, id+".json")`. Прямая конкатенация строк для формирования пути запрещена.

**Failure modes инициализации:**

| Условие | Поведение |
|---|---|
| `os.MkdirAll` вернул ошибку | Конструктор возвращает ошибку. Инициализация прерывается. |

##### 4.2 Конфиг

```yaml
store:
  fs:
    basedir: /var/lib/platform/store
```

##### 4.3 Atomic write (upsert)

Функция atomic write должна использовать именованный возврат ошибки `(err error)`. Без named return переменная `err` в defer-замыкании не связана с возвращаемым значением функции, и cleanup при `return someErr` не выполняется.

Сигнатура:

```go
func (s *Store) upsert(ctx context.Context, path string, obj any) (err error) {
    tmp, err := os.CreateTemp(filepath.Dir(path), ".tmp-*")
    if err != nil {
        return err
    }
    defer func() {
        if err != nil {
            os.Remove(tmp.Name())
        }
    }()
    // encode, sync, close, rename...
}
```

Алгоритм:
1. `os.CreateTemp(filepath.Dir(targetPath), ".tmp-*")` — создать temp-файл в той же директории что и целевой.
2. Сразу после создания: `defer func() { if err != nil { os.Remove(tmp.Name()) } }()` — cleanup при любой ошибке или panic. Использование if-err-branch после каждого шага **запрещено**. Требует named return `(err error)` в сигнатуре.
3. Выставить `created_at`/`updated_at` по правилам §3.1:
   - При создании: `created_at = time.Now().UTC()`
   - При обновлении: `created_at` читается из существующей записи и сохраняется без изменений
   - `updated_at = time.Now().UTC()` — всегда
   - Значения `created_at`/`updated_at` из входящего запроса игнорируются на этом шаге
4. Encode объект в JSON, записать в temp-файл.
5. `f.Sync()` — fsync перед rename.
6. `f.Close()`.
7. `os.Rename(tmpPath, targetPath)`.

Запреты:
- `os.WriteFile` запрещён — не атомарен.
- Temp-файл в `os.TempDir()` запрещён — может быть на другом filesystem.

##### 4.4 Чтение (get)

`os.ReadFile(path)` → JSON decode → вернуть объект.
Если файл не существует (`os.IsNotExist(err)`) → вернуть `store.ErrNotFound`.

##### 4.5 Удаление (delete)

Проверить существование файла → `os.Remove(path)`.
Если файл не существует → вернуть `store.ErrNotFound`.

##### 4.6 Список (list)

`os.ReadDir(dir)` → фильтр файлов с расширением `.json` → последовательный decode каждого с пробросом `ctx`.
При ошибке чтения или decode любого файла — немедленно вернуть ошибку, прекратить обработку остальных. Частичный срез не возвращается.

##### 4.7 Параллельный доступ — явное ограничение

FS-реализация **не защищает конкурентный доступ** в этой итерации.

Требования к коду:
- Struct реализации содержит комментарий `// TODO: add RWMutex for concurrent access`.
- Mutex добавляется как поле struct — не в интерфейс.

---

#### 5. JSONL-роутер

##### 5.0 Конструктор

```go
// package router
func New(entities store.EntityStore, relations store.RelationStore, jobs store.JobStore) *Router
```

Роутер находится в пакете `store/router`. Сторы передаются через конструктор. Роутер не создаёт сторы самостоятельно — это ответственность вызывающего кода.

##### 5.1 Формат входящих запросов

```jsonl
{"op":"upsert","type":"entity","data":{...}}
{"op":"get","type":"relation","id":"<uuid>"}
{"op":"delete","type":"job","id":"<uuid>"}
{"op":"list","type":"entity"}
```

Парсинг: `json.NewDecoder` на `net.Conn`, цикл `Decode()`. `bufio.Scanner` запрещён (ограничение 64KB).

При получении `io.EOF` или `io.ErrUnexpectedEOF` от `Decode()` — роутер выходит из цикла **без логирования ошибки**. Это штатное завершение: клиент закрыл соединение.

##### 5.2 Two-pass decode

1. Первый decode в envelope: `{"op":"...","type":"...","id":"...","data":json.RawMessage}`.
2. Switch по `op` + `type`.
3. Если `op` == `upsert`: проверить что `data` присутствует и не равен `null`. Если отсутствует или `null` — вернуть `INVALID_REQUEST` с текстом `"data is required for upsert"`. Не передавать в second decode.
4. Второй decode `data` → конкретная структура (только для upsert, после прохождения проверки шага 3).

**Разрешение конфликта id:** для `upsert` поле `id` в envelope игнорируется — `id` берётся из `data`. Для `get` и `delete` `id` берётся из envelope.

##### 5.3 Маппинг операций

| `op` | `type` | Метод |
|---|---|---|
| `upsert` | `entity` | `EntityStore.Upsert` |
| `upsert` | `relation` | `RelationStore.Upsert` |
| `upsert` | `job` | `JobStore.Upsert` |
| `get` | `entity` | `EntityStore.Get` |
| `get` | `relation` | `RelationStore.Get` |
| `get` | `job` | `JobStore.Get` |
| `delete` | `entity` | `EntityStore.Delete` |
| `delete` | `relation` | `RelationStore.Delete` |
| `delete` | `job` | `JobStore.Delete` |
| `list` | `entity` | `EntityStore.List` |
| `list` | `relation` | `RelationStore.List` |
| `list` | `job` | `JobStore.List` |

##### 5.4 Ошибки роутера

| Условие | errorCode | error |
|---|---|---|
| Неизвестный `op` | `UNKNOWN_OP` | `"unknown op: <op>"` |
| Неизвестный `type` | `UNKNOWN_TYPE` | `"unknown type: <type>"` |
| Невалидный JSON в запросе | `INVALID_REQUEST` | `"invalid json"` |
| Невалидный JSON в `data` | `INVALID_REQUEST` | `"invalid data"` |
| `data` отсутствует или `null` для `upsert` | `INVALID_REQUEST` | `"data is required for upsert"` |

##### 5.5 Формат всех ответов

Успех: `{"ok":true, ...}`

Ошибка: `{"ok":false,"error":"<human-readable>","errorCode":"<SCREAMING_SNAKE>"}`

Полный список errorCode:

| errorCode | Когда |
|---|---|
| `MISSING_ID` | `id` отсутствует или пустой |
| `NOT_FOUND` | запись не найдена |
| `UNKNOWN_OP` | неизвестная операция |
| `UNKNOWN_TYPE` | неизвестный тип |
| `INVALID_REQUEST` | невалидный JSON, невалидные данные, отсутствующий/null `data` для upsert |
| `WRITE_ERROR` | ошибка записи |
| `READ_ERROR` | ошибка чтения (get) |
| `DELETE_ERROR` | ошибка удаления |
| `LIST_ERROR` | ошибка при list |

##### 5.6 Логирование операций

Логирование — ответственность **роутера**, не FS-реализации. FS-реализация не логирует.

Роутер логирует **после** вызова метода стора, используя поля из входящего запроса:
- `op` — из envelope
- `type` — из envelope
- `id` — из envelope (`get`, `delete`) или из `data.id` (`upsert`); для `list` поле `id` опускается

Успешная операция: `DEBUG op=<op> type=<type> id=<id>` (для `list`: `DEBUG op=list type=<type> count=<N>`)

Ошибка: `ERROR op=<op> type=<type> id=<id> error="<text>"`

---

#### 6. UUID

UUID v7. Генерация клиентская — стор не генерирует id.
Библиотека: `github.com/google/uuid` v1.6+, функция `uuid.NewV7()`.
`uuid.New()` (UUID v4) запрещён.

---

#### 7. Ограничения MVP

- PostgreSQL-реализация — отдельная задача
- Query DSL, cursor pagination — `store/query`
- `FOR UPDATE SKIP LOCKED` — отдельно
- Конкурентная защита FS
- Генерация `id` на стороне стора
- **FS-реализация не проверяет существование `from_id`/`to_id` при upsert relation.** FK-целостность обеспечивается только PG-реализацией через `REFERENCES entities(id) ON DELETE CASCADE`.

---

### nft.md — Нефункциональные требования: store/crud

---
feature: store/crud
version: 1.3.0
created: 2026-04-28T15:02:00Z
updated: 2026-04-28T20:27:48Z
---

#### Надёжность

**НФТ-Н-1. Atomic write при upsert**

Temp + rename. `os.WriteFile` запрещён. При SIGKILL во время upsert — целевой файл либо содержит предыдущую версию, либо не существует.

**НФТ-Н-2. fsync перед rename**

`f.Sync()` вызывается до `os.Rename`.

**НФТ-Н-3. Cleanup temp-файлов при ошибке**

Cleanup реализуется через `defer func() { if err != nil { os.Remove(tmpPath) } }()`, объявляемый сразу после создания temp-файла. Функция atomic write обязана использовать именованный возврат ошибки `(err error)` — без named return переменная `err` в defer не связана с возвращаемым значением, и cleanup при `return someErr` не выполняется. Использование if-err-branch после каждого шага **запрещено**: не защищает от утечки при panic.

**НФТ-Н-4. List не возвращает частичный результат**

При ошибке decode любого файла — операция прекращается немедленно, ни один элемент не возвращается. При будущей параллельной реализации — все горутины отменяются через переданный `ctx`.

**НФТ-Н-5. Failure mode инициализации FS**

Если `os.MkdirAll` возвращает ошибку при старте — конструктор возвращает ошибку. Стор не используется без успешной инициализации. Вызывающий код обязан обработать ошибку конструктора.

---

#### Параллелизм

**НФТ-П-1. Конкурентный доступ не защищается в этой итерации**

Struct содержит `// TODO: add RWMutex for concurrent access`. Mutex добавляется как поле struct без изменения интерфейса.

---

#### Безопасность

**НФТ-Б-1. Fail fast на невалидных входных данных**

Немедленная ошибка при неизвестном `op`, `type`, невалидном JSON, пустом `id`, отсутствующем или null `data` для upsert. Соединение не разрывается из-за одной ошибки — следующий запрос обрабатывается нормально.

**НФТ-Б-2. Стор не генерирует id**

При отсутствии `id` — ошибка `MISSING_ID`.

---

#### Наблюдаемость

**НФТ-О-1. Логирование операций**

Логирование — ответственность **роутера**. FS-реализация не логирует ничего кроме инициализации (§НФТ-О-2).

Роутер логирует после вызова метода стора:

Успешные операции:
- Для `get`, `delete`, `upsert`: `DEBUG op=<op> type=<type> id=<id>`
- Для `list`: `DEBUG op=list type=<type> count=<N>`

Ошибки: `ERROR op=<op> type=<type> id=<id> error="<text>"`

Поле `id` для операции `list` опускается — операция не адресована по id. Уровень DEBUG намеренно: INFO на каждую операцию создаёт избыточный шум.

**НФТ-О-2. Логирование инициализации**

Логирует FS-реализация (конструктор):

Успех: `INFO store/fs initialized basedir=<path>`

Ошибка: `ERROR store/fs init failed basedir=<path> error=<...>`

---

#### Сопровождаемость

**НФТ-С-1.** Три независимых интерфейса.

**НФТ-С-2.** Locking-ready: mutex как поле struct, не в интерфейс.

**НФТ-С-3.** FS-реализация документирована как debug-only.

**НФТ-С-4.** Роутер в отдельном пакете `store/router`. Зависимость роутера от стора — через интерфейсы, не через конкретные реализации. Замена FS на PG не требует изменений в роутере.

---

#### Ограничения реализации

**НФТ-О-Р-1. Зависимости**

- `github.com/google/uuid` v1.6+ — UUID v7.
- Стандартная библиотека для всего остального.
- Внешняя библиотека atomic write (`google/renameio` v2 или `natefinch/atomic`) — допустима, не обязательна.

**НФТ-О-Р-2.** `json.NewDecoder` на `net.Conn`, не `bufio.Scanner`.

**НФТ-О-Р-3.** FS — debug-only, не продакшн.

---

## Changelog

- **CR-1—CR-13**: см. report-002.md
- **CR-14**: §4.3 — добавлено требование named return `(err error)` в сигнатуре функции atomic write + пример сигнатуры. §НФТ-Н-3 — явное объяснение почему named return обязателен.
- **CR-15**: §4.1 — явное требование: путь к файлу через `filepath.Join(dir, id+".json")`, прямая конкатенация строк запрещена.
- **CR-16**: §НФТ-О-1 — зафиксировано: логирование — ответственность роутера, FS не логирует. §5.6 — новый раздел с явными правилами логирования роутера.

## Решения принятые в процессе

- **Named return в atomic write:** обязателен для корректной работы defer-cleanup. Без named return `err` в замыкании не связана с возвращаемым значением функции.
- **filepath.Join для пути к файлу:** обязателен, прямая конкатенация строк запрещена.
- **Логирование — ответственность роутера:** FS не логирует операции. Роутер логирует после вызова метода стора: op, type, id из запроса, результат.

## Открытые вопросы

Нет.
