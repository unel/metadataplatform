# Panic и recover

Panic — для невосстановимых ошибок программиста, не для бизнес-ошибок.

```go
// допустимо — нарушение инварианта при инициализации
func NewServer(cfg Config) *Server {
    if cfg.Port == 0 {
        panic("server: port must be set") // программист забыл настроить
    }
    return &Server{cfg: cfg}
}

// допустимо — в main при старте
func main() {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err) // не panic, но тоже завершает программу
    }
}

// недопустимо — в библиотечном коде при runtime ошибках
func FindEntity(id string) *Entity {
    entity, err := db.Query(...)
    if err != nil {
        panic(err) // плохо — вызывающий код не может обработать
    }
    return entity
}
```
