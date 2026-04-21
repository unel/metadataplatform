# test-fix-2 — store/connection: deadlock в TestHighThroughput

## Провалившийся тест

`TestHighThroughput_1000Lines_NoneDropped` — `connection_adversarial_test.go:78`

## Поведение

Тест висит до срабатывания глобального таймаута (`-timeout 60s`), весь suite падает с паникой.

## Корневая причина

Deadlock между тестом и store:

1. Тест пишет 1000 строк залпом не читая ответы.
2. Store синхронно читает строку → пишет эхо → читает следующую.
3. Socket write-буфер теста заполняется → тест блокируется на записи.
4. Одновременно store пытается записать эхо в обратный буфер, который тоже заполнен.
5. Оба зависают — full-duplex deadlock.

## Stacktrace

```
goroutine 52: stuck at fmt.Fprintf → conn.Write  (connection_adversarial_test.go:89)  ← тест
goroutine 55: stuck at fmt.Fprintf → conn.Write  (handler.go:46)                      ← store
```
