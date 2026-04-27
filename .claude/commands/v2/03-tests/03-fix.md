_Исполнитель: агент **Кроули** (`Кроули`) для adversarial, **Азирафаль** (`Азирафаль`) для happy path._

# test-fix

Исправление замечаний из test-review. Каждый агент чинит свои тесты.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/03-tests/03-fix/README.md` в `tasks/$ARGUMENTS/processes/03-tests/03-fix/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/03-tests/03-fix/README.md`.

## Исполнение

Прочитай `docs/standards/v2/03-tests/03-fix/base-plan.md` и `base-checklist.md`.

Входящие артефакты:
- Последний `tasks/$ARGUMENTS/processes/03-tests/02-review/report-*.md` — замечания
- Тест-файлы из `tasks/$ARGUMENTS/tests/`

Покажи план фикса пользователю. Жди подтверждения.

Исправляй только то что указано в review. Сохраняй структуру `tests/happy/` и `tests/adversarial/`. Держи каждый модуль ≤ 150 строк.

Если замечание требует изменения acceptance — запиши `clarification`, не угадывай.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/03-tests/03-fix/report-NNN.md`:

```markdown
---
purpose: Фикс тестов для $ARGUMENTS
process: 03-tests/03-fix
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | clarification
agent: Кроули + Азирафаль
checklist: все пункты закрыты | открытые: {{список}}
---

## Исправленные замечания

| Замечание | Что сделано |
|---|---|
| ... | ... |

## Неисправленные замечания

| Замечание | Причина | Что нужно |
|---|---|---|
| ... | ... | ... |
```

Обнови `tasks/$ARGUMENTS/processes/03-tests/03-fix/README.md`: добавь `updated: {{datetime}}`.
