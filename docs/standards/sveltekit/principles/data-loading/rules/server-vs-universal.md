# Server load vs Universal load

## +page.server.ts — server-only

Выполняется только на сервере. Имеет доступ к БД, секретам, cookies.

```typescript
// +page.server.ts
import type { PageServerLoad } from './$types'
import { error } from '@sveltejs/kit'

export const load: PageServerLoad = async ({ params, locals }) => {
  const entity = await db.entities.find(params.id)
  if (!entity) error(404, 'Entity not found')
  return { entity }
}
```

**Используй когда:** нужен доступ к БД или секретам; работа с cookies/сессией; form actions.

## +page.ts — universal

Выполняется на сервере при SSR и на клиенте при навигации. Нет доступа к БД напрямую.

```typescript
// +page.ts
import type { PageLoad } from './$types'

export const load: PageLoad = async ({ fetch, params }) => {
  const res = await fetch(`/api/entities/${params.id}`)
  if (!res.ok) error(res.status, 'Failed to load')
  return { entity: await res.json() }
}
```

**Используй когда:** данные из публичного API; нужен доступ к данным и на клиенте при навигации.
