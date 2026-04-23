_Исполнители: агенты **Кроули** (`Кроули`) и **Азирафаль** (`Азирафаль`) — параллельно. Кроули пишет adversarial тесты, Азирафаль — позитивные._

Напиши тесты по acceptance.md.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

## Шаги

1. Прочитай `tasks/$ARGUMENTS/acceptance.md` и `tasks/$ARGUMENTS/status.md`
2. Если `acceptance-review: done` нет в `status.md` — сообщи что сначала нужен `/acceptance-review` и остановись
3. Прочитай `CLAUDE.md` — определи стек тестирования:
   - Go: `testing` + `testify`, интеграционные тесты — `testcontainers-go`
   - Frontend: `vitest` для юнитов, `playwright` (chromium) для e2e, `storybook` для компонентов
4. Для каждого сценария из `acceptance.md` напиши тест:
   - один сценарий = один тест
   - имя теста отражает сценарий (`TestQueryEntityByID`, `test('query entity by id')`)
   - тест проверяет именно то что описано в Then — не больше, не меньше
5. Покажи черновик пользователю, жди подтверждения или правок
6. После подтверждения — запиши тесты в соответствующие файлы
7. Если это переработка после `test-review: failed` — используй `/fix-test $ARGUMENTS` вместо этого шага
8. Обнови `tasks/$ARGUMENTS/status.md` — этап `test-write: done`
9. **Каскадный сброс:** если в `status.md` уже было `test-review: done` — выставь `test-review: pending`, а все последующие шаги (`code-write`, `code-review`, `build`, `test-run`) переведи в `needs-recheck`

## Важно

Не писать тесты которых нет в `acceptance.md`.
Если хочется покрыть что-то ещё — сначала добавить сценарий в acceptance через `/acceptance-write`.
