_Режим: **Харли Куин**. Выполняется оператором напрямую._

# retro-recall

Параллельный recall: все участники фичи дополняют свои notes и complaints.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/06-retro/01-recall/README.md` в `tasks/$ARGUMENTS/processes/06-retro/01-recall/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/06-retro/01-recall/README.md`.

## Исполнение

Прочитай `docs/standards/v2/06-retro/01-recall/base-plan.md` и `base-checklist.md`.

Прочитай директорию `tasks/$ARGUMENTS/` — определи участников по артефактам.

Запусти всех агентов **параллельно** с `mode: bypassPermissions`. Каждый:
- Читает свои артефакты фичи
- Читает `tasks/$ARGUMENTS/notes-user.md` и `complaints-user.md` если существуют
- Дополняет (append) `tasks/$ARGUMENTS/notes-<агент>.md` и `complaints-<агент>.md`

Промпт каждому (подставь имя):

> Ты участвуешь в ретроспективе фичи `$ARGUMENTS`. Прочитай свои артефакты. Прочитай `tasks/$ARGUMENTS/notes-user.md` и `complaints-user.md` если существуют. Дополни `tasks/$ARGUMENTS/notes-<имя>.md` (формат: тег + наблюдение) и `tasks/$ARGUMENTS/complaints-<имя>.md` (сырые жалобы без фильтра). Не удаляй существующие записи.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/06-retro/01-recall/report-NNN.md`:

```markdown
---
purpose: Recall участников для ретро $ARGUMENTS
process: 06-retro/01-recall
run: {{N}}
date: {{datetime}}
created: {{datetime}}
status: done
agent: Харли Куин
---

## Обновлённые файлы

| Агент | notes | complaints | Наблюдений добавлено |
|---|---|---|---|
| ... | ✓/— | ✓/— | N |

## Незаполненные

{{агенты у которых не нашлось наблюдений}}
```
