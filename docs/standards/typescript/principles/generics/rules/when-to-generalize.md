# Когда обобщать

## Плохо — generic без необходимости

```typescript
// T всегда будет string — зачем generic?
function wrap<T>(value: T): { value: T } {
  return { value }
}
```

## Хорошо — логика действительно типонезависима

```typescript
// работает для любого типа одинаково
function first<T>(arr: T[]): T | undefined {
  return arr[0]
}

// constraint — T должен иметь id
function findById<T extends { id: string }>(items: T[], id: string): T | undefined {
  return items.find(item => item.id === id)
}
```

## Именование параметров типа

```typescript
// плохо — непонятно что T и U
function merge<T, U>(a: T, b: U): T & U

// хорошо — говорящие имена
function merge<TBase, TExtension>(base: TBase, ext: TExtension): TBase & TExtension

// исключение — однобуквенные стандартны для простых случаев:
// T — тип элемента, K — ключ, V — значение, E — ошибка
```

## Границы применимости

Generic оправдан когда функция или тип используются с разными типами на практике. Если за всё время жизни кода передаётся только один тип — убери generic.
