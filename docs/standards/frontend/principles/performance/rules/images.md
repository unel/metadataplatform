# Оптимизация изображений

Изображения — главная причина медленной загрузки. Несколько правил дают значимый эффект.

## Форматы

WebP или AVIF вместо JPEG/PNG — в 2-3 раза меньше при том же качестве.

```html
<picture>
  <source srcset="image.avif" type="image/avif" />
  <source srcset="image.webp" type="image/webp" />
  <img src="image.jpg" alt="описание" width="800" height="600" />
</picture>
```

## Явные размеры

Всегда указывай width и height — браузер резервирует место до загрузки, нет CLS.

## Lazy loading

```html
<img src="image.webp" alt="..." loading="lazy" width="400" height="300" />
```

Не применяй к LCP изображению — оно должно загружаться сразу.

## LCP изображение

```html
<link rel="preload" as="image" href="hero.webp" />
<img src="hero.webp" fetchpriority="high" alt="..." />
```
