# Заменяемость подтипов

**Почему:** если подтип нельзя подставить вместо базового без изменения поведения — абстракция ненастоящая. Пользователи абстракции вынуждены знать о подтипах и делать специальные случаи.

## Границы применимости

Любые отношения наследования или реализации интерфейса. Особенно критично для публичных API и точек расширения.

## Если соблюдать рьяно

Иногда полезное специализированное поведение подавляется ради «чистого» соблюдения принципа. Баланс: не добавляй поведение нарушающее контракт, но специализированное поведение сверх контракта — допустимо.

## Если игнорировать

Код с `isinstance` / `type switch` по всей кодовой базе. Новый подтип требует правки во всех этих местах. Абстракция перестаёт работать.

## Когда можно отступить

Если наследование используется для переиспользования кода, а не для полиморфизма — нарушение LSP менее критично, но стоит рассмотреть композицию вместо наследования.

## Плохо

`ReadonlyStorage` реализует `Storage` но нарушает контракт метода `save`:

```
interface Storage:
  save(entity: Entity) -> void
  find(id: ID) -> Entity

class ReadonlyStorage implements Storage:
  save(entity):
    throw NotSupportedError("read-only!")  // нарушает контракт

  find(id):
    return db.find(id)

// пользователь вынужден проверять тип:
function process(storage: Storage, entity):
  if storage instanceof ReadonlyStorage:
    // не сохраняем
  else:
    storage.save(entity)
```

## Хорошо

Разделить интерфейсы по реальным возможностям:

```
interface Reader:
  find(id: ID) -> Entity

interface Writer:
  save(entity: Entity) -> void

interface ReadWriter extends Reader, Writer: ...

class ReadonlyStorage implements Reader:
  find(id): return db.find(id)
  // save не объявлен — нет соблазна его нарушить

class FullStorage implements ReadWriter:
  save(entity): db.save(entity)
  find(id): return db.find(id)
```

Функции принимают только то что им нужно — `Reader` или `Writer`. Никаких проверок типов.
