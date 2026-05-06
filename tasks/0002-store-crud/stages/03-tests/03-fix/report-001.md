---
purpose: Фикс тестов для store/crud
process: 03-tests/03-fix
run: 1
date: 2026-04-30T06:08:33Z
created: 2026-04-30T06:08:33Z
status: done
agent: Кроули + Азирафаль
checklist: все пункты закрыты
---

## Исправленные замечания

| Замечание | Что сделано |
|---|---|
| critical: `newRouterWithLogger` не определена | Де-факто закрыто до фикса — функция уже определена в `happy/helpers_logger_test.go:52` |
| critical: `r.Handle(serverConn)` без ctx в adversarial | `router_helpers_test.go:52` → `r.Handle(context.Background(), serverConn)`, добавлен импорт `"context"` |
| critical: сигнатура Handle в happy-файлах | Happy-файлы уже используют `Handle(ctx, conn)` — правок не потребовалось |
| medium: AC-ROUTER-16 без прямого теста | `router_error_test.go`: добавлен `TestRouter_GetNonexistentEntity_LogsError` — get несуществующего entity → NOT_FOUND + ERROR-лог с полями op/type/id/error |
| minor: `atomic_test.go` 30ms hardcode | Заменён на polling по файлу-маркеру `.ready` (timeout 5s, шаг 5ms); subprocess пишет маркер после `Sync()` |
| minor: `fs_init_test.go` несоответствие logger-интерфейса | Убраны `testLogger`/`logEntry`; `TestFSInit_SuccessfulInit_LogsInfoWithBasedir` переписан на `slog.New(slog.NewTextHandler(&buf, ...))` — тот же паттерн что в adversarial |

## Неисправленные замечания

Нет.
