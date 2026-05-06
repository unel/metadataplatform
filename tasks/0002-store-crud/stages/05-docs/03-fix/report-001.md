---
purpose: Фикс документации для store/crud
process: 05-docs/03-fix
run: 1
date: 2026-05-05T06:12:17Z
created: 2026-05-05T06:12:17Z
see-also: tasks/0002-store-crud/stages/05-docs/02-review/report-001.md
status: done
agent: Танк
checklist: все пункты закрыты
---

## Исправленные замечания

| Замечание | Что сделано |
|---|---|
| `store/router/README.md` — READ_ERROR при upsert (warning) | Добавлено примечание к READ_ERROR: "При upsert означает что существующая запись нечитаема — операция отменена, новые данные не записаны." |
| `store/README.md` — "default {}" в полях структур (warning) | Комментарии `default {}` заменены на точное описание: "nil сериализуется как null, не {}" — для Meta (Entity), Value/Meta (Relation), Payload (Job) |
| `store/README.md` — пример Upsert не компилируется (Nit) | Исправлено: `json.RawMessage(...)` → `json.RawMessage([]byte(...))` |
| `store/router/README.md` — MISSING_ID при отсутствующем id (Nit) | Описание расширено: "...или не передан в запросе get/delete" |
| `store/router/README.md` — примеры ответов без реальных полей (Nit) | Добавлен пример полного ответа get entity со всеми полями |
| `store/fs/README.md` — логирование через роутер (Nit) | Раздел переформулирован: store/fs логирует только инициализацию; операции логируются в store/router |
| `store/fs/TECH.md` — эволюционное последствие Entity receiver (Nit) | Добавлено следствие: при добавлении нового типа хранилища Entity-методы потребуется перенести в обёртку |

## Неисправленные

Нет.

## Примечание для ретро

Некомпилируемый пример кода (`json.RawMessage(...)` без `[]byte`) классифицирован Гриммом как Nit — это занижение. Такой пример вводит разработчика в заблуждение и вызывает ошибку компилятора. Требует пересмотра критериев классификации в ревью документации.
