---
purpose: Ревью тестов для store/crud
process: 03-tests/02-review
run: 1
date: 2026-04-29T18:32:52Z
created: 2026-04-29T18:32:52Z
status: failed
agent: Гримм
checklist: открытые: структура (отсутствующий helper), FIRST (несогласованная сигнатура), покрытие (AC-ROUTER-16 без прямого теста)
---

## Результат

**failed**. Три замечания: одно критическое (отсутствующий символ — тесты не скомпилируются), одно критическое (несогласованная сигнатура `r.Handle` между пакетами), одно medium (AC-ROUTER-16 покрыт только "неявно" без прямой проверки).

Остальное — чисто. Покрытие acceptance полное по всем остальным сценариям. FIRST соблюдён: temp-директории через `t.TempDir()`, детерминизм обеспечен, no shared mutable state. Поведение vs реализация — тесты проверяют контракт, не внутренние вызовы. Размеры файлов в норме. Именование читается как документация.

## Замечания

### happy/router_logging_test.go + happy/router_conn_test.go — отсутствующий символ `newRouterWithLogger`

**Категория:** Нарушение FIRST (Isolated — тест не самодостаточен), структура
**Severity:** critical
**Проблема:** Функция `newRouterWithLogger(t)` вызывается в `router_logging_test.go` (строки 16, 43, 70) и `router_conn_test.go` (строки 21). Нигде в пакете `happy_test` она не определена. `helpers_test.go` содержит только `newFSStore`, `routerConn`, `sendRequest`, `defaultEntity`, `defaultRelation`, `defaultJob`. Возвращаемый тип содержит `log` с методом `findEntry(level, op, kind string)` и полями `id string`, `count int` — структура нигде не объявлена. Пакет не скомпилируется. Ни один тест из `router_logging_test.go` и частично `router_conn_test.go` не запустится.
**Рекомендация:** В test-fix создать файл `store/tests/happy/router_logger_test.go` с определением `newRouterWithLogger`, типом `routerLogEntry` и методом `findEntry`. Либо дополнить `helpers_test.go`.

---

### adversarial/router_helpers_test.go vs happy/router_conn_test.go — несогласованная сигнатура `r.Handle`

**Категория:** Нарушение FIRST (Repeatable — тест зависит от того какой вариант сигнатуры примет реализация)
**Severity:** critical
**Проблема:** В `store/tests/happy/router_conn_test.go` строка 86: `r.Handle(context.Background(), srv)` — два аргумента: `ctx` и `conn`. В `store/tests/adversarial/router_helpers_test.go` строка 53: `go r.Handle(serverConn)` — один аргумент: только `conn`. Сигнатура метода `Handle` контрактно зафиксирована в архитектурных предположениях test-write как `router.Handle(ctx, conn net.Conn)`. Один из двух вызовов не скомпилируется. Вместе они создают противоречие которое code-write не может разрешить без поломки одного из пакетов.
**Рекомендация:** В test-fix привести к единой сигнатуре. Если принято `Handle(ctx context.Context, conn net.Conn)` — исправить adversarial/router_helpers_test.go строка 53. Если `Handle(conn net.Conn)` — исправить happy/router_conn_test.go строка 86 и убрать импорт `context` оттуда же.

---

### AC-ROUTER-16 — покрытие только "неявным" тестом, прямая проверка отсутствует

**Категория:** Пробел в покрытии
**Severity:** medium
**Проблема:** AC-ROUTER-16 требует: при ошибке операции (например, get несуществующего ID) роутер логирует ERROR с полями `op`, `type`, `id`, `error`. В маппинге report-001 указано "неявно покрыт через TestRouter_UnknownOp/UnknownType". Но эти тесты проверяют только ответ клиенту (`UNKNOWN_OP`/`UNKNOWN_TYPE`) — они не захватывают лог и не проверяют поля `error`. Прямого теста который создаёт ситуацию "операция вернула ошибку стора" (например, `get` несуществующего ID через роутер) и верифицирует ERROR-лог с нужными полями — нет.
**Рекомендация:** В test-fix добавить тест в adversarial: роутер получает `{"op":"get","type":"entity","id":"ent-999"}`, в logBuf проверяем наличие `ERROR`, полей `op=get`, `type=entity`, `id=ent-999`, `error=`. Паттерн уже есть в `TestRouter_InvalidJSON_ClosesConnection_LogsParseError` — использовать аналогичный.

---

### atomic_test.go — тест зависит от тайминга (30ms hardcode)

**Категория:** Нарушение FIRST (Repeatable)
**Severity:** minor
**Проблема:** `time.Sleep(30 * time.Millisecond)` в main process перед SIGKILL — слишком мало для медленных CI машин. Subprocess должен успеть создать temp-файл и дойти до `time.Sleep(500ms)`. Если машина перегружена и subprocess не успевает за 30ms — SIGKILL придёт до создания temp-файла, тест пройдёт по ветке `os.IsNotExist(err)` и не проверит ничего содержательного.
**Рекомендация:** Subprocess должен сигнализировать готовность через файл-маркер в basedir. Main process ждёт маркера до SIGKILL, не фиксированное время.

---

### happy/fs_init_test.go — несоответствие logger-интерфейса между happy и adversarial

**Категория:** Поведение vs реализация
**Severity:** minor
**Проблема:** `testLogger` в `fs_init_test.go` реализует кастомный интерфейс `Log(level, msg string)`. AC-FS-INIT-03 требует поле `basedir=<path>` в структурированном логе. Тест проверяет `log.hasEntry("INFO", basedir)` — то есть что `basedir` содержится в `msg`. Но `slog.Info("store/fs initialized", "basedir", cfg.Basedir)` пишет `basedir` как атрибут, не в `msg`. Если реализация принимает `*slog.Logger` — кастомный `testLogger` не реализует тот же интерфейс. Adversarial-тест `TestFSInit_UncreatableBasedir_LogsErrorWithFields` использует slog+bytes.Buffer — там корректно.
**Рекомендация:** Унифицировать happy/fs_init_test.go на slog+bytes.Buffer как в adversarial, либо согласовать интерфейс логгера который принимает `fs.New`.
