# Action: notes/complaints как обязательный финальный шаг в write-скиллах

## Проблема

Ни один агент за всю фичу store/crud не написал ни одной notes или complaint в ходе работы. Формат донесён только на 06-retro/01-recall. Детали первых прогонов размыты. Рецидив из RETRO connection.

Причина: notes/complaints не являются обязательным выходным артефактом ни одного шага. Агент пишет только если получил инструкцию писать — а в AGENT.md и скиллах её нет.

## Решение

В каждый write-скилл добавить финальный шаг:

```markdown
## Финальный шаг: notes/complaints

Запиши наблюдения по ходу работы:
- `processes/<step>/notes-<agent>.md` — конструктив: что заметил, что предлагаешь, что удивило
- `processes/<step>/complaints-<agent>.md` — сырые жалобы: что раздражало, что было неудобно

Теги: [doc], [propose], [friction], [miss], [rework], [whatever]
Формат свободный. Минимум — одна строка.
```

Скиллы: spec-write, acceptance-write, test-write, code-write, docs-write.

## Шаги

- [ ] Обновить base-plan.md в `docs/standards/v2/01-spec/01-write/`
- [ ] Обновить base-plan.md в `docs/standards/v2/02-acceptance/01-write/`
- [ ] Обновить base-plan.md в `docs/standards/v2/03-tests/01-write/`
- [ ] Обновить base-plan.md в `docs/standards/v2/04-code/01-write/`
- [ ] Обновить base-plan.md в `docs/standards/v2/05-docs/01-write/`

## Источники

Все участники (ретро store/crud, 2026-05-05)
Рецидив из: RETRO store/connection
