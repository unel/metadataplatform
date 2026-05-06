---
process: 01-spec/01-write
run: 1
date: 2026-04-28T15:02:00Z
created: 2026-04-28T15:02:00Z
status: done
agent: Танк
checklist: все пункты закрыты
---

## Spec

### ft.md — Функциональные требования: store/crud

---
feature: store/crud
version: 1.0.0
created: 2026-04-28T15:02:00Z
updated: 2026-04-28T15:02:00Z
---

#### Контекст

`store/crud` — CRUD-слой поверх Unix socket инфраструктуры (`store/connection`). Реализует операции `upsert`, `get`, `delete`, `list` для трёх типов: `entity`, `relation`, `job`.

Компоненты:
- **Go-интерфейсы** — `EntityStore`, `RelationStore`, `JobStore`
- **FS-реализация** — JSON-файлы на диск, для отладки
- **JSONL-роутер** — маппит входящие JSONL-команды на вызовы интерфейсов

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
2. При ошибке чтения или decode **любого** файла — немедленно прекратить обработку остальных и вернуть ошибку. Частичный результат не возвращается.
3. При будущей параллельной реализации — при ошибке в одной горутине отменить все остальные.

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

`basedir` задаётся в конфиге. Субдиректории создаются при инициализации если не существуют.

##### 4.2 Конфиг

```yaml
store:
  fs:
    basedir: /var/lib/platform/store
```

##### 4.3 Atomic write (upsert)

Алгоритм:
1. `os.CreateTemp(filepath.Dir(targetPath), ".tmp-*")` — создать temp-файл в той же директории что и целевой.
2. Encode объект в JSON, записать в temp-файл.
3. `f.Sync()` — fsync перед rename.
4. `f.Close()`.
5. `os.Rename(tmpPath, targetPath)`.
6. При любой ошибке на шагах 2–5 — `os.Remove(tmpPath)`.

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

`os.ReadDir(dir)` → фильтр файлов с расширением `.json` → последовательный decode каждого.
При ошибке чтения или decode любого файла — немедленно вернуть ошибку, прекратить обработку остальных. Частичный срез не возвращается.

##### 4.7 Параллельный доступ — явное ограничение

FS-реализация **не защищает конкурентный доступ** в этой итерации.

Требования к коду:
- Struct реализации содержит комментарий `// TODO: add RWMutex for concurrent access`.
- Mutex добавляется как поле struct — не в интерфейс.

---

#### 5. JSONL-роутер

##### 5.1 Формат входящих запросов

```jsonl
{"op":"upsert","type":"entity","data":{...}}
{"op":"get","type":"relation","id":"<uuid>"}
{"op":"delete","type":"job","id":"<uuid>"}
{"op":"list","type":"entity"}
```

Парсинг: `json.NewDecoder` на `net.Conn`, цикл `Decode()`. `bufio.Scanner` запрещён (ограничение 64KB).

##### 5.2 Two-pass decode

1. Первый decode в envelope: `{"op":"...","type":"...","id":"...","data":json.RawMessage}`.
2. Switch по `op` + `type`.
3. Второй decode `data` → конкретная структура (только для upsert).

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
| `INVALID_REQUEST` | невалидный JSON |
| `WRITE_ERROR` | ошибка записи |
| `READ_ERROR` | ошибка чтения (get) |
| `DELETE_ERROR` | ошибка удаления |
| `LIST_ERROR` | ошибка при list |

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

---

### nft.md — Нефункциональные требования: store/crud

---
feature: store/crud
version: 1.0.0
created: 2026-04-28T15:02:00Z
updated: 2026-04-28T15:02:00Z
---

#### Надёжность

**НФТ-Н-1. Atomic write при upsert**

Temp + rename. `os.WriteFile` запрещён. При SIGKILL во время upsert — целевой файл либо содержит предыдущую версию, либо не существует.

**НФТ-Н-2. fsync перед rename**

`f.Sync()` вызывается до `os.Rename`.

**НФТ-Н-3. Cleanup temp-файлов при ошибке**

`os.Remove(tmpPath)` при любой ошибке в процессе atomic write.

**НФТ-Н-4. List не возвращает частичный результат**

При ошибке decode любого файла — операция прекращается немедленно, ни один элемент не возвращается. При будущей параллельной реализации — все горутины отменяются.

---

#### Параллелизм

**НФТ-П-1. Конкурентный доступ не защищается в этой итерации**

Struct содержит `// TODO: add RWMutex for concurrent access`. Mutex добавляется как поле struct без изменения интерфейса.

---

#### Безопасность

**НФТ-Б-1. Fail fast на невалидных входных данных**

Немедленная ошибка при неизвестном `op`, `type`, невалидном JSON, пустом `id`. Соединение не разрывается из-за одной ошибки — следующий запрос обрабатывается нормально.

**НФТ-Б-2. Стор не генерирует id**

При отсутствии `id` — ошибка `MISSING_ID`.

---

#### Наблюдаемость

**НФТ-О-1. Логирование ошибок**

`ERROR op=<op> type=<type> id=<id> error="<text>"`

**НФТ-О-2. Логирование инициализации**

`INFO store/fs initialized basedir=<path>`

---

#### Сопровождаемость

**НФТ-С-1.** Три независимых интерфейса.

**НФТ-С-2.** Locking-ready: mutex как поле struct, не в интерфейс.

**НФТ-С-3.** FS-реализация документирована как debug-only.

---

#### Ограничения реализации

**НФТ-О-Р-1. Зависимости**

- `github.com/google/uuid` v1.6+ — UUID v7.
- Стандартная библиотека для всего остального.
- Внешняя библиотека atomic write (`google/renameio` v2 или `natefinch/atomic`) — допустима, не обязательна.

**НФТ-О-Р-2.** `json.NewDecoder` на `net.Conn`, не `bufio.Scanner`.

**НФТ-О-Р-3.** FS — debug-only, не продакшн.

---

## Решения принятые в процессе

- **List / ошибка файла:** fail fast — при ошибке decode любого файла операция прекращается немедленно, частичный результат не возвращается; будущий параллельный вариант отменяет все горутины.
- **created_at при upsert:** только при первом создании (check-then-set). При обновлении — из существующей записи. `updated_at` — всегда текущее время.
- **ErrNotFound:** sentinel var, не typed error.
- **Two-pass decode:** envelope с `json.RawMessage`, second decode по конкретному типу.
- **Temp-файл:** `filepath.Dir(target)`, не `os.TempDir()`.

## Открытые вопросы

Нет.
