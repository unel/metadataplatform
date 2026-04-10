# Безопасность типов TypeScript

**Категория:** правила которые предотвращают обход системы типов.

**Примеры категорий:**
- Явный `any` — отключает проверку типов
- Ненулевые утверждения (`!`) — принудительное игнорирование null
- Небезопасные операции с типами

## Примеры правил (не исчерпывающий список)

```json
"@typescript-eslint/no-explicit-any": "error",
"@typescript-eslint/no-non-null-assertion": "error"
```

Если `any` или `!` действительно необходим — отключи с причиной:

```ts
// eslint-disable-next-line @typescript-eslint/no-explicit-any -- внешняя либа без типов
const result = legacyLib.process(data) as any
```
