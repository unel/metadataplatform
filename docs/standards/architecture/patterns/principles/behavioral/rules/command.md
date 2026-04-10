# Command

**Суть:** инкапсулировать запрос как объект, позволяя параметризовать клиентов с разными запросами, ставить запросы в очередь, логировать и поддерживать отмену операций.

**Когда применять:** нужна очередь операций; операции должны поддерживать отмену (undo); операция должна выполняться позже или в другом контексте; аудит и логирование действий.

**Когда не применять:** операция простая и выполняется немедленно — прямой вызов функции проще и понятнее.

## В классическом ООП

Интерфейс команды с методом execute:

```
interface Command:
    execute()
    undo()

class DeleteEntityCommand implements Command:
    constructor(repo, entityId)
    execute(): repo.delete(entityId)
    undo():    repo.restore(entityId)

queue.add(new DeleteEntityCommand(repo, id))
queue.executeAll()
```

## В объектном (Go-style)

Структура с данными, функция выполнения:

```
struct DeleteCommand:
    repo     Repository
    entityID string

function (c DeleteCommand) execute() -> error:
    return c.repo.delete(c.entityID)

commands = []Command{ DeleteCommand{repo, id} }
for _, cmd in commands: cmd.execute()
```

## В функциональном

Команда — просто функция, очередь — срез функций:

```
function makeDeleteCommand(repo, id) -> (() -> error):
    return () -> repo.delete(id)

queue = [
    makeDeleteCommand(repo, "123"),
    makeDeleteCommand(repo, "456"),
]

for cmd in queue: cmd()
```

## В структурном

Структура с данными и явным диспетчером:

```
struct Command:
    type:    string
    payload: map

function dispatch(cmd Command, repo Repository):
    if cmd.type == "delete" -> repo.delete(cmd.payload["id"])
    if cmd.type == "create" -> repo.create(cmd.payload)
```
