---
purpose: Описание процесса Acceptance: Review — точка входа для исполнителя и оркестратора
executor: Гримм
next-on-success: 03-tests/01-write
next-on-failure: 02-acceptance/03-fix
rollback-to: 01-spec/03-fix
---

# Acceptance: Review

**Исполнитель:** Гримм

## Что делает

Проверить что acceptance полностью покрывает spec. Классифицировать проблемы: пробел в acceptance / неопределённость в spec / противоречие в spec / качество.

## Входящие артефакты

- актуальная spec из `01-spec/`
- `02-acceptance/01-write/report-*.md`

## Исходящие артефакты

- `report-NNN.md` — замечания с классификацией, или 'Acceptance чистый'

## Навигация

| Исход | Следующий шаг |
|---|---|
| Успех | `03-tests/01-write` |
| Провал | `02-acceptance/03-fix` |
| Откат / переосмысление | `01-spec/03-fix` — если найдена неопределённость или противоречие в spec |

## Артефакты процесса

- `base-plan.md` — базовый план выполнения
- `base-checklist.md` — базовый чек-лист
