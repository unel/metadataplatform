# Codex Claude Bridge

Этот репозиторий уже содержит workflow для Claude в `.claude/`. Для работы в Codex используем следующую привязку.

## Что считается чем

- `.claude/commands/*.md` -> procedural skills / workflow prompts
- `.claude/agents/*.md` -> role prompts / execution styles
- `CLAUDE.md` -> repo-level orchestration rules
- `.claude/project.yaml` -> transitions, retries, proficiency, status policy

## Маппинг агентов

- `harley` -> оркестрация процесса, выбор следующего шага, контроль `status.md`
- `herman` -> исследование кодовой базы, поиск мест использования, восстановление текущего состояния
- `bo` -> внешний ресерч, документация, примеры, best practices
- `ada` -> реализация кода
- `grimm` -> ревью кода/спек/тестов
- `tank` -> спецификация, edge cases, NFR
- `crowley` -> негативные и edge-case тесты
- `aziraphale` -> happy-path и acceptance тесты

## Как это использовать в Codex

1. Если пользователь ссылается на Claude-команду, сначала открыть файл из `.claude/commands/`.
2. Если в команде указан исполнитель, открыть его файл из `.claude/agents/`.
3. Выполнить задачу средствами Codex, сохранив файловые конвенции и порядок шагов.
4. Не предполагать наличие Claude runtime, `Skill`, `Agent`, `SendMessage` и MCP только потому, что они упомянуты в текстах.

## Ограничение

Это мост интерпретации, а не полный порт Claude Code runtime. Он переносит роли, процессы и файловые соглашения, но не магию их тулчейна.
