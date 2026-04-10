# Resource Hints

```html
<head>
  <!-- preload — загрузить критический ресурс раньше -->
  <link rel="preload" href="/fonts/main.woff2" as="font" type="font/woff2" crossorigin />
  <link rel="preload" href="/hero.webp" as="image" />

  <!-- prefetch — загрузить ресурс для следующей страницы в фоне -->
  <link rel="prefetch" href="/next-page.js" />

  <!-- dns-prefetch — резолвить DNS для внешних доменов заранее -->
  <link rel="dns-prefetch" href="//fonts.googleapis.com" />

  <!-- preconnect — установить соединение заранее -->
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
</head>

<!-- defer — выполнить после парсинга HTML, сохраняет порядок -->
<script src="/app.js" defer></script>

<!-- async — выполнить как только загрузится, не блокирует парсинг -->
<script src="/analytics.js" async></script>

<!-- module — defer по умолчанию -->
<script type="module" src="/app.js"></script>
```

## Правила

- `preload` — только для ресурсов текущей страницы которые нужны сразу
- `defer` — основной JS приложения
- `async` — независимые скрипты (аналитика, чаты)
- Не `preload` всё подряд — конкурирует с критическими ресурсами
