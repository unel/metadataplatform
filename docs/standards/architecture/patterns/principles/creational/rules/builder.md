# Builder

**Суть:** конструировать сложный объект шаг за шагом, отделяя процесс сборки от представления результата. Особенно полезен когда много опциональных параметров.

**Когда применять:** объект имеет много опциональных полей; разные конфигурации одного объекта; конструирование требует валидации или трансформаций.

**Когда не применять:** объект простой — достаточно struct literal или конструктора с парой аргументов.

## В классическом ООП

Цепочка методов на объекте-строителе:

```
class QueryBuilder:
    function where(condition) -> QueryBuilder: ...
    function limit(n) -> QueryBuilder: ...
    function orderBy(field) -> QueryBuilder: ...
    function build() -> Query: ...

query = new QueryBuilder()
    .where("status = 'active'")
    .limit(20)
    .orderBy("created_at")
    .build()
```

## В объектном (Go-style)

Функциональные опции или struct с методами:

```
struct QueryOptions:
    filter, limit, orderBy

function withFilter(f) -> (QueryOptions -> QueryOptions):
    return (opts) -> { ...opts, filter: f }

function withLimit(n) -> (QueryOptions -> QueryOptions):
    return (opts) -> { ...opts, limit: n }

query = buildQuery(data, withFilter("active"), withLimit(20))
```

## В функциональном

Композиция трансформеров:

```
function buildQuery(...transforms):
    opts = defaultOptions
    for transform in transforms:
        opts = transform(opts)
    return opts

query = buildQuery(
    where("status = 'active'"),
    limit(20),
    orderBy("created_at")
)
```

## В структурном

Явная структура конфига передаётся в функцию:

```
struct QueryConfig:
    filter = ""
    limit = 10
    orderBy = "id"

config = QueryConfig{
    filter: "status = 'active'",
    limit: 20,
    orderBy: "created_at"
}
result = executeQuery(data, config)
```
