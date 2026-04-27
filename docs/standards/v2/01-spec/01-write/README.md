---
purpose: Описание процесса Spec: Write — точка входа для исполнителя и оркестратора
executor: Танк
next-on-success: 01-spec/02-review
next-on-failure: —
rollback-to: 00-research/01-interview
---

# Spec: Write

**Исполнитель:** Танк

## Что делает

Написать полную спецификацию фичи: ФТ (happy path, failure modes, форматы) и НФТ (только релевантные категории). Черновик согласуется с пользователем.

## Входящие артефакты

- `TASK.md`
- `PROJECT.md`
- `PLAN.md`
- `00-research/01-interview/report-*.md`
- `00-research/02-web/report-*.md`

## Исходящие артефакты

- `report-NNN.md` — полный текст spec

## Навигация

| Исход | Следующий шаг |
|---|---|
| Успех | `01-spec/02-review` |
| Провал | — (нет отдельного процесса) |
| Откат / переосмысление | `00-research/01-interview` — если входных данных недостаточно для написания spec |

## Артефакты процесса

- `base-plan.md` — базовый план выполнения
- `base-checklist.md` — базовый чек-лист
