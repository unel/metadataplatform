# Fan-out / Fan-in

**Суть:** Fan-out — распределить одну задачу между несколькими параллельными воркерами. Fan-in — собрать результаты от нескольких источников в один поток.

**Когда применять:** задача легко делится на независимые части; нужно использовать несколько CPU-ядер; агрегация результатов из нескольких источников (несколько БД, API).

**Когда не применять:** задачи не независимы — между ними есть общее состояние. Накладные расходы на синхронизацию превышают выигрыш от параллелизма.

## Суть (парадигмонезависимо)

```
// fan-out: одна задача → N воркеров
task -> [worker1, worker2, worker3]

// fan-in: N результатов → один поток
[result1, result2, result3] -> merged stream
```

## В Go-style (каналы)

```
// fan-out
function fanOut(input chan Task, n int) []chan Result:
    outputs = make([]chan Result, n)
    for i = 0; i < n; i++:
        outputs[i] = startWorker(input)
    return outputs

// fan-in
function fanIn(channels ...chan Result) chan Result:
    merged = make(chan Result)
    for _, ch in channels:
        go func(c chan Result):
            for r in c: merged <- r
    return merged

results = fanIn(fanOut(tasks, 4)...)
```

## В функциональном (промисы)

```
// fan-out + fan-in через Promise.all
function parallel(tasks, processFn):
    return Promise.all(tasks.map(processFn))

results = await parallel(items, processItem)
```
