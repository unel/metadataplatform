_Исполнитель: агент **Гримм** (`Гримм`)._

# code-review

Ревью кода: соответствие требованиям, SOLID, безопасность, читаемость, размер модулей.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/04-code/02-review/README.md` в `tasks/$ARGUMENTS/processes/04-code/02-review/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/04-code/02-review/README.md`.

## Исполнение

Прочитай `docs/standards/v2/04-code/02-review/base-plan.md` и `base-checklist.md`.

Входящие артефакты:
- Код реализации (по пути из `04-code/01-write/report-*.md`)
- Последний `tasks/$ARGUMENTS/processes/04-code/01-write/report-*.md` или `03-fix/report-*.md`
- Spec и acceptance из `tasks/$ARGUMENTS/processes/`

Не исправляй код — только описывай проблемы с классификацией critical / warning / note.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/04-code/02-review/report-NNN.md`:

```markdown
---
purpose: Ревью кода для $ARGUMENTS
process: 04-code/02-review
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | failed
agent: Гримм
checklist: все пункты закрыты | открытые: {{список}}
---

## Результат

{{Код чистый. Готово к финальному прогону тестов. | Найдено N замечаний (critical: X, warning: Y, note: Z).}}

## Замечания

### `path/to/file.go:42`

**Категория:** critical | warning | note
**Проблема:** {{описание}}
**Рекомендация:** {{что исправить}}
```

Обнови `tasks/$ARGUMENTS/processes/04-code/02-review/README.md`: добавь `updated: {{datetime}}`.
