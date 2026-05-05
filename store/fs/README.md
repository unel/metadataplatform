# store/fs

FS-реализация store-интерфейсов. Хранит данные как JSON-файлы в поддиректориях `basedir`.

**Debug-only.** Не предназначена для продакшн-использования. Нет защиты от конкурентного доступа (TODO: RWMutex).

## Инициализация

```go
import (
    "log/slog"
    storefs "github.com/unel/metadataplatform/store/fs"
)

s, err := storefs.New(storefs.Config{Basedir: "/var/data/store"}, slog.Default())
if err != nil {
    // не удалось создать поддиректории
}

entityStore   := s.Entities()
relationStore := s.Relations()
jobStore      := s.Jobs()
```

`New` создаёт три поддиректории: `entities/`, `relations/`, `jobs/`. Если директории уже существуют — ошибки нет.

## Структура файлов

```
basedir/
  entities/
    01906acd-dead-7000-beef-000000000001.json
    ...
  relations/
    ...
  jobs/
    ...
```

Каждая запись — один JSON-файл, имя файла = `<id>.json`.

## Поведение операций

**Upsert**: атомарная запись через temp-файл + `os.Rename`. Если запись существует — читает старый `created_at` и сохраняет его. `updated_at` всегда = текущий UTC.

**Get**: читает файл, десериализует. Возвращает `ErrNotFound` если файл отсутствует.

**Delete**: удаляет файл. Возвращает `ErrNotFound` если файл отсутствует. **Не идемпотентен.**

**List**: читает все `.json` файлы в директории. Порядок не гарантирован. Возвращает пустой срез (не nil) если записей нет.

## Безопасность

ID валидируется на пустоту и проверяется на path traversal: ID вроде `../secret` вернут `ErrMissingID`. Это реализовано в `validate.go::recordPath`.

## Логирование

Стор логирует через `slog`. При передаче `nil` в `New` используется `slog.Default()`.

Уровни:
- `Info` — инициализация (store/fs логирует только её)

store/fs не логирует операции напрямую. Логирование операций (Debug — успешные, Error — ошибки) выполняется в `store/router`.
