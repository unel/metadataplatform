# Когда что использовать

## interface — для контрактов и публичных API

```typescript
// хорошо — interface для структур которые могут быть расширены
interface Repository<T> {
  find(id: string): Promise<T>
  save(entity: T): Promise<void>
}

// расширение через declaration merging
interface Window {
  myCustomProp: string
}
```

**Используй interface когда:**
- Описываешь контракт который другие будут реализовывать
- Нужна расширяемость через declaration merging
- Это публичный API библиотеки или модуля

## type — для всего остального

```typescript
// объединения — только type
type Status = 'pending' | 'running' | 'done' | 'failed'

// пересечения
type AdminUser = User & { permissions: string[] }

// утилитарные трансформации
type PartialConfig = Partial<Config>

// условные типы
type NonNullable<T> = T extends null | undefined ? never : T
```

**Используй type когда:**
- Объединения (union) и пересечения (intersection)
- Алиасы для примитивов и кортежей
- Условные и mapped типы
- Трансформации существующих типов

## Главное правило

Договорись с командой и следуй соглашению консистентно. Смешение без причины — хуже любого выбора.
