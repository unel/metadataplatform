# Когда использовать бандлер Bun

## SvelteKit — Vite, не Bun bundler

SvelteKit имеет собственный build pipeline через Vite. Не пытайся заменить его бандлером Bun.

```bash
bun run build  # запускает vite build через SvelteKit
```

## Серверные скрипты и CLI утилиты

```typescript
// build.ts
await Bun.build({
  entrypoints: ['./src/cli.ts'],
  outdir: './dist',
  target: 'bun',
  minify: true,
  sourcemap: 'external'
})
```

## Одиночные файлы — транспиляция без бандлинга

```bash
# Bun выполняет TypeScript напрямую без предварительной сборки
bun run src/worker.ts

# для продакшена — можно собрать в один файл
bun build src/worker.ts --outfile dist/worker.js --target bun
```
