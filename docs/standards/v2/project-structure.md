---
purpose: Канонический стандарт структуры репозитория и workflow-артефактов
version: 1.0.0
created: 2026-05-06T09:55:58Z
updated: 2026-05-06T09:55:58Z
reviewed-by: Гримм (3 раунда, pass 2026-05-06)
---

# Структура проекта

> **Целевая структура.** Текущее состояние репозитория ей не соответствует.
> Переход будет выполнен в рамках отдельной задачи (BL-038).

---

## Терминология

| Термин | Англ. | Пример | Описание |
|---|---|---|---|
| этап | stage | `01-spec`, `04-code` | Группа шагов вокруг одного артефакта или цели |
| шаг | step | `01-write`, `02-review` | Конкретное действие внутри этапа |
| задача | task | `0002-store-crud` | Единица работы с полным workflow-деревом |

Папка `stages/` содержит этапы. Этап содержит шаги.
Если шаг потребует вложенности — вводим «подшаг» (sub-step) отдельным решением, не сейчас.

Полный каталог этапов и шагов (назначение, артефакты, исполнители) — в `docs/standards/v2/stages.md`.
Стандарт структуры шага, статус-лога и репорта — в `docs/standards/v2/stages-standard.md`.

---

## Принцип

Монорепа по компонентам: каждый компонент — отдельная папка со всем что к нему относится: исходники, тесты, Dockerfile, compose-фрагмент, README.
В корне — только то что относится ко всему проекту: описание, общие стандарты, tasks.

---

## Корень репозитория

```
metadataplatform/
│
├── store/              — socket-сервер + библиотека store-домена
├── api/                — HTTP API сервер
├── spawner/            — оркестратор воркеров и хуков
├── platform/           — системный CLI
├── frontend/           — SvelteKit-приложение
│
├── workers/            — внешние воркеры (спавнятся spawner'ом)
├── hooks/              — внешние хуки (спавнятся spawner'ом)
│
├── deploy/             — docker-compose и среды
├── scripts/            — общие dev/ops скрипты
├── docs/               — общая документация и стандарты
├── tasks/              — workflow-артефакты (не код)
│
├── PROJECT.md          — описание проекта, схема компонентов, roadmap
├── CLAUDE.md
├── go.mod              — единый go.mod для всего Go-кода
└── go.sum
```

---

## Компоненты

Go-компоненты не используют `src/` — Go-конвенция: пакеты лежат непосредственно в директории компонента.
`dist/` — build-артефакты (gitignored) — у всех компонентов со сборкой.
`docker/` — опциональна. Если есть — строго по структуре ниже. README.md компонента объясняет наличие или отсутствие.

### Go-компоненты

```
store/
├── cmd/
│   └── main.go             — точка входа (package main)
├── fs/                     — адаптер: хранение в файловой системе
├── router/                 — роутер команд (JSONL → handler)
├── interfaces.go           — package store: интерфейсы
├── types.go                — package store: типы и ошибки
├── *_test.go               — тесты рядом с кодом
├── dist/                   — собранный бинарь (gitignored)
├── docker/
│   ├── Dockerfile
│   └── compose.yml         — сервис для docker-compose include
└── README.md
```

`api/`, `spawner/` — та же структура: `cmd/`, внутренние пакеты, тесты, `dist/`, `docker/`.

`platform/` — CLI-инструмент, без Docker:
```
platform/
├── cmd/
│   └── main.go
├── commands/
├── *_test.go
└── dist/
```

Go-билд: `go build -o <component>/dist/<name> ./<component>/cmd/`

### `frontend/` — SvelteKit

SvelteKit требует `src/` по своей конвенции.

```
frontend/
├── src/
│   ├── lib/
│   └── routes/
├── static/
├── dist/                   — сборка (gitignored)
├── docker/
│   ├── Dockerfile
│   └── compose.yml
├── package.json
├── svelte.config.js
├── tsconfig.json
└── README.md
```

### `workers/` и `hooks/`

Каждый воркер и хук — отдельная поддиректория. Стек произвольный.

```
workers/
├── hash.sha256/
│   └── README.md
└── scanner/
    └── README.md

hooks/
└── on-file-created/
    └── README.md
```

---

## Общая инфраструктура

### `deploy/`

Каждый компонент описывает свой сервис в `docker/compose.yml`.
Корневой `deploy/docker-compose.yml` собирает их через `include` (Compose spec 1.6+).

