---
purpose: Фикс кода для store/crud
process: 04-code/03-fix
run: 1
date: 2026-05-04T17:30:56Z
created: 2026-05-04T17:30:56Z
see-also: tasks/store/crud/processes/04-code/02-review/report-001.md
status: done
agent: Ада Лавлейс
checklist: все пункты закрыты
---

## Исправленные замечания

| Замечание | Файл(ы) | Что сделано |
|---|---|---|
| CRITICAL-1: path traversal через id | `store/fs/validate.go` | `recordPath` возвращает `(string, error)`; после `filepath.Join` проверяет что `filepath.Clean(path)` начинается с `filepath.Clean(basedir)+separator`; при нарушении — `ErrMissingID: invalid id`. `validateID` символы не ограничивает. |
| CRITICAL-2: json omitempty на обязательных полях | `store/types.go` | Убран `omitempty` из `Job.Payload`. `Entity.Meta`, `Relation.Value`, `Relation.Meta` уже были без omitempty. |
| WARNING-1: раскрытие деталей ошибок в default-ветке | `store/router/router.go` | `default` в `mapStoreError` возвращает `"internal error"`. Добавлен TODO-комментарий на будущий debug-режим. |
| WARNING-2: string-based error matching | `store/types.go`, `store/fs/*.go`, `store/router/router.go` | Добавлены sentinel errors: `ErrMissingID`, `ErrReadRecord`, `ErrWriteRecord`, `ErrDeleteRecord`, `ErrListRecords`, `ErrReadExisting`. FS-пакет оборачивает через `%w`. `mapStoreError` переписан на `errors.Is`. `containsAny` удалён. |
| WARNING-3: дублирование upsert (372 строки в одном файле) | `store/fs/` | Разбит на 5 файлов (все ≤150 строк): `store.go` (130), `upsert.go` (55), `io.go` (93), `validate.go` (26), `relation.go` (81), `job.go` (81). Обёртка `entityStore` убрана — `*Store` напрямую реализует `store.EntityStore`. |
| WARNING-4: atomicWrite defer без обработки Close | `store/fs/io.go` | `tmp.Close()` вызывается явно перед `os.Rename`; ошибка `Close()` обрабатывается; `defer` только для cleanup при ошибке. |
| Router (>150 строк) | `store/router/` | Разбит на 5 файлов: `router.go`, `handle_upsert.go`, `handle_get.go`, `handle_delete.go`, `handle_list.go` — все ≤150 строк. |

## Неисправленные замечания

| Замечание | Причина | Что нужно |
|---|---|---|
| NOTE-1: TODO без тикета | Согласованное решение: не трогать код | Вынести в ретро |

## Тесты после фикса

```
ok  github.com/unel/metadataplatform/cmd/store
ok  github.com/unel/metadataplatform/tasks/store/crud/tests/adversarial
ok  github.com/unel/metadataplatform/tasks/store/crud/tests/happy
```

Прошло: все (68/68) | Упало: 0
