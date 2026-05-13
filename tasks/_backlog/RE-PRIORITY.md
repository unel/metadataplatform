---
created: 2026-05-13T18:19:09Z
updated: 2026-05-13T18:19:09Z
---

# RE-задачи: приоритеты

## Tier 1 — Процессные дыры (рецидивы из каждой фичи)

| ID | Что | Что менять |
|---|---|---|
| RE-006 | Notes/complaints как обязательный шаг в write-скиллах | 5 base-plan.md (write) |
| RE-009 | Fix-report обязателен после failed | 5+ base-plan.md (review) |
| RE-012 | in-progress при старте шага, done при финале | base-plan каждого скилла |
| RE-004 | 01-recall агрегирует notes/complaints из всего дерева | 06-retro/01-recall |
| RE-008 | Gate перед 06-retro: все шаги done или явное решение | 06-retro/01-recall |
| RE-007 | Gate в pipeline: upstream не в pending → стоп | CLAUDE.md |

## Tier 2 — Качество артефактов (template improvements)

| ID | Что | Что менять |
|---|---|---|
| RE-001 | API Contracts в шаблоне acceptance | 02-acceptance/01-write |
| RE-003 | Инварианты валидации в acceptance | 02-acceptance/01-write |
| RE-002 | Контракт vs Реализация на Go в spec | 01-spec/01-write |
| RE-014 | Error types секция в spec | 01-spec/01-write |
| RE-015 | spec.md — единый источник истины | 01-spec/{write,fix,review} |

## Tier 3 — Reviewer improvements

| ID | Что | Что менять |
|---|---|---|
| RE-019 | `docs/standards/go/` явно в промпт ревьюера | все */02-review |
| RE-020 | Оркестратор передаёт список файлов ревьюеру | оркестратор CLAUDE.md |
| RE-016 | API-несовместимость → clarification, не фикс в code | 04-code/01-write |
| RE-022 | `go build` отдельным шагом перед `go test` | 03-tests/04-run, 04-code/04-testing |

## Tier 4 — Поведение агентов

| ID | Что | Что менять |
|---|---|---|
| RE-024 | Fail fast как дефолт | AGENT.md каждого агента (8 файлов) |
| RE-018 | Auto-proceed failed→fix без паузы | оркестратор CLAUDE.md |
| RE-021 | Report самодостаточен для context-recovery | base-plan всех write |

## Tier 5 — Стандарты и документация

| ID | Что | Что менять |
|---|---|---|
| RE-025 | Datetime через `date -u` задокументировать | новый standards файл |
| RE-028 | frontmatter стандарт + валидация | новый standards файл |
| RE-011 | `log-note.sh` и `log-complaint.sh` скрипты | новые файлы в scripts/ |

## Tier 6 — Код и рефакторинг

| ID | Что |
|---|---|
| RE-036 | Обновить Go до 1.26.2 |
| RE-040 | Вынести store/cmd в internal-пакеты |

## Tier 7 — RFC и исследование

| ID | Что |
|---|---|
| RE-035 | RFC: Харли + Agent tool |
| RE-037 | Система флоу |
| RE-042 | Флоу для не-фичевых задач |
| RE-047 | Karpathy CLAUDE.md практики |
