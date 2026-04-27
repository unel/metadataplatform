_Исполнитель: агент **Гримм** (`Гримм`)._

# docs-review

Ревью документации: точность, полнота, читаемость.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/05-docs/02-review/README.md` в `tasks/$ARGUMENTS/processes/05-docs/02-review/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/05-docs/02-review/README.md`.

## Исполнение

Прочитай `docs/standards/v2/05-docs/02-review/base-plan.md` и `base-checklist.md`.

Входящие артефакты:
- Документационные файлы (из `05-docs/01-write/report-*.md` — список файлов)
- Последний report из `05-docs/`
- Код реализации — для проверки точности

Не правь документацию — только описывай проблемы с классификацией critical / warning / Nit:.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/05-docs/02-review/report-NNN.md`:

```markdown
---
purpose: Ревью документации для $ARGUMENTS
process: 05-docs/02-review
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | failed
agent: Гримм
checklist: все пункты закрыты | открытые: {{список}}
---

## Результат

{{Документация чистая. | Найдено N замечаний.}}

## Замечания

### `path/to/doc.md`

**Категория:** critical | warning | Nit:
**Проблема:** {{описание}}
**Рекомендация:** {{что исправить}}
```

Обнови `tasks/$ARGUMENTS/processes/05-docs/02-review/README.md`: добавь `updated: {{datetime}}`.
