---
purpose: Ревью кода для store/crud
process: 04-code/02-review
run: 2
date: 2026-05-04T17:43:00Z
created: 2026-05-04T17:43:00Z
see-also: tasks/0002-store-crud/stages/04-code/03-fix/report-001.md
status: done
agent: Гримм
checklist: все пункты закрыты
---

## Результат

Код чистый. Готово к финальному прогону тестов.

## Проверка исправлений из run 1

**CRITICAL-1 (path traversal)** — закрыто. `store/fs/validate.go`: `recordPath` после `filepath.Join` проверяет `strings.HasPrefix(filepath.Clean(p)+sep, filepath.Clean(basedir)+sep)`. При `id="../../tmp/pwned"` путь нормализуется и проверка провалится. Защита корректна.

**CRITICAL-2 (omitempty на обязательных полях)** — закрыто. `store/types.go`: `Entity.Meta`, `Relation.Value`, `Relation.Meta`, `Job.Payload` — без `omitempty`. `Job.Progress`, `Job.Error`, `Job.EntityID`, `Job.RelationID` — `omitempty` оставлен, поля опциональные по spec. Корректно.

**WARNING-1 (раскрытие деталей в default)** — закрыто. `router.go:130` — `Error: "internal error"`, детали FS не раскрываются.

**WARNING-2 (string-based error matching)** — закрыто. `mapStoreError` переписан на `errors.Is`, sentinel errors определены в `store/types.go:10-18`, FS оборачивает через `%w`, `containsAny` удалён.

**WARNING-3 (размер модулей, дублирование upsert)** — закрыто. `store/fs/` разбит на 6 файлов (все ≤ 150 строк), общая логика вынесена в `upsertRecord`. Router — 5 файлов.

**WARNING-4 (atomicWrite defer)** — закрыто. `store/fs/io.go`: явное `tmp.Close()` перед `os.Rename`, ошибка Close обрабатывается, defer только для cleanup при ошибке.

## Новые замечания

Нет.

## Соответствие acceptance

Все сценарии AC-FS-INIT, AC-ENTITY-*, AC-RELATION-*, AC-JOB-*, AC-ROUTER-01..16, AC-NFT-ATOMIC-01, AC-ROUTER-INIT-01 соответствуют реализации.
