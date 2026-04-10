_Исполнитель: агент **Харли** (`harley`) — оркестрирует **Танка** и **Гримма**._

Полный цикл acceptance: write → review.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

## Шаги

1. Запусти `/acceptance-write $ARGUMENTS`
2. Прочитай `.claude/project.yaml`, найди ключ `acceptance-write → acceptance-review`:
   - если `confirm` — напиши "Сценарии написаны, продолжаем к ревью? (y/n)" и жди ответа. Если отказ — остановись.
   - если `auto` или ключ отсутствует — продолжай сразу
3. Запусти `/acceptance-review $ARGUMENTS`
4. Если `acceptance-review: failed` в `status.md` — остановись, оркестратор решит что делать дальше

## Важно

Если `.claude/project.yaml` не существует — вести себя как `confirm` для всех переходов.
