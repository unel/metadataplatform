---
feature: store/crud
generated: 2026-04-29T18:32:52Z
updated: 2026-05-04T12:11:22Z
source: docs/standards/v2/03-tests/02-review/README.md
---

# 03-tests/02-review: Ревью тестов

## Что делает

Гримм проверяет качество тестов: FIRST принципы, покрытие acceptance, структуру файлов, именование. Не запускает тесты — только читает код.

## Входящие артефакты

- Тест-файлы в `tests/happy/` и `tests/adversarial/`
- `03-tests/01-write/report-NNN.md` — отчёт о написании
- `02-acceptance/*/report-NNN.md` — для проверки покрытия

## Исходящие артефакты

- `03-tests/02-review/report-NNN.md` — результаты ревью с классификацией проблем

## Навигация

| Условие | Следующий шаг |
|---|---|
| Нет критичных замечаний | 04-code/01-write |
| Есть замечания к тестам | 03-tests/03-fix |
| Сценарий в acceptance не покрываем тестом | `clarification` → 02-acceptance/03-fix |

## Артефакты процесса

- `base-plan.md` — инструкции для Гримма
- `base-checklist.md` — чек-лист ревью
- `report-NNN.md` — результат ревью (иммутабельный)
