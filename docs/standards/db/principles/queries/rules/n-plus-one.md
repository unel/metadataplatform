# N+1 проблема

N+1 — запрос списка (1 запрос) + запрос для каждого элемента (N запросов). При 1000 элементах — 1001 запрос вместо 2.

## Плохо

```
entities = db.query("SELECT * FROM entities WHERE type = 'file'")
for entity in entities:
    relations = db.query("SELECT * FROM relations WHERE from_id = $1", entity.id)
    // N+1: один запрос на каждую entity
```

## Хорошо

```sql
-- один запрос через JOIN
SELECT e.*, r.*
FROM entities e
LEFT JOIN relations r ON r.from_id = e.id
WHERE e.type = 'file'

-- или два запроса: список + все связанные данные сразу
SELECT * FROM entities WHERE type = 'file';
SELECT * FROM relations WHERE from_id = ANY($1::uuid[]);
```

## Как обнаружить

- Логировать медленные запросы (`log_min_duration_statement`)
- Считать количество запросов на запрос API — больше 5-10 — подозрительно
