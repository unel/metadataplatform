---
purpose: Ревью тестов для store/crud
process: 03-tests/02-review
run: 2
date: 2026-04-30T06:30:00Z
created: 2026-04-30T06:30:00Z
see-also: tasks/0002-store-crud/stages/03-tests/03-fix/report-001.md
status: failed
agent: Гримм
checklist: открытые: несовместимый тип логгера между happy и adversarial пакетами
---

## Результат

**failed**. Три critical из run-1 закрыты корректно: `newRouterWithLogger` определён, сигнатура `Handle(ctx, conn)` унифицирована, AC-ROUTER-16 покрыт прямым тестом `TestRouter_GetNonexistentEntity_LogsError`. Два minor закрыты: atomic-тест переведён на маркер готовности, `fs_init_test.go` переписан на slog+bytes.Buffer.

При фиксе добавился новый critical — несовместимость типов логгера.

## Замечания

### happy/helpers_logger_test.go + adversarial/router_helpers_test.go — несовместимый тип логгера в router.New

**Категория:** нарушение FIRST (Isolated), структура
**Severity:** critical
**Проблема:** `newRouterWithLogger` в `happy/helpers_logger_test.go:57` вызывает `router.New(s, s, s, log)` где `log` — `*captureLogger` с методом `LogOp(level, op, typ, id string, count int, errMsg string)`. В `adversarial/router_helpers_test.go:45` тот же `router.New` вызывается с `logger *slog.Logger`. Два типа несовместимы — один из пакетов гарантированно не скомпилируется.
**Рекомендация:** Унифицировать на `*slog.Logger`. В happy убрать `captureLogger`/`routerLogEntry`/`findEntry`, заменить на `slog.New(slog.NewTextHandler(&buf, ...))` + `assert.Contains`. `newRouterWithLogger` становится не нужен.

---

### happy/router_logging_test.go — комментарий AC-ROUTER-16 не соответствует содержимому

**Категория:** misleading coverage mapping
**Severity:** minor
**Проблема:** Заголовок файла `// AC-ROUTER-14, AC-ROUTER-15, AC-ROUTER-16`. Третий тест `TestRouter_List_Entities_LogsDebugWithCount` покрывает AC-ROUTER-15 (DEBUG с count). AC-ROUTER-16 закрыт в adversarial. Комментарий создаёт ложную трассировку.
**Рекомендация:** Убрать `AC-ROUTER-16` из заголовка, добавить `// AC-ROUTER-16: см. adversarial/router_error_test.go`.

---

### adversarial/entity_upsert_error_test.go — TestEntity_Upsert_WriteError_NoTempFileLeft тихо пропускает проверку

**Категория:** нарушение FIRST (Self-validating)
**Severity:** minor
**Проблема:** Строки 85—87: `if err != nil { return }` после `os.ReadDir(entitiesDir)`. Если chmod не сработал (root) или директория нечитаема по другой причине — тест возвращается без единого assert'а и репортится как passed.
**Рекомендация:** Заменить `return` на `t.Skip("entitiesDir unreadable after chmod — running as root; skipping tmp-file check")`.
