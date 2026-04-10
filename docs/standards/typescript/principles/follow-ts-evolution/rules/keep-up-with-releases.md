# Как следить за релизами

## Источники

- **TypeScript blog** — официальные анонсы каждого релиза с примерами
- **TypeScript GitHub** — milestone'ы и PR для понимания что идёт в следующий релиз
- **Matt Pocock (Total TypeScript)** — разборы новых возможностей с практическими примерами

## Что искать в release notes

- Новые утилитарные типы
- Улучшения вывода типов (inference improvements)
- Новые возможности narrowing
- Изменения в работе generics

## Паттерн для workarounds

Когда пишешь сложный workaround из-за ограничения TS — помечай:

```typescript
// TS-WORKAROUND(4.9): conditional type вместо satisfies которого ещё нет
// когда обновимся до 5.0+ — переписать через satisfies
type CheckConfig<T> = T extends ValidConfig ? T : never
```

При обновлении TS — ищи такие комментарии и упрощай.

## Дополнительные источники

- **Matt Pocock (Total TypeScript)** — практические разборы новых возможностей
- **typescript-weekly.com** — еженедельная подборка
- **github.com/microsoft/TypeScript/milestones** — что идёт в следующий релиз
