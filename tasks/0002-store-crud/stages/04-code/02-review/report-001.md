---
purpose: Ревью кода для store/crud
process: 04-code/02-review
run: 1
date: 2026-05-04T12:44:24Z
created: 2026-05-04T12:44:24Z
see-also:
  - tasks/0002-store-crud/stages/04-code/01-write/report-001.md
status: failed
agent: Гримм
checklist: открытые: path traversal, json omitempty, error leakage, string-based error matching, module size, atomicWrite defer order
---

## Результат

Найдено 7 замечаний (critical: 2, warning: 4, note: 1).

## Замечания

### `store/fs/store.go:297` — Path traversal через id

**Категория:** critical
**Проблема:** `recordPath` строит путь через `filepath.Join(dir, id+".json")`. `validateID` проверяет только пустую строку. Если id содержит `../` — `filepath.Join` нормализует путь и запись выйдет за пределы basedir. Воспроизводится напрямую: `Upsert(ctx, Entity{ID: "../../tmp/pwned"})` создаст `/tmp/pwned.json`.
**Рекомендация:** В `validateID` запретить `/`, `\`, `..` и другие path-небезопасные символы.

---

### `store/types.go:19,30-31,46` — Расхождение json-тегов со spec

**Категория:** critical
**Проблема:** Spec §2.1–2.3 определяет `Entity.Meta`, `Relation.Value`, `Relation.Meta`, `Job.Payload` без omitempty. Код везде добавляет `omitempty`. При nil-значениях поля выпадают из JSON-ответа — клиенты которые ожидают `"meta": null` получают объект без поля `meta`. Нарушение контракта AC-ENTITY-UPSERT-01, AC-RELATION-UPSERT-01, AC-JOB-UPSERT-01.
**Рекомендация:** Убрать `omitempty` из `Meta`, `Value`, `Payload` в types.go.

---

### `store/router/router.go:393-395` — Внутренние детали ошибок в ответе клиенту

**Категория:** warning
**Проблема:** `mapStoreError` в ветке `default` возвращает `msg` (строку из `err.Error()`) напрямую в поле `error` ответа. Клиент может получить `"open /var/lib/platform/store/entities/x.json: permission denied"` — путь к файловой системе, детали реализации. Нарушение `base-plan.md` §8.
**Рекомендация:** В ветке `default` вернуть фиксированную строку `"internal error"` вместо `msg`.

---

### `store/router/router.go:371-409` — mapStoreError парсит строки; containsAny — самодельный strings.Contains

**Категория:** warning
**Проблема:** Тип ошибки определяется через сравнение подстрок в `err.Error()`. Переформулировать сообщение в одном месте — маппинг сломается молча. `containsAny` (строки 398-409) — ручной цикл по байтам вместо `strings.Contains`. Пакет `strings` в импортах файла уже есть.
**Рекомендация:** Использовать typed errors или sentinel errors с `errors.Is`/`errors.As`. Заменить `containsAny` на `strings.Contains`.

---

### `store/fs/store.go` — Модуль 372 строки, логика upsert дублируется трижды

**Категория:** warning
**Проблема:** Лимит по стандарту — 150 строк. Логика upsert (читать существующий → установить timestamps → atomicWrite) скопирована для `*Store`, `relationStore`, `jobStore`. Три копии одного алгоритма будут расходиться при изменениях. `entityStore` существует только чтобы делегировать вызовы — лишний уровень без смысла.
**Рекомендация:** Вынести общую upsert-логику в обобщённую функцию или generic-helper.

---

### `store/fs/store.go:302-325` — atomicWrite: дескриптор tmp живёт после os.Rename

**Категория:** warning
**Проблема:** `defer tmp.Close()` объявлен до cleanup-defer. После успешного `os.Rename` файловый дескриптор закрывается на уже переименованный файл. На Linux — не падение, но дескриптор держится дольше нужного и логика неочевидна при чтении. Ошибка `Close()` игнорируется.
**Рекомендация:** Закрывать `tmp` явно перед `os.Rename`, defer использовать только для cleanup при ошибке.

---

### `store/fs/store.go:32` — TODO без номера тикета

**Категория:** note
**Проблема:** `// TODO: add RWMutex for concurrent access` — без номера тикета. Стандарт требует тикет. Spec §4.7 предписывает именно такую формулировку — противоречие стандарта и spec.
**Рекомендация:** Разрешить на уровне процесса: либо убрать TODO либо добавить тикет. Не блокирует.

---

## Соответствие acceptance

| Блок | Статус |
|---|---|
| AC-FS-INIT | Соответствует |
| AC-ENTITY-UPSERT-01, AC-RELATION-UPSERT-01, AC-JOB-UPSERT-01 | Нарушено (omitempty → поля выпадают из JSON) |
| AC-ENTITY-DELETE, AC-ENTITY-LIST | Соответствует |
| AC-ROUTER-* | Соответствует (кроме INTERNAL_ERROR — раскрытие деталей) |
| AC-NFT-ATOMIC | Соответствует |
