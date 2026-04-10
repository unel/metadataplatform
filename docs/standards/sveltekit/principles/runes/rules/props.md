# $props

`$props` — декларация входных параметров компонента. Всегда типизировать явно.

## Плохо

```svelte
<script lang="ts">
  let { name, onClick } = $props() // нет типов
</script>
```

## Хорошо

```svelte
<script lang="ts">
  interface Props {
    name: string
    count?: number
    onClick: (id: string) => void
  }

  let { name, count = 0, onClick }: Props = $props()
</script>
```

## $bindable для двустороннего связывания

```svelte
<script lang="ts">
  interface Props {
    value: string
  }

  let { value = $bindable() }: Props = $props()
</script>
```
