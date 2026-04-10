# Нативные Bun API

Bun предоставляет более быстрые нативные альтернативы стандартным Node.js операциям.

## Файловая система

```typescript
// Bun — быстрее и удобнее
const file = Bun.file('./data.json')
const data = await file.json()
await Bun.write('./output.json', JSON.stringify(result))

// Node.js — работает, но медленнее
import { readFile, writeFile } from 'node:fs/promises'
const data = JSON.parse(await readFile('./data.json', 'utf-8'))
```

## HTTP сервер

```typescript
// Bun.serve — встроенный HTTP сервер без зависимостей
Bun.serve({
  port: 3000,
  fetch(req) {
    return new Response('Hello')
  }
})
```

## Хэширование и криптография

```typescript
// Bun.hash — быстрое хэширование
const hash = Bun.hash('hello world')

// Bun.CryptoHasher — совместимо с Web Crypto API
const hasher = new Bun.CryptoHasher('sha256')
hasher.update('hello')
const digest = hasher.digest('hex')
```

## Переменные окружения

```typescript
// предпочтительно
const dbUrl = Bun.env.DB_URL

// работает но менее идиоматично для Bun
const dbUrl = process.env.DB_URL
```
