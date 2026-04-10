# Facade

**Суть:** предоставить простой интерфейс к сложной подсистеме. Скрыть детали взаимодействия с несколькими компонентами за единой точкой входа.

**Когда применять:** подсистема сложная и вызывающий код не должен знать детали; нужно упростить типичные сценарии использования; изоляция вызывающего кода от изменений внутри подсистемы.

**Когда не применять:** фасад становится god-объектом с сотней методов — лучше разбить на несколько специализированных интерфейсов. Не прячь за фасадом то что должно быть доступно напрямую.

## В классическом ООП

Класс координирует несколько подсистем:

```
class OrderFacade:
    constructor(inventory, payment, notification)

    placeOrder(order):
        inventory.reserve(order.items)
        payment.charge(order.total)
        notification.send(order.userId, "Order placed")
```

## В объектном (Go-style)

Структура с зависимостями, метод координирует несколько сервисов:

```
struct OrderService:
    inventory InventoryService
    payment   PaymentService
    notify    NotificationService

function (s OrderService) placeOrder(order) -> error:
    if err = s.inventory.reserve(order.items); err != nil -> return err
    if err = s.payment.charge(order.total); err != nil -> return err
    return s.notify.send(order.userID, "Order placed")
```

## В функциональном

Функция высшего порядка принимает зависимости и возвращает скомпонованную операцию:

```
function makePlaceOrder(reserve, charge, notify):
    return (order) ->
        reserve(order.items)
        charge(order.total)
        notify(order.userId, "Order placed")

placeOrder = makePlaceOrder(inventory.reserve, payment.charge, notification.send)
```

## В структурном

Функция принимает все зависимости явно:

```
function placeOrder(order, inventory, payment, notification):
    inventory.reserve(order.items)
    payment.charge(order.total)
    notification.send(order.userId, "Order placed")
```
