---
purpose: Выжимка интервью с пользователем — входящий артефакт для 00-research/02-web и 01-spec/01-write
process: 00-research/01-interview
run: 1
date: 2026-04-28T11:10:00Z
created: 2026-04-28T11:10:00Z
status: done
agent: Танк
checklist: все пункты закрыты
---

## Что хочет пользователь

CRUD поверх готовой socket-инфраструктуры (`store/connection`). Три типа: `entity`, `relation`, `job`. Четыре операции: `upsert`, `delete`, `get` (по id), `list` (все записи типа без фильтрации).

Go-интерфейсы: `EntityStore`, `RelationStore`, `JobStore`. Одна реализация — FS (JSON-файлы на диск), для отладки.

JSONL-роутер маппит `op` из входящего запроса на вызов интерфейса.

## Детали поведения

### Интерфейсы

**upsert:**
- `id` обязателен в запросе — если отсутствует, возвращать ошибку клиента
- стор не генерирует id самостоятельно
- `created_at` и `updated_at` стор выставляет сам, значения из `data` игнорируются
- `updated_at` при upsert — всегда текущее время, даже если данные не изменились

**get:**
- по `id`; запись не найдена → `nil, ErrNotFound`

**delete:**
- по `id`; запись не найдена → ошибка (не идемпотентный)
- `id` обязателен — отсутствие `id` → ошибка клиента

**list:**
- возвращает все записи данного типа без фильтрации
- фильтрация — это `store/query`, не эта задача

### UUID

UUID v7 обязателен во всех реализациях, включая FS.

### FS-реализация

- Структура директорий: `entities/<uuid>.json`, `relations/<uuid>.json`, `jobs/<uuid>.json`
- Atomic write: temp-файл + rename (защита от повреждения при crash)
- Конкурентный доступ: не защищается в этой итерации
- Требование к коду: структура должна позволять легко добавить locking позже (через интерфейс или отдельный слой); обязательна заметка в коде и/или документации

### JSONL-роутер

- Неизвестный `op` → ошибка (fail fast)
- Неизвестный `type` (не `entity` / `relation` / `job`) → ошибка
- Формат всех ошибок: `{"ok":false,"error":"...","errorCode":"..."}`
- Формат успеха: `{"ok":true,...}` (по PROJECT.md)

## Ограничения scope

- PostgreSQL-реализация → отдельная задача
- Query DSL (фильтры, JSONB, full-text, `has_relation`) → `store/query`
- Cursor pagination → `store/query`
- `FOR UPDATE SKIP LOCKED` для job polling → отдельно
- Конкурентная защита FS → отдельно

## Критерии успеха

- Go-интерфейсы определены для всех трёх типов
- FS-реализация проходит операции upsert/delete/get/list для каждого типа
- JSONL-роутер корректно маппит op → метод и возвращает ошибки в формате `{"ok":false,"error":"...","errorCode":"..."}`
- Все ошибочные случаи (missing id, unknown op, unknown type, not found, delete non-existent) возвращают явную ошибку, не молчат

## Что искать в web-research

- UUID v7 generation в Go: актуальные библиотеки (google/uuid v1.6+, gofrs/uuid и др.), поддержка v7
- Atomic file write в Go: паттерн temp + rename, существующие библиотеки (natefinch/atomic и др.)
- Паттерны Go-интерфейсов для storage layer: как структурировать под будущий locking без оверинжиниринга
- `ErrNotFound` паттерн в Go storage layers: sentinel error vs typed error

## Открытые вопросы

Нет.

## Retro notes

[retro] Пользователь зафиксировал: агенты не применили принцип fail fast как дефолт при проектировании роутера — потребовалось явное указание. В будущих задачах: fail fast для неизвестных/невалидных входных данных должен быть дефолтным решением, не опцией.
