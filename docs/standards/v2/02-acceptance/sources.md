---
purpose: Источники по написанию acceptance criteria — использованы при написании стандартов 02-acceptance
created: 2026-04-26T15:00
---

# Источники: acceptance criteria

## Лучшие практики и антипаттерны

- [Writing Acceptance Criteria Like a Pro — Business Analysis Experts](https://www.businessanalysisexperts.com/writing-acceptance-criteria-agile-user-story/)
- [How to Write Effective Gherkin Acceptance Criteria — TestQuality](https://testquality.com/how-to-write-effective-gherkin-acceptance-criteria/)
- [Acceptance Criteria: Purposes, Types, Examples and Best Practices — AltexSoft](https://www.altexsoft.com/blog/acceptance-criteria-purposes-formats-and-best-practices/)
- [7 Ways to Improve Your Acceptance Criteria for Fewer Bugs — Medium](https://medium.com/@mattcalder0901/7-ways-to-improve-your-acceptance-criteria-for-fewer-bugs-91fbd5a68985)
- [Given-When-Then Acceptance Criteria for Better User Stories — ParallelHQ](https://www.parallelhq.com/blog/given-when-then-acceptance-criteria)

## Антипаттерны

- [Common pitfalls when defining acceptance criteria — LinkedIn](https://www.linkedin.com/advice/3/what-some-common-pitfalls-avoid-when-3c)
- [Writing Effective Acceptance Criteria — Lane8](https://www.lane8.com.au/post/writing-effective-acceptance-criteria)
- [What is Acceptance Criteria? — Atlassian](https://www.atlassian.com/work-management/project-management/acceptance-criteria)

## Ключевые выводы

- **User-centric**: описывает что пользователь наблюдает, не как система это делает внутри
- **Measurable**: "возвращает за 200мс", не "быстро"; "показывает ошибку X", не "удобно"
- **Atomic**: один сценарий = одна проверяемая вещь; без And-цепочек в When
- **Independent**: сценарии не зависят от порядка и результатов друг друга
- **Given**: только контекст (предусловие), не действие
- **When**: одно действие, не несколько
- **Then**: наблюдаемый результат с точки зрения пользователя/вызывающей стороны
