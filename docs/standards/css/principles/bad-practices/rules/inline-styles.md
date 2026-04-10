# Инлайн-стили для постоянных стилей

Не используй `style=""` для стилей которые не меняются динамически.

**Почему:** инлайн-стили имеют наивысшую специфичность (кроме `!important`) и не переопределяются из CSS-файлов. Стили разбросаны по шаблонам — нет единого места для изменения.

## Когда допустимо

Динамические значения которые вычисляются в runtime:

```html
<div style="width: {progress}%">...</div>
<div style="--accent-color: {userColor}">...</div>
```

## Плохо

```html
<button style="background: blue; padding: 8px 16px; border-radius: 4px;">
  Click
</button>
```

## Хорошо

```html
<button class="button-primary">Click</button>
```
