_Исполнитель: агент **Танк** (`tank`)._

Составь тестовые сценарии по спеке и запиши в acceptance.md.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

## Шаги

1. Прочитай `tasks/$ARGUMENTS/spec.md` и `tasks/$ARGUMENTS/status.md`
2. Если `spec-review: done` нет в `status.md` — сообщи что сначала нужен `/spec-review` и остановись
3. Составь сценарии в формате Given/When/Then для:
   - каждого happy path из ФТ
   - каждого граничного условия
   - каждого corner-case упомянутого в spec-review
   - ключевых НФТ которые можно проверить тестом (производительность, параллелизм и т.п.)
4. Покажи черновик пользователю, жди подтверждения или правок
5. После подтверждения — запиши в `tasks/$ARGUMENTS/acceptance.md`
6. Если это переработка после `acceptance-review: failed` — используй `/fix-acceptance $ARGUMENTS` вместо этого шага
7. Обнови `tasks/$ARGUMENTS/status.md` — этап `acceptance-write: done`
8. **Каскадный сброс:** если в `status.md` уже было `acceptance-review: done` — выставь `acceptance-review: pending`, а все последующие шаги (`test-write`, `test-review`, `code-write`, `code-review`, `build`, `test-run`) переведи в `needs-recheck`

## Формат acceptance.md

```markdown
# Acceptance — store/query

## Сценарий: запрос entity по id
**Given** entity с id "abc" существует в БД
**When** отправляем {"op":"query","type":"entity","filter":{"id":"abc"}}
**Then** получаем {"ok":true,"data":{...}} с данными этой entity

## Сценарий: запрос несуществующей entity
**Given** entity с id "xyz" не существует
**When** отправляем {"op":"query","type":"entity","filter":{"id":"xyz"}}
**Then** получаем {"ok":true,"data":[]}
```

## Важно

Сценарии должны быть конкретными — никаких расплывчатых "Then система работает корректно".
