# Магические числа в z-index

Не используй произвольные числа в z-index — используй именованные уровни через custom properties.

**Почему:** `z-index: 9999` порождает `z-index: 99999`. Через время никто не знает что над чем должно быть. Система слоёв теряет смысл.

## Плохо

```css
.modal { z-index: 1000; }
.tooltip { z-index: 9999; }
.dropdown { z-index: 500; }
```

## Хорошо

```css
:root {
  --z-dropdown: 100;
  --z-modal: 200;
  --z-tooltip: 300;
}

.modal { z-index: var(--z-modal); }
.tooltip { z-index: var(--z-tooltip); }
.dropdown { z-index: var(--z-dropdown); }
```
