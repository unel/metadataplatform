# Интерфейсы

В Go интерфейсы удовлетворяются неявно. Маленький интерфейс — мощный инструмент.

**Принцип:** интерфейс определяет потребитель, не реализация.

## Плохо — большой интерфейс на стороне реализации

```go
// пакет storage определяет огромный интерфейс
type Storage interface {
    FindEntity(id string) (*Entity, error)
    SaveEntity(e *Entity) error
    DeleteEntity(id string) error
    FindRelation(id string) (*Relation, error)
    SaveRelation(r *Relation) error
    // ... 20 методов
}
```

## Хорошо — маленький интерфейс на стороне потребителя

```go
// пакет job нужен только поиск entity — он и определяет интерфейс
type EntityFinder interface {
    FindEntity(ctx context.Context, id string) (*Entity, error)
}

func NewJobRunner(finder EntityFinder) *JobRunner { ... }
```

## Когда embedding оправдан

```go
// расширение стандартного интерфейса
type ReadWriter interface {
    io.Reader
    io.Writer
}

// НЕ для "наследования" реализации
type Base struct { ... }
type Child struct {
    Base // избегай если Child не является Base семантически
}
```
