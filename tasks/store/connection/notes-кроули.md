[miss] acceptance не упомянула full-duplex семантику Unix socket-буфера; FT-04 описывает последовательную отправку без оговорки что клиент должен параллельно читать ответы; из этой дыры вырос deadlock в TestHighThroughput
[miss] acceptance не описывает поведение при MaxConnections: что происходит когда лимит достигнут, какой уровень лога, получает ли клиент EOF или тишину; тест не написан — основания не было
[miss] нет сценария для двойного сигнала (Ctrl+C дважды); CR-12 нашёл реальный panic в main.go; один SIGTERM/SIGINT покрыт, повторный в ожидании 5-секундного таймаута — нет
[miss] acceptance NFT-05 описывает два сценария (EOF и IO-ошибка) в одном блоке, не объясняя как вызвать настоящую IO-ошибку на Unix socket; SetLinger(0) не работает на Unix; тест покрывает только clean EOF для обоих путей
[rework] TestStore_ClientClose_PartialRead_LoggedAtDebug переписан дважды: сначала с SetLinger(0) (не работает на Unix), потом без; SetLinger на Unix — архитектурный факт, не edge case
[rework] TestStore_GracefulShutdown_WaitsForActiveConnections переписан: изначально t.Log вместо t.Fatal при преждевременном завершении; тест был нефальсифицируемым
[rework] TestHighThroughput_1000Lines_NoneDropped переписан из-за deadlock; синхронная запись 1000 строк без параллельного чтения — классическая ловушка full-duplex; в FT-04 всего 3 строки — буфер не забивается, проблема не видна
[friction] два test-review цикла, 4 итерации фиксов; половина замечаний — про structured logging (conn_id в attrs, не в msg); это надо документировать в стандартах, не угадывать
[friction] dialStore добавляет time.Sleep(25ms) после коннекта — "чтобы acceptLoop успел залогировать"; детерминированное ожидание частично решает, но dialStore несёт неявный delay
[propose] стандарты тестирования должны явно описывать: structured logging → искать через attrs, не через содержимое msg; два одинаковых замечания в двух разных ревью — сигнал
[propose] acceptance для асинхронных протоколов должна явно указывать: "клиент читает ответы параллельно с записью"; без этого тест пишется наивно и при достаточном объёме замирает
[doc] Unix socket + SO_LINGER(0): не даёт RST в отличие от TCP; NFT-05 фактически не покрыта настоящим тестом — только чистый EOF
[doc] MaxConnections появился по code-review, отсутствует в спеке и acceptance; поведение при лимите — undocumented
[whatever] CR-9 (deadlock graceful shutdown) — критичный баг; TestStore_GracefulShutdown_ForcesCloseAfter5Seconds зеленел при сломанном коде из-за 6-секундного таймаута обёртки; классика: тест проверяет форму, а не содержание
[miss] FT-08/FT-09 — coverage gap по SIGINT зафиксирован в двух ревью и двух fix-report, но не закрыт; ограничение архитектуры embedded runner: оба сигнала идут через close(stopCh), signal.Notify-ветка не тестируется; баг в SIGINT-обработчике будет невидим
[friction] test-run-report называет буфер Unix socket "~8 КБ", test-fix-3 называет "~208 КБ"; в обоих случаях буфер конечен и deadlock воспроизводится — но расхождение в цифрах говорит что никто не смотрел реальный kernel-default; SO_SNDBUF/SO_RCVBUF на unix socket не тривиальны
[doc] go test паттерн ./store/... не работает если store — скомпилированный бинарник в корне; корректный паттерн ./cmd/store/...; это не очевидно и reproducible ошибка в CI если кто-то скопирует команду из test-run-report
[doc] тест который "выглядит как assertion но является time.Sleep" — антипаттерн; waitForMsg("conn_id") в NFT-01 возвращал false всегда (conn_id в attrs, не в msg) и работал как задержка; два ревью понадобилось чтобы это поймать
[propose] медленные тесты (>2s) должны помечаться testing.Short() skip как конвенция с первой итерации, не как замечание ревью; TestStore_GracefulShutdown_ForcesCloseAfter5Seconds добавил skip только в fix-4
