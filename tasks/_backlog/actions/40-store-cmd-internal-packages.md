# Action: вынести логику store/cmd/ в internal-пакеты

## Проблема

`store/cmd/` содержит несколько файлов в `package main` (`config.go`, `handler.go`, `store.go`), которые логически отвечают за разные вещи — конфиг, обработку соединений, управление lifecycle сервера. Сгруппировать их по смыслу нельзя подпапками: Go требует один пакет на директорию.

## Решение

Вынести логику в отдельные пакеты:

- `store/internal/config/` — парсинг и валидация конфига
- `store/internal/handler/` — обработчик соединения (JSONL reader/writer)
- `store/internal/server/` — lifecycle сокет-сервера (listen, accept, shutdown)

`store/cmd/main.go` становится тонким wiring-слоем: инициализация + wire-up.

## Источники

Обсуждение RE-038, 2026-05-06: отложено как отдельный рефактор кода.
