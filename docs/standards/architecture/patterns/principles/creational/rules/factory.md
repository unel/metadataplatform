# Factory

**Суть:** скрыть логику выбора и создания конкретной реализации за единой точкой входа. Вызывающий код знает что хочет получить — но не знает как это создаётся.

**Когда применять:** конкретный тип определяется в runtime; конструирование сложное или требует внешних зависимостей; нужно изолировать вызывающий код от деталей создания.

**Когда не применять:** объект создаётся одним способом и это не изменится — фабрика добавляет сложность без пользы.

## В классическом ООП

Метод или класс возвращает экземпляр через интерфейс:

```
interface Logger { log(msg) }
class FileLogger implements Logger { ... }
class StdoutLogger implements Logger { ... }

class LoggerFactory:
    static create(type) -> Logger:
        if type == "file" -> return new FileLogger()
        if type == "stdout" -> return new StdoutLogger()
```

## В объектном (Go-style)

Функция возвращает значение через интерфейс, конкретный тип скрыт:

```
interface Logger:
    log(msg)

function newLogger(type) -> Logger:
    if type == "file" -> return FileLogger{...}
    if type == "stdout" -> return StdoutLogger{...}
```

## В функциональном

Фабрика — функция возвращающая функцию:

```
function makeLogger(type) -> (msg -> void):
    if type == "file" -> return (msg) -> writeToFile(msg)
    if type == "stdout" -> return (msg) -> print(msg)

log = makeLogger("file")
log("hello")
```

## В структурном

Функция-конструктор возвращает структуру с явными указателями на функции:

```
struct Logger:
    logFn: (msg -> void)

function newLogger(type) -> Logger:
    if type == "file" -> return Logger{ logFn: writeToFile }
    if type == "stdout" -> return Logger{ logFn: print }
```
