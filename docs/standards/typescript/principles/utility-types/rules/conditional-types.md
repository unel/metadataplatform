# Conditional Types

`T extends U ? X : Y` — тип зависит от условия. Мощный инструмент для type-level логики.

**Когда применять:** нужно разное поведение типа в зависимости от переданного типа-аргумента. Встроенных утилит недостаточно.

**Когда не применять:** если задача решается встроенными утилитами — не изобретай conditional type. Сложные вложенные условия — over-engineering.

```typescript
// простой случай — тип результата зависит от входного типа
type IsArray<T> = T extends any[] ? true : false
type A = IsArray<string[]> // true
type B = IsArray<string>   // false

// infer — извлечение типа из структуры
type UnpackArray<T> = T extends (infer Item)[] ? Item : T
type C = UnpackArray<string[]> // string
type D = UnpackArray<string>   // string

// распределение по union
type ToArray<T> = T extends any ? T[] : never
type E = ToArray<string | number> // string[] | number[]
```

## Границы применимости

Если conditional type сложнее трёх уровней вложенности — добавь комментарий объясняющий что происходит.
