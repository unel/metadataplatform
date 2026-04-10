# Singleton

**Суть:** гарантировать что у типа есть только один экземпляр, и предоставить глобальную точку доступа к нему.

**Когда применять:** действительно глобальное состояние (connection pool, конфиг приложения, registry); создание дорогое и должно происходить один раз.

**Когда не применять:** почти всегда есть лучшая альтернатива — dependency injection. Синглтон затрудняет тестирование и скрывает зависимости.

## В классическом ООП

Статический метод возвращает единственный экземпляр:

```
class Database:
    static instance = nil

    static getInstance() -> Database:
        if instance == nil -> instance = new Database()
        return instance
```

## В объектном (Go-style)

Пакетная переменная инициализируется один раз (sync.Once или init):

```
var db *Database
var once sync.Once

function getDB() -> *Database:
    once.Do(() -> db = newDatabase())
    return db
```

## В функциональном

Замыкание хранит единственный экземпляр:

```
getDB = (() ->
    db = createDatabase()
    return () -> db
)()

// использование
db = getDB()
```

## В структурном

Глобальная переменная, инициализируемая явно при старте:

```
var GlobalDB *Database

function initDB(config):
    GlobalDB = openDatabase(config)

// в main:
initDB(config)
// везде остальном:
GlobalDB.query(...)
```

## Предпочтительная альтернатива

Передавай зависимость явно через аргументы или конструктор — это проще тестировать и понимать:

```
function handleRequest(db Database, r Request): ...
```
