_Исполнитель: агент **Танк** (`Танк`)._

# docs-fix

Исправление замечаний из docs-review.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/05-docs/03-fix/README.md` в `tasks/$ARGUMENTS/processes/05-docs/03-fix/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/05-docs/03-fix/README.md`.

## Исполнение

Прочитай `docs/standards/v2/05-docs/03-fix/base-plan.md` и `base-checklist.md`.

Входящие артефакты:
- Последний `tasks/$ARGUMENTS/processes/05-docs/02-review/report-*.md` — замечания
- Документационные файлы

Покажи план фикса пользователю. Жди подтверждения.

Исправляй только то что указано в review. Приоритет: critical → warning → Nit:.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/05-docs/03-fix/report-NNN.md`:

```markdown
---
purpose: Фикс документации для $ARGUMENTS
process: 05-docs/03-fix
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done
agent: Танк
checklist: все пункты закрыты | открытые: {{список}}
---

## Исправленные замечания

| Замечание | Что сделано |
|---|---|
| ... | ... |

## Неисправленные

| Замечание | Причина |
|---|---|
| ... | ... |
```

Обнови `tasks/$ARGUMENTS/processes/05-docs/03-fix/README.md`: добавь `updated: {{datetime}}`.
