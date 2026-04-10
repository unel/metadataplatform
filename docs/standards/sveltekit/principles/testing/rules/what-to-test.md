# Что тестировать

## Load функции

```typescript
// +page.server.test.ts
import { load } from './+page.server'
import { describe, it, expect, vi } from 'vitest'

describe('load', () => {
  it('возвращает entity по id', async () => {
    const mockDB = { entities: { find: vi.fn().mockResolvedValue({ id: '1', name: 'Test' }) } }
    const result = await load({ params: { id: '1' }, locals: { db: mockDB } } as any)
    expect(result.entity.name).toBe('Test')
  })

  it('бросает 404 если entity не найден', async () => {
    const mockDB = { entities: { find: vi.fn().mockResolvedValue(null) } }
    await expect(load({ params: { id: '999' }, locals: { db: mockDB } } as any))
      .rejects.toMatchObject({ status: 404 })
  })
})
```

## Компоненты

```typescript
import { render, screen } from '@testing-library/svelte'
import EntityCard from './EntityCard.svelte'

it('отображает название entity', () => {
  render(EntityCard, { props: { entity: { id: '1', name: 'Test Entity' } } })
  expect(screen.getByText('Test Entity')).toBeInTheDocument()
})
```

## Что НЕ тестировать отдельно

- Роутинг SvelteKit — это зона ответственности фреймворка
- Реактивность рун — тестируй поведение компонента, не механизм рун
