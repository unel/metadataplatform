# Action: scripts/feature-status.sh

## Проблема

Статус фичи выясняется вручную: открывать каждый status-log.md по очереди. Барьер высок — агенты пишут статусы в неправильных местах (корневой status.md вместо status-log.md процессов). Восстановление состояния на пустом контексте занимает много времени и контекста.

## Решение

Написать `scripts/feature-status.sh <feature>`:

```
$ scripts/feature-status.sh store/crud

store/crud
  00-research/01-interview  ✓ done    2026-04-26T14:23:11Z
  00-research/02-web        ✓ done    2026-04-26T14:51:03Z
  01-spec/01-write          ✓ done    2026-04-27T09:12:44Z
  01-spec/02-review         ✗ failed  2026-04-27T11:03:22Z
  ...
```

Показывает: шаг, статус, timestamp последней записи в status-log.md.

## Шаги

- [ ] Написать `scripts/feature-status.sh`
- [ ] Добавить в README.md проекта раздел "Инструменты"

## Источники

Пользователь, Харли (ретро store/crud, 2026-05-05)
