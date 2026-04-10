# Безопасные миграции

## Добавление NOT NULL колонки

Нельзя добавить NOT NULL колонку без default на непустую таблицу — блокирует таблицу на время backfill.

```sql
-- шаг 1: добавить nullable
ALTER TABLE entities ADD COLUMN new_field text;

-- шаг 2: заполнить (отдельный деплой или миграция)
UPDATE entities SET new_field = 'default' WHERE new_field IS NULL;

-- шаг 3: добавить constraint (отдельная миграция)
ALTER TABLE entities ALTER COLUMN new_field SET NOT NULL;
```

## Удаление колонки

```sql
-- шаг 1: убрать использование из кода (деплой)
-- шаг 2: пометить в БД (опционально)
COMMENT ON COLUMN entities.old_field IS 'deprecated, will be removed';
-- шаг 3: удалить после нескольких деплоев
ALTER TABLE entities DROP COLUMN old_field;
```

## Создание индекса без блокировки

```sql
-- всегда CONCURRENTLY для продакшена
CREATE INDEX CONCURRENTLY IF NOT EXISTS entities_type_idx ON entities(type);
```

## Идемпотентность

```sql
ALTER TABLE entities ADD COLUMN IF NOT EXISTS new_field text;
CREATE INDEX IF NOT EXISTS entities_type_idx ON entities(type);
DROP INDEX IF EXISTS old_index_name;
```
