_Исполнитель: агент **Танк** (`Танк`)._

Процесс `00-research / 01-interview`. Выясни у пользователя что он хочет и зачем.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## Шаг 1. Инициализация

Создай директорию `tasks/$ARGUMENTS/processes/00-research/01-interview/` если не существует.

Скопируй `README.md` из `docs/standards/v2/00-research/01-interview/README.md` в `tasks/$ARGUMENTS/processes/00-research/01-interview/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/00-research/01-interview/README.md`.

Запиши в `status-log.md`:
```
---
purpose: Хронология статусов процесса interview — оркестратор читает для понимания текущего состояния
---
created: {{datetime}}
updated: {{datetime}}

# {{datetime}} — pending
```

---

## Шаг 2. Генерация плана

Прочитай:
- `docs/standards/v2/00-research/01-interview/base-plan.md`
- `tasks/$ARGUMENTS/TASK.md`
- `PROJECT.md`
- `PLAN.md`

Сгенерируй `tasks/$ARGUMENTS/processes/00-research/01-interview/plan.md`:
```markdown
---
purpose: План процесса interview — исполнитель читает перед стартом
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-plan: 00-research-interview v{{version}}
---

{{адаптированный план: конкретные пробелы в понимании задачи, сгруппированные вопросы}}
```

---

## Шаг 3. Генерация чек-листа

Прочитай:
- `docs/standards/v2/00-research/01-interview/base-checklist.md`
- расширения из `docs/standards/v2/extensions/` по контексту
- сгенерированный `plan.md`

Запиши `tasks/$ARGUMENTS/processes/00-research/01-interview/checklist.md`:
```markdown
---
purpose: Чек-лист прогона interview — исполнитель отмечает по ходу выполнения
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-checklist: 00-research-interview v{{version}}
extensions:
  - {{extension}} v{{version}}
plan: {{datetime плана}}
---

{{базовый чек-лист + пункты из плана}}
```

Обнови `status-log.md`:
```
# {{datetime}} — in-progress
→ plan.md, checklist.md
```

---

## Шаг 4. Разговор с пользователем

Задавай вопросы сгруппированно. Не спрашивай то что уже есть в `TASK.md`.

Фиксируй ответы. При необходимости уточняй — одно уточнение за раз.

После ответов — сформулируй синтез и подтверди с пользователем.

---

## Шаг 5. Завершение

Подсчитай N, запиши `report-NNN.md`:

```markdown
---
purpose: Выжимка интервью с пользователем — входящий артефакт для 00-research/02-web и 01-spec/01-write
process: 00-research/01-interview
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:  # если информация из предыдущих прогонов актуальна
status: done
agent: Танк
checklist: все пункты закрыты | открытые: {{список}}
---

## Что хочет пользователь

{{конкретное описание: что делает фича, зачем}}

## Ограничения scope

{{что явно не входит}}

## Критерии успеха

{{как поймём что готово}}

## Что искать в web-research

{{конкретные направления для Бо}}

## Открытые вопросы

{{если что-то осталось неясным}}
```

Дозапиши в `status-log.md`:
```
# {{datetime}} — done
→ report-NNN.md
```

