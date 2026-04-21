# Test Fix 4 — store/connection (Азирафаль)

## Источник

Замечания из `test-review-2.md` — позитивные тесты.

---

## Исправленные замечания

### [minor] NFT-01 — `waitForMsg("conn_id")` заменён на `waitForConnID`

**Файл:** `cmd/store/connection_test.go`, функция `TestStore_ClientDisconnects_OtherClientsUnaffected`

`logs.waitForMsg("conn_id", 500*time.Millisecond)` всегда возвращал false и работал как `time.Sleep`, потому что `conn_id` живёт в structured attrs, а не в текстовом msg. Это вводило в заблуждение: вызов выглядел как смысловое ожидание, но им не был.

Заменён на `logs.waitForConnID(500*time.Millisecond)` — метод из testhelpers, который ищет именно в attrs. Теперь ожидание честное: тест действительно ждёт пока store залогирует подключение с conn_id.

### [minor] `TestStore_GracefulShutdown_ForcesCloseAfter5Seconds` — добавлен `testing.Short()` skip

**Файл:** `cmd/store/connection_test.go`, функция `TestStore_GracefulShutdown_ForcesCloseAfter5Seconds`

Тест ждёт 5–7 секунд (принудительное закрытие по таймауту shutdown). В CI это заметная задержка. Добавлен skip при `-short` флаге с пояснительным сообщением об ожидаемом времени выполнения.

---

## Не исправлено (намеренно)

### [medium] FT-08/FT-09 — SIGINT фактически не тестируется

Ограничение архитектуры embedded runner: оба сигнала моделируются закрытием `stopCh`. Для настоящего покрытия нужен либо black-box тест через `os.Process.Signal`, либо явный сигнальный канал в runner. Нового сценария в `acceptance.md` нет — тест не добавляю. Ограничение задокументировано в комментарии к FT-09 в тестовом файле.
