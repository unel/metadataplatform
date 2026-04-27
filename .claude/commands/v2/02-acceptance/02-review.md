_Исполнитель: агент **Гримм** (`Гримм`)._

Процесс `02-acceptance / 02-review`. Проверь что сценарии полностью покрывают спеку.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## Шаг 1. Инициализация

Создай директорию `tasks/$ARGUMENTS/processes/02-acceptance/02-review/` если не существует.

Скопируй `README.md` из `docs/standards/v2/02-acceptance/02-review/README.md` в `tasks/$ARGUMENTS/processes/02-acceptance/02-review/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/02-acceptance/02-review/README.md`.

Запиши в `status-log.md`:
```
# {{datetime}} — pending
```

---

## Шаг 2. Генерация плана

Прочитай:
- `docs/standards/v2/02-acceptance/02-review/base-plan.md`
- актуальную spec (последний report из `01-spec/`)
- acceptance из последнего `tasks/$ARGUMENTS/processes/02-acceptance/01-write/report-*.md`

Сгенерируй `tasks/$ARGUMENTS/processes/02-acceptance/02-review/plan.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-plan: 02-acceptance-review v{{version}}
---

{{адаптированный план}}
```

---

## Шаг 3. Генерация чек-листа

Прочитай:
- `docs/standards/v2/02-acceptance/02-review/base-checklist.md`
- расширения из `docs/standards/v2/extensions/`
- `plan.md`

Запиши `tasks/$ARGUMENTS/processes/02-acceptance/02-review/checklist.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-checklist: 02-acceptance-review v{{version}}
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

## Шаг 4. Выполнение

Выполняй по `plan.md`, отмечай пункты в `checklist.md`.

Классифицируй проблемы:
- **Пробел в acceptance** — поведение описано в spec, сценарий не написан
- **Неопределённость в spec** — сценарий невозможно написать, поведение не описано → нужен откат к `01-spec/03-fix`
- **Противоречие в spec** — два требования несовместимы → нужен откат к `01-spec/03-fix`
- **Проблема качества** — сценарий расплывчатый или тривиальный

Не исправлять acceptance или spec самостоятельно — только описывать проблемы.

---

## Шаг 5. Завершение

Подсчитай N, запиши `report-NNN.md`:

```markdown
---
process: 02-acceptance/02-review
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:  # если информация из предыдущих прогонов актуальна
status: done | failed | clarification
agent: Гримм
checklist: все пункты закрыты | открытые: {{список}}
---

## Результат

{{если нет проблем: "Acceptance чистый. Готово к test-write."}}

## Замечания

### AR-1 [пробел в acceptance | неопределённость в spec | противоречие в spec | проблема качества]

{{сценарий или раздел spec}}

{{объяснение проблемы}}

## Решения принятые в процессе

- ...
```

Дозапиши в `status-log.md`:
```
# {{datetime}} — done | failed
→ report-NNN.md
```

Выведи итог: статус, количество замечаний по типам. Если есть проблемы типа "неопределённость/противоречие в spec" — отметь явно что нужен откат к `01-spec/03-fix`.

