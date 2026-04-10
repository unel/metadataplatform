# Структура документа

```html
<!DOCTYPE html>
<html lang="ru">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <title>Заголовок страницы</title>
</head>
<body>
  <header>
    <nav aria-label="Основная навигация">
      <ul>
        <li><a href="/">Главная</a></li>
      </ul>
    </nav>
  </header>

  <main>
    <h1>Заголовок страницы</h1>
    <article>
      <h2>Раздел</h2>
      <p>Содержимое</p>
    </article>
    <aside>
      <h2>Дополнительно</h2>
    </aside>
  </main>

  <footer>...</footer>
</body>
</html>
```

## Правила иерархии заголовков

- Один `h1` на страницу — главный заголовок
- Не пропускай уровни: после `h2` идёт `h3`, не `h4`
- Заголовки отражают структуру контента, не стиль
