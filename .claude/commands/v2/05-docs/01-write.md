_Исполнитель: агент **Танк** (`Танк`)._

# docs-write

Написание документации к завершённой фиче.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/05-docs/01-write/README.md` в `tasks/$ARGUMENTS/processes/05-docs/01-write/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/05-docs/01-write/README.md`.

## Исполнение

Прочитай `docs/standards/v2/05-docs/01-write/base-plan.md` и `base-checklist.md`.

Входящие артефакты:
- Spec и acceptance из `tasks/$ARGUMENTS/processes/`
- Код реализации
- `PROJECT.md` — соглашения по документации

Пиши для читателя: что делает фича, как использовать (с примерами), поведение при ошибках, нетривиальные решения в TECH.md если нужно.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/05-docs/01-write/report-NNN.md`:

```markdown
---
purpose: Написание документации для $ARGUMENTS
process: 05-docs/01-write
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | clarification
agent: Танк
checklist: все пункты закрыты | открытые: {{список}}
---

## Созданные / обновлённые файлы

| Файл | Что добавлено |
|---|---|
| ... | ... |
```

Обнови `tasks/$ARGUMENTS/processes/05-docs/01-write/README.md`: добавь `updated: {{datetime}}`.
