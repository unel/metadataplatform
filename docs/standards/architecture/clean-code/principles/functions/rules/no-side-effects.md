# Без скрытых побочных эффектов

Функция делает ровно то что написано в её имени — и ничего больше.

**Почему:** скрытый побочный эффект нарушает контракт функции. Вызывающий код не ожидает изменений за пределами явного результата — это источник труднообнаруживаемых багов.

## Границы применимости

Функции с описательными именами (check, validate, get, find). Функции с императивными именами (save, send, process) — побочный эффект является их целью.

## Если соблюдать рьяно

Функции становятся чисто функциональными везде. В системах с состоянием (БД, файловая система, сеть) это приводит к неудобным архитектурным решениям.

## Если игнорировать

`checkPassword` инициализирует сессию. `validateOrder` списывает деньги. Тесты падают неожиданно. Функции нельзя вызвать повторно без побочных последствий.

## Когда можно отступить

Явные side-effect функции: `saveAndPublish`, `deleteAndNotify` — если имя честно описывает оба действия и они неразделимы по смыслу.

## Плохо

Имя говорит "проверь" — но функция ещё и меняет состояние:

```
function checkPassword(username, password) -> bool:
    valid = hash(password) == users[username].passwordHash
    if valid -> session.initialize(username)  // скрытый побочный эффект
    return valid
```

## Хорошо

Каждая функция делает ровно то что обещает имя:

```
function checkPassword(username, password) -> bool:
    return hash(password) == users[username].passwordHash

function login(username, password) -> error:
    if not checkPassword(username, password) -> return ErrInvalidCredentials
    return session.initialize(username)
```
