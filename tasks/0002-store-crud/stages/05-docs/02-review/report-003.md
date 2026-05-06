---
purpose: Ревью документации для store/crud
process: 05-docs/02-review
run: 3
date: 2026-05-05T06:26:51Z
created: 2026-05-05T06:26:51Z
see-also: tasks/0002-store-crud/stages/05-docs/03-fix/report-002.md
status: done
agent: Гримм
checklist: все пункты закрыты
---

## Результат

Документация чистая.

## Проверено

**store/README.md** — Примеры корректны: сигнатуры `Get`, `Upsert` соответствуют `interfaces.go`. Sentinel errors — все семь присутствуют в `types.go`. Поведенческие гарантии (timestamps управляются стором, идемпотентность Upsert, Delete не идемпотентен, ErrMissingID на пустом ID) — соответствуют `store/fs/store.go` и `store/fs/validate.go`.

**store/fs/README.md** — Поведение всех четырёх операций соответствует коду. `nil` в `New` → `slog.Default()` — корректно (проверено в `store.go:38-40`). List возвращает пустой срез — подтверждается явной проверкой `if results == nil { results = []store.Entity{} }`. Логирование операций в `store/router` — корректно.

**store/fs/TECH.md** — ADR точные. Атомарная запись, сохранение `created_at`, path traversal protection, Entity-обёртки — все описания соответствуют `io.go`, `upsert.go`, `validate.go`, `store.go`. Утверждение про `s.Entities()` возвращающий сам `s` — подтверждено кодом.

**store/router/README.md** — Формат запроса/ответа соответствует `router.go` и `handle_*.go`. Коды ошибок — полный маппинг через `mapStoreError` совпадает с таблицей. Поведение при parse error (`conn.Close(); return`) — корректно. Примечание об omitempty для `name`, `description`, `subtype` — соответствует JSON-тегам в `types.go`. Исправление из run 2 принято корректно: убраны лишние поля, добавлено примечание.

## Замечания

### `store/router/README.md` — поля без omitempty не упомянуты рядом с примечанием об omitempty

**Категория:** Nit:
**Проблема:** Поле `meta` у Entity (и `value`, `meta` у Relation, `payload` у Job) не имеет `omitempty` и всегда присутствует в JSON-ответе, даже как `null`. Примечание об omitempty описывает только `name`/`description`/`subtype`, но умалчивает про поля без omitempty. Читатель может решить что nil meta тоже пропускается.
**Рекомендация:** Добавить рядом примечание: "Поля `meta`, `value`, `payload` всегда присутствуют в ответе, при nil — как `null`." Не блокирует.
