# Однонаправленные зависимости

**Почему:** циклические зависимости между слоями означают что их нельзя тестировать и деплоить независимо. Изменение в одном месте волнами распространяется во все стороны. Нижний слой знающий о верхнем — это нарушение инкапсуляции.

## Границы применимости

Архитектурные слои внутри одного бинаря и между бинарями системы. Внутри одного слоя взаимные зависимости допустимы.

## Если соблюдать рьяно

Иногда удобная маленькая зависимость «снизу вверх» убирается слишком дорогой абстракцией. Оцени стоимость — если нарушение локальное и изолированное, иногда прагматично допустить.

## Если игнорировать

Store начинает вызывать spawner напрямую для триггера джобов. Spawner вызывает API для получения конфига. Всё зависит от всего — impossible to test, impossible to change.

## Когда можно отступить

Event-driven архитектура размывает слои: нижний слой эмитит событие, верхний слушает. Это не нарушение — зависимость через broker, не напрямую.

## Слои проекта

```
workers / hooks          ← знают только свой контракт (CLI args + stdout)
       ↑
spawner / api / platform ← знают о store, не знают о конкретных workers
       ↑
     store               ← знает только о БД, не знает ни о чём выше
       ↑
  PostgreSQL
```

## Плохо — store вызывает spawner

```
// store/entity.go
func (s *Store) CreateEntity(entity Entity) error {
  err := s.db.Insert(entity)
  if err != nil { return err }

  // store не должен знать о spawner
  s.spawnerClient.Trigger("entity.created", entity.ID)
  return nil
}
```

Store теперь зависит от spawner → нельзя запустить store без spawner → нельзя тестировать store изолированно.

## Хорошо — ответственность на верхнем слое

```
// api/handler.go — верхний слой знает об обоих
func handleCreateEntity(w, r) {
  entity := parseBody(r)

  // сначала сохраняем
  err := store.CreateEntity(entity)
  if err != nil { ... }

  // потом триггерим — на уровне где знают о spawner
  spawner.Trigger("entity.created", entity.ID)
}

// store/entity.go — чистый, без зависимостей на spawner
func (s *Store) CreateEntity(entity Entity) error {
  return s.db.Insert(entity)
}
```
