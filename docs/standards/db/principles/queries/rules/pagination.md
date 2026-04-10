# Pagination

OFFSET медленный на больших таблицах — БД всё равно читает все предыдущие строки. Cursor-based pagination работает одинаково быстро при любом размере таблицы.

## Плохо — OFFSET

```sql
SELECT * FROM entities ORDER BY id LIMIT 20 OFFSET 10000;
-- БД читает 10020 строк, возвращает 20
```

## Хорошо — Cursor

```sql
-- первая страница
SELECT * FROM entities ORDER BY id LIMIT 20;

-- следующая страница (cursor = последний id предыдущей)
SELECT * FROM entities WHERE id > $cursor ORDER BY id LIMIT 20;
```

## Ограничения cursor pagination

- Только с `ORDER BY` по уникальной колонке
- Нельзя перейти на произвольную страницу
- Подходит для бесконечного скролла, не для "перейти на страницу N"
