_Исполнитель: агент **Ада Лавлейс** (`Ада Лавлейс`)._

# code-write

Написание реализации. Цель: заставить тесты пройти (TDD Green).

## Инициализация

Скопируй `README.md` из `docs/standards/v2/04-code/01-write/README.md` в `tasks/$ARGUMENTS/processes/04-code/01-write/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/04-code/01-write/README.md`.

## Исполнение

Прочитай `docs/standards/v2/04-code/01-write/base-plan.md` и `base-checklist.md`.

Входящие артефакты:
- Тесты: `tasks/$ARGUMENTS/tests/happy/` и `tasks/$ARGUMENTS/tests/adversarial/` — определяют контракт
- Финальный acceptance: последний report из `tasks/$ARGUMENTS/processes/02-acceptance/`
- Spec: последний report из `tasks/$ARGUMENTS/processes/01-spec/`
- `PROJECT.md` — стек, соглашения по структуре кода

Пиши минимально необходимый код чтобы тесты прошли. Модули ≤ 150 строк. Безопасность по OWASP.

После написания — запусти тесты. Ожидаемый результат: все зелёные.
Если тест требует поведения вне acceptance — `clarification`, не угадывай.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/04-code/01-write/report-NNN.md`:

```markdown
---
purpose: Написание кода для $ARGUMENTS
process: 04-code/01-write
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | failed | clarification
agent: Ада Лавлейс
checklist: все пункты закрыты | открытые: {{список}}
---

## Созданные / изменённые файлы

{{список}}

## Результаты прогона тестов

Прошло: N | Упало: N

## Примечания

{{если что-то пришлось решать без явного указания в spec/acceptance}}
```

Обнови `tasks/$ARGUMENTS/processes/04-code/01-write/README.md`: добавь `updated: {{datetime}}`.
