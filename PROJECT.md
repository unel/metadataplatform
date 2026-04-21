# Project Context — Metadata Platform

## Статус проекта

Greenfield проект на этапе проектирования. Исходного кода пока нет. Полная техническая спецификация — в `metadata-platform-tech-doc.md`.

## Стек

- **Frontend**: SvelteKit 5 с рунами, на Bun
- **Backend**: Go 1.22 — отдельные бинари, оркестрация через docker-compose
- **БД**: PostgreSQL с `jsonb`, GIN-индексами и full-text search через `tsvector`
- **HTTP фреймворк**: Huma (автогенерация OpenAPI спеки), поверх `chi`
- **Store layer**: sqlc (генерация Go кода из SQL)
- **Конфигурация**: YAML для всего, ENV переменные перебивают YAML

## Структура бинарей

```
cmd/
  store/    — хранилище, DSL поверх БД, слушает Unix socket(ы)
  spawner/  — оркестратор, поллит store, спавнит процессы по правилам
  api/      — HTTP сервер, читает/пишет через store
  platform/ — CLI, читает/пишет через store, триггерит spawner

workers/    — внешние исполняемые файлы, спавнятся spawner'ом по jobs
hooks/      — внешние исполняемые файлы, спавнятся spawner'ом по событиям
```

Схема БД накатывается init SQL скриптом через `/docker-entrypoint-initdb.d/`.

## Базовая модель данных

- **Entity** — любой объект. Поля: `id` (UUID v7), `type`, `subtype`, `name`, `description`, `meta jsonb`, `created_at`, `updated_at`.
- **Relation** — направленная связь между двумя entity, несёт JSON payload в `value jsonb`. Поля: `id`, `from_id`, `to_id`, `type`, `subtype`, `value jsonb`, `meta jsonb`, `created_at`, `updated_at`, `search_tsv tsvector`.
- **Job** — задача обработки. Поля: `id`, `entity_id`, `relation_id`, `kind`, `worker`, `status`, `progress jsonb`, `error text`, `payload jsonb`, `created_at`, `updated_at`.

Ключевой принцип: relations несут сгруппированные факты в JSON payload — не нужно разбивать связанные поля на отдельные строки.

## Entity типы (MVP)

| type | subtype | meta |
|------|---------|------|
| `file` | `video`, `image`, `audio`, `text`, `executable`, `unknown` | `path`, `size`, `mtime` (unix ts), `mimetype` |

Preview и mask — не отдельные типы entity. Это `file`-сущности, связанные с исходным файлом через `relation(type=derivedFrom, subtype=preview/mask)`, детали — в meta relation.

## Job статусы

`pending` → `running` → `done` / `failed`

`progress jsonb` хранит `{"done": 450, "total": 1000}`. Воркер пишет обновления прогресса в JSONL-поток батчами.

## Store (`cmd/store/`)

Единственный бинарь работающий с БД напрямую. Слушает Unix socket(ы), принимает команды на чтение и запись.

Протокол — JSONL per-line запрос/ответ:
```
→ {"op":"upsert","type":"entity","data":{...}}
← {"ok":true,"id":"uuid"}
→ {"op":"query","type":"job","filter":{"status":"pending"}}
← {"ok":true,"data":[...]}
→ {"op":"delete","type":"entity","id":"..."}
← {"ok":true}
```

Конфиг сокетов — какой сокет какие команды принимает (ENV перебивает YAML):
```yaml
store:
  db_url: postgres://localhost/platform  # перебивается $DB_URL
  sockets:
    - path: /run/platform/store.sock
      ops: [query, upsert, delete]
    - path: /run/platform/store-readonly.sock
      ops: [query]
    - path: /run/platform/store-write.sock
      ops: [upsert, delete]
```

## Spawner (`cmd/spawner/`)

Оркестратор — поллит секторы данных в store, спавнит процессы по правилам из YAML, читает JSONL stdout и отправляет обратно в store. Объединяет логику runner + hookrunner.

```yaml
spawn_rules:
  - watch: jobs
    filter: {status: pending}
    run: workers/${type}.${subtype}
    args: [--entity-id=${entity_id}]
    limits:
      global: 8
      by_type: 2       # опционально
      by_subtype: 1    # опционально

  - watch: events
    on: entity.created
    match:
      - type: file
        subtype: video
      - type: file
        meta:
          path:
            glob: "/media/**"
          size:
            gt: 10485760
          mtime:
            gt: 1704067200
    run: hooks/on-file-created
    args: [--entity-id=${entity_id}]
```

- Один инстанс, защита через PostgreSQL advisory lock
- Блокировка jobs: `SELECT ... FOR UPDATE SKIP LOCKED`
- Шаблоны в `args`: `${VAR}` и `${VAR:-default}` через либу `drone/envsubst`
- Строковые поля meta: операторы `glob`, `regex`
- Числовые и временные поля: `gt`, `lt`, `gte`, `lte`, `eq`
- Match — список OR-групп, внутри каждой AND
- Принимает команды на явный триггер через сокет:
```
→ {"op":"spawn","worker":"hash.sha256","payload":{...}}
← {"ok":true,"job_id":"..."}
```

## Архитектура воркеров

Воркеры и хуки — внешние исполняемые файлы в подмонтированных папках. Spawner матчит job по `<type>.<subtype>` к файлу `workers/<type>.<subtype>`.

