# Как проверять поддержку браузерами

Перед использованием новой CSS-фичи — проверяй поддержку.

## Инструменты

- **caniuse.com** — детальная таблица поддержки по браузерам и версиям
- **web.dev/baseline** — статус "Baseline" означает что фича поддерживается во всех современных браузерах минимум 2.5 года
- **MDN Browser Compatibility** — таблица в документации каждой фичи

## Правило

- Статус **Baseline Widely Available** → можно использовать без оговорок
- Статус **Baseline Newly Available** → можно использовать с проверкой целевой аудитории
- Не в Baseline → нужен @supports или fallback

## @supports для прогрессивного улучшения

```css
.container {
  display: flex; /* fallback */
}

@supports (display: grid) {
  .container {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  }
}
```
