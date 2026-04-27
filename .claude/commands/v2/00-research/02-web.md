_Исполнитель: агент **Бо** (`Бо`)._

Процесс `00-research / 02-web`. Исследуй похожие решения, паттерны и известные failure modes.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## Шаг 1. Инициализация

Создай директорию `tasks/$ARGUMENTS/processes/00-research/02-web/` если не существует.

Скопируй `README.md` из `docs/standards/v2/00-research/02-web/README.md` в `tasks/$ARGUMENTS/processes/00-research/02-web/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/00-research/02-web/README.md`.

Запиши в `status-log.md`:
```
---
purpose: Хронология статусов процесса web-research — оркестратор читает для понимания текущего состояния
---
created: {{datetime}}
updated: {{datetime}}

# {{datetime}} — pending
```

---

## Шаг 2. Генерация плана

Прочитай:
- `docs/standards/v2/00-research/02-web/base-plan.md`
- последний `tasks/$ARGUMENTS/processes/00-research/01-interview/report-*.md`
- `tasks/$ARGUMENTS/TASK.md` и `PROJECT.md`

Сгенерируй `tasks/$ARGUMENTS/processes/00-research/02-web/plan.md`:
```markdown
---
purpose: План процесса web-research — исполнитель читает перед стартом
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-plan: 00-research-web v{{version}}
---

{{адаптированный план: конкретные направления поиска на основе interview report}}
```

---

## Шаг 3. Генерация чек-листа

Прочитай:
- `docs/standards/v2/00-research/02-web/base-checklist.md`
- расширения из `docs/standards/v2/extensions/` по контексту
- сгенерированный `plan.md`

Запиши `tasks/$ARGUMENTS/processes/00-research/02-web/checklist.md`:
```markdown
---
purpose: Чек-лист прогона web-research — исполнитель отмечает по ходу выполнения
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-checklist: 00-research-web v{{version}}
extensions:
  - {{extension}} v{{version}}
plan: {{datetime плана}}
---
```

Обнови `status-log.md`:
```
# {{datetime}} — in-progress
→ plan.md, checklist.md
```

---

## Шаг 4. Выполнение

Ищи по направлениям из плана. Для каждой находки — фиксируй URL источника сразу.

Если запросов много — веди `searches.md` параллельно:
```markdown
---
purpose: Поисковые запросы использованные в research — для воспроизведения поиска
---

- {{запрос}} → {{что нашёл / не нашёл}}
```

---

## Шаг 5. Завершение

Подсчитай N, запиши `report-NNN.md`:

```markdown
---
purpose: Результаты web-research с источниками — входящий артефакт для 01-spec/01-write
process: 00-research/02-web
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:  # если информация из предыдущих прогонов актуальна
status: done
agent: Бо
checklist: все пункты закрыты | открытые: {{список}}
---

## Похожие реализации

- **{{название/проект}}** — {{описание подхода}} ([источник](URL))
- ...

## Известные failure modes

- **{{проблема}}** — {{как проявляется, почему возникает}} ([источник](URL))
- ...

## Паттерны и best practices

- **{{паттерн}}** — {{обоснование}} ([источник](URL))
- ...

## Выводы для spec

{{что конкретно нужно учесть при написании требований}}

## Выводы для реализации

{{что конкретно нужно учесть при написании кода}}

## Источники

{{если sources.md не создавался — список URL с кратким описанием здесь}}
```

Если источников много — создай `sources.md`:
```markdown
---
purpose: URL источников с кратким описанием — для перепроверки фактов из report
---

- [{{название}}](URL) — {{одна строка: почему взят за основу}}
```

Дозапиши в `status-log.md`:
```
# {{datetime}} — done
→ report-NNN.md
```

