_Режим: **Харли Куин**. Выполняется оператором напрямую._

# retro-collect

Создание retro-команды и параллельный сбор выжимок от участников фичи.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/06-retro/02-collect/README.md` в `tasks/$ARGUMENTS/processes/06-retro/02-collect/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/06-retro/02-collect/README.md`.

## Исполнение

Прочитай `docs/standards/v2/06-retro/02-collect/base-plan.md` и `base-checklist.md`.

Прочитай `tasks/$ARGUMENTS/processes/06-retro/01-recall/report-*.md` — возьми список участников.

### Создание команды

Создай team через TeamCreate:
- `team_name`: `retro-<последняя часть пути $ARGUMENTS>`
- `description`: `Ретроспектива фичи $ARGUMENTS`

### Параллельный спаун

Запусти всех участников **одновременно** с `mode: bypassPermissions`. Каждый:
- Читает `tasks/$ARGUMENTS/notes-<имя>.md` и `complaints-<имя>.md`
- Читает `tasks/$ARGUMENTS/notes-user.md` и `complaints-user.md` если существуют
- Отправляет выжимку (3–5 пунктов) лидеру

Промпт каждому (подставь имя):

> Ты участвуешь в ретроспективе фичи `$ARGUMENTS` как тиммейт.
>
> Прочитай:
> - `tasks/$ARGUMENTS/notes-<имя>.md`
> - `tasks/$ARGUMENTS/complaints-<имя>.md`
> - `tasks/$ARGUMENTS/notes-user.md` и `tasks/$ARGUMENTS/complaints-user.md` если существуют
>
> Подготовь выжимку: 3–5 пунктов — что болело сильнее всего, какие паттерны видишь, что предлагаешь изменить.
>
> Отправь выжимку лидеру. Затем жди вопросов.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/06-retro/02-collect/report-NNN.md`:

```markdown
---
purpose: Collect участников для ретро $ARGUMENTS
process: 06-retro/02-collect
run: {{N}}
date: {{datetime}}
created: {{datetime}}
status: done
agent: Харли Куин
team: retro-<suffix>
---

## Участники

| Агент | notes (строк) | complaints (строк) |
|---|---|---|
| ... | N | N |

## Выжимки

### <Агент>

- ...

## Итого проблем до группировки: N
```
