---
created: 2026-05-06T04:38:54Z
updated: 2026-05-13T21:09:00Z
---

# Global Backlog

Глобальный бэклог всех action items из ретро.
ID стабилен — не меняется при переносах между фичами.
Происхождение видно из ссылки на action-файл.

Ревьюится в начале каждой фичи (`00-research/01-interview`).
Пополняется в конце каждого ретро (`06-retro/04-solve`).

## Открытые

- [ ] **RE-001** — API Contracts в шаблоне acceptance-write → [crud/01-api-contracts-acceptance](../0002-store-crud/artifacts/actions/01-api-contracts-acceptance.md)
- [ ] **RE-002** — Разделы "Контракт" и "Реализация на Go" в шаблоне spec → [crud/02-spec-contract-vs-impl](../0002-store-crud/artifacts/actions/02-spec-contract-vs-impl.md)
- [ ] **RE-003** — Пункт "инварианты валидации" в шаблоне acceptance → [crud/03-invariants-in-acceptance](../0002-store-crud/artifacts/actions/03-invariants-in-acceptance.md)
- [ ] **RE-004** — 01-recall: find по всем notes/complaints в processes/ → [crud/04-recall-aggregation](../0002-store-crud/artifacts/actions/04-recall-aggregation.md)
- [x] **RE-005** — Backlog механизм: 04-solve пополняет BACKLOG.md, 00-research ревьюит → [crud/05-backlog-mechanism](../0002-store-crud/artifacts/actions/05-backlog-mechanism.md) *(закрыто 2026-05-06)*
- [x] **RE-006** — Notes/complaints как обязательный финальный шаг в write-скиллах → [crud/06-notes-complaints-mandatory](../0002-store-crud/artifacts/actions/06-notes-complaints-mandatory.md) *(закрыто 2026-05-13)*
- [ ] **RE-007** — Gate-блокировка pipeline: вышестоящие шаги не в pending → [crud/07-pipeline-gate](../0002-store-crud/artifacts/actions/07-pipeline-gate.md)
- [ ] **RE-008** — Gate перед 06-retro: все предыдущие шаги done или явное решение → [crud/08-retro-gate](../0002-store-crud/artifacts/actions/08-retro-gate.md)
- [x] **RE-009** — Fix-report обязателен перед переводом в done после failed → [crud/09-fix-report-mandatory](../0002-store-crud/artifacts/actions/09-fix-report-mandatory.md) *(закрыто 2026-05-13)*
- [x] **RE-010** — scripts/feature-status.sh: компактный статус всех шагов → [crud/10-feature-status-script](../0002-store-crud/artifacts/actions/10-feature-status-script.md) *(закрыто 2026-05-11)*
- [x] **RE-011** — scripts/log-note.sh и log-complaint.sh: снизить барьер к ведению notes → [crud/11-log-note-script](../0002-store-crud/artifacts/actions/11-log-note-script.md) *(закрыто 2026-05-13)*
- [x] **RE-012** — Стандарт start/end в status-log: in-progress при старте, done при завершении → [crud/12-status-log-timestamps](../0002-store-crud/artifacts/actions/12-status-log-timestamps.md) *(закрыто 2026-05-13)*
- [x] **RE-013** — Раздел "Голос" в AGENT.md каждого агента + убрать over-constraining из SOUL.md → [crud/13-agent-voice](../0002-store-crud/artifacts/actions/13-agent-voice.md) *(закрыто 2026-05-13)*
- [ ] **RE-014** — Раздел "Error types" в шаблоне spec → [crud/14-error-types-in-spec](../0002-store-crud/artifacts/actions/14-error-types-in-spec.md)
- [ ] **RE-015** — spec.md на уровне группы: единый источник истины → [crud/15-spec-md-single-source](../0002-store-crud/artifacts/actions/15-spec-md-single-source.md)
- [ ] **RE-016** — Правило clarification 03-tests/04-code: несовместимость API → clarification → [crud/16-tests-code-boundary](../0002-store-crud/artifacts/actions/16-tests-code-boundary.md)
- [ ] **RE-017** — Валидация каскадного сброса при изменении upstream → [crud/17-cascade-reset-validation](../0002-store-crud/artifacts/actions/17-cascade-reset-validation.md)
- [ ] **RE-018** — Auto-proceed failed→fix без паузы когда следующий шаг однозначен → [crud/18-auto-proceed-failed-fix](../0002-store-crud/artifacts/actions/18-auto-proceed-failed-fix.md)
- [ ] **RE-019** — docs/standards/go/ явно в промпт ревьюера → [crud/19-go-standards-path](../0002-store-crud/artifacts/actions/19-go-standards-path.md)
- [ ] **RE-020** — Оркестратор передаёт список файлов ревьюеру → [crud/20-reviewer-file-list](../0002-store-crud/artifacts/actions/20-reviewer-file-list.md)
- [ ] **RE-021** — Стандарт документирования шага для context-recovery → [crud/21-step-documentation-standard](../0002-store-crud/artifacts/actions/21-step-documentation-standard.md)
- [ ] **RE-022** — go build ./... отдельным шагом: отличать compile-fail от test-fail → [crud/22-go-build-separate-step](../0002-store-crud/artifacts/actions/22-go-build-separate-step.md)
- [x] **RE-023** — RFC: Танк/Бо без Write tools → [crud/23-rfc-write-tools](../0002-store-crud/artifacts/actions/23-rfc-write-tools.md) *(закрыто 2026-05-13)*
- [ ] **RE-024** — Fail fast как дефолт в AGENT.md и write-скиллах → [crud/24-fail-fast-default](../0002-store-crud/artifacts/actions/24-fail-fast-default.md)
- [ ] **RE-025** — Стандарт datetime через date -u задокументировать → [crud/25-datetime-standard](../0002-store-crud/artifacts/actions/25-datetime-standard.md)
- [ ] **RE-026** — Управление контекстом при восстановлении сессии: стратегия сжатия, не только STATUS.md → [crud/26-compact-feature-summary](../0002-store-crud/artifacts/actions/26-compact-feature-summary.md)
- [x] **RE-027** — Стандарт структуры tasks/<feature>/: канонический layout → [crud/27-folder-structure-standard](../0002-store-crud/artifacts/actions/27-folder-structure-standard.md) *(закрыто 2026-05-06)*
- [ ] **RE-028** — Стандарт frontmatter + валидация обязательных полей → [crud/28-frontmatter-standard](../0002-store-crud/artifacts/actions/28-frontmatter-standard.md)
- [ ] **RE-029** — `// inv:` стандарт документирования инвариантов → [conn/01-inv-comments](../0001-store-connection/artifacts/actions/01-inv-comments.md)
- [ ] **RE-030** — Spec sync gate в code-review: изменение конфига → обновить spec+acceptance → [conn/02-spec-sync-gate](../0001-store-connection/artifacts/actions/02-spec-sync-gate.md)
- [ ] **RE-031** — slog/waitForMsg стандарт в TDD → [conn/03-slog-tdd-standard](../0001-store-connection/artifacts/actions/03-slog-tdd-standard.md)
- [ ] **RE-032** — Чек-лист сетевых фич в acceptance: full-duplex, backpressure → [conn/04-network-acceptance-checklist](../0001-store-connection/artifacts/actions/04-network-acceptance-checklist.md)
- [ ] **RE-033** — Тест-симулякр: внешний таймаут < внутреннего → [conn/05-test-simulacrum-checklist](../0001-store-connection/artifacts/actions/05-test-simulacrum-checklist.md)
- [ ] **RE-034** — main.go проверка в code-review: все файлы пакета в каждом раунде → [conn/07-code-review-all-files](../0001-store-connection/artifacts/actions/07-code-review-all-files.md)
- [x] **RE-035** — RFC: Харли + Agent tool для автономной оркестрации → [conn/08-harley-agent-tool-rfc](../0001-store-connection/artifacts/actions/08-harley-agent-tool-rfc.md) *(закрыто 2026-05-13)*
- [x] **RE-047** — Адоптировать практики из karpathy CLAUDE.md в скиллы и стандарты → [conn/09-adopt-karpathy](../0001-store-connection/artifacts/actions/09-adopt-karpathy-claude-md.md) *(закрыто 2026-05-13)*
- [ ] **RE-036** — Обновить Go до 1.26.2 → [general/36-go-upgrade](actions/36-go-upgrade.md)
- [x] **RE-037** — Система флоу: формат описания, инициализация скелета, next/previous-step в frontmatter → [general/37-flow-system](actions/37-flow-system.md) *(закрыто 2026-05-13)*
- [x] **RE-038** — Мигрировать репозиторий на целевую структуру проекта → [general/38-project-structure-migration](actions/38-project-structure-migration.md) *(закрыто 2026-05-06)*
- [x] **RE-039** — Написать `docs/standards/v2/stages.md`: каталог всех этапов и шагов workflow → [general/39-stages-spec](actions/39-stages-spec.md) *(закрыто 2026-05-06)*
- [ ] **RE-040** — Вынести логику `store/cmd/` в internal-пакеты (config, handler, server) → [general/40-store-cmd-internal-packages](actions/40-store-cmd-internal-packages.md)
- [ ] **RE-041** — Авторство в начале файлов: кто создавал и дополнял (агент + этап) → [general/41-module-authorship](actions/41-module-authorship.md)
- [ ] **RE-042** — Флоу для задач не привязанных к фиче + какие ещё флоу могут быть → [general/42-non-feature-workflow](actions/42-non-feature-workflow.md)
- [ ] **RE-043** — scripts/validate-status-log-meta.sh: проверка frontmatter status-log.md → [general/43-validate-status-log-meta](actions/43-validate-status-log-meta.md)
- [ ] **RE-044** — scripts/add-status-log-meta.sh: добавление frontmatter в существующие status-log.md → [general/44-migrate-status-log-meta](actions/44-migrate-status-log-meta.md)
- [x] **RE-045** — Инструкции агентам: создавать brief/report до вызова begin/end-step → [general/45-agent-briefs-instructions](actions/45-agent-briefs-instructions.md) *(закрыто 2026-05-13)*
- [x] **RE-046** — Интегрировать скрипты в инструкции к скиллам и агентам → [general/46-integrate-scripts-into-instructions](actions/46-integrate-scripts-into-instructions.md) *(закрыто 2026-05-13)*

