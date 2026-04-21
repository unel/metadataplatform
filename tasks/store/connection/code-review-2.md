# Code Review 2 — store/connection

Повторное ревью после фикса CR-1..CR-8 из code-review-1.md.

## Статус исправлений из code-review-1.md

- CR-1 (Race на WaitGroup): исправлен. acceptLoop имеет локальный wg, defer wg.Wait().
- CR-2 (двунаправленный stopCh): исправлен. RunStoreWithExitCode принимает <-chan struct{}.
- CR-3 (двойной applyDefaults): частично. RunStoreWithExitCode больше не вызывает applyDefaults напрямую — но loadConfig вызывает его (config.go:41), и RunStore тоже (store.go:19). Если RunStore вызывается с конфигом из loadConfig — defaults применяются дважды. Идемпотентно, не ломает. Принята как is.
- CR-4 (side effect wg в acceptLoop): исправлен.
- CR-5 (неограниченные соединения): исправлен. Семафор, MaxConnections в конфиге.
- CR-6 (waitForSocket мёртвый код): исправлен.
- CR-7 (cancel() дублируется): исправлен.
- CR-8 (мёртвое поле logs): исправлен.

---

## Новые замечания

### CR-9: Graceful shutdown не форсирует закрытие соединений после таймаута [critical]

**Файл:** `cmd/store/store.go`, строки 67–74

```go
select {
case <-done:
case <-time.After(5 * time.Second):
    <-done   // ← дедлок если conn-горутины не завершились
}
cancel()
```

После срабатывания таймаута код блокируется на `<-done` — то есть всё равно ждёт завершения conn-горутин. Но `cancel()` вызывается только после `<-done` — значит он никогда не будет вызван пока горутины живы. Контекст не отменяется, `conn.Close()` в горутинах-наблюдателях не срабатывает, горутины не завершаются. Дедлок.

По спеке: "Если по истечении 5 секунд остались активные соединения — они принудительно закрываются". Принудительное закрытие обеспечивается именно через `cancel()` → `ctx.Done()` → `conn.Close()` в handleConn. Но cancel нужно вызывать до второго `<-done`, а не после.

Правильная последовательность:
```go
select {
case <-done:
    cancel()
case <-time.After(5 * time.Second):
    cancel()  // форсируем закрытие активных соединений
    <-done    // ждём завершения горутин после форсированного закрытия
}
```

Тест `TestStore_GracefulShutdown_ForcesCloseAfter5Seconds` при этом не ловит проблему — он ждёт через `sp.Stop()` который имеет собственный таймаут 6 секунд и завершает тест не из-за store, а из-за таймаута обёртки. Реальный баг маскируется тестом.

---

### CR-10: `ReadLine` — документально устаревший API без обоснования [minor]

**Файл:** `cmd/store/handler.go`, строка 46

`bufio.ReadLine` помечена в документации Go как "low-level primitive; most callers should use ReadBytes or Scanner instead". Код работает корректно, но использование устаревшего низкоуровневого API без комментария "почему не Scanner" создаёт вопрос для следующего читателя. При замене handleMessage на реальный роутер (следующая фича) этот код будут трогать — вопрос возникнет именно тогда.

Если выбор осознан (ReadLine даёт более точный контроль над isPrefix-логикой), достаточно одного комментария. Если нет — лучше Scanner с `sc.Buffer(make([]byte, maxLineBytes+1), maxLineBytes+1)`.

---

### CR-11: Ошибки `ln.Close()` при shutdown не логируются [minor]

**Файл:** `cmd/store/store.go`, строка 56–58

```go
for i, ln := range listeners {
    ln.Close()
    os.Remove(cfg.Store.Sockets[i].Path)
}
```

Ошибки `ln.Close()` и `os.Remove()` игнорируются. `os.Remove` при штатном завершении обычно не фейлится. `ln.Close()` тоже. Но если фейлится — никто не узнает. Минорно, но противоречит стандарту error-handling/dont-ignore.

---

## Итог

Одно критичное замечание (CR-9): graceful shutdown по таймауту не работает — `cancel()` никогда не вызывается пока активные соединения живы. Дедлок. Спека нарушена.

Два минорных (CR-10, CR-11): не блокируют, но стоит адресовать.
