---
purpose: Стандарт процесса как первоклассной единицы — перенесён в документацию
created: 2026-04-26
updated: 2026-04-27
---

# Action: стандарт процесса

Стандарт перенесён в: `docs/standards/v2/process-standard.md`

## Источник

Пользователь (ретро store/connection, 2026-04-26):

> "Каждый процесс должен прогоняться по чек-листу и сам иметь чек-листы, артефакты и отдельный статус-лог"

## Что решает

Большинство action points этого ретро — симптомы одной болезни: процессы не являются явными единицами.

Без этой структуры:
- fix-report не пишется — нет артефактной структуры процесса `code-fix`
- main.go выпадает из ревью — нет чек-листа процесса `code-review`
- тест-симулякр проходит — нет выходного чек-листа `test-review`
- история итераций теряется — нет статус-лога каждого прогона
- агент повторяет одни и те же ошибки — notes/complaints не накапливаются

## Связанные action points

- `01-inv-comments.md` — чек-лист и артефакт для code-review
- `02-spec-sync-gate.md` — выходной чек-лист для code-review
- `03-slog-tdd-standard.md` — выходной чек-лист для test-fix
- `04-network-acceptance-checklist.md` — входной чек-лист для acceptance-write
- `05-test-simulacrum-checklist.md` — выходной чек-лист для test-review
- `06-fix-report-mandatory.md` — обязательный артефакт для всех fix-процессов
- `07-code-review-all-files.md` — чек-лист для code-review
- `08-harley-agent-tool-rfc.md` — архитектурное ограничение на процессы оркестрации
- `09-adopt-karpathy-claude-md.md` — `verify:` поля = явные выходные чек-листы шагов
