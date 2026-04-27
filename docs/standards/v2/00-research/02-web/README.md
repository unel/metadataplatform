---
purpose: Описание процесса Research: Web — точка входа для исполнителя и оркестратора
executor: Бо
next-on-success: 01-spec/01-write
next-on-failure: —
rollback-to: 00-research/01-interview
---

# Research: Web

**Исполнитель:** Бо

## Что делает

Исследовать внешний контекст: похожие реализации, известные failure modes, устоявшиеся паттерны. Каждый факт — с URL источника.

## Входящие артефакты

- `00-research/01-interview/report-*.md`
- `TASK.md`
- `PROJECT.md`

## Исходящие артефакты

- `report-NNN.md` — реализации / failure modes / паттерны с URL
- `searches.md` (опц.) — поисковые запросы
- `sources.md` (опц.) — URL с описаниями

## Навигация

| Исход | Следующий шаг |
|---|---|
| Успех | `01-spec/01-write` |
| Провал | — (нет отдельного процесса) |
| Откат / переосмысление | `00-research/01-interview` — если в ходе поиска возникли вопросы к пользователю |

## Артефакты процесса

- `base-plan.md` — базовый план выполнения
- `base-checklist.md` — базовый чек-лист
