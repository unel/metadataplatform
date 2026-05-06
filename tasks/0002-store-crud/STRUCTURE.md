# STRUCTURE — 0002-store-crud

Задача выполнена по workflow v2. Полная структура этапов в `stages/`.

## Артефакты

```
artifacts/
├── retro.md          — итог ретро
└── actions/          — action items из ретро (→ BACKLOG.md)
```

> `artifacts/spec.md` и `artifacts/acceptance.md` не созданы: в старом workflow
> единственным источником истины были report'ы шагов. Нужно извлечь финальные
> версии из `stages/01-spec/03-fix/` и `stages/02-acceptance/03-fix/`.

## Этапы

```
stages/
├── 00-research/      01-interview, 02-web
├── 01-spec/          01-write, 02-review, 03-fix
├── 02-acceptance/    01-write, 02-review, 03-fix
├── 03-tests/         01-write, 02-review, 03-fix, 04-run
├── 04-code/          01-write, 02-review, 03-fix, 04-testing
├── 05-docs/          01-write, 02-review, 03-fix
└── 06-retro/         01-recall, 02-collect, 03-analyze, 04-solve, 05-write
```

## Тесты

Тесты живут в компоненте `store/`, разделены по принципу ответственности:

```
store/fs/*_happy_test.go / *_adversarial_test.go     — fs CRUD тесты (package fs_test)
store/integration/*_happy_test.go / *_adversarial_test.go — роутер + fs (package integration_test)
```
