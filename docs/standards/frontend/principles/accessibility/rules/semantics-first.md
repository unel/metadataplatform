# Семантика прежде ARIA

Используй нативные HTML элементы — они уже несут семантику и поведение. ARIA — последнее средство когда нативного HTML недостаточно.

**Правило:** первое правило ARIA — не использовать ARIA.

## Плохо

```html
<div role="button" tabindex="0" onclick="submit()">Отправить</div>
<div role="navigation">...</div>
<span role="heading" aria-level="1">Заголовок</span>
```

## Хорошо

```html
<button type="submit">Отправить</button>
<nav>...</nav>
<h1>Заголовок</h1>
```

## Когда ARIA нужен

Кастомные компоненты без нативного аналога: combobox, tree, data grid, tooltip.

```html
<div role="combobox" aria-expanded="true" aria-haspopup="listbox">
  <input aria-autocomplete="list" aria-controls="listbox-id" />
</div>
```
