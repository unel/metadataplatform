# Вынос логики из компонента

Сложная логика не должна жить в .svelte файле. Выноси в отдельные модули.

**Когда выносить:** бизнес-логика; сложные вычисления; переиспользуемое состояние между компонентами.

## Паттерн: rune-функция (createXxx)

```typescript
// lib/search.svelte.ts
export function createSearch(initialQuery = '') {
  let query = $state(initialQuery)
  let results = $state<Result[]>([])
  let loading = $state(false)

  async function search() {
    loading = true
    results = await fetchResults(query)
    loading = false
  }

  return { 
    get query() { return query },
    set query(v) { query = v },
    get results() { return results },
    get loading() { return loading },
    search
  }
}
```

```svelte
<script lang="ts">
  import { createSearch } from '$lib/search.svelte'
  const search = createSearch()
</script>
```
