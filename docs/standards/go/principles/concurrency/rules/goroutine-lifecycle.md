# Жизненный цикл горутин

Каждая горутина должна иметь явный способ завершения. Утечка горутин — частый источник утечки памяти.

```go
// плохо — горутина никогда не завершится
go func() {
    for {
        process()
    }
}()

// хорошо — завершение через context
func startWorker(ctx context.Context, jobs <-chan Job) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return // завершаемся при отмене context
            case job, ok := <-jobs:
                if !ok {
                    return // завершаемся при закрытии канала
                }
                process(job)
            }
        }
    }()
}

// errgroup — горутины с обработкой ошибок
g, ctx := errgroup.WithContext(ctx)
for _, job := range jobs {
    job := job // захват переменной
    g.Go(func() error {
        return process(ctx, job)
    })
}
if err := g.Wait(); err != nil {
    return err
}
```
