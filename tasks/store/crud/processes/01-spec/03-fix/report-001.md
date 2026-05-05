---
process: 01-spec/03-fix
run: 1
date: 2026-04-28T15:20:00Z
created: 2026-04-28T15:20:00Z
see-also: tasks/store/crud/processes/01-spec/01-write/report-001.md
status: done
agent: Танк
checklist: все пункты закрыты
---

## Spec (обновлённая)

### ft.md — Функциональные требования: store/crud

---
feature: store/crud
version: 1.1.0
created: 2026-04-28T15:02:00Z
updated: 2026-04-28T15:20:00Z
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
2. Сразу после создания: `defer func() { if err != nil { os.Remove(tmpPath) } }()` — cleanup при любой ошибке или panic. Использование if-err-branch после каждого шага **запрещено**.
3. Encode объект в JSON, записать в temp-файл.
4. `f.Sync()` — fsync перед rename.
5. `f.Close()`.
6. `os.Rename(tmpPath, targetPath)`.

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
version: 1.1.0
created: 2026-04-28T15:02:00Z
updated: 2026-04-28T15:20:00Z
---

#### Надёжность

**НФТ-Н-1. Atomic write при upsert**

Temp + rename. `os.WriteFile` запрещён. При SIGKILL во время upsert — целевой файл либо содержит предыдущую версию, либо не существует.

**НФТ-Н-2. fsync перед rename**

`f.Sync()` вызывается до `os.Rename`.

**НФТ-Н-3. Cleanup temp-файлов при ошибке**

Cleanup реализуется через `defer func() { if err != nil { os.Remove(tmpPath) } }()`, объявляемый сразу после создания temp-файла. Использование if-err-branch после каждого шага **запрещено**: не защищает от утечки при panic.

**НФТ-Н-4. List не возвращает частичный результат**

При ошибке decode любого файла — операция прекращается немедленно, ни один элемент не возвращается. При будущей параллельной реализации — все горутины отменяются через переданный `ctx`.

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

**НФТ-О-1. Логирование ошибок и успешных операций**

Ошибки: `ERROR op=<op> type=<type> id=<id> error="<text>"`

Успешные операции: `DEBUG op=<op> type=<type> id=<id>` — уровень DEBUG намеренно. INFO на каждую операцию не используется: создаёт избыточный шум. Логирование успехов на INFO — осознанно отключено.

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

## Changelog

- **CR-1**: В "Решениях" добавлено обоснование неидемпотентности Delete — spawner при retry получает `NOT_FOUND` и обязан трактовать как ошибку.
- **CR-2**: §3.4 п.2 — порядок не гарантирован, тесты используют `assert.ElementsMatch`.
- **CR-3a**: §2.2 — комментарий в struct Relation: `search_tsv` исключён намеренно, PG-триггер.
- **CR-3b**: §7 — явное ограничение: FS не проверяет `from_id`/`to_id`, FK только в PG.
- **CR-4**: §5.1 — `io.EOF`/`io.ErrUnexpectedEOF` — штатное завершение, выход без логирования.
- **CR-5**: §5.2 шаг 3 — проверка `data` на null/отсутствие; §5.4 — строка в таблицу ошибок.
- **CR-6**: §НФТ-О-1 — DEBUG-логирование успехов с обоснованием.
- **CR-7**: §3.4 п.3 — ctx пробрасывается в операции List уже сейчас.
- **CR-8**: §4.3 шаг 2 — defer с захватом ошибки, запрет if-err-branch; §НФТ-Н-3 — переписан.

## Решения принятые в процессе

- **List / ошибка файла:** fail fast, частичный результат не возвращается.
- **created_at при upsert:** только при первом создании. `updated_at` — всегда текущее время.
- **ErrNotFound:** sentinel var, не typed error.
- **Two-pass decode:** envelope с `json.RawMessage`, second decode по конкретному типу.
- **Temp-файл:** `filepath.Dir(target)`, не `os.TempDir()`.
- **Delete не идемпотентен — намеренно.** Spawner при retry получит `NOT_FOUND` и обязан трактовать как ошибку — желаемый эффект не является основанием считать вызов успешным.
- **List / порядок:** не гарантирован. FS даёт лексикографический порядок UUID v7 — деталь реализации, не контракт.
- **search_tsv:** исключён из Go-struct Relation намеренно. PG управляет триггером. Поле не попадает в прикладной код.
- **FS / FK при upsert relation:** FS не проверяет `from_id`/`to_id`. Debug-реализация, целостность не гарантируется. Для продакшн — PG с FK-констрейнтами.

## Открытые вопросы

Нет.
