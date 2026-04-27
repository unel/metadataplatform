_Режим: **Харли Куин**. Выполняется оператором напрямую._

Полный цикл спецификации для фичи: ft → nft → review.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

## Шаги

1. Запусти `/spec-ft $ARGUMENTS`
2. Прочитай `.claude/project.yaml`, найди секцию `transitions`, ключ `spec-ft → spec-nft`:
   - если значение `confirm` — напиши "ФТ готовы, продолжаем к НФТ? (y/n)" и жди ответа. Если отказ — остановись.
   - если значение `auto` или ключ отсутствует — продолжай сразу
3. Запусти `/spec-nft $ARGUMENTS`
4. Прочитай `.claude/project.yaml`, найди ключ `spec-nft → spec-review`:
   - если `confirm` — напиши "НФТ готовы, продолжаем к ревью? (y/n)" и жди ответа. Если отказ — остановись.
   - если `auto` или ключ отсутствует — продолжай сразу
5. Запусти `/spec-review $ARGUMENTS`
6. Если `spec-review: failed` в `status.md`:
   - Прочитай `max_retries` для `spec-review` из `.claude/project.yaml`
   - Прочитай `spec-iterations` из `tasks/$ARGUMENTS/status.md` (если нет — считай 0)
   - Если `spec-iterations ≥ max_retries` — стоп, показать все провалы, ждать инструкций пользователя
   - Иначе: инкрементируй `spec-iterations` в `status.md`, вызови `/fix-spec $ARGUMENTS`, затем вернись к шагу 5
7. Обнови `tasks/$ARGUMENTS/status.md` — этап `spec: done`

## Важно

Если `.claude/project.yaml` не существует — вести себя как `confirm` для всех переходов (безопасный дефолт).
