# store/router

JSONL-роутер: принимает соединение `net.Conn`, читает команды построчно, вызывает store-интерфейсы, возвращает ответы.

Роутер не зависит от конкретной реализации стора — работает через `store.EntityStore`, `store.RelationStore`, `store.JobStore`.

## Инициализация

```go
import (
    "github.com/unel/metadataplatform/store/router"
)

r := router.New(entityStore, relationStore, jobStore, logger)

// Для каждого входящего соединения:
go r.Handle(ctx, conn)
```

## Протокол

JSONL per-line: один JSON-объект на строку. Соединение поддерживает несколько запросов до EOF.

### Формат запроса

```json
{"op": "<op>", "type": "<type>", "id": "<id>", "data": {...}}
```

| Поле | Обязательность | Описание |
|---|---|---|
| `op` | всегда | `upsert`, `get`, `delete`, `list` |
| `type` | всегда | `entity`, `relation`, `job` |
| `id` | для `get`, `delete` | ID записи |
| `data` | для `upsert` | объект для записи |

### Формат ответа

Все ответы содержат `"ok": true/false`.

**Успех upsert:**
```json
{"ok": true, "id": "01906acd-dead-7000-beef-000000000001"}
```

**Успех get / list:**
```json
{"ok": true, "data": {...}}
{"ok": true, "data": [...]}
```

**Пример полного ответа get entity:**
```json
{"ok": true, "data": {"id": "01906acd-dead-7000-beef-000000000001", "type": "file", "subtype": "video", "meta": {"path": "/data/file.mp4"}, "created_at": "2024-01-15T10:00:00Z", "updated_at": "2024-01-15T10:00:00Z"}}
```

Поля `name`, `description`, `subtype` отсутствуют в ответе когда пустые (omitempty).

**Успех delete:**
```json
{"ok": true}
```

**Ошибка:**
```json
{"ok": false, "errorCode": "NOT_FOUND", "error": "not found"}
```

## Операции

### upsert

```json
{"op": "upsert", "type": "entity", "data": {"id": "...", "type": "file", "subtype": "video", "meta": {"path": "/data/file.mp4"}}}
```

`data` обязателен. Если `data` отсутствует — `INVALID_REQUEST`.

Timestamps (`created_at`, `updated_at`) в `data` игнорируются — выставляются стором.

### get

```json
{"op": "get", "type": "entity", "id": "01906acd-dead-7000-beef-000000000001"}
```

Возвращает запись в `data`. Если не найдена — `NOT_FOUND`.

### delete

```json
{"op": "delete", "type": "entity", "id": "01906acd-dead-7000-beef-000000000001"}
```

Возвращает `{"ok": true}` при успехе. Если не найдена — `NOT_FOUND`. **Не идемпотентен.**

### list

```json
{"op": "list", "type": "entity"}
```

Возвращает все записи в `data` как массив. Порядок не гарантирован.

## Коды ошибок

| errorCode | Причина |
|---|---|
| `INVALID_REQUEST` | Отсутствует обязательное поле (`data` для upsert), невалидный JSON |
| `UNKNOWN_OP` | Неизвестная операция |
| `UNKNOWN_TYPE` | Неизвестный тип (`entity`/`relation`/`job`) |
| `NOT_FOUND` | Запись не найдена |
| `MISSING_ID` | ID пустой, содержит path traversal или не передан в запросе get/delete |
| `READ_ERROR` | Ошибка чтения записи из хранилища. При upsert означает что существующая запись нечитаема — операция отменена, новые данные не записаны. |
| `WRITE_ERROR` | Ошибка записи в хранилище |
| `DELETE_ERROR` | Ошибка удаления из хранилища |
| `LIST_ERROR` | Ошибка листинга хранилища |
| `INTERNAL_ERROR` | Прочие внутренние ошибки |

## Поведение при parse error

Если входящий JSON невалиден — соединение закрывается немедленно. Поток считается испорченным.

## Совместимость

Протокол совпадает с форматом из `PROJECT.md`. Роутер — единственная точка входа в стор через сокет; прямые вызовы интерфейсов возможны в тестах и внутренних компонентах.
