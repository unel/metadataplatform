---
purpose: Описание процесса написания тестов — точка входа для исполнителей и оркестратора
executor: Кроули + Азирафаль (параллельно)
---

# 03-tests/01-write: Написание тестов

## Что делает

Два агента пишут тесты параллельно по acceptance:
- **Азирафаль** — happy path и contract тесты (`tests/happy/`)
- **Кроули** — adversarial тесты: edge cases, failure modes, граничные условия (`tests/adversarial/`)

## Входящие артефакты

- `02-acceptance/01-write/report-NNN.md` или `02-acceptance/03-fix/report-NNN.md` — финальный acceptance
- `01-spec/*/report-NNN.md` — для понимания архитектурного контекста

## Исходящие артефакты

- Тест-файлы в `tests/happy/` и `tests/adversarial/`
- `03-tests/01-write/report-NNN.md` — результаты прогона, список что покрыто

## Артефакты процесса

- `base-plan.md` — инструкции для исполнителей
- `base-checklist.md` — чек-лист качества тестов
- `report-NNN.md` — результат прогона (иммутабельный)
