# Decorator

**Суть:** динамически добавлять объекту новое поведение, оборачивая его в объект с тем же интерфейсом. Альтернатива наследованию для расширения функциональности.

**Когда применять:** нужно добавить поведение (логирование, кэширование, авторизацию) не меняя исходный код; поведение должно добавляться и убираться динамически; несколько независимых расширений которые можно комбинировать.

**Когда не применять:** расширение одно и постоянное — проще встроить напрямую. Глубокая цепочка декораторов затрудняет отладку.

## В классическом ООП

Декоратор реализует тот же интерфейс и оборачивает исходный объект:

```
interface Repository:
    find(id) -> Entity

class CachingRepository implements Repository:
    constructor(inner: Repository, cache: Cache)

    find(id) -> Entity:
        if cache.has(id) -> return cache.get(id)
        result = inner.find(id)
        cache.set(id, result)
        return result
```

## В объектном (Go-style)

Структура реализует интерфейс и делегирует к внутреннему объекту:

```
interface Repository:
    find(id) -> Entity

struct CachingRepository:
    inner Repository
    cache Cache

function (r CachingRepository) find(id) -> Entity:
    if r.cache.has(id) -> return r.cache.get(id)
    result = r.inner.find(id)
    r.cache.set(id, result)
    return result
```

## В функциональном

Функция-обёртка добавляет поведение вокруг исходной функции:

```
function withCaching(find, cache) -> (id -> Entity):
    return (id) ->
        if cache.has(id) -> return cache.get(id)
        result = find(id)
        cache.set(id, result)
        return result

cachedFind = withCaching(repo.find, cache)
```

## В структурном

Функция-обёртка принимает исходную функцию и возвращает расширенную:

```
function cachingDecorator(findFn, cache):
    return function(id):
        cached = cache.get(id)
        if cached != nil -> return cached
        result = findFn(id)
        cache.set(id, result)
        return result
```
