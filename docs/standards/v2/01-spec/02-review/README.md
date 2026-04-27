---
purpose: Описание процесса Spec: Review — точка входа для исполнителя и оркестратора
executor: Гримм
next-on-success: 02-acceptance/01-write
next-on-failure: 01-spec/03-fix
rollback-to: 01-spec/01-write
---

# Spec: Review

**Исполнитель:** Гримм

## Что делает

Проверить спеку на противоречия, полноту ФТ/НФТ, corner cases. Каждое замечание классифицируется: critical / warning / note.

## Входящие артефакты

- `01-spec/01-write/report-*.md` или `01-spec/03-fix/report-*.md`
- `PROJECT.md`

## Исходящие артефакты

- `report-NNN.md` — замечания с классификацией, или 'Спека чистая'

## Навигация

| Исход | Следующий шаг |
|---|---|
| Успех | `02-acceptance/01-write` |
| Провал | `01-spec/03-fix` |
| Откат / переосмысление | `01-spec/01-write` — если spec настолько неполна что ревью невозможно |

## Артефакты процесса

- `base-plan.md` — базовый план выполнения
- `base-checklist.md` — базовый чек-лист
