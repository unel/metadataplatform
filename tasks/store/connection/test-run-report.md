# test-run-report — store/connection

Команда: `go test -v -race -timeout 120s ./cmd/store/... ./store/...`

## Статус: FAILED

## Упавшие тесты

### TestHighThroughput_1000Lines_NoneDropped — TIMEOUT / DEADLOCK

Файл: `cmd/store/connection_adversarial_test.go:78`

Тест завис и убит глобальным таймаутом (120s).

Корневая причина: full-duplex deadlock между тестом и store.

1. Тест пишет 1000 строк залпом в основной горутине, не читая ответы.
2. Store читает строку → пишет эхо → читает следующую (синхронно).
3. Буфер unix socket (~8 КБ) переполняется — тест блокируется на `fmt.Fprintf` (строка 89).
4. Store пытается записать эхо, но буфер в обратную сторону тоже заполнен — store блокируется на `fmt.Fprintf` (handler.go:46).
5. Оба ждут друг друга. Вечно.

Стектрейс:
```
goroutine 52: stuck at fmt.Fprintf → conn.Write  (connection_adversarial_test.go:89)  ← тест
goroutine 42: stuck at fmt.Fprintf → conn.Write  (handler.go:46)                      ← store
```

Фикс: в тесте отправку вынести в отдельную горутину, читать ответы параллельно.

## Прошедшие тесты

```
PASS: TestMaxLineBytes_ExactlyAtLimit_Accepted
PASS: TestMaxLineBytes_OneByteBeyondLimit_Rejected
PASS: TestMaxLineBytes_MinimalLimit_OneByte
```

Остальные тесты не запустились — suite упал паникой по таймауту после зависания на третьем тесте.

## Примечание по `./store/...`

`store` в корне проекта — скомпилированный бинарник, не Go-пакет. Паттерн `./store/...` возвращает ошибку `lstat ./store/: not a directory`. Тесты запускались только по `./cmd/store/...`.
