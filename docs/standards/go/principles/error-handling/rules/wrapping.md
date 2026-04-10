# Оборачивание ошибок

```go
// оборачивай с контекстом через %w — сохраняет цепочку для errors.Is/As
if err != nil {
    return fmt.Errorf("find entity %s: %w", id, err)
}

// проверка конкретной ошибки
if errors.Is(err, sql.ErrNoRows) {
    return nil, ErrNotFound
}

// sentinel errors — на уровне пакета
var (
    ErrNotFound   = errors.New("not found")
    ErrConflict   = errors.New("conflict")
)

// typed error — когда нужна дополнительная информация
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error: %s — %s", e.Field, e.Message)
}

// извлечение typed error
var valErr *ValidationError
if errors.As(err, &valErr) {
    // работаем с valErr.Field, valErr.Message
}
```
