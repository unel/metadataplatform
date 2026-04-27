_Исполнитель: агент **Танк** (`Танк`)._

Процесс `01-spec / 01-write`. Собери полную спецификацию фичи (ФТ + НФТ).

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## Шаг 1. Инициализация

Создай директорию `tasks/$ARGUMENTS/processes/01-spec/01-write/` если не существует.

Скопируй `README.md` из `docs/standards/v2/01-spec/01-write/README.md` в `tasks/$ARGUMENTS/processes/01-spec/01-write/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/01-spec/01-write/README.md`.

Запиши в `status-log.md`:
```
# {{datetime}} — pending
```

---

## Шаг 2. Генерация плана

Прочитай:
- `docs/standards/v2/01-spec/01-write/base-plan.md` — базовый план
- `tasks/$ARGUMENTS/TASK.md` — описание задачи
- `PROJECT.md` — архитектура, стек, ограничения
- последний `tasks/$ARGUMENTS/processes/00-research/01-interview/report-*.md` (если существует)
- последний `tasks/$ARGUMENTS/processes/00-research/02-web/report-*.md` (если существует)

Сгенерируй `tasks/$ARGUMENTS/processes/01-spec/01-write/plan.md`:
- Скопируй базовый план
- Адаптируй под конкретную задачу: убери нерелевантные шаги, добавь специфичные
- Укажи в мете версию базового плана

```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-plan: 01-spec-write v{{version}}
---

{{адаптированный план}}
```

---

## Шаг 3. Генерация чек-листа

Прочитай:
- `docs/standards/v2/01-spec/01-write/base-checklist.md` — базовый чек-лист
- нужные расширения из `docs/standards/v2/extensions/` — определи по контексту задачи
- сгенерированный `plan.md` — добавь чек-пункты отражающие шаги плана

Запиши `tasks/$ARGUMENTS/processes/01-spec/01-write/checklist.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-checklist: 01-spec-write v{{version}}
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

Выполняй по `plan.md`, отмечай пункты в `checklist.md`.

При неоднозначности в задаче — не выбирай молча. Опиши контекст, варианты с плюсами/минусами, жди решения пользователя.

Если входных данных недостаточно — запиши в `status-log.md`:
```
# {{datetime}} — clarification
Неясно: {{что именно}}. Ожидаю уточнения.
```
И сообщи пользователю. После уточнения — продолжи выполнение.

---

## Шаг 5. Завершение

Покажи черновик spec пользователю. Жди подтверждения или правок.

После подтверждения запиши `report-001.md` (при повторном прогоне — `report-002.md` и т.д.):

```markdown
---
process: 01-spec/01-write
run: 1
date: {{datetime}}
created: {{datetime}}
see-also:  # если информация из предыдущих прогонов актуальна
status: done | failed | clarification
agent: Танк
checklist: все пункты закрыты | открытые: {{список}}
---

## Spec

{{полный текст спецификации}}

## Решения принятые в процессе

- ...

## Открытые вопросы

- ...
```

Дозапиши в `status-log.md`:
```
# {{datetime}} — done
→ report-001.md
```

