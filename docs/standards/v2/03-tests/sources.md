---
purpose: Источники по TDD и написанию качественных тестов — использованы при написании стандартов 03-tests
created: 2026-04-26T15:00
---

# Источники: TDD и качество тестов

## FIRST принципы

- [F.I.R.S.T principles of testing — Medium (Tasdik Rahman)](https://medium.com/@tasdikrahman/f-i-r-s-t-principles-of-testing-1a497acda8d6)
- [FIRST Principles as Solid Rules for Tests — DZone](https://dzone.com/articles/first-principles-solid-rules-for-tests)
- [Unit Tests Are FIRST — The Pragmatic Programmers](https://medium.com/pragmatic-programmers/unit-tests-are-first-fast-isolated-repeatable-self-verifying-and-timely-a83e8070698e)
- [F.I.R.S.T. principles of testing — aalonso.dev](https://aalonso.dev/blog/2024/f-i-r-s-t-principles-of-testing/)

## TDD практики

- [Essential practices for writing better tests in TDD — WWT](https://www.wwt.com/blog/essential-practices-for-writing-better-tests-in-tdd)
- [5 steps of test-driven development — IBM Developer](https://developer.ibm.com/articles/5-steps-of-test-driven-development/)
- [Test-driven development: principles, tools & pitfalls — Statsig](https://www.statsig.com/perspectives/tdd-principles-tools-pitfalls)

## Ключевые выводы

- **FIRST**: Fast, Isolated, Repeatable, Self-validating, Timely
- **Behavior not implementation**: тест проверяет наблюдаемое поведение, не внутренние детали — стабилен при рефакторинге
- **One test = one thing**: один тест проверяет один сценарий из acceptance
- **Own setup**: каждый тест сам готовит данные, не зависит от других тестов
- **Deterministic**: никаких random, time.Now() без инъекции, сетевых вызовов без контроля
- **Descriptive name**: имя теста = описание сценария + ожидаемый результат
- **Red-Green-Refactor**: тест пишется до кода, сначала падает, потом проходит
