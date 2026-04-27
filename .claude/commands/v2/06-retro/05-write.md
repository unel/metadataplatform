_Режим: **Харли Куин**. Выполняется оператором напрямую._

# retro-write

Запись финального RETRO.md по итогам обсуждения.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/06-retro/03-write/README.md` в `tasks/$ARGUMENTS/processes/06-retro/03-write/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/06-retro/03-write/README.md`.

## Исполнение

Прочитай `docs/standards/v2/06-retro/03-write/base-plan.md` и `base-checklist.md`.

Прочитай `tasks/$ARGUMENTS/processes/06-retro/02-discuss/report-*.md`.

Запиши `tasks/$ARGUMENTS/RETRO.md` по шаблону из base-plan: что шло хорошо, проблемы по категориям, таблица решений, action items, открытые вопросы.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/06-retro/03-write/report-NNN.md`:

```markdown
---
purpose: Запись RETRO.md для $ARGUMENTS
process: 06-retro/03-write
run: {{N}}
date: {{datetime}}
created: {{datetime}}
status: done
agent: Харли Куин
---

## Результат

RETRO.md записан: `tasks/$ARGUMENTS/RETRO.md`

Проблем: N (по категориям: процессные N, инструментальные N, требования N, командные N)
Решений: N
Action items: N
```
