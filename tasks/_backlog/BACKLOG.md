---
created: 2026-05-06T04:38:54Z
updated: 2026-05-06T04:38:54Z
---

# Global Backlog

Глобальный бэклог всех action items из ретро.
ID стабилен — не меняется при переносах между фичами.
Происхождение видно из ссылки на action-файл.

Ревьюится в начале каждой фичи (`00-research/01-interview`).
Пополняется в конце каждого ретро (`06-retro/04-solve`).

## Открытые

- [ ] **BL-001** — API Contracts в шаблоне acceptance-write → [crud/01-api-contracts-acceptance](../store/crud/processes/06-retro/actions/01-api-contracts-acceptance.md)
- [ ] **BL-002** — Разделы "Контракт" и "Реализация на Go" в шаблоне spec → [crud/02-spec-contract-vs-impl](../store/crud/processes/06-retro/actions/02-spec-contract-vs-impl.md)
- [ ] **BL-003** — Пункт "инварианты валидации" в шаблоне acceptance → [crud/03-invariants-in-acceptance](../store/crud/processes/06-retro/actions/03-invariants-in-acceptance.md)
- [ ] **BL-004** — 01-recall: find по всем notes/complaints в processes/ → [crud/04-recall-aggregation](../store/crud/processes/06-retro/actions/04-recall-aggregation.md)
- [x] **BL-005** — Backlog механизм: 04-solve пополняет BACKLOG.md, 00-research ревьюит → [crud/05-backlog-mechanism](../store/crud/processes/06-retro/actions/05-backlog-mechanism.md) *(закрыто 2026-05-06)*
- [ ] **BL-006** — Notes/complaints как обязательный финальный шаг в write-скиллах → [crud/06-notes-complaints-mandatory](../store/crud/processes/06-retro/actions/06-notes-complaints-mandatory.md)
- [ ] **BL-007** — Gate-блокировка pipeline: вышестоящие шаги не в pending → [crud/07-pipeline-gate](../store/crud/processes/06-retro/actions/07-pipeline-gate.md)
- [ ] **BL-008** — Gate перед 06-retro: все предыдущие шаги done или явное решение → [crud/08-retro-gate](../store/crud/processes/06-retro/actions/08-retro-gate.md)
- [ ] **BL-009** — Fix-report обязателен перед переводом в done после failed → [crud/09-fix-report-mandatory](../store/crud/processes/06-retro/actions/09-fix-report-mandatory.md)
- [ ] **BL-010** — scripts/feature-status.sh: компактный статус всех шагов → [crud/10-feature-status-script](../store/crud/processes/06-retro/actions/10-feature-status-script.md)
- [ ] **BL-011** — scripts/log-note.sh и log-complaint.sh: снизить барьер к ведению notes → [crud/11-log-note-script](../store/crud/processes/06-retro/actions/11-log-note-script.md)
- [ ] **BL-012** — Стандарт start/end в status-log: in-progress при старте, done при завершении → [crud/12-status-log-timestamps](../store/crud/processes/06-retro/actions/12-status-log-timestamps.md)
- [ ] **BL-013** — Раздел "Голос" в AGENT.md каждого агента → [crud/13-agent-voice](../store/crud/processes/06-retro/actions/13-agent-voice.md)
- [ ] **BL-014** — Раздел "Error types" в шаблоне spec → [crud/14-error-types-in-spec](../store/crud/processes/06-retro/actions/14-error-types-in-spec.md)
- [ ] **BL-015** — spec.md на уровне группы: единый источник истины → [crud/15-spec-md-single-source](../store/crud/processes/06-retro/actions/15-spec-md-single-source.md)
- [ ] **BL-016** — Правило clarification 03-tests/04-code: несовместимость API → clarification → [crud/16-tests-code-boundary](../store/crud/processes/06-retro/actions/16-tests-code-boundary.md)
- [ ] **BL-017** — Валидация каскадного сброса при изменении upstream → [crud/17-cascade-reset-validation](../store/crud/processes/06-retro/actions/17-cascade-reset-validation.md)
- [ ] **BL-018** — Auto-proceed failed→fix без паузы когда следующий шаг однозначен → [crud/18-auto-proceed-failed-fix](../store/crud/processes/06-retro/actions/18-auto-proceed-failed-fix.md)
- [ ] **BL-019** — docs/standards/go/ явно в промпт ревьюера → [crud/19-go-standards-path](../store/crud/processes/06-retro/actions/19-go-standards-path.md)
- [ ] **BL-020** — Оркестратор передаёт список файлов ревьюеру → [crud/20-reviewer-file-list](../store/crud/processes/06-retro/actions/20-reviewer-file-list.md)
- [ ] **BL-021** — Стандарт документирования шага для context-recovery → [crud/21-step-documentation-standard](../store/crud/processes/06-retro/actions/21-step-documentation-standard.md)
- [ ] **BL-022** — go build ./... отдельным шагом: отличать compile-fail от test-fail → [crud/22-go-build-separate-step](../store/crud/processes/06-retro/actions/22-go-build-separate-step.md)
- [ ] **BL-023** — RFC: Танк/Бо без Write tools → [crud/23-rfc-write-tools](../store/crud/processes/06-retro/actions/23-rfc-write-tools.md)
- [ ] **BL-024** — Fail fast как дефолт в AGENT.md и write-скиллах → [crud/24-fail-fast-default](../store/crud/processes/06-retro/actions/24-fail-fast-default.md)
- [ ] **BL-025** — Стандарт datetime через date -u задокументировать → [crud/25-datetime-standard](../store/crud/processes/06-retro/actions/25-datetime-standard.md)
- [ ] **BL-026** — Compact summary STATUS.md: один файл вместо всех status-log → [crud/26-compact-feature-summary](../store/crud/processes/06-retro/actions/26-compact-feature-summary.md)
- [ ] **BL-027** — Стандарт структуры tasks/<feature>/: канонический layout → [crud/27-folder-structure-standard](../store/crud/processes/06-retro/actions/27-folder-structure-standard.md)
- [ ] **BL-028** — Стандарт frontmatter + валидация обязательных полей → [crud/28-frontmatter-standard](../store/crud/processes/06-retro/actions/28-frontmatter-standard.md)
- [ ] **BL-029** — `// inv:` стандарт документирования инвариантов → [conn/c01-inv-comments](../store/crud/processes/06-retro/actions/c01-inv-comments.md)
- [ ] **BL-030** — Spec sync gate в code-review: изменение конфига → обновить spec+acceptance → [conn/c02-spec-sync-gate](../store/crud/processes/06-retro/actions/c02-spec-sync-gate.md)
- [ ] **BL-031** — slog/waitForMsg стандарт в TDD → [conn/c03-slog-tdd-standard](../store/crud/processes/06-retro/actions/c03-slog-tdd-standard.md)
- [ ] **BL-032** — Чек-лист сетевых фич в acceptance: full-duplex, backpressure → [conn/c04-network-acceptance-checklist](../store/crud/processes/06-retro/actions/c04-network-acceptance-checklist.md)
- [ ] **BL-033** — Тест-симулякр: внешний таймаут < внутреннего → [conn/c05-test-simulacrum-checklist](../store/crud/processes/06-retro/actions/c05-test-simulacrum-checklist.md)
- [ ] **BL-034** — main.go проверка в code-review: все файлы пакета в каждом раунде → [conn/c06-code-review-all-files](../store/crud/processes/06-retro/actions/c06-code-review-all-files.md)
- [ ] **BL-035** — RFC: Харли + Agent tool для автономной оркестрации → [conn/c07-harley-agent-tool-rfc](../store/crud/processes/06-retro/actions/c07-harley-agent-tool-rfc.md)

## Закрытые

| ID | Название | Закрыто | Причина |
|---|---|---|---|
| BL-005 | Backlog механизм | 2026-05-06 | реализовано |
