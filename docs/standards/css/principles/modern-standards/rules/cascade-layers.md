# Cascade Layers (@layer)

Используй @layer для явного управления порядком каскада — base, components, utilities.

**Почему:** без @layer специфичность — источник неожиданных переопределений. @layer даёт предсказуемый порядок независимо от специфичности селекторов.

## Как проверять актуальность

caniuse.com/css-cascade-5 — поддержка широкая (все современные браузеры).

## Пример

```css
@layer base, components, utilities;

@layer base {
  * { box-sizing: border-box; }
  body { margin: 0; }
}

@layer components {
  .button { padding: var(--spacing-sm) var(--spacing-md); }
}

@layer utilities {
  .mt-auto { margin-top: auto; }
}
```

Утилиты всегда побеждают компоненты — независимо от специфичности селекторов.
