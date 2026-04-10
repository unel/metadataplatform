# Form Actions

```typescript
// +page.server.ts
import { fail, redirect } from '@sveltejs/kit'
import type { Actions } from './$types'

export const actions: Actions = {
  create: async ({ request, locals }) => {
    const data = await request.formData()
    const name = data.get('name')

    if (!name || typeof name !== 'string') {
      return fail(400, { error: 'Name is required' })
    }

    const entity = await db.entities.create({ name })
    redirect(303, `/entities/${entity.id}`)
  }
}
```

```svelte
<!-- +page.svelte -->
<script lang="ts">
  import { enhance } from '$app/forms'
  import type { ActionData } from './$types'

  let { form }: { form: ActionData } = $props()
</script>

<form method="POST" action="?/create" use:enhance>
  <input name="name" />
  {#if form?.error}
    <p class="error">{form.error}</p>
  {/if}
  <button type="submit">Create</button>
</form>
```

## Когда использовать fetch вместо form actions

- Сложная интерактивность без перезагрузки данных страницы
- Загрузка файлов с прогрессом
- Realtime операции
