# Test Review — store/connection

## Статус: failed

### [medium] FT-08/FT-09 — SIGTERM и SIGINT тестируют одно и то же
`connection_test.go:212` — оба теста проходят через `stopCh`, SIGINT фактически не отличается от SIGTERM. Тест не проверяет что именно SIGINT обработан отдельно.

### [medium] NFT-06 — WaitsForActiveConnections нефальсифицируем
`connection_test.go:590` — если store не ждёт активные соединения и завершается сразу, тест всё равно проходит зелёным. Нет assertion что store был жив в момент между SIGTERM и закрытием соединения.

### [medium] FT-13 — time.Sleep вместо детерминированного ожидания
`connection_test.go:343` — `time.Sleep(100ms)` для ожидания логов флакает на медленных машинах. Нужен `waitFor` с условием и таймаутом.

### [minor] FT-07 — conn_id ищется в неправильном поле лога
`connection_test.go:171` — `waitForMsg("conn_id")` ищет подстроку в `msg`, но `conn_id` живёт в structured fields (`attrs`). Нужно использовать `connIDs()` или аналог.

### [minor] NFT-04 — не проверяется лог ошибки записи
`connection_test.go:519` — тест проверяет что store продолжает работу после ошибки записи, но не проверяет что ошибка залогирована (это требование из acceptance NFT-04).

### [minor] NFT-05 — SetLinger(0) не работает на Unix socket
`connection_test.go:568` — `SetLinger(0)` применимо только к TCP. На Unix socket`ах для имитации IO-ошибки (а не чистого EOF) нужен другой подход.

### [minor] EDGE-01b — счётчик логов нестабилен в цикле
`connection_test.go:709` — в цикле по whitespace-строкам счётчик логов может сдвинуться из-за задержки логирования предыдущей итерации. Нужна изоляция соединений между итерациями.

### [minor] testhelpers_test.go — newTestScanner бессмысленная обёртка
`testhelpers_test.go:315` — `newTestScanner` это тривиальная обёртка над `bufio.NewScanner` с неверным комментарием. Либо убрать, либо обёртка должна добавлять реальную ценность (буфер, таймаут).
