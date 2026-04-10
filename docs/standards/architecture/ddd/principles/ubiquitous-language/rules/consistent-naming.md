# Согласованные имена

**Почему:** когда один концепт называется по-разному — возникает неявный translation layer в голове разработчика. `item` в API, `record` в БД, `entry` в бизнес-логике — это одно и то же, но каждый раз нужно думать. Ошибки появляются именно на этих границах перевода.

## Границы применимости

Доменные концепты — названия из бизнес-модели. Технические термины (`handler`, `middleware`, `config`) не доменные — для них ubiquitous language не применяется.

## Если соблюдать рьяно

Иногда контекст диктует другой термин — например, HTTP API говорит `resource`, SQL говорит `row`. Это нормально на самой границе адаптера. Нарушение — когда чужой термин просачивается внутрь доменной логики.

## Если игнорировать

Кодовая база становится вавилонской башней: `entity` в store, `item` в API handler, `record` в тестах — всё об одном. Поиск по коду не работает. Новый разработчик тратит время на выяснение что это одно и то же.

## Когда можно отступить

На границе с внешней системой где термин фиксирован (например, Stash API говорит `scene` — в адаптере допустимо, но внутрь проходит как `entity` с `subtype=stash_scene`).

## Плохо

```
// в store
type Entity struct { ... }
func (s *Store) FindEntity(id string) (*Entity, error)

// в API handler — тот же концепт, другое имя
type Item struct { ... }
func handleGetItem(w, r) {
  record := store.FindEntity(r.id)  // Entity → record
  item := toItem(record)            // record → Item
  json.Encode(item)
}

// в тесте
assert response.body.entry.type == "file"  // Item → entry
```

## Хорошо

```
// везде одно слово
type Entity struct { ... }

func (s *Store) FindEntity(id string) (*Entity, error)

func handleGetEntity(w, r) {
  entity := store.FindEntity(r.id)
  json.Encode(entity)
}

// тест
assert response.body.entity.type == "file"
```
