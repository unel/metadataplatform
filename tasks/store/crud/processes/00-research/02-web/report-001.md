---
purpose: Результаты web-research с источниками — входящий артефакт для 01-spec/01-write
process: 00-research/02-web
run: 1
date: 2026-04-28T14:37:17Z
created: 2026-04-28T14:37:17Z
status: done
agent: Бо
checklist: все пункты закрыты
---

## Похожие реализации

- **golang-scribble (nanobox-io/golang-scribble)** — файловая JSON БД: одна директория на коллекцию, один файл на запись. Concurrency явно помечена как TODO — "Better support for concurrency". Показывает что наивная реализация без locking — норма для debug-хранилища, но нужна заметка. ([источник](https://github.com/nanobox-io/golang-scribble))

- **jorzel/go-repository-pattern** — эталонный пример repository pattern в Go с несколькими реализациями (in-memory, Redis, external API) за одним интерфейсом. Интерфейс определяется на стороне потребителя, не провайдера. ([источник](https://github.com/jorzel/go-repository-pattern))

- **google/renameio** — стандарт де-факто для atomic file write в Go продакшне. TempFile создаёт файл в той же директории что и цель, CloseAtomicallyReplace делает rename. ([источник](https://pkg.go.dev/github.com/google/renameio/v2))

- **natefinch/atomic** — более простой API: WriteFile и ReplaceFile. Та же идея temp+rename, кроссплатформенная обёртка. ([источник](https://github.com/natefinch/atomic))

## Известные failure modes

- **UUID v7 без монотонности при batch-генерации** — google/uuid до issue #148 не гарантировал порядок UUID при генерации нескольких штук в одну миллисекунду. При FS-хранилище не критично (сортировка по id не требуется), но важно знать. ([источник](https://github.com/google/uuid/issues/148))

- **os.Rename: "invalid cross-device link"** — если tempdir (`os.TempDir()`) на другом filesystem чем целевой файл, rename падает. Решение: создавать temp-файл в той же директории что и цель. Реальный кейс: hashicorp/consul-template issue #58. ([источник](https://github.com/hashicorp/consul-template/issues/58))

- **os.WriteFile не атомарен** — явно задокументировано в golang/go issue #56173. При сбое в середине записи файл окажется частично перезаписан. ([источник](https://github.com/golang/go/issues/56173))

- **bufio.Scanner максимальный размер токена 64KB** — при JSONL-парсинге через Scanner: если одна строка > 64KB, Scanner падает с `token too long`. Рекомендация: `json.NewDecoder` напрямую на `net.Conn`. ([источник](https://chriswilcox.dev/blog/2024/04/09/Scan-vs-Read-in-bufio.html))

- **Concurrent writes к FS без locking** — несколько горутин пишут в один и тот же файл — data corruption. Даже с atomic write конкурентная запись одного id двумя горутинами может затереть данные друг друга. ([источник](https://blog.gopheracademy.com/advent-2014/safe-json-file-db-in-go/))

- **NFS + atomic rename не атомарен** — google/renameio явно документирует: rename(2) не атомарен на NFS с несколькими клиентами. Для локального FS — без проблем. ([источник](https://pkg.go.dev/github.com/google/renameio/v2))

## Паттерны и best practices

- **UUID v7: google/uuid v1.6+** — `uuid.NewV7()` добавлен в v1.5.0, текущий stable v1.6.0 (январь 2024). Для FS-реализации монотонность не критична. ([источник](https://pkg.go.dev/github.com/google/uuid))

- **Atomic write: temp в той же директории** — `os.CreateTemp(filepath.Dir(target), ".tmp-*")` → write → `Sync()` → `Close()` → `os.Rename(tmp, target)`. `Sync()` обязателен — без него данные могут остаться в page cache при crash. ([источник](https://michael.stapelberg.ch/posts/2017-01-28-golang_atomically_writing/))

- **ErrNotFound — sentinel var** — `var ErrNotFound = errors.New("not found")` на уровне пакета. Caller проверяет через `errors.Is(err, store.ErrNotFound)`. Typed error — только когда нужен контекст (какой id не найден), для нас sentinel достаточен. ([источник](https://go.dev/blog/go1.13-errors))

- **Two-pass JSON decode через json.RawMessage** — для JSONL-роутера с переменной структурой `data`. Первый decode: `{"op":"...", "type":"...", "data": json.RawMessage}`. Второй decode: `data` → конкретная структура после switch по `op`+`type`. ([источник](https://eagain.net/articles/go-dynamic-json/))

- **Repository interface — маленький и сфокусированный** — Interface Segregation: `EntityStore` отдельно, `RelationStore` отдельно, `JobStore` отдельно. Позволяет заменить реализацию одного типа независимо. ([источник](https://maddevs.io/blog/solid-interface-segregation-principle-in-golang/))

- **Locking-ready дизайн без оверинжиниринга** — mutex встраивается в struct реализации, не в интерфейс. Интерфейс остаётся чистым. Альтернатива: отдельный `LockedEntityStore` как обёртка-декоратор. ([источник](https://pkg.go.dev/sync))

- **json.NewDecoder для JSONL-роутера** — для `net.Conn` предпочтительнее `json.NewDecoder(conn)` с циклом `Decode()` чем `bufio.Scanner` + `Unmarshal` — нет ограничения в 64KB. ([источник](https://dev.to/taqkarim/you-might-not-be-using-json-decoder-correctly-in-golang-12mb))

## Выводы для spec

1. **Atomic write**: явно требовать temp+rename для upsert. `os.WriteFile` запрещён. Temp-файл — в той же директории что и цель.
2. **Locking**: зафиксировать что concurrent access не защищается в этой итерации, но структура допускает добавление mutex без изменения интерфейса. Требовать заметку `// TODO: add RWMutex for concurrent access` в коде.
3. **ErrNotFound**: sentinel var `var ErrNotFound = errors.New("not found")` в пакете store. Роутер транслирует в `{"ok":false,"error":"not found","errorCode":"NOT_FOUND"}`.
4. **JSONL-роутер**: two-pass decode. Неизвестный `op` или `type` — немедленная ошибка, не fallthrough.
5. **Интерфейсы**: три отдельных (`EntityStore`, `RelationStore`, `JobStore`). Методы: `Upsert`, `Get`, `Delete`, `List`.
6. **List ошибки**: зафиксировать в spec поведение при ошибке чтения одного файла — пропустить или вернуть error?

## Выводы для реализации

1. `uuid.NewV7()` из `github.com/google/uuid` v1.6+. Не `uuid.New()` (это v4).
2. Atomic write: `os.CreateTemp(filepath.Dir(path), ".tmp-*")` → encode → `f.Sync()` → `f.Close()` → `os.Rename(tmpPath, path)`. При ошибке — `os.Remove(tmpPath)`.
3. `f.Sync()` обязателен перед rename.
4. Directory structure: `{basedir}/entities/{uuid}.json`, `{basedir}/relations/{uuid}.json`, `{basedir}/jobs/{uuid}.json`.
5. `List` — `os.ReadDir(dir)`, filter `.json`, decode каждый. Поведение при ошибке одного файла — уточнить в spec.
6. JSONL-роутер: `json.NewDecoder` на `net.Conn`, loop `decoder.Decode(&envelope)`. Не `bufio.Scanner`.
7. `created_at` только при первом создании (check-then-set), `updated_at` — всегда при upsert через `time.Now().UTC()`.
