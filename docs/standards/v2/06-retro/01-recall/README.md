---
purpose: Описание процесса recall — агенты вспоминают и записывают наблюдения по фиче
executor: Харли Куин (оркестратор)
next-on-success: 06-retro/02-discuss
next-on-failure: 06-retro/01-recall
rollback-to: ~
---

# 06-retro/01-recall: Recall

## Что делает

Харли запускает всех участников фичи параллельно. Каждый читает свои артефакты и дополняет `notes-<агент>.md` (теги: rework / friction / miss / propose / doc / whatever) и `complaints-<агент>.md` (сырые жалобы без фильтра).

## Входящие артефакты

- Все артефакты фичи: spec, acceptance, код, тесты, docs, review-репорты
- Существующие `notes-*.md` и `complaints-*.md` (дополняются, не перезаписываются)
- `notes-user.md` и `complaints-user.md` если написаны пользователем

## Исходящие артефакты

- Обновлённые `notes-<агент>.md` и `complaints-<агент>.md` для каждого участника
- `06-retro/01-recall/report-NNN.md`

## Навигация

| Условие | Следующий шаг |
|---|---|
| Все агенты завершили recall | 06-retro/02-discuss |

## Артефакты процесса

- `base-plan.md` — инструкции для Харли
- `base-checklist.md` — чек-лист
- `report-NNN.md` — результат (иммутабельный)
