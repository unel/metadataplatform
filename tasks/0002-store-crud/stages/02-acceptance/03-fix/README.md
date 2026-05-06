---
purpose: Описание процесса Acceptance: Fix — точка входа для исполнителя и оркестратора
executor: Танк
next-on-success: 02-acceptance/02-review
next-on-failure: —
rollback-to: 01-spec/03-fix
feature: store/crud
generated: 2026-04-29T13:36:15Z
source: docs/standards/v2/02-acceptance/03-fix/README.md
---

# Acceptance: Fix

**Исполнитель:** Танк

## Что делает

Исправить acceptance по замечаниям ревью. Только пробелы и проблемы качества — проблемы spec передаются оркестратору для отката.

## Входящие артефакты

- `02-acceptance/02-review/report-*.md`
- актуальный acceptance из `01-write` или предыдущего `03-fix`

## Исходящие артефакты

- `report-NNN.md` — обновлённый acceptance + changelog

## Навигация

| Исход | Следующий шаг |
|---|---|
| Успех | `02-acceptance/02-review` |
| Провал | — (нет отдельного процесса) |
| Откат / переосмысление | `01-spec/03-fix` — если замечание касается неопределённости или противоречия в spec |

## Артефакты процесса

- `base-plan.md` — базовый план выполнения
- `base-checklist.md` — базовый чек-лист
