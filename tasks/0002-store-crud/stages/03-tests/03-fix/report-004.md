---
purpose: Фикс тестов для store/crud (run 4)
process: 03-tests/03-fix
run: 4
date: 2026-05-04T07:30:54Z
created: 2026-05-04T07:30:54Z
see-also: tasks/0002-store-crud/stages/03-tests/02-review/report-004.md
status: done
agent: Азирафаль
---

## Исправленные замечания

| Замечание | Что сделано |
|---|---|
| router_logging_test.go s.Upsert → s.Entities().Upsert | Заменено в двух местах: строка 45 (TestRouter_Get_Entity_LogsDebugWithOpTypeID) и строки 68–69 (TestRouter_List_Entities_LogsDebugWithCount). `s` имеет тип `*fs.Store` и не реализует EntityStore напрямую — нужна обёртка `.Entities()`. |
| Дублирующее имя теста TestEntity_Get_EmptyID_ReturnsError | Удалён из happy/entity_get_test.go. Комментарий в заголовке файла обновлён: AC-ENTITY-GET-03 теперь явно отсылает к adversarial-пакету. Неиспользуемые импорты (`errors`, `store`) убраны. |
| Маппинг AC-ENTITY-LIST-05 | В tasks/0002-store-crud/stages/03-tests/01-write/report-001.md добавлен TestEntity_List_IgnoresNonJsonFiles рядом с TestEntity_List_TmpFilesPresent_IgnoresTmpFiles для AC-ENTITY-LIST-05. |

## Результат

`go test ./store/tests/happy/...` — ok.
