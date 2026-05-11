---
created: 2026-05-06T04:38:54Z
updated: 2026-05-11T20:11:02Z
---

# Global Backlog

Глобальный бэклог всех action items из ретро.
ID стабилен — не меняется при переносах между фичами.
Происхождение видно из ссылки на action-файл.

Ревьюится в начале каждой фичи (`00-research/01-interview`).
Пополняется в конце каждого ретро (`06-retro/04-solve`).

## Открытые

- [ ] **BL-001** — API Contracts в шаблоне acceptance-write → [crud/01-api-contracts-acceptance](../0002-store-crud/artifacts/actions/01-api-contracts-acceptance.md)
- [ ] **BL-002** — Разделы "Контракт" и "Реализация на Go" в шаблоне spec → [crud/02-spec-contract-vs-impl](../0002-store-crud/artifacts/actions/02-spec-contract-vs-impl.md)
- [ ] **BL-003** — Пункт "инварианты валидации" в шаблоне acceptance → [crud/03-invariants-in-acceptance](../0002-store-crud/artifacts/actions/03-invariants-in-acceptance.md)
- [ ] **BL-004** — 01-recall: find по всем notes/complaints в processes/ → [crud/04-recall-aggregation](../0002-store-crud/artifacts/actions/04-recall-aggregation.md)
- [x] **BL-005** — Backlog механизм: 04-solve пополняет BACKLOG.md, 00-research ревьюит → [crud/05-backlog-mechanism](../0002-store-crud/artifacts/actions/05-backlog-mechanism.md) *(закрыто 2026-05-06)*
- [ ] **BL-006** — Notes/complaints как обязательный финальный шаг в write-скиллах → [crud/06-notes-complaints-mandatory](../0002-store-crud/artifacts/actions/06-notes-complaints-mandatory.md)
- [ ] **BL-007** — Gate-блокировка pipeline: вышестоящие шаги не в pending → [crud/07-pipeline-gate](../0002-store-crud/artifacts/actions/07-pipeline-gate.md)
- [ ] **BL-008** — Gate перед 06-retro: все предыдущие шаги done или явное решение → [crud/08-retro-gate](../0002-store-crud/artifacts/actions/08-retro-gate.md)
- [ ] **BL-009** — Fix-report обязателен перед переводом в done после failed → [crud/09-fix-report-mandatory](../0002-store-crud/artifacts/actions/09-fix-report-mandatory.md)
- [x] **BL-010** — scripts/feature-status.sh: компактный статус всех шагов → [crud/10-feature-status-script](../0002-store-crud/artifacts/actions/10-feature-status-script.md) *(закрыто 2026-05-11)*
- [ ] **BL-011** — scripts/log-note.sh и log-complaint.sh: снизить барьер к ведению notes → [crud/11-log-note-script](../0002-store-crud/artifacts/actions/11-log-note-script.md)
- [ ] **BL-012** — Стандарт start/end в status-log: in-progress при старте, done при завершении → [crud/12-status-log-timestamps](../0002-store-crud/artifacts/actions/12-status-log-timestamps.md)
- [ ] **BL-013** — Раздел "Голос" в AGENT.md каждого агента → [crud/13-agent-voice](../0002-store-crud/artifacts/actions/13-agent-voice.md)
- [ ] **BL-014** — Раздел "Error types" в шаблоне spec → [crud/14-error-types-in-spec](../0002-store-crud/artifacts/actions/14-error-types-in-spec.md)
- [ ] **BL-015** — spec.md на уровне группы: единый источник истины → [crud/15-spec-md-single-source](../0002-store-crud/artifacts/actions/15-spec-md-single-source.md)
- [ ] **BL-016** — Правило clarification 03-tests/04-code: несовместимость API → clarification → [crud/16-tests-code-boundary](../0002-store-crud/artifacts/actions/16-tests-code-boundary.md)
- [ ] **BL-017** — Валидация каскадного сброса при изменении upstream → [crud/17-cascade-reset-validation](../0002-store-crud/artifacts/actions/17-cascade-reset-validation.md)
- [ ] **BL-018** — Auto-proceed failed→fix без паузы когда следующий шаг однозначен → [crud/18-auto-proceed-failed-fix](../0002-store-crud/artifacts/actions/18-auto-proceed-failed-fix.md)
- [ ] **BL-019** — docs/standards/go/ явно в промпт ревьюера → [crud/19-go-standards-path](../0002-store-crud/artifacts/actions/19-go-standards-path.md)
- [ ] **BL-020** — Оркестратор передаёт список файлов ревьюеру → [crud/20-reviewer-file-list](../0002-store-crud/artifacts/actions/20-reviewer-file-list.md)
- [ ] **BL-021** — Стандарт документирования шага для context-recovery → [crud/21-step-documentation-standard](../0002-store-crud/artifacts/actions/21-step-documentation-standard.md)
- [ ] **BL-022** — go build ./... отдельным шагом: отличать compile-fail от test-fail → [crud/22-go-build-separate-step](../0002-store-crud/artifacts/actions/22-go-build-separate-step.md)
- [ ] **BL-023** — RFC: Танк/Бо без Write tools → [crud/23-rfc-write-tools](../0002-store-crud/artifacts/actions/23-rfc-write-tools.md)
- [ ] **BL-024** — Fail fast как дефолт в AGENT.md и write-скиллах → [crud/24-fail-fast-default](../0002-store-crud/artifacts/actions/24-fail-fast-default.md)
- [ ] **BL-025** — Стандарт datetime через date -u задокументировать → [crud/25-datetime-standard](../0002-store-crud/artifacts/actions/25-datetime-standard.md)
- [ ] **BL-026** — Управление контекстом при восстановлении сессии: стратегия сжатия, не только STATUS.md → [crud/26-compact-feature-summary](../0002-store-crud/artifacts/actions/26-compact-feature-summary.md)
- [x] **BL-027** — Стандарт структуры tasks/<feature>/: канонический layout → [crud/27-folder-structure-standard](../0002-store-crud/artifacts/actions/27-folder-structure-standard.md) *(закрыто 2026-05-06)*
- [ ] **BL-028** — Стандарт frontmatter + валидация обязательных полей → [crud/28-frontmatter-standard](../0002-store-crud/artifacts/actions/28-frontmatter-standard.md)
- [ ] **BL-029** — `// inv:` стандарт документирования инвариантов → [conn/c01-inv-comments](../0002-store-crud/artifacts/actions/c01-inv-comments.md)
- [ ] **BL-030** — Spec sync gate в code-review: изменение конфига → обновить spec+acceptance → [conn/c02-spec-sync-gate](../0002-store-crud/artifacts/actions/c02-spec-sync-gate.md)
- [ ] **BL-031** — slog/waitForMsg стандарт в TDD → [conn/c03-slog-tdd-standard](../0002-store-crud/artifacts/actions/c03-slog-tdd-standard.md)
- [ ] **BL-032** — Чек-лист сетевых фич в acceptance: full-duplex, backpressure → [conn/c04-network-acceptance-checklist](../0002-store-crud/artifacts/actions/c04-network-acceptance-checklist.md)
- [ ] **BL-033** — Тест-симулякр: внешний таймаут < внутреннего → [conn/c05-test-simulacrum-checklist](../0002-store-crud/artifacts/actions/c05-test-simulacrum-checklist.md)
- [ ] **BL-034** — main.go проверка в code-review: все файлы пакета в каждом раунде → [conn/c06-code-review-all-files](../0002-store-crud/artifacts/actions/c06-code-review-all-files.md)
- [ ] **BL-035** — RFC: Харли + Agent tool для автономной оркестрации → [conn/c07-harley-agent-tool-rfc](../0002-store-crud/artifacts/actions/c07-harley-agent-tool-rfc.md)
- [ ] **BL-036** — Обновить Go до 1.26.2 → [general/36-go-upgrade](actions/36-go-upgrade.md)
- [ ] **BL-037** — Система флоу: формат описания, инициализация скелета, next/previous-step в frontmatter → [general/37-flow-system](actions/37-flow-system.md)
- [x] **BL-038** — Мигрировать репозиторий на целевую структуру проекта → [general/38-project-structure-migration](actions/38-project-structure-migration.md) *(закрыто 2026-05-06)*
- [x] **BL-039** — Написать `docs/standards/v2/stages.md`: каталог всех этапов и шагов workflow → [general/39-stages-spec](actions/39-stages-spec.md) *(закрыто 2026-05-06)*
- [ ] **BL-040** — Вынести логику `store/cmd/` в internal-пакеты (config, handler, server) → [general/40-store-cmd-internal-packages](actions/40-store-cmd-internal-packages.md)
- [ ] **BL-041** — Авторство в начале файлов: кто создавал и дополнял (агент + этап) → [general/41-module-authorship](actions/41-module-authorship.md)
- [ ] **BL-042** — Флоу для задач не привязанных к фиче + какие ещё флоу могут быть → [general/42-non-feature-workflow](actions/42-non-feature-workflow.md)
- [ ] **BL-043** — scripts/validate-status-log-meta.sh: проверка frontmatter status-log.md → [general/43-validate-status-log-meta](actions/43-validate-status-log-meta.md)
- [ ] **BL-044** — scripts/add-status-log-meta.sh: добавление frontmatter в существующие status-log.md → [general/44-migrate-status-log-meta](actions/44-migrate-status-log-meta.md)
- [ ] **BL-045** — Инструкции агентам: создавать brief/report до вызова begin/end-step → [general/45-agent-briefs-instructions](actions/45-agent-briefs-instructions.md)
- [ ] **BL-046** — Интегрировать скрипты в инструкции к скиллам и агентам → [general/46-integrate-scripts-into-instructions](actions/46-integrate-scripts-into-instructions.md)

## Закрытые

| ID | Название | Закрыто | Причина |
|---|---|---|---|
| BL-005 | Backlog механизм | 2026-05-06 | реализовано |
| BL-010 | scripts/feature-status.py | 2026-05-11 | `scripts/feature-status.py` |
| BL-027 | Стандарт структуры tasks/<feature>/ | 2026-05-06 | `docs/standards/v2/project-structure.md` |
| BL-038 | Мигрировать репозиторий на целевую структуру проекта | 2026-05-06 | выполнено |
| BL-039 | Каталог этапов и шагов stages.md | 2026-05-06 | `docs/standards/v2/stages.md` |
