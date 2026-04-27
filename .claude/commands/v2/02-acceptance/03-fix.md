_Исполнитель: агент **Танк** (`Танк`)._

Процесс `02-acceptance / 03-fix`. Исправь acceptance по замечаниям ревью.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## Шаг 1. Инициализация

Создай директорию `tasks/$ARGUMENTS/processes/02-acceptance/03-fix/` если не существует.

Скопируй `README.md` из `docs/standards/v2/02-acceptance/03-fix/README.md` в `tasks/$ARGUMENTS/processes/02-acceptance/03-fix/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/02-acceptance/03-fix/README.md`.

Запиши в `status-log.md`:
```
# {{datetime}} — pending
```

---

## Шаг 2. Генерация плана

Прочитай:
- `docs/standards/v2/02-acceptance/03-fix/base-plan.md`
- актуальную spec
- актуальный acceptance: последний `02-acceptance/03-fix/report-*.md`, если нет — `02-acceptance/01-write/report-*.md`
- замечания: последний `02-acceptance/02-review/report-*.md`

Сгенерируй `tasks/$ARGUMENTS/processes/02-acceptance/03-fix/plan.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-plan: 02-acceptance-fix v{{version}}
---

{{план: для каждого AR-N — что меняется}}
```

---

## Шаг 3. Генерация чек-листа

Запиши `tasks/$ARGUMENTS/processes/02-acceptance/03-fix/checklist.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-checklist: 02-acceptance-fix v{{version}}
extensions:
  - {{extension}} v{{version}}
plan: {{datetime плана}}
---
```

Обнови `status-log.md`:
```
# {{datetime}} — in-progress
```

---

## Шаг 4. Согласование

Покажи план исправлений пользователю. Жди подтверждения.

Проблемы типа "неопределённость/противоречие в spec" не трогай — они требуют отката к `01-spec/03-fix`. Если они есть — запиши в status-log:
```
# {{datetime}} — clarification
AR-N требует исправления spec: {{описание}}. Ожидаю откат к 01-spec/03-fix.
```

---

## Шаг 5. Выполнение

Исправляй точечно по замечаниям.

---

## Шаг 6. Завершение

Запиши `report-NNN.md`:

```markdown
---
process: 02-acceptance/03-fix
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:  # если информация из предыдущих прогонов актуальна
status: done | clarification
agent: Танк
checklist: все пункты закрыты | открытые: {{список}}
---

## Acceptance (обновлённый)

{{полный текст acceptance}}

## Changelog

- AR-1: {{что изменено}}
- AR-2: не изменялось — {{причина}}

## Открытые вопросы

- ...
```

Обнови статусы:
- `02-acceptance/02-review` → `pending`
- `03-tests`, `04-code`, `05-docs` → `stale` (если выполнены)

Дозапиши в `status-log.md`:
```
# {{datetime}} — done
→ report-NNN.md
```

