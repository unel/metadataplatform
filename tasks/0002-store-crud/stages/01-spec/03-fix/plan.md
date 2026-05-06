---
generated: 2026-04-28T15:20:00Z
created: 2026-04-28T15:20:00Z
updated: 2026-04-28T15:20:00Z
base-plan: 01-spec-fix v1.0.0
---

# План исправлений: store/crud spec

## Источники

- Спека: `tasks/0002-store-crud/stages/01-spec/01-write/report-001.md`
- Замечания: `tasks/0002-store-crud/stages/01-spec/02-review/report-001.md`

## CR-1 [warning] — Delete не идемпотентен: нет обоснования

**Что меняем:** раздел "Решения принятые в процессе" (ft.md).
**Как:** добавить явную запись — Delete не идемпотентен намеренно; spawner при retry получает `NOT_FOUND` и обязан трактовать это как ошибку.

## CR-2 [warning] — Порядок List не определён

**Что меняем:** §3.4 List (ft.md).
**Как:** добавить явный пункт — порядок не гарантирован. Тесты используют `assert.ElementsMatch`.

## CR-3a [warning] — `search_tsv` в Go-struct Relation

**Что меняем:** §2.2 Relation (ft.md).
**Как:** добавить комментарий в struct — поле исключено намеренно; PG управляет через триггер.

## CR-3b [warning] — Проверка `from_id`/`to_id` при upsert relation в FS

**Что меняем:** §7 Ограничения MVP (ft.md).
**Как:** добавить явную запись — FS не проверяет FK; целостность только в PG.

## CR-4 [note] — `io.EOF` при decode

**Что меняем:** §5.1 (ft.md).
**Как:** добавить абзац — `io.EOF`/`io.ErrUnexpectedEOF` штатное завершение; выход из цикла без логирования.

## CR-5 [note] — `data: null` при upsert

**Что меняем:** §5.2, §5.4 (ft.md).
**Как:** шаг 3 в two-pass decode — проверка data на null/отсутствие → `INVALID_REQUEST`.

## CR-6 [note] — Логирование успешных операций

**Что меняем:** §НФТ-О-1 (nft.md).
**Как:** добавить DEBUG-логирование успешных операций с обоснованием.

## CR-7 [note] — ctx пробрасывается в List

**Что меняем:** §3.4 List (ft.md).
**Как:** добавить явный пункт — ctx пробрасывается уже сейчас для совместимости будущей параллельной реализации.

## CR-8 [note] — Cleanup через defer

**Что меняем:** §4.3 (ft.md) + §НФТ-Н-3 (nft.md).
**Как:** cleanup через `defer func() { if err != nil { os.Remove(tmpPath) } }()`. Запрет if-err-branch.
