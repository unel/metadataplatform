# Данные вместо логики

**Почему:** логика в коде — это код который нужно читать, понимать, тестировать, поддерживать. Данные в таблице — это данные которые можно просто смотреть. Добавление нового случая в таблицу не требует понимания логики, только добавления строки.

## Границы применимости

Когда управляющая структура выбирает значение, а не выполняет сложную разную логику. Если каждая ветка if делает принципиально разное — таблица не поможет.

## Если соблюдать рьяно

Не всякий switch стоит заменять таблицей. Два-три случая — switch читается лучше. Таблица оправдана когда случаев много или они будут добавляться.

## Если игнорировать

Бизнес-правила зарастают в if/else-деревья. Добавление нового типа — правка в нескольких местах. Тест на «какой worker у типа file.video» — почти невозможен без запуска всего.

## Direct access — ключ напрямую

```
// плохо — логика закодирована в switch
function workerFor(entityType, entitySubtype):
  if entityType == "file":
    if entitySubtype == "video": return "hash.video"
    if entitySubtype == "audio": return "hash.audio"
    if entitySubtype == "image": return "hash.image"
  if entityType == "job":
    return "job.runner"
  return "default"

// хорошо — таблица
WORKER_TABLE = {
  "file.video":  "hash.video",
  "file.audio":  "hash.audio",
  "file.image":  "hash.image",
  "job":         "job.runner",
}

function workerFor(entityType, entitySubtype):
  key = entityType + "." + entitySubtype
  return WORKER_TABLE.get(key) ?? WORKER_TABLE.get(entityType) ?? "default"
```

Добавление нового типа — одна строка в таблице.

## Staircase access — диапазоны

```
// плохо
function taxRate(income):
  if income < 10000:  return 0.0
  if income < 30000:  return 0.1
  if income < 70000:  return 0.2
  return 0.3

// хорошо — верхние границы диапазонов + ставки
TAX_BRACKETS = [
  {upTo: 10000,  rate: 0.0},
  {upTo: 30000,  rate: 0.1},
  {upTo: 70000,  rate: 0.2},
  {upTo: MAX_INT, rate: 0.3},
]

function taxRate(income):
  for bracket in TAX_BRACKETS:
    if income < bracket.upTo:
      return bracket.rate
```

Изменение ставок или границ — только данные, логика не трогается.

## В проекте

Spawner rules в YAML — это table-driven method на уровне конфигурации. Добавление нового правила спавна — строка в YAML, не правка кода spawner.