**Вызов воркера/хука:**
- Параметры — CLI аргументы: плоские поля `job.payload` как `--key=value`
- Payload джоб держим плоским; сложные данные воркер читает сам через API
- stdout — JSONL поток команд (spawner читает построчно в реальном времени)
- stderr — логи контейнера

**Формат JSONL (stdout воркера):**
```jsonl
{"op":"upsert","type":"entity","data":{...}}
{"op":"upsert","type":"relation","data":{...}}
{"op":"upsert","type":"job","data":{"id":"...","progress":{"done":10,"total":1000}}}
{"op":"delete","type":"entity","id":"..."}
```

## Platform (`cmd/platform/`)

Системный CLI — читает/пишет через store socket, явно триггерит spawner через его сокет. Используется из хуков и вручную:
```bash
platform entity create --type=file --subtype=video --meta-path=/data/file.mp4
platform job create --kind=hash --worker=hash.sha256 --entity-id=<uuid>
platform spawn hash.sha256 --entity-id=<uuid>
```

## Схема БД

```sql
create table entities (
  id uuid primary key,
  type text not null,
  subtype text,
  name text,
  description text,
  meta jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
create index entities_type_idx on entities (type, subtype);
create index entities_meta_gin_idx on entities using gin (meta);

create table relations (
  id uuid primary key,
  from_id uuid not null references entities(id) on delete cascade,
  to_id uuid not null references entities(id) on delete cascade,
  type text not null,
  subtype text,
  value jsonb not null default '{}'::jsonb,
  meta jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  search_tsv tsvector
);
create index relations_from_idx on relations (from_id);
create index relations_to_idx on relations (to_id);
create index relations_type_idx on relations (type, subtype);
create index relations_value_gin_idx on relations using gin (value);
create index relations_meta_gin_idx on relations using gin (meta);
create index relations_search_tsv_gin_idx on relations using gin (search_tsv);

create table jobs (
  id uuid primary key,
  entity_id uuid references entities(id) on delete cascade,
  relation_id uuid references relations(id) on delete cascade,
  kind text not null,
  worker text not null,
  status text not null default 'pending',
  progress jsonb,
  error text,
  payload jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);
```

## Store DSL — формат query

```json
{"op": "query", "type": "entity", "filter": {...}, "cursor": "abc", "limit": 20, "order_by": "id"}
```

**Filter операторы:**
- Скалярные: `{"type": "file", "subtype": "video"}`
- Meta: `{"meta": {"path": {"glob": "/media/**"}, "size": {"gt": 10485760}, "mtime": {"gt": 1704067200}}}`
- Full-text: `{"search": "текст запроса"}`
- Вложенный lookup (1 уровень): `{"has_relation": {"type": "fact", "subtype": "hash.sha256", "value": {"hash": "..."}}}`
- OR-группы: filter как массив `[{...}, {...}]`

Pagination: cursor-based, cursor только с `order_by: id`.

## Search

- Структурный: `jsonb` containment (`@>`) и JSON path с GIN-индексами
- Full-text: `tsvector` + GIN на `relations.search_tsv` (агрегат `type`, `subtype`, `value`, `meta`)
- Pagination: cursor-based (`WHERE id > $cursor ORDER BY id LIMIT $limit`)

## Go: особенности реализации

- `jsonb` колонки маппятся через `driver.Valuer` / `sql.Scanner`
- `map[string]any` для динамических payload; typed structs когда схема стабильна

## Примеры payload в relations

**Hash**: `type=fact, subtype=hash:sha256` → `value: {"hash": "..."}`

**Stash info**: `type=fact, subtype=stash:info` → `value: {"sceneId": "987", "views": 42, "rating": 4.7}`, `meta: {"source": "stash@home"}`

**Derived preview**: `type=derivedFrom, subtype=preview` → `value: {"generator": "preview-worker:v1", "format": "jpeg", "width": 320, "height": 180}`

## Frontend

- **SvelteKit 5** с рунами, отдельный сервис в docker-compose на **Bun**
- Экраны MVP:
  - **Поиск / список entity** — один экран, фильтры по type/subtype/meta/full-text
  - **Карточка entity** — основная инфа + фоновая подгрузка количества relations по типам + таб с полным списком relations
  - **Карточка relation** — отдельная страница
  - **Монитор jobs** — список с polling (реалтайм SSE/WS — в roadmap)
  - **REPL** — JSON editor с подсветкой синтаксиса, `POST /commands`, ответ рядом

`POST /commands` — отладочный endpoint в API, проксирует команду напрямую в store socket.

## MVP scope

Реализовать первым делом:
- CRUD entities, relations, jobs (через store)
- Scanner worker (обходит ФС, создаёт `file` entity + `storedIn` relation)
- Hash worker (`hash.sha256`)
- Stash enricher
- Store + Spawner + API + Platform
- Поиск по `type/subtype`, JSONB containment, full-text

## Roadmap

**Нормальный приоритет:**
- Projections — денормализованные карточки entity (таблица `projections`, пересборка по событиям)
- Абстрактные доменные типы: `film`, `serial`, `book`
- Тип `external` (subtype: `stash_scene` и др.) — нужно продумать модель
- Коллекции — нужно уточнить модель (entity vs relation-паттерн)
- WebSocket/SSE для прогресса джоб в API

**Низкий приоритет:**
- Job статусы `paused`, `interrupted`
- Прогресс обработки конкретного файла (сейчас только job-уровень)
- Preview generator worker
- Image masks / OCR / embeddings
- Relation graph view
- Partial indexes под горячие search paths
