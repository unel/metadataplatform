# Action: стандарт frontmatter и валидация

## Проблема

Поля frontmatter в md-файлах процессов то добавляются, то пропадают. Структура непоследовательна: в одном файле `date` + `created`, в другом только `date`, в третьем ни того ни другого.

## Решение

Зафиксировать обязательные поля для каждого типа файла:

**report-NNN.md:**
```yaml
purpose, process, run, date, created, status, agent
```

**status-log.md:** формат строк `# datetime — статус`

**notes-<agent>.md / complaints-<agent>.md:** свободный формат, теги обязательны

Валидация в `scripts/feature-status.sh --validate`.

## Шаги

- [ ] Написать `docs/standards/v2/frontmatter.md`
- [ ] Добавить валидацию в scripts

## Источники

Пользователь (ретро store/crud, 2026-05-05)
