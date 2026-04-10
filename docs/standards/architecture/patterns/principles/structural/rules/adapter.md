# Adapter

**Суть:** преобразовать интерфейс одного объекта в интерфейс который ожидает другой. Позволяет объектам с несовместимыми интерфейсами работать вместе.

**Когда применять:** интеграция с внешней библиотекой или сервисом с чужим интерфейсом; изоляция кода от деталей внешней системы; замена одной реализации на другую без изменения вызывающего кода.

**Когда не применять:** интерфейсы совместимы или легко унифицируются — адаптер лишний слой. Не используй адаптер чтобы скрыть плохой дизайн своего кода.

## В классическом ООП

Класс-обёртка реализует нужный интерфейс и делегирует к адаптируемому объекту:

```
interface Logger:
    log(level, msg)

class ThirdPartyLogger:
    write(msg)  // другой интерфейс

class LoggerAdapter implements Logger:
    constructor(thirdParty: ThirdPartyLogger)

    log(level, msg):
        thirdParty.write("[{level}] {msg}")
```

## В объектном (Go-style)

Структура реализует нужный интерфейс, внутри — адаптируемый объект:

```
interface Logger:
    log(level, msg)

struct LoggerAdapter:
    inner ThirdPartyLogger

function (a LoggerAdapter) log(level, msg):
    a.inner.write("[{level}] {msg}")
```

## В функциональном

Функция-обёртка приводит сигнатуру к нужному виду:

```
function adaptLogger(thirdPartyWrite) -> (level, msg -> void):
    return (level, msg) -> thirdPartyWrite("[{level}] {msg}")

log = adaptLogger(thirdParty.write)
log("error", "something failed")
```

## В структурном

Функция-переходник принимает данные в одном формате, вызывает другую в нужном:

```
function logViaThirdParty(level, msg):
    thirdPartyWrite("[{level}] {msg}")
```
