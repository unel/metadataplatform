# jsonb операторы

## Основные операторы

| Оператор | Что делает | Использует индекс |
|----------|-----------|-------------------|
| `@>` | containment — объект содержит этот фрагмент | GIN ✓ |
| `<@` | обратный containment | GIN ✓ |
| `?` | ключ существует | GIN ✓ |
| `->` | извлечь значение как jsonb | — |
| `->>` | извлечь значение как text | — |
| `#>` | извлечь по пути как jsonb | — |
| `#>>` | извлечь по пути как text | — |

## Плохо — фильтрация через text cast

```sql
-- не использует GIN индекс, медленно на больших таблицах
SELECT * FROM entities WHERE meta->>'path' = '/media/video.mp4';
```

## Хорошо — containment через @>

```sql
-- использует GIN индекс
SELECT * FROM entities WHERE meta @> '{"path": "/media/video.mp4"}';

-- с параметром (безопасно от инъекций)
SELECT * FROM entities WHERE meta @> jsonb_build_object('path', $1);
```

## JSON path для сложных условий

```sql
-- размер больше 10MB
SELECT * FROM entities WHERE meta @? '$.size ? (@ > 10485760)';

-- путь начинается с /media/
SELECT * FROM entities WHERE meta @? '$.path ? (@ starts with "/media/")';

-- комбинация условий
SELECT * FROM entities
WHERE meta @@ '$.size > 10485760 && $.mimetype == "video/mp4"';
```

## Обновление отдельного поля

```sql
-- плохо — перезаписывает весь объект
UPDATE entities SET meta = '{"path": "/new"}' WHERE id = $1;

-- хорошо — обновляет только нужное поле
UPDATE entities SET meta = meta || jsonb_build_object('path', $1) WHERE id = $2;

-- для вложенных путей
UPDATE entities SET meta = jsonb_set(meta, '{nested,key}', $1) WHERE id = $2;
```

## Когда ->> оправдан

Извлечение значения для отображения или сравнения с уже известным типом, когда индекс не нужен:

```sql
-- просто читаем значение, поиска нет
SELECT meta->>'name', meta->>'path' FROM entities WHERE id = $1;
```
