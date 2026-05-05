# Action: раздел "API Contracts" в шаблоне acceptance-write

## Проблема

Два тест-агента (Кроули и Азирафаль) писали тесты к несуществующему API и угадывали сигнатуры независимо. Кроули предположил `*slog.Logger`, Азирафаль — кастомный `captureLogger`. Оба несовместимы. Кроме того, оба предположили что `*fs.Store` реализует три интерфейса с методом `Get` разных возвращаемых типов — в Go невозможно. Ада обнаружила это в 04-code. Итог: 4+ лишних раунда review/fix.

Action item из RETRO store/connection ("шаблон спеки: обязательный пример лог-строки") не дошёл до этой фичи — рецидив.

## Решение

Добавить раздел "API Contracts" в шаблон `docs/standards/v2/02-acceptance/01-write/`:

```markdown
## API Contracts

### Конструкторы
- `New(logger <тип>, ...) (*Store, error)` — или TBD: решает code-агент, тест-агенты используют интерфейс

### Интерфейсы
- `Store` реализует: <список> — или TBD

### Типы ошибок
- `ErrNotFound`, `ErrInvalidID`, ... — или TBD

### Логгер
- Тип: `*slog.Logger` / кастомный интерфейс / TBD
- Если TBD — тест-агенты используют минимальный mock

### Важно
Если поле TBD — тест-агент должен явно задокументировать допущение в [doc]-теге notes.
```

## Шаги

- [ ] Добавить раздел "API Contracts" в `docs/standards/v2/02-acceptance/01-write/base-plan.md`
- [ ] Добавить пункт в `docs/standards/v2/02-acceptance/01-write/base-checklist.md`: "API Contracts заполнен или TBD явно помечен"
- [ ] Обновить шаблон acceptance-write в скилле

## Источники

Азирафаль, Кроули, Гримм (ретро store/crud, 2026-05-05)
Рецидив из: RETRO store/connection, action 01-inv-comments.md
