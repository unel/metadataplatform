# Почему strict обязателен

`strict: true` включает группу флагов которые вместе дают реальную защиту от типичных ошибок.

**Что включает strict:**
- `strictNullChecks` — null и undefined не совместимы с другими типами
- `strictFunctionTypes` — корректная контравариантность параметров функций
- `strictBindCallApply` — типизированный bind/call/apply
- `noImplicitAny` — запрет неявного any
- `noImplicitThis` — запрет неявного this
- `alwaysStrict` — `"use strict"` в каждом файле

**Почему нельзя отключать отдельные флаги:**
Каждый флаг закрывает реальный класс ошибок. Отключение = осознанное согласие на этот класс багов. Если флаг мешает — разберись почему, не отключай.

**Если отключение всё же необходимо:**
```json
// tsconfig.json — с обоснованием рядом
{
  "strict": true,
  "strictPropertyInitialization": false // legacy код, рефакторинг в roadmap
}
```