## Закрытые

| ID | Название | Закрыто | Причина |
|---|---|---|---|
| RE-005 | Backlog механизм | 2026-05-06 | реализовано |
| RE-010 | scripts/feature-status.py | 2026-05-11 | `scripts/feature-status.py` |
| RE-027 | Стандарт структуры tasks/<feature>/ | 2026-05-06 | `docs/standards/v2/project-structure.md` |
| RE-038 | Мигрировать репозиторий на целевую структуру проекта | 2026-05-06 | выполнено |
| RE-039 | Каталог этапов и шагов stages.md | 2026-05-06 | `docs/standards/v2/stages.md` |
| RE-013 | Раздел "Голос" в AGENT.md каждого агента | 2026-05-13 | `## Голос` добавлен в 8 AGENT.md; "Signs of drift" разжаты во всех SOUL.md |
| RE-023 | RFC: Танк/Бо без Write tools | 2026-05-13 | Write+Edit добавлены Бо и Герману; у Танка уже был |
| RE-006 | Notes/complaints как обязательный шаг в write-скиллах | 2026-05-13 | Секция "Финальный шаг" добавлена в 5 base-plan.md (spec, acceptance, tests, code, docs) |
| RE-009 | Fix-report трейсабилити через in-response-to | 2026-05-13 | `scripts/set-step-status.py` валидирует in-response-to в 03-fix; brief-NNN.md в stages-standard; 5 base-plan.md обновлены |
| RE-011 | scripts/log-note и log-complaint | 2026-05-13 | `scripts/log-note.py`, `scripts/log-complaint.py`; интегрированы в AGENT.md и base-plan.md |
| RE-012 | Стандарт start/end в status-log | 2026-05-13 | `scripts/set-step-status.py` валидирует порядок; `## Скрипты` добавлен в 15 base-plan.md |
| RE-037 | Система флоу | 2026-05-13 | `docs/standards/v2/flows/`, `scripts/scaffold-feature.py`, frontmatter в status-log.md |
| RE-045 | Инструкции агентам про брифы | 2026-05-13 | `## Скрипты` добавлен в 8 AGENT.md с порядком brief→in-progress→report→done |
| RE-046 | Интеграция скриптов в скиллы и агентов | 2026-05-13 | `## Скрипты` в 8 AGENT.md + 15 base-plan.md; `## Финальный шаг` с командами |
| RE-035 | Харли + Agent tool для оркестрации | 2026-05-13 | решено через `_Режим: Харли Куин_` в CLAUDE.md — main thread с Agent tool |
| RE-047 | Karpathy practices в скиллы | 2026-05-13 | все 7 механик в base-plan.md: surgical changes, assumptions, strong criteria, Red run, verify |
