# Zero Values и инициализация

Go инициализирует переменные нулевыми значениями. Проектируй типы так чтобы нулевое значение было полезным.

```go
// хорошо — нулевое значение Buffer сразу готово к работе
var buf bytes.Buffer
buf.WriteString("hello")

// хорошо — нулевое значение Mutex не требует инициализации
var mu sync.Mutex
mu.Lock()

// плохо — требует явной инициализации
type Config struct {
    items map[string]string // нулевое значение nil — запись вызовет панику
}

// хорошо — конструктор гарантирует корректное начальное состояние
func NewConfig() *Config {
    return &Config{items: make(map[string]string)}
}
```
