# REST Conventions

## Именование ресурсов

```
GET    /entities          — список
GET    /entities/{id}     — один элемент
POST   /entities          — создание
PUT    /entities/{id}     — полное обновление
PATCH  /entities/{id}     — частичное обновление
DELETE /entities/{id}     — удаление

GET    /entities/{id}/relations  — вложенный ресурс
```

Существительные, множественное число, kebab-case для составных слов (`/entity-types`).

## Статус коды

```
200 OK              — успешный GET/PATCH/PUT
201 Created         — успешный POST с созданием
204 No Content      — успешный DELETE или POST без тела ответа
400 Bad Request     — невалидный запрос (ошибка клиента)
401 Unauthorized    — не аутентифицирован
403 Forbidden       — аутентифицирован, но нет прав
404 Not Found       — ресурс не найден
409 Conflict        — конфликт состояния (дубликат, устаревшая версия)
422 Unprocessable   — валидные данные но бизнес-правило нарушено
429 Too Many Req    — rate limiting
500 Internal Error  — ошибка сервера
```

## Формат ошибок

Единый формат по всему API:

```json
{
  "error": "entity_not_found",
  "message": "Entity with id '123' not found",
  "details": { "id": "123" }
}
```
