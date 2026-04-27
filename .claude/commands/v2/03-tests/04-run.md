_Исполнитель: агент **Азирафаль** (`Азирафаль`)._

# test-run

Запуск полного тест-suite. Интерпретация результатов. Документирование падений.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/03-tests/04-run/README.md` в `tasks/$ARGUMENTS/processes/03-tests/04-run/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/03-tests/04-run/README.md`.

## Исполнение

Прочитай `docs/standards/v2/03-tests/04-run/base-plan.md` и `base-checklist.md`.

Проверь что review прошёл: последний `tasks/$ARGUMENTS/processes/03-tests/02-review/report-*.md` со статусом `done`.

Прочитай `PROJECT.md` — команда для запуска тестов.

### Определи контекст запуска

**Первый прогон (до реализации — TDD Red):**
Проверь: есть ли уже код реализации (директория, файлы по спеке)?
- Если нет — это первый прогон. Запусти тесты. Ожидаемый результат: **все тесты падают**.
- Если тест **прошёл** без реализации — красный флаг: тест ничего не проверяет (всегда true, неправильный assert, тестирует мок). Это `failed` → откат к `03-tests/03-fix`.
- Если все красные → `done`, передаём в `04-code/01-write`.

**Финальный прогон (после реализации — TDD Green):**
Запусти тесты:
- Полный suite: `tests/happy/` и `tests/adversarial/`
- С детальным выводом (-v или аналог)
- С race detector если поддерживается (-race в Go)

Для каждого падения — определи причину: ошибка в тесте / проблема в acceptance / ошибка в коде. Не угадывай правильное поведение если acceptance неясен — эскалируй.

## Отчёт

Запиши `tasks/$ARGUMENTS/processes/03-tests/04-run/report-NNN.md`:

```markdown
---
purpose: Запуск тестов для $ARGUMENTS
process: 03-tests/04-run
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
context: first-run | final-run
status: done | failed | clarification
agent: Азирафаль
checklist: все пункты закрыты | открытые: {{список}}
---

## Итог

Контекст: первый прогон (Red) | финальный прогон (Green)
Прошло: N | Упало: N | Пропущено: N

## Аномалии первого прогона (если применимо)

_Тесты прошедшие без реализации — красный флаг_

| Тест | Почему подозрительно |
|---|---|
| ... | ... |

## Упавшие тесты (финальный прогон)

| Тест | Сценарий acceptance | Причина | Что делать |
|---|---|---|---|
| ... | ... | ... | ... |

## Полный вывод

{{вывод тест-раннера или ключевые строки}}
```

Обнови `tasks/$ARGUMENTS/processes/03-tests/04-run/README.md`: добавь `updated: {{datetime}}`.
