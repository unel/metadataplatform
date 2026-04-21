# Code Fix 1 — store/connection

## Источник

Задание: исправить 5 багов в реализации фичи store/connection.

---

## Результат анализа

Из 5 описанных багов 4 уже присутствовали в исправленном виде в коде на момент задания. Реально потребовал правки только баг 2.

---

## Баг 1 — Stale socket file

**Статус: уже исправлен**

`listenUnix` в `cmd/store/store.go` (строки 98–101) уже содержит `os.Remove(path)` перед `net.Listen`. Файл не трогался.

---

## Баг 2 — Partial read EOF не логируется

**Статус: исправлен**

**Файл:** `cmd/store/handler.go`

`io.ErrUnexpectedEOF` ранее не имел отдельной ветки — падал в общий `log.add("INFO", "read error: ...")`. Добавлена явная ветка после обработки `io.EOF`:

```go
if errors.Is(err, io.ErrUnexpectedEOF) {
    log.add("DEBUG", "connection closed (unexpected EOF, partial line)", attrs)
    return
}
```

Уровень DEBUG соответствует спеке: "нормальное закрытие (клиент прислал EOF) — уровень DEBUG".

---

## Баг 3 — Graceful shutdown не ждёт соединения

**Статус: уже исправлен**

`RunStore` в `cmd/store/store.go` использует `wg.Wait()` с таймаутом 5 секунд через `select { case <-done: ... case <-time.After(5*time.Second): ... }`. Активные handler-горутины удерживают `wg` открытым. Тест `TestStore_GracefulShutdown_WaitsForActiveConnections` проходит. Файл не трогался.

---

## Баг 4 — Пустая строка логируется

**Статус: уже исправлен**

`handleConn` в `cmd/store/handler.go` (строки 84–86) проверяет `strings.TrimSpace(line) == ""` и делает `continue` без лога и ответа. Файл не трогался по этому пункту.

---

## Баг 5 — Exit code 0 при отсутствии директории сокета

**Статус: уже исправлен**

При ошибке `net.Listen` в `RunStore` возвращается `false`, `RunStoreWithExitCode` возвращает `1`, `main` вызывает `os.Exit(1)`. Тест `TestStore_SocketDirDoesNotExist_FailsWithError` проходит. Файл не трогался.

---

## Изменённые файлы

- `cmd/store/handler.go` — добавлена ветка `io.ErrUnexpectedEOF` с DEBUG логом
