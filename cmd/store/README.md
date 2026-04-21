# store

Unix-socket сервер для приёма и обработки JSONL-сообщений от других компонентов платформы (api, spawner).

## Сборка

```sh
go build -o store ./cmd/store
```

## Запуск

Сервис требует конфиг-файл. Путь задаётся флагом или переменной окружения:

```sh
# через флаг
./store --config /etc/platform/config.yaml

# через переменную окружения
CONFIG_PATH=/etc/platform/config.yaml ./store
```

Если ни флаг, ни переменная не заданы — сервис завершается с ошибкой и кодом 1.

### Пример конфига

```yaml
store:
  db_url: postgres://localhost/platform   # не используется в текущей версии
  max_line_bytes: 1048576                 # 1 МБ (дефолт)
  max_connections: 1024                   # дефолт
  sockets:
    - path: /tmp/store.sock
      ops: [query, upsert, delete]
```

Поле `sockets` обязательно и не может быть пустым — иначе `exit(1)`.

## Как подключиться вручную

### nc (netcat-openbsd)

```sh
nc -U /tmp/store.sock
```

### socat

```sh
socat - UNIX-CONNECT:/tmp/store.sock
```

## Пример сессии

Запускаем store:

```sh
./store --config config.yaml
# INFO listening on socket /tmp/store.sock
```

В другом терминале подключаемся и отправляем строки:

```sh
$ nc -U /tmp/store.sock
{"op":"query","table":"assets"}
{"op":"query","table":"assets"}

{"op":"upsert","table":"assets","data":{"id":1}}
{"op":"upsert","table":"assets","data":{"id":1}}
```

Сервис возвращает каждую строку эхом. Пустые строки игнорируются — ответ на них не приходит.

Завершение сессии: `Ctrl+D` (EOF).

## Остановка

Сервис корректно завершается по `SIGTERM` или `SIGINT` (`Ctrl+C`):

```
INFO shutting down
```

Активным соединениям даётся 5 секунд на завершение, затем они закрываются принудительно. Сокет-файлы удаляются при остановке.

## Логирование

Сервис использует стандартный `log/slog` (structured logging), вывод в stderr.

| Уровень | Что логируется |
|---------|----------------|
| INFO    | старт каждого сокета, новые подключения (с `conn_id`), ошибки чтения/записи, переполнение буфера, неожиданные отключения, shutdown |
| DEBUG   | каждая полученная строка (с `conn_id`), нормальное закрытие соединения (EOF) |

Для включения DEBUG-логов установите переменную окружения (стандартный механизм slog не настраивается флагом — при необходимости настройте через обёртку или `slog.SetDefault`).
