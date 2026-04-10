# Альтернативы any

## unknown вместо any для неизвестных данных

```typescript
// плохо
function parse(data: any) {
  return data.name // нет проверки
}

// хорошо — unknown заставляет проверять перед использованием
function parse(data: unknown): string {
  if (typeof data === 'object' && data !== null && 'name' in data) {
    return String((data as { name: unknown }).name)
  }
  throw new Error('Invalid data')
}
```

## Type guards для сужения типов

```typescript
function isUser(value: unknown): value is User {
  return typeof value === 'object' && value !== null && 'id' in value
}

const data: unknown = fetchFromAPI()
if (isUser(data)) {
  console.log(data.id) // здесь data — User
}
```

## never для исчерпывающих проверок

```typescript
type Status = 'pending' | 'done' | 'failed'

function handle(status: Status) {
  switch (status) {
    case 'pending': return handlePending()
    case 'done': return handleDone()
    case 'failed': return handleFailed()
    default:
      const _exhaustive: never = status // ошибка если добавить новый статус без обработки
  }
}
```

## Когда any допустим

- Внешние библиотеки без типов (с комментарием)
- Миграция legacy кода (временно, с комментарием и задачей в roadmap)
- Низкоуровневые утилиты где тип действительно не важен

Всегда с `eslint-disable` комментарием и причиной.
