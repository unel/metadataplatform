# Code Fix 5 — CR-13: isCleanClose семантически некорректна

## Замечание

CR-13 из `code-review-3.md`: функция `isCleanClose` объединяла `syscall.EPIPE` и `net.ErrClosed` под одним именем и одним уровнем логирования (DEBUG). Но это принципиально разные случаи: `EPIPE` — клиент закрыл соединение со своей стороны, `net.ErrClosed` — мы сами закрыли `conn` из watcher-горутины по `ctx.Done`. Принудительный shutdown логировался как "connection closed (EOF)" на уровне DEBUG — событие терялось при диагностике.

## Изменения

**`cmd/store/handler.go`**

- Удалена функция `isCleanClose`.
- Добавлены две функции с точной семантикой:
  - `isClientClose(err)` — только `syscall.EPIPE`, clean FIN от клиента
  - `isForcedClose(err)` — только `net.ErrClosed`, принудительное закрытие с нашей стороны
- В точке вызова (`fmt.Fprintf` write error) `if/else` заменён на `switch`:
  - `isClientClose` → DEBUG "connection closed (EOF)" (без изменений)
  - `isForcedClose` → INFO "connection forcibly closed" (было DEBUG, стало INFO с внятным текстом)
  - default → INFO "write error: ..." (без изменений)
