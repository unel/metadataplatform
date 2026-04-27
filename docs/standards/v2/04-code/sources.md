---
purpose: Источники по code review, именованию и принципам чистого кода — использованы при написании стандарта 04-code
created: 2026-04-26T16:30
---

# Источники: code review и чистый код

## Code Review best practices

- [What to look for in a code review — Google Engineering Practices](https://google.github.io/eng-practices/review/reviewer/looking-for.html)
- [The Standard of Code Review — Google Engineering Practices](https://google.github.io/eng-practices/review/reviewer/standard.html)
- [Code Review Best Practices — Palantir Blog](https://blog.palantir.com/code-review-best-practices-19e02780015f)
- [Best Practices for Peer Code Review — SmartBear (Cisco data)](https://smartbear.com/learn/code-review/best-practices-for-peer-code-review/)

## Именование

- [Summary of Clean Code — Robert C. Martin (глава 2: Meaningful Names)](https://gist.github.com/wojteklu/73c6914cc446146b8b533c0988cf8d29)
- [Writing Clean Code — Naming Variables, Functions, Methods, Classes — Medium](https://medium.com/@mikhailhusyev/writing-clean-code-naming-variables-functions-methods-and-classes-6074a6796c7b)

## Single Level of Abstraction (SLAP)

- [Single Level of Abstraction (SLA) — Principles Wiki](http://principles-wiki.net/principles:single_level_of_abstraction)
- [SLAP: Single Level of Abstraction Principle — DEV Community](https://dev.to/le0nidas/slap-single-level-of-abstraction-principle-4o98)
- [Your Methods Should be "Single Level of Abstraction" Long — TechYourChance](https://www.techyourchance.com/single-level-of-abstraction-principle/)
- [Stepdown Rule — stevengong.co](https://stevengong.co/notes/Stepdown-Rule)

## Ключевые выводы

### Именование (Martin, Clean Code, гл. 2)

- **Intention-Revealing**: имя говорит *зачем* существует, *что делает*, *как используется*. Если имя требует комментария — имя неправильное
- **Avoid Misinformation**: не используй устоявшийся термин в нестандартном смысле (`accountList` — только если это реально List)
- **Meaningful Distinctions**: различие в имени должно раскрывать различие в смысле; `ProductInfo` vs `ProductData` — не различие
- **Pronounceable**: имена должны произноситься вслух — если нельзя произнести, нельзя обсудить
- **Searchable**: избегай магических чисел и однобуквенных переменных — они не ищутся grep'ом
- **One Word Per Concept**: выбери одно слово для одного понятия и держись его. Либо `fetch`, либо `get`, но не оба вперемешку
- **Глагол для функций**, существительное для классов/переменных; булевые — `is`/`has`/`can` префикс
- **Без encoding**: венгерская нотация и `IInterface` префиксы — мусор

### Single Level of Abstraction / SLAP (Martin, Clean Code, стр. 36; Ford, The Productive Programmer)

> "Each method should be written in terms of a single level of abstraction."

Метод не смешивает *что* делается (бизнес-оркестрация) с *как* это делается (SQL, HTTP, форматирование). Нарушение видно по: комментариям `// Step 1 / Step 2`, смеси высокоуровневых вызовов с низкоуровневыми деталями.

**Stepdown Rule** (Clean Code, гл. 3): функции в файле располагаются сверху вниз в порядке убывания уровня абстракции — читая файл, спускаешься на один уровень за раз.

### Code Review (Google, Palantir, SmartBear)

- **Оптимальный объём за сессию**: 200–400 LOC; быстрее 500 LOC/час — дефекты пропускаются (SmartBear/Cisco)
- **Nit: префикс**: non-blocking замечания помечаются `Nit:` — отделяет blocker от style suggestion (Google)
- **Complexity**: "cannot be understood quickly" — достаточный повод запросить упрощение (Google)
- **Over-engineering**: решение несуществующих будущих проблем — отдельная категория замечаний (Google)
- **TODO без тикета**: не принимается; обязателен номер issue (Palantir)
- **Комментарии объясняют WHY**, не WHAT: хорошо написанный код самодокументирован (Google)
