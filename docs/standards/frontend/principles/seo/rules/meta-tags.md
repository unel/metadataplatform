# Мета-теги

В SvelteKit через `<svelte:head>` на каждой странице:

```svelte
<svelte:head>
  <title>{entity.name} — MetadataPlatform</title>
  <meta name="description" content={entity.description} />
  <meta property="og:title" content={entity.name} />
  <meta property="og:description" content={entity.description} />
  <link rel="canonical" href={canonicalUrl} />
</svelte:head>
```

## Правила

- `title` уникален для каждой страницы, до 60 символов
- `description` описывает содержимое, до 160 символов
- Не дублируй title в description
- Canonical обязателен если один контент доступен по нескольким URL
