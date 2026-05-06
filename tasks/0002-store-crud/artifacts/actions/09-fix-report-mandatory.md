# Action: fix-report обязателен после каждого failed

## Проблема

report-001.md для `04-code/01-write` не создан. Spec-fix-report отсутствует ни для connection, ни для crud. История итераций теряется. Рецидив из RETRO connection (action 06-fix-report-mandatory.md) — не реализован.

## Решение

В скиллах review: после результата `failed` — обязательно создать `<step>/fix-report-NNN.md` с описанием что именно не так. Без этого файла нельзя двигать статус в done.

Формат fix-report:
```markdown
---
step: <шаг>
run: N
verdict: failed
---
## Что не так
<список замечаний>
## Что нужно исправить
<конкретные действия>
```

## Шаги

- [ ] Добавить требование fix-report в `docs/standards/v2/*/02-review/base-plan.md` (все группы)
- [ ] Добавить пункт в чек-листы review: "fix-report создан если verdict failed"

## Источники

Танк, Ада, Гримм (ретро store/crud, 2026-05-05)
Рецидив из: `tasks/store/connection/RETRO/actions/06-fix-report-mandatory.md`
