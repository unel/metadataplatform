# Action: spec.md как единый источник истины на уровне группы

## Проблема

Две копии спеки: полный текст в `01-spec/01-write/report-001.md` и обновлённый в `01-spec/03-fix/report-NNN.md`. Агенты читают разные версии, неизвестно какая актуальна.

## Решение

`01-spec/spec.md` — единственный файл спеки:
- `01-write` создаёт его
- `03-fix` обновляет на месте
- `02-review` и все downstream читают именно его

Reports (`report-NNN.md`) остаются как лог прогонов и дельта-изменений, не как копии спеки.

## Шаги

- [ ] Обновить `docs/standards/v2/01-spec/01-write/base-plan.md` — создавать `spec.md` вместо полного текста в report
- [ ] Обновить `docs/standards/v2/01-spec/03-fix/base-plan.md` — обновлять `spec.md` на месте
- [ ] Обновить `docs/standards/v2/01-spec/02-review/base-plan.md` — читать `spec.md`

## Источники

Пользователь (ретро store/crud, 2026-05-05)
