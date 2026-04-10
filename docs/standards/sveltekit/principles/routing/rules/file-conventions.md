# Файловые конвенции

```
src/routes/
  +page.svelte          — страница
  +page.ts              — universal load (клиент + сервер)
  +page.server.ts       — server-only load, form actions
  +layout.svelte        — layout для всего поддерева
  +layout.ts            — load для layout
  +error.svelte         — страница ошибки
  (auth)/               — route group — общий layout без изменения URL
    login/+page.svelte
    register/+page.svelte
  entities/
    [id]/               — динамический сегмент
      +page.svelte
    [...rest]/          — catch-all сегмент
```

## Принципы

- `+page.svelte` получает данные только через `data` prop из load функции — не делает fetch сам
- `+layout.svelte` оборачивает страницы — не содержит логику конкретных страниц
- Route groups `(name)` — для группировки с общим layout без влияния на URL
