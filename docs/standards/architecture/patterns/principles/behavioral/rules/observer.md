# Observer

**Суть:** определить зависимость "один ко многим" между объектами так что при изменении одного все зависимые уведомляются автоматически.

**Когда применять:** изменение одного объекта требует изменения других, число которых неизвестно; объекты должны уведомлять других не зная кто они; событийные системы, реактивные UI.

**Когда не применять:** подписчики мало и они известны — прямой вызов проще. Глубокие цепочки событий трудно отлаживать.

## В классическом ООП

Subject хранит список Observer'ов и уведомляет их:

```
interface Observer:
    update(event)

class EventBus:
    subscribers = []

    subscribe(observer: Observer):
        subscribers.add(observer)

    publish(event):
        for s in subscribers: s.update(event)
```

## В объектном (Go-style)

Канал или callback-функции вместо интерфейса наблюдателя:

```
type Handler func(event Event)

struct EventBus:
    handlers []Handler

function (b *EventBus) subscribe(h Handler):
    b.handlers = append(b.handlers, h)

function (b *EventBus) publish(event Event):
    for _, h in b.handlers: h(event)
```

## В функциональном

Поток событий и подписка через функции:

```
function createEventBus():
    subscribers = []
    return {
        subscribe: (fn) -> subscribers.push(fn),
        publish:   (event) -> subscribers.forEach(fn -> fn(event))
    }

bus = createEventBus()
bus.subscribe((e) -> log(e))
bus.publish({ type: "created", id: "123" })
```

## В структурном

Массив функций-обработчиков, явный вызов при событии:

```
struct EventBus:
    handlers: [](event -> void)

function notify(bus, event):
    for handler in bus.handlers:
        handler(event)
```
