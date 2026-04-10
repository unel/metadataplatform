# Worker Pool

**Суть:** ограниченный набор воркеров обрабатывает очередь задач. Контролирует параллелизм — не создаёт горутину на каждую задачу.

**Когда применять:** задачи поступают быстрее чем обрабатываются; нужно ограничить нагрузку на внешние ресурсы (БД, сеть, CPU); spawner с лимитом параллельных воркеров.

**Когда не применять:** задач мало и они короткие — overhead пула не оправдан. Задачи имеют сильно разное время выполнения — нужна более умная балансировка.

## Суть (парадигмонезависимо)

```
tasks = queue of work items
workers = pool of N concurrent executors

for each task in tasks:
    wait for free worker
    worker.process(task)
```

## В Go-style (каналы)

```
function workerPool(numWorkers, tasks chan Task) chan Result:
    results = make(chan Result)

    for i = 0; i < numWorkers; i++:
        go func():
            for task in tasks:
                results <- process(task)

    return results
```

## В функциональном (промисы)

```
function workerPool(tasks, concurrency, processFn):
    semaphore = new Semaphore(concurrency)
    return Promise.all(
        tasks.map(task ->
            semaphore.acquire()
                .then(() -> processFn(task))
                .finally(() -> semaphore.release())
        )
    )
```
