_Исполнитель: агент **Танк** (`Танк`)._

Исправь acceptance criteria по замечаниям ревью.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

## Шаги

1. Найди последний `tasks/$ARGUMENTS/acceptance-review-N.md` (наибольший N) — это замечания текущей итерации
2. Прочитай `tasks/$ARGUMENTS/acceptance.md` и `tasks/$ARGUMENTS/spec.md`
3. Определи M = количество существующих файлов `tasks/$ARGUMENTS/acceptance-fix-*.md` + 1
4. Покажи пользователю план исправлений: что именно и как будет изменено по каждому замечанию
5. Жди подтверждения или правок
6. Внеси исправления в `tasks/$ARGUMENTS/acceptance.md`
7. Запиши в `tasks/$ARGUMENTS/acceptance-fix-M.md`:
   ```
   # Фикс acceptance — итерация M

   - [ID замечания]: что исправлено и как
   - [ID замечания]: что исправлено и как
   - [ID замечания]: не исправлялось — [причина]
   ```
8. Обнови `tasks/$ARGUMENTS/status.md`:
   - `acceptance-review: pending`
   - все последующие шаги (`test-write`, `test-review`, `code-write`, `code-review`, `build`, `test-run`) → `needs-recheck`

## Важно

Если замечание типа «Неопределённость в спеке» — не выдумывай поведение, остановись и уточни через `/spec $ARGUMENTS`.
