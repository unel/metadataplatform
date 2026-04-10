# Единая точка доступа к данным

**Почему:** когда несколько компонентов ходят в БД напрямую, каждый реализует свой слой доступа. Изменение схемы → правки везде. Кэширование, транзакции, пагинация — дублируются или реализуются по-разному. Store централизует это в одном месте.

## Границы применимости

Любой доступ к персистентному хранилищу. Кэш в памяти внутри одного процесса — не репозиторий.

## Если соблюдать рьяно

Иногда прямой SQL-запрос в 3 строки решает задачу за секунду, а routing через store socket добавляет overhead. Для отладочных инструментов и одноразовых скриптов прямой доступ допустим.

## Если игнорировать

Схема БД «протекает» во все компоненты. Переименование колонки — правки в 10 местах. Нет единого места для оптимизации запросов. Тестирование каждого компонента требует реальной БД.

## Когда можно отступить

`cmd/platform` — CLI для ручных операций. Если нужна нестандартная операция недоступная через store API — допустимо. Но в продакшен коде — только через store.

## Плохо — spawner ходит в БД напрямую

```
// cmd/spawner/poller.go
// spawner знает о PostgreSQL, SQL, схеме
db := postgres.Connect(config.DBURL)

rows, _ := db.Query(`
  SELECT id, worker, payload
  FROM jobs
  WHERE status = 'pending'
  ORDER BY created_at
  FOR UPDATE SKIP LOCKED
  LIMIT 10
`)

// теперь spawner зависит от схемы jobs
// изменение схемы → правим spawner
```

## Хорошо — spawner работает через store socket

```
// cmd/spawner/poller.go
// spawner знает только о протоколе store
jobs, err := storeClient.Query(StoreQuery{
  Type:   "job",
  Filter: Filter{Status: "pending"},
  Limit:  10,
})

// детали SQL скрыты в store
// изменение схемы → правим только store
```

## Структура в проекте

```
cmd/store/     — единственный владелец SQL и PostgreSQL
  internal/
    db/        — sqlc-генерированный код
    query/     — сложные запросы
  main.go      — JSONL-сервер на Unix socket

cmd/api/       — только storeClient.Query/Upsert/Delete
cmd/spawner/   — только storeClient.Query/Upsert
workers/       — только stdout JSONL (spawner проксирует в store)
```
