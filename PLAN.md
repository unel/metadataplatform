# Plan

## В работе

- `store/connection` — Unix socket сервер, echo + логирование

## Backlog

### store
- `store/log-level` — настройка уровня логирования через конфиг/ENV
- `store/protocol` — разбор и исполнение JSONL-операций (query/upsert/delete), фильтрация по ops
- `store/db-connection` — подключение к PostgreSQL, пул соединений
- `store/entity-crud` — CRUD entities через store DSL
- `store/relation-crud` — CRUD relations
- `store/job-crud` — CRUD jobs

### api
- `api/entities` — HTTP CRUD для entities
- `api/relations` — HTTP CRUD для relations
- `api/jobs` — HTTP CRUD для jobs
- `api/commands` — POST /commands, отладочный прокси в store

### spawner
- `spawner/jobs-watch` — поллинг pending jobs, спавн workers
- `spawner/events-watch` — поллинг событий, спавн hooks
- `spawner/socket` — приём явных триггеров через сокет

### platform
- `platform/cli` — entity/job/spawn команды

### workers
- `workers/scanner` — обход ФС, создание file entities
- `workers/hash-sha256` — хэширование файлов
- `workers/stash-enricher` — обогащение из Stash

### frontend
- `frontend/search` — поиск и список entities
- `frontend/entity-card` — карточка entity
- `frontend/jobs-monitor` — монитор jobs
- `frontend/repl` — JSON REPL
