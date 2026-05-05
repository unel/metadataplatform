---
source: tasks/store/crud/processes/06-retro/04-solve/report-002.md
actions: tasks/store/crud/processes/06-retro/actions/
date: 2026-05-05T18:20:12Z
updated: 2026-05-05T18:20:12Z
---

# Backlog из ретро store/crud

Проверять в начале `00-research` следующей фичи.
Детали каждого пункта: `tasks/store/crud/processes/06-retro/actions/`

## Высокий приоритет (до старта следующей фичи)

- [ ] **API Contracts в шаблоне acceptance-write** — сигнатуры конструкторов, интерфейсов, логгера, error types; TBD явно помечается → [01-api-contracts-acceptance.md](../store/crud/processes/06-retro/actions/01-api-contracts-acceptance.md)
- [ ] **Разделы "Контракт" и "Реализация на Go" в шаблоне spec** — требования отдельно от Go-специфики → [02-spec-contract-vs-impl.md](../store/crud/processes/06-retro/actions/02-spec-contract-vs-impl.md)
- [ ] **Пункт "инварианты валидации" в шаблоне acceptance** — обязательно для операций с пользовательскими данными → [03-invariants-in-acceptance.md](../store/crud/processes/06-retro/actions/03-invariants-in-acceptance.md)
- [ ] **01-recall: find по всем notes/complaints в processes/** — не терять файлы из промежуточных шагов → [04-recall-aggregation.md](../store/crud/processes/06-retro/actions/04-recall-aggregation.md)
- [ ] **Backlog механизм** — 04-solve создаёт `tasks/_backlog/`; 00-research проверяет backlog предыдущей фичи → [05-backlog-mechanism.md](../store/crud/processes/06-retro/actions/05-backlog-mechanism.md)
- [ ] **Notes/complaints как обязательный финальный шаг** — в каждом write-скилле → [06-notes-complaints-mandatory.md](../store/crud/processes/06-retro/actions/06-notes-complaints-mandatory.md)
- [ ] **Gate-блокировка pipeline** — вышестоящие шаги не в pending перед запуском следующего → [07-pipeline-gate.md](../store/crud/processes/06-retro/actions/07-pipeline-gate.md)
- [ ] **Gate перед 06-retro** — все предыдущие шаги done или явное решение → [08-retro-gate.md](../store/crud/processes/06-retro/actions/08-retro-gate.md)
- [ ] **Fix-report обязателен** — без него нельзя двигать статус в done после failed → [09-fix-report-mandatory.md](../store/crud/processes/06-retro/actions/09-fix-report-mandatory.md)
- [ ] **scripts/feature-status.sh** — компактный статус всех шагов → [10-feature-status-script.md](../store/crud/processes/06-retro/actions/10-feature-status-script.md)
- [ ] **scripts/log-note.sh и log-complaint.sh** — снизить барьер к ведению notes → [11-log-note-script.md](../store/crud/processes/06-retro/actions/11-log-note-script.md)
- [ ] **Стандарт start/end в status-log** — in-progress при старте, done при завершении → [12-status-log-timestamps.md](../store/crud/processes/06-retro/actions/12-status-log-timestamps.md)
- [ ] **Раздел "Голос" в AGENT.md каждого агента** — характер, стиль, [whatever] и юмор как норма → [13-agent-voice.md](../store/crud/processes/06-retro/actions/13-agent-voice.md)

## Средний приоритет

- [ ] **Раздел "Error types" в шаблоне spec** — типы ошибок до code-write → [14-error-types-in-spec.md](../store/crud/processes/06-retro/actions/14-error-types-in-spec.md)
- [ ] **spec.md на уровне группы** — единый источник истины, write создаёт, fix обновляет → [15-spec-md-single-source.md](../store/crud/processes/06-retro/actions/15-spec-md-single-source.md)
- [ ] **Правило clarification 03-tests/04-code** — несовместимость API → clarification, не фикс в 04-code → [16-tests-code-boundary.md](../store/crud/processes/06-retro/actions/16-tests-code-boundary.md)
- [ ] **Валидация каскадного сброса** — проверка downstream статусов при изменении upstream → [17-cascade-reset-validation.md](../store/crud/processes/06-retro/actions/17-cascade-reset-validation.md)
- [ ] **Auto-proceed failed→fix** — без паузы когда следующий шаг однозначен → [18-auto-proceed-failed-fix.md](../store/crud/processes/06-retro/actions/18-auto-proceed-failed-fix.md)
- [ ] **docs/standards/go/ явно в промпт ревьюера** — не ссылаться на несуществующую extensions/ → [19-go-standards-path.md](../store/crud/processes/06-retro/actions/19-go-standards-path.md)
- [ ] **Оркестратор передаёт список файлов ревьюеру** — не угадывать через ls → [20-reviewer-file-list.md](../store/crud/processes/06-retro/actions/20-reviewer-file-list.md)
- [ ] **Стандарт документирования шага для context-recovery** — каждый артефакт самодостаточен → [21-step-documentation-standard.md](../store/crud/processes/06-retro/actions/21-step-documentation-standard.md)
- [ ] **go build ./... отдельным шагом** — отличать compile-fail от test-fail → [22-go-build-separate-step.md](../store/crud/processes/06-retro/actions/22-go-build-separate-step.md)
- [ ] **RFC: Танк/Бо без Write tools** — решить как намеренный дизайн или добавить инструменты → [23-rfc-write-tools.md](../store/crud/processes/06-retro/actions/23-rfc-write-tools.md)

## Низкий приоритет

- [ ] **Fail fast как дефолт** — явный пункт в AGENT.md и write-скиллах → [24-fail-fast-default.md](../store/crud/processes/06-retro/actions/24-fail-fast-default.md)
- [ ] **Стандарт datetime через date -u** — задокументировать в стандартах → [25-datetime-standard.md](../store/crud/processes/06-retro/actions/25-datetime-standard.md)
- [ ] **Compact summary STATUS.md** — один файл вместо всех status-log → [26-compact-feature-summary.md](../store/crud/processes/06-retro/actions/26-compact-feature-summary.md)
- [ ] **Стандарт структуры tasks/`<feature>`/** — зафиксировать канонический layout → [27-folder-structure-standard.md](../store/crud/processes/06-retro/actions/27-folder-structure-standard.md)
- [ ] **Стандарт frontmatter + валидация** — обязательные поля, проверка в скриптах → [28-frontmatter-standard.md](../store/crud/processes/06-retro/actions/28-frontmatter-standard.md)

## Перенесено из RETRO connection (невыполненные)

_Оригиналы: `tasks/store/connection/RETRO/actions/`_

- [ ] **`// inv:` стандарт** — документирование инвариантов в lifecycle/concurrency блоках → [c01-inv-comments.md](../store/crud/processes/06-retro/actions/c01-inv-comments.md)
- [ ] **Spec sync gate в code-review** — изменение конфига/поведения → обновить spec.md + acceptance.md → [c02-spec-sync-gate.md](../store/crud/processes/06-retro/actions/c02-spec-sync-gate.md)
- [ ] **slog/waitForMsg стандарт в TDD** — structured logs + grep-конвенция → [c03-slog-tdd-standard.md](../store/crud/processes/06-retro/actions/c03-slog-tdd-standard.md)
- [ ] **Чек-лист сетевых фич в acceptance** — full-duplex, backpressure, тестируемость → [c04-network-acceptance-checklist.md](../store/crud/processes/06-retro/actions/c04-network-acceptance-checklist.md)
- [ ] **Тест-симулякр пункт в test-review** — внешний таймаут < внутреннего → [c05-test-simulacrum-checklist.md](../store/crud/processes/06-retro/actions/c05-test-simulacrum-checklist.md)
- [ ] **main.go проверка в code-review** — все файлы пакета в каждом раунде → [c06-code-review-all-files.md](../store/crud/processes/06-retro/actions/c06-code-review-all-files.md)
- [ ] **RFC: Харли + Agent tool** — автономная оркестрация → [c07-harley-agent-tool-rfc.md](../store/crud/processes/06-retro/actions/c07-harley-agent-tool-rfc.md)

## Выполнено

*(заполняется в 00-research следующей фичи)*
