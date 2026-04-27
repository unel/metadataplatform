---
purpose: Описание процесса написания кода — точка входа для Ады и оркестратора
executor: Ада Лавлейс
next-on-success: 04-code/02-review
next-on-failure: 04-code/03-fix
rollback-to: 03-tests/03-fix
---

# 04-code/01-write: Написание кода

## Что делает

Ада пишет минимально необходимую реализацию чтобы тесты прошли (TDD Green). Код соответствует spec и acceptance, не содержит ничего лишнего.

## Входящие артефакты

- `03-tests/04-run/report-NNN.md` — первый прогон (все тесты красные)
- Тест-файлы `tests/happy/` и `tests/adversarial/` — определяют контракт
- `01-spec/*/report-NNN.md` и `02-acceptance/*/report-NNN.md` — для понимания требований

## Исходящие артефакты

- Код реализации в соответствующей директории проекта
- `04-code/01-write/report-NNN.md`

## Навигация

| Условие | Следующий шаг |
|---|---|
| Код написан, тесты зелёные | 04-code/02-review |
| Тесты требуют корректировки | `clarification` → 03-tests/03-fix |
| Acceptance неполный или противоречивый | `clarification` → 02-acceptance/03-fix |

## Артефакты процесса

- `base-plan.md` — инструкции для Ады
- `base-checklist.md` — чек-лист качества кода
- `report-NNN.md` — результат (иммутабельный)
