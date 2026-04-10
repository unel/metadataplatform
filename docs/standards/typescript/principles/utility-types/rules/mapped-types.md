# Mapped Types

`{ [K in keyof T]: ... }` — трансформация всех полей типа по одному шаблону.

**Когда применять:** нужно применить одну трансформацию ко всем полям типа. Встроенных Partial/Readonly/Required недостаточно.

```typescript
// базовый паттерн
type Nullable<T> = { [K in keyof T]: T[K] | null }
type NullableUser = Nullable<User>

// с переименованием ключей (as)
type Getters<T> = {
  [K in keyof T as `get${Capitalize<string & K>}`]: () => T[K]
}
type UserGetters = Getters<User>
// { getId: () => string, getName: () => string, ... }

// фильтрация полей через never
type OnlyStrings<T> = {
  [K in keyof T as T[K] extends string ? K : never]: T[K]
}
```

## Границы применимости

Mapped types мощны но снижают читаемость. Если трансформация применяется к одному конкретному типу — лучше написать тип явно.
