# $derived

`$derived` — вычисляемое значение которое автоматически пересчитывается при изменении зависимостей.

**Когда применять:** любое значение которое вычисляется из других реактивных данных.

**Когда не применять:** если вычисление имеет побочные эффекты — используй `$effect`. Если значение не зависит от реактивных данных — просто константа.

## Плохо — $effect для вычислений

```svelte
<script lang="ts">
  let items = $state<Item[]>([])
  let filtered = $state<Item[]>([])

  $effect(() => {
    filtered = items.filter(i => i.active) // $effect для вычисления — антипаттерн
  })
</script>
```

## Хорошо

```svelte
<script lang="ts">
  let items = $state<Item[]>([])
  let filtered = $derived(items.filter(i => i.active))
  let count = $derived(filtered.length)
</script>
```

## $derived.by для сложных вычислений

```svelte
<script lang="ts">
  let data = $state<Record<string, number>>({})

  let summary = $derived.by(() => {
    const values = Object.values(data)
    return {
      total: values.reduce((a, b) => a + b, 0),
      avg: values.length ? values.reduce((a, b) => a + b, 0) / values.length : 0
    }
  })
</script>
```
