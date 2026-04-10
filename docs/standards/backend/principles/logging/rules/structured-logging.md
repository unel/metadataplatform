# Structured Logging

Логи — JSON, не строки. Машина читает поля, человек читает message.

## Плохо

```
[ERROR] Failed to process entity abc123: connection refused at db:5432
```

## Хорошо

```json
{
  "level": "error",
  "time": "2024-01-15T10:30:00Z",
  "request_id": "req-uuid",
  "entity_id": "abc123",
  "operation": "process_entity",
  "error": "connection refused",
  "addr": "db:5432",
  "msg": "failed to process entity"
}
```

## Уровни

- **DEBUG** — детали для разработки, в продакшене выключен
- **INFO** — значимые события (запрос получен, job запущен, job завершён)
- **WARN** — нештатная ситуация, система справилась (retry, fallback)
- **ERROR** — ошибка требующая внимания, операция не выполнена

## Что логировать на INFO

- Старт и остановка сервиса
- Начало и конец обработки job'а
- Входящие запросы (метод, путь, статус, время)

## Что НЕ логировать

- Пароли, токены, секреты
- Персональные данные (email, имена) без необходимости
- Каждую итерацию цикла на INFO
