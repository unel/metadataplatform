# tsvector и tsquery

## Формирование tsvector

```sql
-- агрегируем несколько полей с весами
-- A — самый высокий вес (title), B — ниже (body)
UPDATE relations SET search_tsv =
  setweight(to_tsvector('simple', coalesce(type, '')), 'A') ||
  setweight(to_tsvector('simple', coalesce(subtype, '')), 'B') ||
  setweight(to_tsvector('simple', coalesce(value::text, '')), 'C') ||
  setweight(to_tsvector('simple', coalesce(meta::text, '')), 'D');

-- через триггер — обновление автоматически при INSERT/UPDATE
CREATE TRIGGER update_search_tsv
  BEFORE INSERT OR UPDATE ON relations
  FOR EACH ROW EXECUTE FUNCTION update_search_tsv_fn();
```

## Поиск

```sql
-- plainto_tsquery: строку без спецсимволов → запрос AND
SELECT * FROM relations
WHERE search_tsv @@ plainto_tsquery('simple', $1);

-- websearch_to_tsquery: понимает "фразу в кавычках" и -исключение
SELECT * FROM relations
WHERE search_tsv @@ websearch_to_tsquery('simple', $1);

-- НЕ использовать to_tsquery с пользовательским вводом напрямую —
-- требует синтаксис tsquery, пользователи его не знают
```

## Ранжирование результатов

```sql
SELECT
  id,
  ts_rank(search_tsv, query) AS rank
FROM relations,
  websearch_to_tsquery('simple', $1) query
WHERE search_tsv @@ query
ORDER BY rank DESC
LIMIT 20;
```

## Подсветка совпадений

```sql
-- ts_headline возвращает фрагмент текста с выделенными совпадениями
SELECT ts_headline('simple', value::text, query) AS snippet
FROM relations,
  websearch_to_tsquery('simple', $1) query
WHERE search_tsv @@ query;
```

## Конфигурация языка

`simple` — без стемминга, подходит для имён собственных, идентификаторов, смешанного контента. `russian`/`english` — стемминг, подходит для текстовых документов.

В рамках одного приложения используй одну конфигурацию везде — иначе результаты поиска будут непредсказуемы.
