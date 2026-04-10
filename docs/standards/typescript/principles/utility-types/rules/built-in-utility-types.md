# Встроенные утилитарные типы

Список не исчерпывающий — TypeScript регулярно добавляет новые. Следи за релизами.

## Трансформация полей

```typescript
// Partial<T> — все поля опциональные
type UpdateUser = Partial<User>

// Required<T> — все поля обязательные
type FullConfig = Required<Config>

// Readonly<T> — все поля readonly
type ImmutableUser = Readonly<User>
```

## Выбор полей

```typescript
// Pick<T, K> — только указанные поля
type UserPreview = Pick<User, 'id' | 'name'>

// Omit<T, K> — все поля кроме указанных
type CreateUserDTO = Omit<User, 'id' | 'createdAt'>
```

## Типы из значений

```typescript
// Record<K, V> — объект с ключами типа K и значениями типа V
type StatusMap = Record<Status, string>

// ReturnType<T> — тип возвращаемого значения функции
type QueryResult = ReturnType<typeof buildQuery>

// Parameters<T> — тип параметров функции как кортеж
type QueryParams = Parameters<typeof buildQuery>
```

## Работа с промисами

```typescript
// Awaited<T> — тип результата промиса (разворачивает вложенные Promise)
type UserResult = Awaited<Promise<User>>           // User
type NestedResult = Awaited<Promise<Promise<User>>> // User

async function fetchUser(): Promise<User> { ... }
type FetchedUser = Awaited<ReturnType<typeof fetchUser>> // User
```

## Работа с объединениями

```typescript
// Extract<T, U> — типы из T которые совместимы с U
type StringOrNumber = string | number | boolean
type OnlyPrimitive = Extract<StringOrNumber, string | number> // string | number

// Exclude<T, U> — типы из T которые не совместимы с U
type WithoutBoolean = Exclude<StringOrNumber, boolean> // string | number

// NonNullable<T> — убирает null и undefined
type DefinedUser = NonNullable<User | null | undefined> // User
```
