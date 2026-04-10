# Barricade Pattern

**Почему:** смешивать валидацию внешних данных с обработкой внутренних инвариантов — значит не доверять собственному коду. Barricade разделяет эти два режима явно: за барьером защищаешься от внешнего мира, внутри — от собственных ошибок через assertions.

## Границы применимости

Любая система с внешним вводом. Барьер — это не обязательно один класс, это архитектурный слой.

## Если соблюдать рьяно

Чёткая граница барьера иногда неудобна — хочется добавить «немного» валидации внутри для удобства. Если это случается систематически — возможно барьер проведён не там.

## Если игнорировать

Код становится параноидальным или беспечным: либо проверки null везде до бесконечности, либо panic в неожиданных местах.

## Барьер в проекте

```
Внешний мир (HTTP, сокет, CLI)
          ↓
    [ BARRICADE ]         ← API handler, JSONL-парсер, flag-парсер
          ↓
  Валидированные данные   ← внутри: assertions, не error handling
          ↓
   Store, Spawner, Domain logic
```

## Assertions vs Error handling

```
// Error handling — за барьером, внешние данные
function handleRequest(raw):
  if raw.type == "":
    return Error(400, "type required")   // ожидаемая ситуация, обрабатываем

// Assertion — внутри барьера, внутренний инвариант
function processEntity(entity: Entity):
  assert entity.type != "", "processEntity called with empty type — bug in caller"
  // если это сработало — это баг разработчика, не пользователя
```

## Антипаттерн: assertion как error handling

```
// плохо: assertion для управления потоком
function findJob(id):
  job = db.find(id)
  assert job != null, "job not found"  // это нормальная ситуация, не баг!
  return job

// хорошо: error handling для ожидаемых случаев
function findJob(id):
  job = db.find(id)
  if job == null:
    return null, Error("job not found: " + id)
  return job, nil
```

## Антипаттерн: assertion без сообщения

```
// плохо: непонятно что нарушено
assert items != null

// хорошо: явно что за инвариант
assert items != null, "processItems: items slice must not be nil, check caller"
```
