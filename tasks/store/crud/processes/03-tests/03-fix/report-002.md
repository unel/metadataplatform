---
purpose: Фикс тестов для store/crud
process: 03-tests/03-fix
run: 2
date: 2026-04-30T06:50:00Z
created: 2026-04-30T06:50:00Z
status: done
agent: Кроули + Азирафаль
checklist: все пункты закрыты
---

## Исправленные замечания

| Замечание | Что сделано |
|---|---|
| critical: captureLogger несовместим с *slog.Logger | helpers_logger_test.go: captureLogger/findEntry убраны, добавлен newRouterWithSlogBuf; router_conn_test.go + router_happy_test.go: newRouterWithLogger → newRouterWithSlogBuf |
| minor: AC-ROUTER-16 в заголовке router_logging_test.go | заголовок исправлен: AC-ROUTER-16 убран, добавлена ссылка на adversarial |
| minor: entity_upsert_error_test.go return без assert | entity_upsert_error_test.go:88 return → t.Skip(...) |

## Неисправленные замечания

Нет.
