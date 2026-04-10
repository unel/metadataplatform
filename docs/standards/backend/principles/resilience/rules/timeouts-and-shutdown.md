# Таймауты и Graceful Shutdown

## Таймауты

Каждый внешний вызов должен иметь таймаут. Без таймаута один зависший upstream блокирует весь сервис.

```
// примеры таймаутов (конкретные значения зависят от SLA)
http client timeout:  30s
db query timeout:     10s
db connection timeout: 5s
external api timeout: 15s
```

## Graceful Shutdown

При получении SIGTERM сервис должен:
1. Перестать принимать новые запросы
2. Дождаться завершения текущих (с таймаутом)
3. Закрыть соединения с БД и другими ресурсами
4. Завершиться с кодом 0

```
signal SIGTERM received
→ stop accepting new connections
→ wait for in-flight requests (max 30s)
→ close db pool
→ exit(0)
```

## Retry

Retry только для идемпотентных операций и временных ошибок (сеть, таймаут). Не делай retry при 4xx — это ошибка клиента.

Exponential backoff: 1s → 2s → 4s → 8s с jitter чтобы не создавать thundering herd.
