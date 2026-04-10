# $effect

`$effect` — для побочных эффектов: DOM-манипуляции, подписки на внешние системы, синхронизация с localStorage.

**Когда применять:** нужно что-то сделать при изменении состояния помимо обновления UI — запись в localStorage, управление фокусом, интеграция с внешней библиотекой.

**Когда не применять:** вычисление значений — это `$derived`. Синхронизация двух частей состояния — пересмотри модель данных.

## Плохо — $effect для вычислений

```svelte
<script lang="ts">
  let query = $state('')
  let results = $state<Result[]>([])

  $effect(() => {
    results = search(query) // должен быть $derived если search синхронный
  })
</script>
```

## Хорошо — $effect для реальных побочных эффектов

```svelte
<script lang="ts">
  let theme = $state<'light' | 'dark'>('light')

  $effect(() => {
    localStorage.setItem('theme', theme)
    document.documentElement.dataset.theme = theme
  })
</script>
```

## Cleanup

```svelte
<script lang="ts">
  let active = $state(false)

  $effect(() => {
    if (!active) return
    const handler = () => console.log('scroll')
    window.addEventListener('scroll', handler)
    return () => window.removeEventListener('scroll', handler) // cleanup
  })
</script>
```
