# Правило: `// inv:` комментарии

## Зачем

CR-7 переноса cancel() был сделан как косметический — "убрать дублирование". CR-9 нашёл дедлок: после переноса cancel() вызывался после блокирующего `<-activeConns`. Инвариант "cancel() должен вызываться до любого блокирующего ожидания в shutdown-ветках" нигде не был зафиксирован — ни в тестах, ни в комментариях.

`// inv:` — это машиночитаемый маркер инварианта. Он:
- делает неявное явным
- выживает после рефакторинга (его сложно случайно убрать)
- даёт ревьюверу точку входа: "если трогаешь этот блок — прочитай инвариант"

## Формат

```go
// inv: <утверждение> — <объяснение если неочевидно>
```

Утверждение — краткое, императивное, без "должен быть" если можно без него.

## Примеры

### Shutdown path: cancel() до блокирующего ожидания

```go
func (s *Server) shutdown(ctx context.Context, done <-chan struct{}, activeConns <-chan struct{}) error {
    // inv: cancel() must run before any blocking <-done in shutdown branches
    // to ensure context propagation reaches active goroutines
    select {
    case <-done:
        cancel()
        return nil
    case <-time.After(timeout):
        cancel() // must precede <-activeConns
        <-activeConns
        return ErrShutdownTimeout
    }
}
```

### Мьютекс: порядок захвата

```go
// inv: mu must be held when reading or writing s.active
// callers must not hold connMu when acquiring mu — lock order: mu → connMu
func (s *Server) setActive(v bool) {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.active = v
}
```

### Идемпотентность

```go
// inv: Close() is idempotent — safe to call multiple times
// second call returns ErrAlreadyClosed, does not panic
func (c *Conn) Close() error {
    c.once.Do(func() {
        close(c.done)
    })
    ...
}
```

### WaitGroup: Add до запуска

```go
// inv: wg.Add(1) must be called before go func() — not inside
// otherwise wg.Wait() may return before goroutine starts
wg.Add(1)
go func() {
    defer wg.Done()
    worker(ctx)
}()
```

## Где ставить

- Перед блоком кода (select, if, for), к которому инвариант относится
- Не на уровне файла — только локально, рядом с применением
- Если инвариант распространяется на несколько функций — продублировать у каждой точки входа

## Что не является `// inv:`

- Объяснение что делает код (`// inv:` — это про порядок и предположения, не про логику)
- TODO или FIXME
- Документация публичного API — для этого godoc

## Правило ревью

Ревьювер, видя изменение в lifecycle/concurrency/shutdown блоке с `// inv:`, **обязан** сформулировать: как фикс сохраняет (или нарушает) инвариант. Это не опционально — без этого замечание не считается полным.
