# store

Пакет `store` определяет типы данных и интерфейсы хранилища для Metadata Platform.

Три сущности: Entity (любой объект), Relation (направленная связь между двумя Entity), Job (задача обработки). Для каждой — интерфейс CRUD: Upsert, Get, Delete, List.

```go
// Получить Entity по ID
entity, err := entityStore.Get(ctx, "01906acd-dead-7000-beef-000000000001")
if errors.Is(err, store.ErrNotFound) {
    // entity не существует
}

// Создать или обновить Entity
err = entityStore.Upsert(ctx, store.Entity{
    ID:      "01906acd-dead-7000-beef-000000000001",
    Type:    "file",
    Subtype: "video",
    Meta:    json.RawMessage([]byte(`{"path": "/data/file.mp4", "size": 104857600}`)),
})
```

## Структура

| Пакет | Что делает |
|---|---|
| `store` | Типы, интерфейсы, sentinel errors |
| `store/fs` | FS-реализация (debug-only, данные как JSON-файлы) |
| `store/router` | JSONL-роутер: маппит входящие команды на интерфейсы стора |

## Sentinel errors

| Ошибка | Значение |
|---|---|
| `store.ErrNotFound` | Запись не найдена |
| `store.ErrMissingID` | ID пустой или невалидный |
| `store.ErrReadRecord` | Ошибка чтения записи |
| `store.ErrWriteRecord` | Ошибка записи |
| `store.ErrDeleteRecord` | Ошибка удаления |
| `store.ErrListRecords` | Ошибка листинга |
| `store.ErrReadExisting` | Ошибка чтения существующей записи при upsert |

Проверяй через `errors.Is` — ошибки могут быть обёрнуты через `%w`.

## Типы данных

### Entity

```go
type Entity struct {
    ID          string          // UUID v7
    Type        string          // обязательный: "file", "film" и др.
    Subtype     string          // опциональный
    Name        string          // опциональный
    Description string          // опциональный
    Meta        json.RawMessage // произвольный JSON; nil сериализуется как null, не {}
    CreatedAt   time.Time       // выставляется стором, не клиентом
    UpdatedAt   time.Time       // выставляется стором, не клиентом
}
```

### Relation

```go
type Relation struct {
    ID        string          // UUID v7
    FromID    string          // ID исходной Entity
    ToID      string          // ID целевой Entity
    Type      string          // обязательный: "derivedFrom", "fact" и др.
    Subtype   string          // опциональный
    Value     json.RawMessage // payload связи; nil сериализуется как null, не {}
    Meta      json.RawMessage // метаданные связи; nil сериализуется как null, не {}
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Job

```go
type Job struct {
    ID         string          // UUID v7
    EntityID   string          // опциональный
    RelationID string          // опциональный
    Kind       string          // тип задачи: "hash" и др.
    Worker     string          // исполнитель: "hash.sha256" и др.
    Status     string          // pending | running | done | failed
    Progress   json.RawMessage // {"done": 450, "total": 1000}
    Error      string          // заполняется при Status=failed
    Payload    json.RawMessage // аргументы воркера; nil сериализуется как null, не {}
    CreatedAt  time.Time
    UpdatedAt  time.Time
}
```

## Важно

- **Timestamps управляются стором**: передавать `CreatedAt`/`UpdatedAt` в `Upsert` бессмысленно — стор перезапишет их.
- **Upsert идемпотентен**: повторный вызов с тем же ID обновляет запись, сохраняя оригинальный `CreatedAt`.
- **Delete не идемпотентен**: вызов на несуществующем ID вернёт `store.ErrNotFound`.
- **ID обязателен**: пустой ID на любой операции вернёт `store.ErrMissingID`.
