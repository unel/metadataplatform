---
purpose: Описание процесса Acceptance: Write — точка входа для исполнителя и оркестратора
executor: Танк
next-on-success: 02-acceptance/02-review
next-on-failure: —
rollback-to: 01-spec/03-fix
---

# Acceptance: Write

**Исполнитель:** Танк

## Что делает

Составить тестовые сценарии по спеке в формате Given/When/Then. Покрыть все happy paths, failure modes, граничные условия и проверяемые НФТ.

## Входящие артефакты

- актуальная spec из `01-spec/`
- `PROJECT.md`

## Исходящие артефакты

- `report-NNN.md` — полный текст acceptance сценариев

## Навигация

| Исход | Следующий шаг |
|---|---|
| Успех | `02-acceptance/02-review` |
| Провал | — (нет отдельного процесса) |
| Откат / переосмысление | `01-spec/03-fix` — если обнаружен failure mode которого нет в spec |

## Артефакты процесса

- `base-plan.md` — базовый план выполнения
- `base-checklist.md` — базовый чек-лист
