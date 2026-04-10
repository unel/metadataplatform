# Full-text search в PostgreSQL

PostgreSQL реализует FTS через `tsvector` (индексированный вектор документа) и `tsquery` (запрос с операторами). GIN-индекс по `tsvector` обеспечивает быстрый поиск без seq scan.

В проекте: `relations.search_tsv` — агрегат полей `type`, `subtype`, `value`, `meta`.
