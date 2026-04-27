_Исполнитель: агент **Гримм** (`Гримм`)._

Процесс `01-spec / 02-review`. Проверь спеку на противоречия, сложности и corner-cases.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## Шаг 1. Инициализация

Создай директорию `tasks/$ARGUMENTS/processes/01-spec/02-review/` если не существует.

Скопируй `README.md` из `docs/standards/v2/01-spec/02-review/README.md` в `tasks/$ARGUMENTS/processes/01-spec/02-review/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/01-spec/02-review/README.md`.

Запиши в `status-log.md`:
```
# {{datetime}} — pending
```

---

## Шаг 2. Генерация плана

Прочитай:
- `docs/standards/v2/01-spec/02-review/base-plan.md` — базовый план
- последний `tasks/$ARGUMENTS/processes/01-spec/01-write/report-*.md` — spec
- `PROJECT.md` — архитектура, ограничения

Сгенерируй `tasks/$ARGUMENTS/processes/01-spec/02-review/plan.md`:
- Скопируй базовый план
- Адаптируй под конкретную задачу: убери нерелевантные шаги, добавь специфичные
- Укажи в мете версию базового плана

```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-plan: 01-spec-review v{{version}}
---

{{адаптированный план}}
```

---

## Шаг 3. Генерация чек-листа

Прочитай:
- `docs/standards/v2/01-spec/02-review/base-checklist.md` — базовый чек-лист
- нужные расширения из `docs/standards/v2/extensions/` — определи по контексту задачи
- сгенерированный `plan.md` — добавь чек-пункты отражающие шаги плана

Запиши `tasks/$ARGUMENTS/processes/01-spec/02-review/checklist.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-checklist: 01-spec-review v{{version}}
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

При неоднозначности — не выбирай молча. Опиши контекст, варианты с плюсами/минусами, жди решения.

Если spec неполна или противоречива настолько, что ревью невозможно — запиши в `status-log.md`:
```
# {{datetime}} — clarification
Неясно: {{что именно}}. Ожидаю уточнения.
```
И сообщи пользователю.

---

## Шаг 5. Завершение

Подсчитай N: количество существующих `report-*.md` в этой директории + 1.

Запиши `report-NNN.md`:

```markdown
---
process: 01-spec/02-review
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:  # если информация из предыдущих прогонов актуальна
status: done | failed | clarification
agent: Гримм
checklist: все пункты закрыты | открытые: {{список}}
---

## Результат

{{если нет проблем: "Спека чистая. Готово к acceptance."}}
{{если есть проблемы: список замечаний}}

## Замечания

### CR-1 [critical | warning | note]

{{цитата или ссылка на пункт spec}}

{{объяснение проблемы}}

### CR-2 ...

## Решения принятые в процессе

- ...

## Открытые вопросы

- ...
```

Дозапиши в `status-log.md`:
```
# {{datetime}} — done | failed
→ report-NNN.md
```

Выведи пользователю итог: статус, количество замечаний по категориям.

