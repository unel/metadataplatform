# Как следить за обновлениями

## Источники

- **postgresql.org/docs/release/** — release notes каждой версии
- **pganalyze.com/blog** — разборы новых возможностей и performance improvements
- **planet.postgresql.org** — агрегатор блогов PostgreSQL community
- **github.com/postgres/postgres** — changelog и discussions

## Что искать

- Новые JSON возможности (JSON_TABLE появился в PG 16, улучшенный jsonpath в PG 14)
- Новые типы индексов и улучшения планировщика
- Изменения в поведении существующих функций
- Новые расширения в contrib

## Паттерн для устаревших подходов

```sql
-- OUTDATED(PG 13): ручное формирование tsvector конкатенацией
-- актуальный подход: использовать generated column (PG 12+)
-- рефакторинг: tasks/feature/tsvector-generated-column
ALTER TABLE relations
  ADD COLUMN search_tsv tsvector GENERATED ALWAYS AS (
    to_tsvector('simple', coalesce(type,'') || ' ' || coalesce(subtype,''))
  ) STORED;
```

## При обновлении PostgreSQL

1. Читай release notes — особенно раздел "Incompatibilities"
2. Проверяй устаревшие функции и операторы
3. Запускай `EXPLAIN ANALYZE` на ключевых запросах — планировщик улучшается, иногда меняется поведение
