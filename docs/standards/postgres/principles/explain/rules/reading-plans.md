# Чтение планов выполнения

## Базовый синтаксис

```sql
EXPLAIN (ANALYZE, BUFFERS, FORMAT TEXT)
SELECT * FROM entities WHERE meta @> '{"type": "video"}';
```

`ANALYZE` — реально выполняет запрос и показывает actual rows/time.
`BUFFERS` — показывает сколько блоков прочитано из кэша и с диска.

## Структура вывода

```
Bitmap Heap Scan on entities  (cost=8.42..32.17 rows=10 width=240)
                               (actual time=0.124..0.187 rows=8 loops=1)
  Recheck Cond: (meta @> '{"type": "video"}'::jsonb)
  ->  Bitmap Index Scan on entities_meta_gin_idx
        (cost=0.00..8.42 rows=10 width=0)
        (actual time=0.112..0.112 rows=8 loops=1)
        Index Cond: (meta @> '{"type": "video"}'::jsonb)
Planning Time: 0.3 ms
Execution Time: 0.4 ms
```

**cost=8.42..32.17** — оценка планировщика (стартовая..полная). Не миллисекунды.
**rows=10** — оценка планировщика. **actual rows=8** — реальное.
**loops=1** — сколько раз узел выполнялся.

## Что искать

**Seq Scan на большой таблице** — планировщик не нашёл подходящий индекс или оценил его невыгодным:
```
Seq Scan on entities  (cost=0.00..5234.00 rows=100000 width=240)
```
→ проверь есть ли индекс, подходит ли условие для его использования.

**Расхождение rows vs actual rows** — статистика устарела:
```
rows=1000  (actual rows=89341)
```
→ `ANALYZE entities;` для пересбора статистики.

**Nested Loop с большим actual rows** — может быть признаком отсутствия индекса на join колонке:
```
Nested Loop  (actual rows=50000 loops=1)
  -> Seq Scan on relations
```
→ проверь индекс на `from_id`/`to_id`.

## Когда планировщик игнорирует индекс

Планировщик может предпочесть Seq Scan если:
- Таблица маленькая (seq scan дешевле)
- Selectivity низкая (индекс вернёт >20% строк)
- Статистика устарела

Форсировать индекс через `SET enable_seqscan = off` — только для диагностики, не в продакшене.
