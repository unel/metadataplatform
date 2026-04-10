# Mobile-First

Пиши базовые стили для мобильного экрана, расширяй для большего через `min-width`.

**Почему:** мобильный дизайн — ограниченное пространство, это сложнее. Проще добавлять возможности для большого экрана, чем убирать для маленького. Меньше переопределений.

## Плохо — desktop-first

```css
.sidebar {
  width: 300px;
  float: left;
}

@media (max-width: 768px) {
  .sidebar {
    width: 100%;
    float: none;
  }
}
```

## Хорошо — mobile-first

```css
.sidebar {
  width: 100%; /* мобильный — базовый */
}

@media (min-width: 768px) {
  .sidebar {
    width: 300px;
    float: left;
  }
}
```
