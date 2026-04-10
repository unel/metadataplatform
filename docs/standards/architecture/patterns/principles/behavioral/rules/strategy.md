# Strategy

**Суть:** определить семейство алгоритмов, инкапсулировать каждый и сделать их взаимозаменяемыми. Позволяет менять алгоритм независимо от клиентов которые его используют.

**Когда применять:** несколько вариантов одного алгоритма; алгоритм выбирается в runtime; нужно изолировать бизнес-логику от деталей реализации алгоритма.

**Когда не применять:** алгоритм один и не меняется — стратегия добавляет лишнюю индирекцию. Не используй если разница между стратегиями — один параметр.

## В классическом ООП

Интерфейс стратегии, конкретные реализации, контекст принимает стратегию:

```
interface SortStrategy:
    sort(data) -> data

class QuickSort implements SortStrategy: ...
class MergeSort implements SortStrategy: ...

class Sorter:
    constructor(strategy: SortStrategy)
    sort(data): return strategy.sort(data)
```

## В объектном (Go-style)

Интерфейс + структура с полем-стратегией:

```
interface Sorter:
    sort(data) -> data

struct DataProcessor:
    sorter Sorter

function (p DataProcessor) process(data):
    return p.sorter.sort(data)
```

## В функциональном

Функция высшего порядка принимает алгоритм как аргумент:

```
function process(data, sortFn):
    return sortFn(data)

process(data, quickSort)
process(data, mergeSort)
```

## В структурном

Указатель на функцию в структуре конфига:

```
struct ProcessorConfig:
    sortFn: (data -> data)

function process(data, config):
    return config.sortFn(data)
```
