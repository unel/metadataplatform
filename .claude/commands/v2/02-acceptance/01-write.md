_Исполнитель: агент **Танк** (`Танк`)._

Процесс `02-acceptance / 01-write`. Составь тестовые сценарии по спеке.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## Шаг 1. Инициализация

Создай директорию `tasks/$ARGUMENTS/processes/02-acceptance/01-write/` если не существует.

Скопируй `README.md` из `docs/standards/v2/02-acceptance/01-write/README.md` в `tasks/$ARGUMENTS/processes/02-acceptance/01-write/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/02-acceptance/01-write/README.md`.

Запиши в `status-log.md`:
```
# {{datetime}} — pending
```

---

## Шаг 2. Генерация плана

Прочитай:
- `docs/standards/v2/02-acceptance/01-write/base-plan.md` — базовый план
- актуальную spec: последний `tasks/$ARGUMENTS/processes/01-spec/03-fix/report-*.md`, если нет — `tasks/$ARGUMENTS/processes/01-spec/01-write/report-*.md`
- `PROJECT.md`

Сгенерируй `tasks/$ARGUMENTS/processes/02-acceptance/01-write/plan.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-plan: 02-acceptance-write v{{version}}
---

{{адаптированный план}}
```

---

## Шаг 3. Генерация чек-листа

Прочитай:
- `docs/standards/v2/02-acceptance/01-write/base-checklist.md`
- нужные расширения из `docs/standards/v2/extensions/`
- сгенерированный `plan.md`

Запиши `tasks/$ARGUMENTS/processes/02-acceptance/01-write/checklist.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-checklist: 02-acceptance-write v{{version}}
extensions:
  - {{extension}} v{{version}}
plan: {{datetime плана}}
---

{{базовый чек-лист + пункты из плана + расширения}}
```

Обнови `status-log.md`:
```
# {{datetime}} — in-progress
Расширения: {{список}}. → plan.md, checklist.md
```

---

## Шаг 4. Выполнение

Составляй сценарии в формате Given/When/Then.

**Критически важно:** если при составлении сценариев обнаружен новый failure mode которого нет в spec — не угадывай. Запиши:
```
# {{datetime}} — clarification
Обнаружен failure mode не описанный в spec: {{описание}}.
Требуется откат к 01-spec/03-fix.
```
Сообщи пользователю и остановись.

---

## Шаг 5. Завершение

Покажи черновик acceptance пользователю. Жди подтверждения или правок.

После подтверждения подсчитай N и запиши `report-NNN.md`:

```markdown
---
process: 02-acceptance/01-write
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:  # если информация из предыдущих прогонов актуальна
status: done | clarification
agent: Танк
checklist: все пункты закрыты | открытые: {{список}}
---

## Acceptance

{{полный текст acceptance сценариев}}

## Решения принятые в процессе

- ...

## Открытые вопросы

- ...
```

Дозапиши в `status-log.md`:
```
# {{datetime}} — done
→ report-NNN.md
```

