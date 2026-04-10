# FOR UPDATE SKIP LOCKED

Паттерн для очередей задач: несколько воркеров берут строки параллельно без конфликтов.

## Проблема без SKIP LOCKED

```sql
-- воркер 1 и воркер 2 одновременно:
SELECT id FROM jobs WHERE status = 'pending' LIMIT 1;
-- оба получат одну строку → дублирование обработки
```

## Правильный паттерн

```sql
-- каждый воркер берёт себе строку атомарно
-- SKIP LOCKED — не ждать, просто пропустить занятые строки
SELECT id, payload FROM jobs
WHERE status = 'pending'
ORDER BY created_at
LIMIT 1
FOR UPDATE SKIP LOCKED;
```

Этот SELECT внутри транзакции. Воркер обрабатывает строку, потом:

```sql
UPDATE jobs SET status = 'running', updated_at = now() WHERE id = $1;
COMMIT;
```

## Advisory locks — singleton-гарантия

Используется когда нужен ровно один инстанс приложения (spawner в этом проекте):

```sql
-- пытается взять advisory lock с id = 12345
-- возвращает true если успешно, false если уже занят
SELECT pg_try_advisory_lock(12345);
```

Блокировка держится всё время жизни соединения. При разрыве соединения — автоматически освобождается.

```go
// При старте spawner
var locked bool
err := db.QueryRow("SELECT pg_try_advisory_lock($1)", spawnerlockID).Scan(&locked)
if !locked {
    log.Fatal("another spawner instance is running")
}
// держим соединение живым, не закрываем
```

## Дедлоки

Дедлок возникает когда транзакция A ждёт блокировку транзакции B, а B ждёт A.

Предотвращение: всегда блокируй строки в одном порядке. Если нужно обновить entities и relations — сначала entity, потом relation, везде в приложении одинаково.
