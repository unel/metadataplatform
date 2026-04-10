# Context

`context.Context` — первый аргумент любой функции которая делает I/O или может быть отменена.

```go
// всегда первым аргументом, всегда называется ctx
func FindEntity(ctx context.Context, id string) (*Entity, error) { ... }

// не храни context в структуре
type Service struct {
    ctx context.Context // плохо
}

// передавай явно в каждый вызов
func (s *Service) Process(ctx context.Context, id string) error {
    entity, err := s.store.FindEntity(ctx, id)
    ...
}

// отмена и таймауты
ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
defer cancel() // всегда вызывай cancel

// значения в context — только для cross-cutting данных (request-id, trace)
// не для бизнес-параметров
```