```
deploy/
├── docker-compose.yml        — собирает compose.yml компонентов через include
└── docker-compose.dev.yml    — dev-оверрайды
```

```yaml
# deploy/docker-compose.yml
include:
  - path: ../store/docker/compose.yml
  - path: ../api/docker/compose.yml
  - path: ../spawner/docker/compose.yml
  - path: ../frontend/docker/compose.yml
```

### `scripts/`

_Планируется (BL-010, BL-011) — пока не существуют._

```
scripts/
├── feature-status.sh     — статус всех шагов задачи (BL-010)
├── log-note.sh           — добавить заметку агента (BL-011)
└── log-complaint.sh      — добавить жалобу агента (BL-011)
```

### `docs/`

```
docs/
├── standards/
│   ├── go/               — Go-стандарты (concurrency, errors, testing…)
│   ├── frontend/         — SvelteKit / TypeScript
│   ├── architecture/     — общие архитектурные принципы
│   └── v2/               — стандарты workflow v2
├── skills-guide/         — документация по скиллам и агентам
├── rfc/                  — RFC-документы
└── deps/                 — графы зависимостей между задачами
```

---

## Структура `tasks/`

```
tasks/
├── _backlog/
│   ├── BACKLOG.md
│   └── actions/              — action items не привязанные к задаче
│
├── 0001-store-connection/
├── 0002-store-crud/
└── NNNN-<slug>/
```

Задачи нумеруются глобально в порядке создания. Нет вложенности по компонентам — номер даёт хронологию, slug даёт смысл.

### Структура одной задачи

```
NNNN-<slug>/
│
├── stages/
│   ├── <NN-этап>/            — напр. 01-spec, 04-code, 06-retro
│   │   └── <NN-шаг>/         — напр. 01-write, 02-review, 03-fix
│   │       ├── status-log.md
│   │       ├── README.md
│   │       ├── report-NNN.md
│   │       ├── notes-<agent>.md
│   │       └── complaints-<agent>.md
│   └── 06-retro/
│       └── <NN-шаг>/
│           └── ...
│
├── artifacts/
│   ├── spec.md               — единственная актуальная спека
│   ├── acceptance.md         — единственный актуальный acceptance
│   ├── retro.md              — итоговый документ ретро (06-retro/05-write)
│   └── actions/              — action items из ретро
│       └── NN-<slug>.md
│
├── STRUCTURE.md              — что где лежит в этой задаче
└── README.md                 — зачем задача
```

**Правила:**

- `artifacts/spec.md` и `artifacts/acceptance.md` — единственный источник истины. `write` создаёт, `fix` обновляет на месте. `report-NNN.md` — дельта прогона, не полная копия.
- Go-тесты пишутся сразу в код компонента. В report — путь к файлам от корня репо.
- `notes-<agent>.md` и `complaints-<agent>.md` — агент пишет в папку своего шага по мере работы. При агрегации на `06-retro/01-recall` агент читает все notes/complaints по всем шагам задачи, переосмысляет и пишет синтез в `stages/06-retro/01-recall/notes-<agent>.md` и `complaints-<agent>.md` — со ссылками на первоисточники.
- Статус задачи — только через `scripts/feature-status.sh`, который читает `status-log.md` каждого шага. Ручной STATUS.md не ведётся.

### Конвенция ссылок на артефакты

Короткая форма — в display text (читаемо), полный путь — в href (кликабельно).

| Короткая форма | Полный путь |
|---|---|
| `<slug>/act/<action-slug>` | `tasks/NNNN-<slug>/artifacts/actions/<action-slug>.md` |
| `<slug>/spec` | `tasks/NNNN-<slug>/artifacts/spec.md` |
| `<slug>/acc` | `tasks/NNNN-<slug>/artifacts/acceptance.md` |

Пример (пути актуальны после выполнения BL-038):
```markdown
- [ ] **BL-001** — API Contracts в acceptance → [store-crud/act/01-api-contracts](../0002-store-crud/artifacts/actions/01-api-contracts.md)
```

---

## Отложено

- Миграция репо на эту структуру — BL-038 (включает: переименование tasks/, artifacts/, stages/, обновление PROJECT.md, судьба `_backlog/retro-*.md`).
- Миграция `tasks/0001-store-connection` (legacy без `stages/`) и `tasks/0002-store-crud` (частичное соответствие: нет `artifacts/`, нет `STRUCTURE.md`) — в рамках BL-038.
