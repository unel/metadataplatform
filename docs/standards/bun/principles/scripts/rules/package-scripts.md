# Стандартные scripts

```json
{
  "scripts": {
    "dev": "bun --hot src/index.ts",
    "build": "bun build src/index.ts --outdir dist --target bun",
    "test": "bun test",
    "test:watch": "bun test --watch",
    "test:coverage": "bun test --coverage",
    "lint": "eslint src",
    "typecheck": "tsc --noEmit",
    "check": "bun run lint && bun run typecheck"
  }
}
```

## Сложные скрипты — в отдельные файлы

```json
{
  "scripts": {
    "db:migrate": "bun run scripts/migrate.ts",
    "db:seed": "bun run scripts/seed.ts",
    "generate": "bun run scripts/generate.ts"
  }
}
```

```typescript
// scripts/migrate.ts
import { readdir } from 'node:fs/promises'

const migrations = await readdir('./migrations')
for (const file of migrations.sort()) {
  console.log(`Running migration: ${file}`)
  // ...
}
```

## --hot для разработки

`bun --hot` перезагружает модули без перезапуска процесса. Быстрее чем `--watch` для серверного кода.
