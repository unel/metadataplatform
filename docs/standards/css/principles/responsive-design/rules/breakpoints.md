# Breakpoints

Breakpoints через custom properties — не магические числа разбросанные по коду.

**Почему:** единое место для изменения. Консистентность между компонентами.

## Плохо

```css
@media (min-width: 768px) { ... }
@media (min-width: 769px) { ... } /* опечатка? или намеренно? */
@media (min-width: 1024px) { ... }
```

## Хорошо

```css
:root {
  --bp-sm: 480px;
  --bp-md: 768px;
  --bp-lg: 1024px;
  --bp-xl: 1280px;
}

/* В SvelteKit — через JS-переменные или preprocessor */
@media (min-width: 768px) { ... } /* ссылайся на документированные значения */
```
