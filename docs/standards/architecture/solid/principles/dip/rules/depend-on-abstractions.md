# Зависи от абстракций

**Почему:** когда высокоуровневый модуль жёстко связан с конкретной реализацией, замена реализации (другая БД, другой транспорт, другой алгоритм) ломает весь модуль. Зависимость от абстракции — это зависимость от контракта, реализация может меняться свободно.

## Границы применимости

Межмодульные зависимости, особенно пересекающие архитектурные слои (бизнес-логика → инфраструктура). Внутри одного слоя жёсткие зависимости допустимы если они не мешают тестированию.

## Если соблюдать рьяно

Каждая зависимость превращается в интерфейс. Logger, Config, UUID generator — всё за интерфейсом. Оверинжиниринг в большинстве случаев.

## Если игнорировать

Бизнес-логика содержит SQL, HTTP-вызовы, обращения к файловой системе. Юнит-тесты невозможны без поднятия всей инфраструктуры. Замена PostgreSQL на что-то другое — переписывание половины кода.

## Когда можно отступить

Утилитарные зависимости без сайд-эффектов (математика, строковые операции, константы) — абстрагировать не нужно. Конкретная реализация без альтернатив и без нужды в тестировании изолированно — тоже.

## Плохо

Высокоуровневый модуль создаёт низкоуровневые объекты сам:

```
class JobRunner:
  function run(jobId):
    // жёстко зависит от конкретных реализаций
    db = PostgresDB("postgres://localhost/platform")
    store = EntityStore(db)
    job = store.findJob(jobId)

    logger = FileLogger("/var/log/runner.log")
    logger.info("running job", jobId)

    // бизнес-логика перемешана с инфраструктурой
    ...
```

Замена БД → правим `JobRunner`. Тест → нужна реальная БД и файловая система.

## Хорошо

Зависимости инжектируются через конструктор:

```
interface JobStore:
  findJob(id) -> Job
  updateJob(job) -> void

interface Logger:
  info(message, ...args) -> void

class JobRunner:
  constructor(store: JobStore, logger: Logger):
    self.store = store
    self.logger = logger

  function run(jobId):
    job = self.store.findJob(jobId)
    self.logger.info("running job", jobId)
    ...

// сборка — только в точке входа (main)
db = PostgresDB(config.dbUrl)
store = StoreClient(db)
logger = FileLogger(config.logPath)
runner = JobRunner(store, logger)

// тест — подставляем заглушки
runner = JobRunner(FakeStore(), NullLogger())
```

В проекте: `cmd/api/`, `cmd/spawner/` — точки сборки. Бизнес-логика в `internal/` зависит только от интерфейсов, конкретные реализации (store socket client, postgres) инжектируются в `main`.
