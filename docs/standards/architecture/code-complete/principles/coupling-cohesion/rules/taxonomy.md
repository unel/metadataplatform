# Таксономия cohesion и coupling

**Почему:** без конкретной шкалы сложно оценить качество модуля. «Плохое зацепление» — расплывчато. «Это semantic coupling» — конкретно и указывает что именно исправить.

## Cohesion — уровни с примерами

**Functional** (лучший) — один модуль, одна задача:
```
module HashCalculator:
  function sha256(data: bytes) -> string
  // всё внутри служит одной цели
```

**Sequential** — выход одного шага — вход другого:
```
module FileProcessor:
  function read(path) -> bytes
  function parse(bytes) -> Entity
  function validate(entity) -> Result
  // шаги связаны данными, порядок важен
```

**Communicational** — операции над одними данными, порядок не важен:
```
module EntityStats:
  function count(entities) -> int
  function averageSize(entities) -> float
  function oldestCreatedAt(entities) -> time
```

**Temporal** — выполняются в одно время, данные не связаны:
```
module AppInit:
  function initLogger()
  function initDB()
  function initCache()
  // связаны только временем запуска
```

**Logical** (плохой) — похожие операции выбираются флагом:
```
// признак: параметр-флаг определяет что делает функция
function processEntity(entity, mode: "hash" | "scan" | "enrich"):
  if mode == "hash": ...
  if mode == "scan": ...
  // это три разные функции, не одна
```

**Coincidental** (худший) — случайный набор:
```
module Utils:
  function formatDate()
  function connectDB()
  function sha256()
  function renderTemplate()
  // ничего общего
```

## Coupling — уровни с примерами

**Simple-data-parameter** (лучший) — примитивные данные:
```
function hashFile(path: string, algorithm: string) -> string
```

**Simple-object** — передаётся объект одного типа:
```
function processJob(job: Job) -> Result
```

**Object-parameter** — передаётся объект, используются его поля:
```
function processContext(ctx: Context) -> Result
// Context содержит много данных, используем несколько
```

**Semantic** (худший) — зависимость от внутренней семантики чужого модуля:
```
// spawner знает что store обрабатывает команды в порядке FIFO
// и намеренно отправляет upsert перед query чтобы "увидеть" свежие данные
spawner.storeClient.send(upsert)
spawner.storeClient.send(query)  // полагаемся на внутреннее поведение store
// изменение store сломает spawner без изменения интерфейса
```
