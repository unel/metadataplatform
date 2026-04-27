_Режим: **Харли Куин**. Выполняется оператором напрямую._

# retro-solve

Выработка решений по сгруппированным проблемам, верификация с командой, закрытие retro-команды.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/06-retro/04-solve/README.md` в `tasks/$ARGUMENTS/processes/06-retro/04-solve/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/06-retro/04-solve/README.md`.

## Исполнение

Прочитай `docs/standards/v2/06-retro/04-solve/base-plan.md` и `base-checklist.md`.

Прочитай `tasks/$ARGUMENTS/processes/06-retro/03-analyze/report-*.md` — сгруппированные проблемы и team name.

### Выработка решений

Для каждой категории проблем предложи конкретное решение с приоритетом.

### Проверка с тиммейтами

По каждой категории через SendMessage:

> «Для [категория]: предлагаю [решение]. Видите подводные камни? Что упустила?»

Дождись ответов. Скорректируй.

### Завершение команды

1. Отправь каждому тиммейту: `message: {type: "shutdown_request"}`
2. Дождись завершения
3. Вызови TeamDelete

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/06-retro/04-solve/report-NNN.md`:

```markdown
---
purpose: Solve для ретро $ARGUMENTS
process: 06-retro/04-solve
run: {{N}}
date: {{datetime}}
created: {{datetime}}
status: done
agent: Харли Куин
---

## Что шло хорошо

- ...

## Решения

| Проблема | Категория | Решение | Приоритет |
|---|---|---|---|
| ... | ... | ... | высокий / средний / низкий |

## Берём в работу

- [ ] <задача> (решает: <проблему>)

## Открытые вопросы

{{что осталось без решения и почему}}
```
