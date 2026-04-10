# Как следить за обновлениями

## Источники

- **go.dev/doc/go1.N** — release notes каждой версии
- **go.dev/blog** — официальный блог, разборы новых возможностей
- **go.dev/talks** — доклады команды Go
- **Effective Go** — базовые идиомы, периодически обновляется
- **github.com/golang/go/issues** — что идёт в следующий релиз

## Что искать

- Новые языковые возможности (generics появились в 1.18, range over func в 1.22)
- Новые пакеты в стандартной библиотеке (например `slices`, `maps`, `cmp` в 1.21)
- Изменения в рекомендуемых паттернах (например: новый slog вместо log в 1.21)
- Улучшения инструментария (go vet, staticcheck)

## Паттерн для устаревших паттернов

```go
// OUTDATED(Go 1.20): ручная реализация слайсовых утилит
// актуальный подход: пакет slices из стандартной библиотеки (Go 1.21+)
// рефакторинг: tasks/feature/stdlib-slices-migration
func contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}
```

## При обновлении Go

1. Читай release notes полностью — Go редко ломает совместимость, но бывает
2. Ищи OUTDATED комментарии
3. Запускай `go vet` и `staticcheck` — они знают о новых идиомах
