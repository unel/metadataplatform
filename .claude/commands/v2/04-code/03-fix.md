_Исполнитель: агент **Ада Лавлейс** (`Ада Лавлейс`)._

# code-fix

Исправление замечаний из code-review. Тесты должны оставаться зелёными.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/04-code/03-fix/README.md` в `tasks/$ARGUMENTS/processes/04-code/03-fix/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/04-code/03-fix/README.md`.

## Исполнение

Прочитай `docs/standards/v2/04-code/03-fix/base-plan.md` и `base-checklist.md`.

Входящие артефакты:
- Последний `tasks/$ARGUMENTS/processes/04-code/02-review/report-*.md` — замечания
- Код реализации

Покажи план фикса пользователю. Жди подтверждения.

Исправляй только то что указано в review. Сначала critical, потом warning, потом note. Держи модули ≤ 150 строк. После фикса — запусти тесты, ожидаемый результат: всё зелёное.

Если замечание требует изменения acceptance — запиши `clarification`, не угадывай.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/04-code/03-fix/report-NNN.md`:

```markdown
---
purpose: Фикс кода для $ARGUMENTS
process: 04-code/03-fix
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | clarification
agent: Ада Лавлейс
checklist: все пункты закрыты | открытые: {{список}}
---

## Исправленные замечания

| Замечание | Файл | Что сделано |
|---|---|---|
| ... | ... | ... |

## Неисправленные замечания

| Замечание | Причина | Что нужно |
|---|---|---|
| ... | ... | ... |

## Тесты после фикса

Прошло: N | Упало: N
```

Обнови `tasks/$ARGUMENTS/processes/04-code/03-fix/README.md`: добавь `updated: {{datetime}}`.
