# Когда создавать индекс

## Обязательно

- Внешние ключи — без индекса JOIN и ON DELETE CASCADE медленные
- Колонки в WHERE часто выполняемых запросов
- Колонки в ORDER BY если нужна сортировка без seq scan
- `jsonb` колонки с поиском по содержимому (`@>`) — GIN индекс
- `tsvector` колонки — GIN индекс

## Специальные индексы

```sql
-- GIN для jsonb containment и full-text
CREATE INDEX entities_meta_gin ON entities USING gin(meta);
CREATE INDEX relations_search_gin ON relations USING gin(search_tsv);

-- Partial index — только для подмножества строк
CREATE INDEX jobs_pending_idx ON jobs(created_at)
  WHERE status = 'pending';

-- Composite index — порядок колонок важен
CREATE INDEX entities_type_subtype ON entities(type, subtype);
```

## Когда НЕ создавать

- Таблицы с редкими запросами и частыми записями
- Колонки с очень низкой кардинальностью (boolean, статус с 2-3 значениями) — seq scan быстрее
- Индекс по колонке которая уже покрыта составным индексом

## В продакшене — всегда CONCURRENTLY

```sql
CREATE INDEX CONCURRENTLY entities_type_idx ON entities(type);
```
