# Test Review 2 — store/connection

## Статус: passed (с оговорками)

Проверено относительно замечаний из test-review.md (первое ревью, статус failed).

---

## Исправленные замечания из test-review.md

- [fixed] FT-13 — `time.Sleep` убран, используется `waitForNConnIDs` с таймаутом
- [fixed] FT-07 — conn_id теперь ищется через `waitForConnID()` (смотрит в attrs)
- [fixed] NFT-06 — тест фальсифицируем: если store завершится до sleep(200ms) — t.Fatal сработает
- [fixed] NFT-04 — проверка наличия лог-записи об ошибке добавлена
- [fixed] NFT-05 — SetLinger(0) убран, clean EOF имитируется штатным закрытием соединения
- [fixed] EDGE-01b — baseline фиксируется после `waitForNConnIDs(i+1)`, изоляция между итерациями обеспечена
- [fixed] newTestScanner — обёртка теперь имеет реальную ценность (буфер 4 МБ, задокументировано)

---

## Оставшиеся замечания

### [medium] FT-08/FT-09 — SIGINT фактически не тестируется

`connection_test.go:214` — оба теста проходят через `close(runner.stopCh)`. Разница между SIGTERM и SIGINT не проверяется: ни сигнал не посылается, ни отдельный обработчик не вызывается. Тест честно это документирует в комментарии, но это остаётся coverage gap — настоящая реализация с `signal.Notify` может иметь баг только в SIGINT-ветке и тест его не поймает.

Это ограничение архитектуры embedded runner, не тестов. Для полного покрытия нужен либо black-box тест через `os.Process.Signal`, либо явная заглушка сигнального канала в runner.

### [minor] NFT-01 — waitForMsg("conn_id") ищет в msg, а не в attrs

`connection_test.go:419` — `logs.waitForMsg("conn_id", 500*time.Millisecond)` служит как delay, но если conn_id живёт только в attrs (структурированные поля), этот вызов всегда вернёт false и просто отработает как `time.Sleep(500ms)`. Никакого real assertion нет — по сути `time.Sleep` маскирует ожидание. Не баг, но вводит в заблуждение.

### [minor] TestStore_GracefulShutdown_ForcesCloseAfter5Seconds — медленный тест без аннотации

`connection_test.go:636` — тест по сценарию ждёт 5+ секунд (принудительное закрытие по таймауту). Не помечен `t.Parallel()` или `testing.Short()`. В CI это 5-7 секунд на каждый прогон. Добавить хотя бы `testing.Short()` skip или задокументировать ожидаемое время.

---

## Покрытие acceptance.md

Все 22 сценария (FT-01..14, NFT-01..06, EDGE-01, 01b, 02, 03, 04) покрыты тестами в connection_test.go. Adversarial-файл добавляет сверх acceptance — это правильно.

Тестов без соответствующего acceptance-сценария не обнаружено.
