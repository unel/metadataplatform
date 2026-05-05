# Action: go build ./... отдельным шагом перед go test

## Проблема

"setup failed" одинаково для двух разных ситуаций: (1) Red phase — код не написан, тест ожидаемо не компилируется; (2) реальная ошибка компиляции из-за несовместимости API. Нельзя отличить ожидаемое от неожиданного.

## Решение

В `03-tests/04-run` и `04-code/04-testing` добавить отдельный шаг:
```bash
go build ./...   # compile check — отдельная запись в status-log
go test ./...    # test run
```

Статусы в status-log:
- `compile-ok` / `compile-fail` — отдельно
- `test-pass` / `test-fail` — отдельно

## Шаги

- [ ] Обновить `docs/standards/v2/03-tests/04-run/base-plan.md`
- [ ] Обновить `docs/standards/v2/04-code/04-testing/base-plan.md`

## Источники

Кроули (ретро store/crud, 2026-05-05)
Рецидив из: RETRO store/connection
