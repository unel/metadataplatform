# Bun Test API

Совместим с Jest — большинство Jest тестов работают без изменений.

```typescript
import { describe, it, expect, mock, beforeEach } from 'bun:test'

describe('EntityService', () => {
  let service: EntityService
  let mockStore: ReturnType<typeof mock>

  beforeEach(() => {
    mockStore = mock(() => Promise.resolve({ id: '123', type: 'file' }))
    service = new EntityService({ find: mockStore })
  })

  it('возвращает entity по id', async () => {
    const entity = await service.find('123')
    expect(entity.id).toBe('123')
    expect(mockStore).toHaveBeenCalledWith('123')
  })

  it('бросает если entity не найдена', async () => {
    mockStore.mockImplementation(() => Promise.resolve(null))
    expect(service.find('999')).rejects.toThrow('not found')
  })
})
```

## Отличия от Jest

- `mock()` вместо `jest.fn()`
- `spyOn()` из `bun:test`
- Нет `jest.mock()` для модулей — используй dependency injection
- `--watch` поддерживается: `bun test --watch`

## Coverage

```bash
bun test --coverage
# или с порогом
bun test --coverage --coverage-threshold 80
```
