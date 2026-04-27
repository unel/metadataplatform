_Исполнитель: агент **Гримм** (`Гримм`)._

# test-review

Ревью тестов: FIRST, покрытие, поведение vs реализация, структура файлов.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/03-tests/02-review/README.md` в `tasks/$ARGUMENTS/processes/03-tests/02-review/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/03-tests/02-review/README.md`.

## Исполнение

Прочитай `docs/standards/v2/03-tests/02-review/base-plan.md` и `base-checklist.md`.

Входящие артефакты:
- `tasks/$ARGUMENTS/tests/happy/` — happy path тесты
- `tasks/$ARGUMENTS/tests/adversarial/` — adversarial тесты
- Последний `tasks/$ARGUMENTS/processes/03-tests/01-write/report-*.md` или `03-fix/report-*.md`
- Финальный acceptance из `tasks/$ARGUMENTS/processes/02-acceptance/`

Не исправляй тесты — только описывай проблемы с классификацией.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/03-tests/02-review/report-NNN.md`:

```markdown
---
purpose: Ревью тестов для $ARGUMENTS
process: 03-tests/02-review
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | failed | clarification
agent: Гримм
checklist: все пункты закрыты | открытые: {{список}}
---

## Результат

{{Тесты чистые. Готово к code-write. | Найдено N замечаний.}}

## Замечания

### <Тест / файл>

**Категория:** пробел в покрытии | нарушение FIRST | проверка реализации | неопределённость в acceptance
**Проблема:** {{описание}}
**Рекомендация:** {{что исправить}}
```

Обнови `tasks/$ARGUMENTS/processes/03-tests/02-review/README.md`: добавь `updated: {{datetime}}`.
