# STRUCTURE — 0001-store-connection

Задача выполнена до введения workflow v2. Структура legacy: нет `stages/`, рабочие файлы хранятся плоско в корне.

## Артефакты

```
artifacts/
├── spec.md           — спека задачи
├── acceptance.md     — acceptance criteria
├── retro.md          — итог ретро
└── actions/          — action items из ретро
```

## Исторические документы (корень)

```
notes-<agent>.md      — заметки агентов
complaints-<agent>.md — жалобы агентов
code-review-N.md      — ревью кода
code-fix-N.md         — отчёты фиксов кода
test-review.md        — ревью тестов
test-fix-N.md         — отчёты фиксов тестов
test-run-report.md    — отчёт прогона тестов
status.md             — итоговый статус (legacy, не status-log.md)
```
