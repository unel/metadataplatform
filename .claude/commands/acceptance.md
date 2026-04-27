_Режим: **Харли Куин**. Выполняется оператором напрямую._

Полный цикл acceptance: write → review.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

## Шаги

1. Запусти `/acceptance-write $ARGUMENTS`
2. Прочитай `.claude/project.yaml`, найди ключ `acceptance-write → acceptance-review`:
   - если `confirm` — напиши "Сценарии написаны, продолжаем к ревью? (y/n)" и жди ответа. Если отказ — остановись.
   - если `auto` или ключ отсутствует — продолжай сразу
3. Запусти `/acceptance-review $ARGUMENTS`
4. Если `acceptance-review: failed` в `status.md`:
   - Прочитай `max_retries` для `acceptance-review` из `.claude/project.yaml`
   - Прочитай `acceptance-iterations` из `tasks/$ARGUMENTS/status.md` (если нет — считай 0)
   - Если `acceptance-iterations ≥ max_retries` — стоп, показать все провалы, ждать инструкций пользователя
   - Иначе: инкрементируй `acceptance-iterations` в `status.md`, вызови `/fix-acceptance $ARGUMENTS`, затем вернись к шагу 3

## Важно

Если `.claude/project.yaml` не существует — вести себя как `confirm` для всех переходов.
