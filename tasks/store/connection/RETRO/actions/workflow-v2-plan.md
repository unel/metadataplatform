---
purpose: План построения workflow v2 — текущий прогресс и что осталось
created: 2026-04-26T15:30
updated: 2026-04-26T18:00
---

# Workflow v2 — план и прогресс

## Что сделано

### Стандарты `docs/standards/v2/`

- [x] `00-research/01-interview/` — base-plan v1.1.0 (ладдеринг, 5 почему, противоречия), base-checklist, README, sources
- [x] `00-research/02-web/` — base-plan, base-checklist, README
- [x] `01-spec/01-write/` — base-plan (читает research reports), base-checklist, README
- [x] `01-spec/02-review/` — base-plan, base-checklist, README
- [x] `01-spec/03-fix/` — base-plan, base-checklist, README
- [x] `02-acceptance/01-write/` — base-plan v1.1.0 (Given/When/Then строгость, user-centric, измеримость, anti-patterns), base-checklist v1.1.0, README
- [x] `02-acceptance/02-review/` — base-plan, base-checklist, README
- [x] `02-acceptance/03-fix/` — base-plan, base-checklist, README
- [x] `02-acceptance/sources.md`
- [x] `03-tests/01-write/` — base-plan (FIRST, happy/adversarial папки, ≤150 строк), base-checklist, README
- [x] `03-tests/02-review/` — base-plan, base-checklist, README
- [x] `03-tests/03-fix/` — base-plan, base-checklist, README
- [x] `03-tests/04-run/` — base-plan (Red-only, тест без impl = красный флаг), base-checklist, README
- [x] `03-tests/sources.md`
- [x] `04-code/01-write/` — base-plan (TDD Green, ≤150 строк, SOLID, OWASP), base-checklist, README
- [x] `04-code/02-review/` — base-plan v1.1.0 (именование 9 правил, SLAP/Stepdown Rule, critical/warning/Nit:), base-checklist v1.1.0, README
- [x] `04-code/03-fix/` — base-plan, base-checklist, README
- [x] `04-code/04-testing/` — base-plan (Green run после write+review), base-checklist, README
- [x] `04-code/sources.md`
- [x] `05-docs/01-write/` — base-plan v1.1.0 (Diataxis 4 типа, ADR формат, docs-as-code), base-checklist, README
- [x] `05-docs/02-review/` — base-plan, base-checklist, README
- [x] `05-docs/03-fix/` — base-plan, base-checklist, README
- [x] `05-docs/sources.md`
- [x] `06-retro/01-recall/` — base-plan, base-checklist, README
- [x] `06-retro/02-collect/` — base-plan, base-checklist, README
- [x] `06-retro/03-analyze/` — base-plan, base-checklist, README
- [x] `06-retro/04-solve/` — base-plan, base-checklist, README
- [x] `06-retro/05-write/` — base-plan (→ читает 04-solve report), base-checklist, README

### Скиллы `.claude/commands/v2/`

- [x] `00-research/01-interview.md`
- [x] `00-research/02-web.md`
- [x] `01-spec/01-write.md`
- [x] `01-spec/02-review.md`
- [x] `01-spec/03-fix.md`
- [x] `02-acceptance/01-write.md`
- [x] `02-acceptance/02-review.md`
- [x] `02-acceptance/03-fix.md`
- [x] `03-tests/01-write.md`
- [x] `03-tests/02-review.md`
- [x] `03-tests/03-fix.md`
- [x] `03-tests/04-run.md`
- [x] `04-code/01-write.md`
- [x] `04-code/02-review.md`
- [x] `04-code/03-fix.md`
- [x] `04-code/04-testing.md`
- [x] `05-docs/01-write.md`
- [x] `05-docs/02-review.md`
- [x] `05-docs/03-fix.md`
- [x] `06-retro/01-recall.md`
- [x] `06-retro/02-collect.md`
- [x] `06-retro/03-analyze.md`
- [x] `06-retro/04-solve.md`
- [x] `06-retro/05-write.md`

### Стандарт процесса

