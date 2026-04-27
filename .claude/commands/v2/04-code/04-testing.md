_Исполнитель: агент **Азирафаль** (`Азирафаль`)._

# code-testing

Финальный прогон тестов после code-review. TDD Green: все тесты должны пройти.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/04-code/04-testing/README.md` в `tasks/$ARGUMENTS/processes/04-code/04-testing/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/04-code/04-testing/README.md`.

## Исполнение

Прочитай `docs/standards/v2/04-code/04-testing/base-plan.md` и `base-checklist.md`.

Проверь что code-review прошёл: последний report из `tasks/$ARGUMENTS/processes/04-code/02-review/` или `04-code/03-fix/` со статусом `done`.

Прочитай `PROJECT.md` — команда для запуска тестов.

Запусти полный suite:
- `tests/happy/` и `tests/adversarial/`
- С детальным выводом (-v или аналог)
- С race detector если поддерживается (-race в Go)

Ожидаемый результат: **все зелёные**. Если тест упал — определи источник: код / тест / acceptance. Не меняй тесты чтобы они прошли. Не угадывай при неясном acceptance — эскалируй.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/04-code/04-testing/report-NNN.md`:

```markdown
---
purpose: Финальный прогон тестов для $ARGUMENTS
process: 04-code/04-testing
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | failed | clarification
agent: Азирафаль
checklist: все пункты закрыты | открытые: {{список}}
---

## Итог

Прошло: N | Упало: N | Пропущено: N

## Упавшие тесты

| Тест | Источник (код / тест / acceptance) | Что делать |
|---|---|---|
| ... | ... | ... |

## Полный вывод

{{вывод тест-раннера или ключевые строки}}
```

Обнови `tasks/$ARGUMENTS/processes/04-code/04-testing/README.md`: добавь `updated: {{datetime}}`.
