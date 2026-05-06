# Action: мигрировать репозиторий на целевую структуру проекта

## Проблема

Текущая структура репо не соответствует стандарту из `docs/standards/v2/project-structure.md`:
- `cmd/store/` и `store/` в корне вместо `store/cmd/` и `store/`
- `tasks/store/connection/` и `tasks/store/crud/` вместо `tasks/NNNN-slug/`
- `artifacts/` отсутствует в живых задачах: spec/acceptance лежат не там, actions — в `processes/06-retro/actions/` вместо `artifacts/actions/`
- `store/connection` — legacy-формат без `processes/`
- `store/crud` — частично соответствует: есть `processes/`, но нет `artifacts/`, нет `STRUCTURE.md`, notes/complaints местами не в шагах

## Шаги (нужно уточнить при планировании)

- [ ] Переименовать `processes/` → `stages/` во всех задачах и обновить все ссылки (CLAUDE.md, скиллы, агенты)

- [ ] Переместить `cmd/store/` → `store/cmd/`, обновить импорты и go.mod
- [ ] Переименовать `tasks/store/connection` → `tasks/0001-store-connection`
- [ ] Переименовать `tasks/store/crud` → `tasks/0002-store-crud`
- [ ] В `0002-store-crud`: создать `artifacts/`, переместить spec/acceptance, перенести actions из `processes/06-retro/actions/` → `artifacts/actions/`, добавить `STRUCTURE.md`
- [ ] В `0001-store-connection`: решить — legacy или мигрировать под v2 с `stages/`
- [ ] Обновить PROJECT.md: схема компонентов расходится с целевой структурой (`cmd/store/` vs `store/cmd/`)
- [ ] Переместить `tasks/0001-store-connection/RETRO.md` → `artifacts/retro.md`
- [ ] Обновить `CLAUDE.md`: заменить `processes/` → `stages/` и путь `tasks/$FEATURE/processes/$PROCESS/status-log.md` → новый формат
- [ ] Определить судьбу `tasks/_backlog/retro-store-crud.md` — не покрыт стандартом
- [ ] Обновить ссылки в BACKLOG.md после переименований

## Источники

Обсуждение структуры проекта, 2026-05-06
