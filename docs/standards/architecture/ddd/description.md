# Domain-Driven Design

Принципы DDD применимые к проекту.

## Ubiquitous Language
Единый словарь терминов используемый везде — в коде, документации, разговорах.
В этом проекте: `entity`, `relation`, `job`, `projection`, `worker`, `hook`, `spawner`.
Признак нарушения: один и тот же concept называется по-разному в разных местах.

## Bounded Context
Чёткие границы между компонентами — каждый компонент знает только то что ему нужно.
В проекте: `store` не знает о `spawner`, `worker` не знает о БД напрямую.
Признак нарушения: компонент импортирует или вызывает то что лежит за его границей.

## Layered Architecture
Зависимости идут только в одном направлении: domain → store → api.
Признак нарушения: нижний слой знает о верхнем.

## Repository Pattern
Store — единая точка доступа к данным, скрывает детали хранилища.
Признак нарушения: компонент работает с БД напрямую минуя store.

## Источники

- Поиск: "DDD principles examples", "domain driven design bounded context"
- context7: `resolve library domain driven design`
