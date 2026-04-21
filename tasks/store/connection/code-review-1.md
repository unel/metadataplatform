# Code Review — store/connection

## Критичные

### CR-1: Race на WaitGroup при shutdown под нагрузкой

**Файл:** `cmd/store/store.go`, строки 53–70

`acceptLoop` делает `wg.Add(1)` внутри себя для каждого принятого соединения (`store.go:114`). При shutdown: `ln.Close()` вызывается, `acceptLoop` должна выйти из `ln.Accept()` с ошибкой. Но если планировщик поставил `acceptLoop` на паузу между `ln.Accept()` (успешный вызов, соединение принято) и `wg.Add(1)`, а в это время все ранее запущенные handler-горутины завершились и `wg.Wait()` вернул управление — последующий `wg.Add(1)` из acceptLoop является UB по документации sync.WaitGroup: "If a WaitGroup is reused to wait for several independent sets of events, new Add calls must happen after all previous Wait calls have returned."

На практике вероятность мала (Accept возвращает ошибку после Close до следующего wg.Add), но это зависит от планировщика и порядка выполнения горутин. Под нагрузкой при одновременном потоке входящих соединений и сигнале shutdown race реален.

**Фикс:** acceptLoop должна проверять контекст/stopCh перед `wg.Add`, либо использовать другой механизм синхронизации (например, отдельный semaphore для счёта активных conn-горутин, не зависящий от общего wg).

---

## Важные

### CR-2: `RunStoreWithExitCode` принимает `chan struct{}` вместо `<-chan struct{}`

**Файл:** `cmd/store/store.go`, строка 75

`RunStore` правильно объявляет `stopCh <-chan struct{}`. `RunStoreWithExitCode` объявляет `stopCh chan struct{}` — двунаправленный канал. Функция только читает из него. Несоответствие сигнатур: внешний код может случайно закрыть канал из этой функции, хотя это не её ответственность.

**Фикс:** изменить сигнатуру на `stopCh <-chan struct{}`. Это ломает совместимость только в сторону корректности.

### CR-3: `applyDefaults` вызывается дважды — в `RunStoreWithExitCode` и в `RunStore`

**Файл:** `cmd/store/store.go`, строки 19 и 85

`RunStoreWithExitCode` вызывает `applyDefaults` (`store.go:85`), затем передаёт cfg в `RunStore`, который снова вызывает `applyDefaults` (`store.go:19`). Сейчас это harmless (функция идемпотентна). Если логика defaults усложнится — двойной вызов станет источником трудноуловимых багов.

**Фикс:** убрать вызов `applyDefaults` из `RunStore` — или из `RunStoreWithExitCode`. Одна точка ответственности.

### CR-4: `acceptLoop` принимает `*sync.WaitGroup` и делает `wg.Add` — скрытый side effect

**Файл:** `cmd/store/store.go`, строки 103–119

Функция `acceptLoop` по сигнатуре выглядит как "запустить цикл приёма соединений". На деле она мутирует внешний WaitGroup, добавляя в него счётчики для каждой conn-горутины. Это нарушение no-side-effects и single-responsibility: вызывающий код (`RunStore`) не видит что acceptLoop изменяет его wg.

Дополнительно: вызывающий код (`store.go:43`) делает `wg.Add(1)` для самой acceptLoop-горутины, а acceptLoop внутри делает `wg.Add(1)` для каждого handler. Это работает, но ответственность за управление wg размазана между двумя уровнями.

**Фикс:** acceptLoop возвращает или управляет своим локальным wg, либо принимает callback `onConn func(net.Conn, uint64)` и не знает ничего про глобальный wg.

### CR-5: Неограниченное число параллельных соединений

**Файл:** `cmd/store/store.go`, `acceptLoop`

Стандарт проекта (concurrency/worker-pool): "количество параллельных воркеров ограничено и настраиваемо". Текущая реализация создаёт горутину на каждое входящее соединение без лимита. При flood-атаке или ошибочном поведении клиентов store создаст неограниченное число горутин вплоть до OOM.

Спека MVP явно не требует лимита соединений, но стандарт проекта требует. Следующая фича (DB-соединение) усугубит: каждая горутина откроет DB-соединение.

**Фикс:** добавить параметр конфига `store.max_connections` с дефолтом, реализовать semaphore-based ограничение в acceptLoop.

---

## Минорные

### CR-6: `waitForSocket` объявлена в testhelpers_test.go но нигде не используется

**Файл:** `cmd/store/testhelpers_test.go`, строка 307

Мёртвый код в тестах. В Go функции в тестовых файлах не вызывают ошибку компиляции при неиспользовании, но это мусор.

### CR-7: `cancel()` вызывается в обеих ветках select при shutdown

**Файл:** `cmd/store/store.go`, строки 65 и 68

`cancel()` вызывается и в ветке `<-done`, и в ветке `<-time.After(...)`. Context.CancelFunc идемпотентна, это safe. Но выглядит как дублирование — лучше вынести вызов за select или добавить комментарий что это намеренно.

### CR-8: `StoreProcess.logs []string` объявлено в struct но не используется

**Файл:** `cmd/store/testhelpers_test.go`, строка 22

В `StoreProcess` есть поле `logs []string` которое нигде не заполняется и не читается (реальные логи идут через `sp.runner.logBuf`). Мёртвое поле.
