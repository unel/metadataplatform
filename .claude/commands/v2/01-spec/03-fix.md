_Исполнитель: агент **Танк** (`Танк`)._

Процесс `01-spec / 03-fix`. Исправь спеку по замечаниям ревью.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## Шаг 1. Инициализация

Создай директорию `tasks/$ARGUMENTS/processes/01-spec/03-fix/` если не существует.

Скопируй `README.md` из `docs/standards/v2/01-spec/03-fix/README.md` в `tasks/$ARGUMENTS/processes/01-spec/03-fix/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/01-spec/03-fix/README.md`.

Запиши в `status-log.md`:
```
# {{datetime}} — pending
```

---

## Шаг 2. Генерация плана

Прочитай:
- `docs/standards/v2/01-spec/03-fix/base-plan.md` — базовый план
- актуальную spec: сначала последний `tasks/$ARGUMENTS/processes/01-spec/03-fix/report-*.md`, если нет — последний `tasks/$ARGUMENTS/processes/01-spec/01-write/report-*.md`
- замечания: последний `tasks/$ARGUMENTS/processes/01-spec/02-review/report-*.md`
- `PROJECT.md` если замечания архитектурные

Сгенерируй `tasks/$ARGUMENTS/processes/01-spec/03-fix/plan.md`:
- Скопируй базовый план
- Адаптируй: перечисли конкретные CR-N из замечаний, для каждого — что меняется
- Укажи в мете версию базового плана

```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-plan: 01-spec-fix v{{version}}
---

{{адаптированный план}}
```

---

## Шаг 3. Генерация чек-листа

Прочитай:
- `docs/standards/v2/01-spec/03-fix/base-checklist.md` — базовый чек-лист
- нужные расширения из `docs/standards/v2/extensions/`
- сгенерированный `plan.md`

Запиши `tasks/$ARGUMENTS/processes/01-spec/03-fix/checklist.md`:
```markdown
---
generated: {{datetime}}
created: {{datetime}}
updated: {{datetime}}
base-checklist: 01-spec-fix v{{version}}
extensions:
  - {{extension}} v{{version}}
plan: {{datetime плана}}
---

{{базовый чек-лист + пункты из плана + расширения}}
```

Обнови `status-log.md`:
```
# {{datetime}} — in-progress
→ plan.md, checklist.md
```

---

## Шаг 4. Согласование

Покажи пользователю план исправлений: по каждому CR-N — что именно и как будет изменено.

Жди подтверждения или правок. Если замечание требует архитектурного решения — остановись:

```
# {{datetime}} — clarification
CR-N требует архитектурного решения: {{описание}}. Ожидаю решения.
```

---

## Шаг 5. Выполнение

Исправляй точечно — только то что указано в замечаниях. Не переписывай spec целиком.

Отмечай пункты в `checklist.md` по мере выполнения.

---

## Шаг 6. Завершение

Подсчитай N: количество существующих `report-*.md` в этой директории + 1.

Запиши `report-NNN.md`:

```markdown
---
process: 01-spec/03-fix
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:  # если информация из предыдущих прогонов актуальна
status: done | clarification
agent: Танк
checklist: все пункты закрыты | открытые: {{список}}
---

## Spec (обновлённая)

{{полный текст обновлённой спецификации}}

## Changelog

- CR-1: {{что изменено}}
- CR-2: {{что изменено}}
- CR-3: не изменялось — {{причина}}

## Решения принятые в процессе

- ...

## Открытые вопросы

- ...
```

Обнови статусы последующих процессов:
- `01-spec/02-review` → `pending`
- `02-acceptance`, `03-tests`, `04-code`, `05-docs` → `stale` (если уже были выполнены)

Дозапиши в `status-log.md`:
```
# {{datetime}} — done
→ report-NNN.md
```

