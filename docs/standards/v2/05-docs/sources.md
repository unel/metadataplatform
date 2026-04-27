---
purpose: Источники по написанию технической документации — использованы при написании стандарта 05-docs
created: 2026-04-26T16:45
---

# Источники: техническая документация

## Diataxis / Divio — система четырёх типов

- [Diátaxis documentation framework](https://diataxis.fr/) — официальный сайт, автор Daniele Procida (Canonical)
- [Divio Documentation System](https://docs.divio.com/documentation-system/) — старое имя, те же материалы
- [Diataxis adoption by Canonical/Ubuntu](https://ubuntu.com/blog/diataxis-a-new-foundation-for-canonical-documentation) — принят как стандарт Ubuntu, Python Software Foundation

## Write the Docs

- [Documentation Principles — Write the Docs](https://www.writethedocs.org/guide/writing/docs-principles/)
- [Docs as Code — Write the Docs](https://www.writethedocs.org/guide/docs-as-code/)
- [Software Documentation Guide — Write the Docs](https://www.writethedocs.org/guide/index.html)

## Google

- [Google Developer Documentation Style Guide](https://developers.google.com/style)
- [Active Voice — Google Style Guide](https://developers.google.com/style/voice)

## README best practices

- [Make a README](https://www.makeareadme.com/)
- [README Best Practices — jehna/GitHub](https://github.com/jehna/readme-best-practices)
- [How to Write a Good README — freeCodeCamp](https://www.freecodecamp.org/news/how-to-write-a-good-readme-file/)

## ADR — Architecture Decision Records

- [Documenting Architecture Decisions — Michael Nygard, Cognitect (2011)](https://www.cognitect.com/blog/2011/11/15/documenting-architecture-decisions)
- [ADR Templates — adr.github.io](https://adr.github.io/adr-templates/)
- [ADR examples — joelparkerhenderson/GitHub](https://github.com/joelparkerhenderson/architecture-decision-record)
- [ADR Best Practices — AWS Architecture Blog](https://aws.amazon.com/blogs/architecture/master-architecture-decision-records-adrs-best-practices-for-effective-decision-making/)

## Ключевые выводы

### Diataxis — четыре типа документации (Procida)

Система из двух осей: **action vs cognition** × **learning vs working**.

| Тип | Когда писать | Что содержит |
|---|---|---|
| **How-to guide** | Есть задача которую нужно решить | Шаги к конкретному результату; предполагает базовые знания |
| **Reference** | Нужна точная нейтральная информация | API, параметры, коды ошибок — без советов и мнений |
| **Explanation** | Нужно объяснить "почему так" | Трейдоффы, архитектурные решения, контекст; единственный тип где уместно мнение |
| **Tutorial** | Нужно ввести новичка | Конкретные действия к конкретному результату; не объясняет — делает |

**Критически важно**: держать типы раздельными. Tutorial не объясняет почему. Reference не даёт советы.

### Write the Docs

- **Incorrect docs хуже чем no docs** — устаревшая документация активно вредит
- **Docs as Code** — документация в репозитории рядом с кодом, проходит review
- **DRY не работает в docs** — частичное дублирование бизнес-логики норма
- **Progressive disclosure** — обзор сначала, детали по ссылке

### README

- Первые три предложения: что это, зачем, как работает
- Usage examples с кодом — обязательно, раньше чем installation
- Антипаттерн: build/dev instructions на первом месте — отпугивают

### ADR (Nygard, 2011)

- **Append-only**: не редактируют, при изменении — новый ADR, старый помечают "superseded by #N"
- **Context — самое важное поле**: через год никто не вспомнит почему так решили
- **One decision — one record**
- Живёт в `docs/adr/` или `doc/decisions/` — в репозитории, не в Confluence
