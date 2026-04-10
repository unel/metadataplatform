# Container Queries (@container)

Предпочитай `@container` вместо `@media` для адаптивности компонентов.

**Почему:** медиа-запросы реагируют на размер viewport — компонент не знает в каком контексте он используется. Контейнерные запросы реагируют на размер родителя — компонент адаптируется к своему контейнеру независимо от viewport.

## Как проверять актуальность

caniuse.com/css-container-queries — Baseline Widely Available, можно использовать.

## Плохо — медиа-запрос для компонента

```css
/* карточка ломается если её поместить в узкую колонку на широком экране */
@media (min-width: 768px) {
  .card { display: flex; }
}
```

## Хорошо — контейнерный запрос

```css
.card-wrapper {
  container-type: inline-size;
}

.card { display: block; }

@container (min-width: 400px) {
  .card { display: flex; }
}
```
