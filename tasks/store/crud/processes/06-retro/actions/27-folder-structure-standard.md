# Action: стандарт структуры tasks/<feature>/

## Проблема

Структура папок поплыла: notes/complaints агентов лежали в `processes/`, `processes/03-tests/01-write/` — вместо правильных мест. Файлы появляются в неожиданных местах.

## Решение

Зафиксировать стандарт в `docs/standards/v2/`:

```
tasks/<feature>/
  processes/
    <step>/
      status-log.md
      README.md
      report-NNN.md
      notes-<agent>.md       ← здесь, в своём шаге
      complaints-<agent>.md  ← здесь, в своём шаге
  RETRO/
    actions/
  _backlog/  ← нет, это tasks/_backlog/retro-<feature>.md
  README.md
  RETRO.md
  STATUS.md
```

## Шаги

- [ ] Написать `docs/standards/v2/structure.md` — каноническая структура фичи
- [ ] Добавить валидацию в `scripts/feature-status.sh`

## Источники

Пользователь (ретро store/crud, 2026-05-05)
