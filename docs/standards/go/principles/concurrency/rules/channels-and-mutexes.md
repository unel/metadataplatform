# Каналы и мьютексы

```go
// канал закрывает отправитель
func producer(jobs chan<- Job) {
    defer close(jobs) // отправитель закрывает
    for _, j := range work {
        jobs <- j
    }
}

// мьютекс — защита структуры с разделяемым состоянием
type Cache struct {
    mu    sync.RWMutex
    items map[string]Item
}

func (c *Cache) Get(key string) (Item, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    item, ok := c.items[key]
    return item, ok
}

func (c *Cache) Set(key string, item Item) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.items[key] = item
}

// sync.Once — инициализация один раз
var (
    instance *Service
    once     sync.Once
)

func getInstance() *Service {
    once.Do(func() {
        instance = &Service{}
    })
    return instance
}
```
