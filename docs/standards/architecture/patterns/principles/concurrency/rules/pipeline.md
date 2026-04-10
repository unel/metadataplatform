# Pipeline

**Суть:** данные проходят через последовательность этапов обработки, каждый этап независим и работает конкурентно с соседними. Выход одного этапа — вход следующего.

**Когда применять:** последовательная обработка данных где каждый этап независим; стриминг больших объёмов данных без загрузки всего в память; ETL-процессы, обработка медиа.

**Когда не применять:** этапы сильно зависят друг от друга по состоянию. Один этап значительно медленнее остальных — он становится узким местом для всего пайплайна.

## Суть (парадигмонезависимо)

```
source -> [stage1] -> [stage2] -> [stage3] -> sink

каждый stage: читает из входного канала, пишет в выходной
```

## В Go-style (каналы)

```
function parse(raw chan []byte) chan Record:
    out = make(chan Record)
    go func():
        for data in raw: out <- parseRecord(data)
    return out

function enrich(records chan Record) chan Record:
    out = make(chan Record)
    go func():
        for r in records: out <- addMetadata(r)
    return out

// композиция:
raw = readFiles(paths)
parsed = parse(raw)
enriched = enrich(parsed)
store(enriched)
```

## В функциональном

```
function pipeline(...stages):
    return (source) -> stages.reduce((stream, stage) -> stage(stream), source)

process = pipeline(parse, enrich, validate, store)
process(readFiles(paths))
```
