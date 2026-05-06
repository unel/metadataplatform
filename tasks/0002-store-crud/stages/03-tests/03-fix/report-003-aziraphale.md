---
purpose: Фикс happy-тестов для store/crud (clarification от Ады)
process: 03-tests/03-fix
run: 3
date: 2026-05-04T06:28:38Z
created: 2026-05-04T06:28:38Z
status: done
agent: Азирафаль
---

## Исправленные файлы

- `store/tests/happy/relation_test.go` — все вызовы `s.Get/Upsert/Delete/List` заменены на `s.Relations().Get/Upsert/Delete/List`
- `store/tests/happy/job_test.go` — все вызовы `s.Get/Upsert/Delete/List` заменены на `s.Jobs().Get/Upsert/Delete/List`
- `store/tests/happy/router_happy_test.go` — `router.New(s, s, s, nil)` → `router.New(s.Entities(), s.Relations(), s.Jobs(), nil)`; compile-time assertions переписаны (только EntityStore для `*fs.Store`); `s.Upsert(ctx, entity)` в Arrange → `s.Entities().Upsert(...)`
- `store/tests/happy/router_conn_test.go` — `router.New(s, s, s, nil)` → `router.New(s.Entities(), s.Relations(), s.Jobs(), nil)`
- `store/tests/happy/helpers_logger_test.go` — `router.New(s, s, s, logger)` → `router.New(s.Entities(), s.Relations(), s.Jobs(), logger)`
- `store/tests/happy/nft_test.go` — два дополнительных фикса: `t.Context()` → `context.Background()` (Go 1.22 не поддерживает `t.Context`); `assert.Equal(t, 7, ...)` → `assert.Equal(t, uuid.Version(7), ...)` (строгое сравнение типов)

## Что было не так

`*fs.Store` имеет три метода-обёртки: `Entities() store.EntityStore`, `Relations() store.RelationStore`, `Jobs() store.JobStore`. Сами методы `Get/Upsert/Delete/List` на `*fs.Store` реализованы только для Entity — они есть напрямую на типе. Для Relation и Job — только через приватные типы `relationStore` и `jobStore`, возвращаемые соответственными методами.

Следовательно: `var _ store.RelationStore = (*fs.Store)(nil)` — не компилируется (у `*fs.Store` нет метода `Get(ctx, string) (Relation, error)`). Аналогично для JobStore. А вызовы `s.Upsert(ctx, rel)` и `s.Get(ctx, rel.ID)` в тестах relation/job — тоже не компилируются.

Дополнительно обнаружены два независимых дефекта в `nft_test.go`:
- `t.Context()` недоступен в Go 1.22 (установленный компилятор 1.22.2)
- `assert.Equal(t, 7, id.Version())` — testify сравнивает типы строго; `int(7)` != `uuid.Version(7)` (uint8)

## Что сделано

1. `relation_test.go` — в каждом тесте добавлена переменная `rs := s.Relations()` и все операции переведены на `rs`.
2. `job_test.go` — аналогично, `js := s.Jobs()`.
3. `router_happy_test.go` — все три места `router.New(s, s, s, ...)` заменены на `router.New(s.Entities(), s.Relations(), s.Jobs(), ...)`. Compile-time assertion оставлен только для `store.EntityStore` (`var _ store.EntityStore = (*fs.Store)(nil)` — работает). Ассерции для RelationStore и JobStore заменены комментарием — эти гарантии теперь проверяются в `TestRouter_New_AcceptsStoreInterfaces` через присваивание интерфейсным переменным.
4. `router_conn_test.go` — `router.New(s, s, s, nil)` → `router.New(s.Entities(), s.Relations(), s.Jobs(), nil)`.
5. `helpers_logger_test.go` — аналогично.
6. `nft_test.go` — `t.Context()` → `context.Background()` + добавлен импорт `context`; тип исправлен на `uuid.Version(7)`.

Результат: `go test ./store/tests/happy/...` — ok (36 тестов, все pass). Adversarial тесты не затронуты и тоже проходят.

## Clarification

Нет. Все изменения укладываются в рамки существующих acceptance-сценариев — ни один AC не затронут. Изменены только вызовы API в тестах, не поведение которое проверяется.
