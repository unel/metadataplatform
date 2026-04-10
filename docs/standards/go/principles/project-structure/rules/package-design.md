# Дизайн пакетов

```
cmd/
  store/
    main.go         — точка входа, только инициализация и запуск
  spawner/
    main.go
  api/
    main.go

internal/
  store/            — логика store, недоступна снаружи модуля
    handler.go
    query.go
    socket.go
  spawner/
    runner.go
    rules.go
  entity/           — доменные типы
    entity.go
    relation.go
    job.go
```

## Принципы именования пакетов

```go
// хорошо — одно слово, по содержимому
package store
package entity
package query

// плохо — многословно, повторяет путь
package storepackage
package entitymodel
package queryutil

// использование — имя пакета должно дополнять имя типа
store.Handler   // не store.StoreHandler
entity.Entity   // допустимо если в пакете много типов
query.Builder   // не query.QueryBuilder
```

## internal/ vs pkg/

- `internal/` — код специфичный для этого приложения, Go запрещает импорт снаружи модуля
- `pkg/` — переиспользуемый код который может быть импортирован другими модулями
- В этом проекте — всё в `internal/`, публичных библиотек нет
