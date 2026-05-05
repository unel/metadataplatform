# Action: передавать явный путь docs/standards/go/ в промпт ревьюера

## Проблема

Гримм написал "расширения: нет" — `docs/standards/v2/extensions/` не существует. При этом `docs/standards/go/` содержит принципы concurrency, error-handling, interfaces, testing, project-structure — прямо релевантные. Ревьюер не нашёл и не применил стандарты Go.

## Решение

В промпт каждого ревьюера (spec-review, code-review, test-review) передавать явный путь:
```
Стандарты Go: docs/standards/go/
```
Не ссылаться на несуществующую extensions/.

## Шаги

- [ ] Обновить base-plan в `docs/standards/v2/*/02-review/` — добавить явный путь к Go-стандартам
- [ ] Либо: создать `docs/standards/v2/extensions/` как симлинк или index на `docs/standards/go/`

## Источники

Пользователь (ретро store/crud, 2026-05-05)
