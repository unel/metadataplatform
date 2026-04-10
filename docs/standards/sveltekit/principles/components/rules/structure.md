# Структура компонента

Порядок секций в .svelte файле: script → разметка → стили.

```svelte
<script lang="ts">
  // 1. импорты
  import { SomeComponent } from '$lib/components'

  // 2. типы props
  interface Props {
    title: string
    items: Item[]
  }

  // 3. props
  let { title, items }: Props = $props()

  // 4. локальное состояние
  let selected = $state<string | null>(null)

  // 5. вычисляемые значения
  let hasItems = $derived(items.length > 0)

  // 6. функции-обработчики
  function handleSelect(id: string) {
    selected = id
  }
</script>

<!-- разметка -->
<div class="container">
  <h1>{title}</h1>
  {#if hasItems}
    {#each items as item}
      <button onclick={() => handleSelect(item.id)}>{item.name}</button>
    {/each}
  {/if}
</div>

<style>
  .container { ... }
</style>
```
