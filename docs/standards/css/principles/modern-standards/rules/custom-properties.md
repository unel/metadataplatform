# Custom Properties (CSS Variables)

Используй CSS custom properties для всех повторяющихся значений: цвета, отступы, типографика, радиусы, тени.

**Почему:** единое место для токенов дизайна. Изменение темы или значения — правка в одном месте. Значения читаемы и самодокументированы.

## Как проверять актуальность

caniuse.com/css-variables — поддержка universal, можно использовать везде.

## Плохо

```css
.button {
  background: #3b82f6;
  padding: 8px 16px;
  border-radius: 4px;
}

.card {
  border-radius: 4px; /* магическое число продублировано */
}
```

## Хорошо

```css
:root {
  --color-primary: #3b82f6;
  --spacing-sm: 8px;
  --spacing-md: 16px;
  --radius-sm: 4px;
}

.button {
  background: var(--color-primary);
  padding: var(--spacing-sm) var(--spacing-md);
  border-radius: var(--radius-sm);
}

.card {
  border-radius: var(--radius-sm);
}
```
