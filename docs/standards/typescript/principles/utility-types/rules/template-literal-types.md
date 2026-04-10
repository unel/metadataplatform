# Template Literal Types

Строковые типы через шаблоны — позволяют строить типы из строковых комбинаций.

**Когда применять:** типизация строковых паттернов (event names, CSS классы, пути API). Генерация типов из строковых ключей.

```typescript
// базовый паттерн
type EventName = 'click' | 'focus' | 'blur'
type Handler = `on${Capitalize<EventName>}` // 'onClick' | 'onFocus' | 'onBlur'

// типизация путей API
type ApiPath = `/api/${string}`
function fetch(path: ApiPath): Promise<unknown> { ... }
fetch('/api/users')  // ok
fetch('/users')      // ошибка

// комбинации
type Axis = 'x' | 'y'
type Direction = 'top' | 'bottom' | 'left' | 'right'
type Margin = `margin-${Direction}` // 'margin-top' | 'margin-bottom' | ...
```

## Границы применимости

Template literal types легко становятся нечитаемыми при вложенности. Используй когда строковый паттерн действительно важен для типобезопасности.
