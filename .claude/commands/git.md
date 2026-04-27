Git workflow: ветки, коммиты, поиск по истории, бэкап.

Аргументы: `$ARGUMENTS` — действие и параметры, например `commit`, `branch feature/store/query`, `backup`, `logsearch когда менялся файл X`

## /git commit

1. Запусти `git status` и `git diff` — посмотри что изменилось
2. Сгруппируй изменения по смысловым единицам (одна фича/фикс = один коммит)
3. Для каждой группы:
   - составь сообщение по формату: `<type>(<scope>): <что сделано>`
   - типы: `feat`, `fix`, `refactor`, `docs`, `test`, `chore`
   - покажи какие файлы войдут и сообщение — жди подтверждения
   - после подтверждения: `git add <files>` + `git commit -m "..."`

## /git branch

- Ветка на каждую фичу: `feature/<binary>/<feature>`, например `feature/store/query`
- При старте фичи — создать ветку от `main`
- При завершении — показать diff с `main`, напомнить про merge

## /git logsearch

Аргументы: свободное описание что ищем, например:
- `когда последний раз менялся store/query.go`
- `когда добавили функцию ParseFilter`
- `как менялся файл internal/store/query.go со временем`
- `когда удалили поддержку subtype из DSL`

1. Определи тип поиска по описанию:
   - **по файлу** → `git log --follow --oneline <file>`
   - **по строке/функции** → `git log -S "<term>" --oneline` или `git log -G "<regex>"`
   - **история изменений файла** → `git log --follow -p <file>`
   - **по сообщению коммита** → `git log --oneline --grep="<term>"`
2. Запусти подходящую команду
3. Покажи результат осмысленно — не raw git output:
   - когда и что менялось
   - краткое содержание каждого релевантного коммита
4. Предложи `git show <hash>` для деталей если нашёл

## /git backup

1. Прочитай путь бэкапа из `.claude/project.yaml` (`backup.path`)
2. Создай bundle: `git bundle create <backup.path>/<repo>-<YYYY-MM-DD>.bundle --all`
3. Если есть worktrees — скопируй: `cp -r <worktree-path> <backup.path>/worktrees/`
4. Сообщи что сохранено и куда