- [x] `tasks/store/connection/RETRO/actions/00-process-standard.md`
- [x] `tasks/store/connection/RETRO/actions/00b-processes-catalog.md`

### Составные скиллы v2

- [x] `feature.md` v2 — полный флоу: research → spec → acceptance → tests → code → docs → retro; clarification-cascade дописан
- [x] `return-to-feature.md` v2 — чтение status-log.md + stale/clarification логика; git-ветка; сводка перед запуском

## Что осталось

### Ретро `store/connection` — action items (2026-04-27)

Три потока, выполняются параллельно. Ветка: `retro/store/connection`.

#### Поток А — Гримм (ревью-скиллы)
- [x] **02** spec-sync-gate → `docs/standards/v2/04-code/02-review/base-checklist.md` + `base-plan.md`
- [x] **05** тест-симулякры → `docs/standards/v2/03-tests/02-review/base-checklist.md` + `base-plan.md`
- [x] **07** code-review все файлы Go → `docs/standards/v2/04-code/02-review/base-checklist.md` + `docs/standards/go/principles/concurrency/rules/goroutine-lifecycle.md`

#### Поток Б — Танк (docs + acceptance/fix скиллы)
- [x] **03** slog стандарт → `docs/standards/architecture/tdd/principles/structured-logging/` (3 файла) + `docs/standards/v2/03-tests/03-fix/base-plan.md` + `base-checklist.md` + `docs/standards/v2/04-code/03-fix/base-plan.md` + `base-checklist.md` + `docs/standards/v2/03-tests/02-review/base-checklist.md`
- [x] **04** чек-лист сетевых протоколов → `docs/standards/v2/02-acceptance/01-write/base-checklist.md` + `docs/standards/v2/02-acceptance/02-review/base-checklist.md`
- [x] **06** fix-report обязателен → `docs/standards/templates/fix-report.md` (шаблон) + `docs/standards/v2/04-code/03-fix/base-plan.md` + `base-checklist.md` + `docs/standards/v2/03-tests/03-fix/base-plan.md` + `base-checklist.md`

#### Поток В — Ада (Go standards + code-write)
- [x] **01** `// inv:` стандарт → `docs/standards/go/principles/invariants/` (3 файла) + `docs/standards/v2/04-code/02-review/base-checklist.md`
- [x] **09** karpathy practices → `docs/standards/go/principles/implementation/` (3 файла) + `docs/standards/v2/04-code/01-write/base-plan.md` + `docs/standards/v2/04-code/03-fix/base-plan.md`

#### Отдельно — Харли
- [x] **08** RFC: Харли + Agent tool → написать `docs/rfc/harley-agent-tool.md`

#### Отдельно — TODO
- [x] Актуализировать `CLAUDE.md` начиная с секции `## Workflow: статусы задач и каскадный сброс` — переписана под v2 (status-log.md, 6 статусов, clarification алгоритм)
- [x] Отметки о датах — ISO 8601 UTC с суффиксом Z: `date -u +"%Y-%m-%dT%H:%M:%SZ"` → зафиксировано в CLAUDE.md и process-standard.md

## Ключевые мета-правила принятые в v2

1. **Структура папок тестов**: `tests/happy/` (Азирафаль) и `tests/adversarial/` (Кроули)
2. **Размер модуля**: ≤ 150 строк без комментариев — дроби по функциональным группам
3. **Иммутабельность report**: только `created`, без `updated`
4. **see-also вместо стакинга**: ссылки на актуальные предыдущие reports
5. **README при инициализации**: копируется из стандарта в task directory
6. **Правила движения**: в каждом процессе — done when / next / rollback
7. **Given/When/Then строгость**: Given=контекст, When=одно действие, Then=наблюдаемый результат
8. **00-research перед spec**: interview (Танк) + web (Бо) → reports → spec-write
9. **Red/Green split**: `03-tests/04-run` = Red (без impl все падают); `04-code/04-testing` = Green (все проходят)
10. **06-retro split**: recall → collect (team+выжимки) → analyze (группировка+уточнение) → solve (решения+закрытие) → write (RETRO.md)
