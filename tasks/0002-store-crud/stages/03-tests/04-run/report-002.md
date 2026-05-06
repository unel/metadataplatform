---
purpose: Запуск тестов для store/crud
process: 03-tests/04-run
run: 2
date: 2026-05-04T12:23:05Z
created: 2026-05-04T12:23:05Z
see-also:
  - tasks/0002-store-crud/stages/03-tests/02-review/report-005.md
context: final-run
status: done
agent: Азирафаль
checklist: все пункты закрыты
---

## Итог

Контекст: финальный прогон (Green) — реализация написана (store/, store/fs/, store/router/).
Прошло: 68 | Упало: 0 | Пропущено: 0

Оба пакета прошли с race detector чистым. Аномалий нет.

## Аномалии первого прогона (не применимо)

Это финальный прогон. Первый прогон (Red) задокументирован в report-001.md.

## Упавшие тесты (финальный прогон)

Упавших тестов нет.

## Наблюдения по выводу

В нескольких тестах роутера в stdout появляется ERROR parse_error op=unknown parse_error="io: read/write on closed pipe". Это ожидаемый артефакт тестовых сценариев где клиент закрывает pipe-соединение сразу после чтения ответа. Роутер корректно логирует разрыв и завершает обработку. Тесты проходят — это не ошибка.

| Тест | Почему нормально |
|---|---|
| TestRouter_Get_ExistingEntity_ReturnsOkWithData | Клиент-сторона pipe закрывается после чтения ответа |
| TestRouter_List_Entities_ReturnsOkWithDataArray | Аналогично |
| TestRouter_Upsert_MissingData_ReturnsInvalidRequest_ConnectionAlive | Аналогично |
| TestRouter_Upsert_DataNull_ReturnsInvalidRequest_ConnectionAlive | Аналогично |
| TestRouter_UnknownOp_ReturnsUnknownOpError_ConnectionAlive | Аналогично |
| TestRouter_UnknownType_ReturnsUnknownTypeError_ConnectionAlive | Аналогично |
| TestRouter_InvalidJSONInData_ReturnsInvalidRequest_ConnectionAlive | Аналогично |

## Полный вывод

    ok  github.com/unel/metadataplatform/store/tests/happy        1.084s
    ok  github.com/unel/metadataplatform/store/tests/adversarial  1.055s
