---
process: 02-acceptance/01-write
run: 1
date: 2026-04-29T10:51:59Z
created: 2026-04-29T10:51:59Z
spec-version: 1.4.0
spec-source: tasks/0002-store-crud/stages/01-spec/03-fix/report-004.md
status: done
agent: Танк
---

# Acceptance: store/crud

## Блок 1: FS-инициализация

### Сценарий AC-FS-INIT-01: Успешная инициализация FS-стора

**Given** конфиг содержит `basedir` — путь к директории которая не существует
**When** вызван конструктор `fs.New(cfg)`
**Then** возвращает `(*Store, nil)`; директории `entities/`, `relations/`, `jobs/` созданы на диске

---

### Сценарий AC-FS-INIT-02: Ошибка создания директорий при инициализации

**Given** конфиг содержит `basedir` — путь к директории которую нельзя создать (например, нет прав на родительский путь)
**When** вызван конструктор `fs.New(cfg)`
**Then** возвращает `(nil, err)` где `err != nil`; стор недоступен

---

## Блок 2: EntityStore — Upsert

### Сценарий AC-ENTITY-UPSERT-01: Создание новой entity

**Given** FS-стор инициализирован, запись с `id="ent-1"` не существует
**When** вызван `EntityStore.Upsert(ctx, Entity{ID:"ent-1", Type:"service"})`
**Then** возвращает `nil`; файл `entities/ent-1.json` существует на диске; `created_at` и `updated_at` установлены в UTC-время вызова; входящие значения `created_at`/`updated_at` в объекте игнорированы

---

### Сценарий AC-ENTITY-UPSERT-02: Обновление существующей entity

**Given** FS-стор инициализирован, запись с `id="ent-1"` существует с `created_at=T0`
**When** вызван `EntityStore.Upsert(ctx, Entity{ID:"ent-1", Type:"service", Name:"updated"})`
**Then** возвращает `nil`; файл `entities/ent-1.json` содержит обновлённый `Name`; `created_at` равен `T0` (сохранён из существующей записи); `updated_at` обновлён в UTC-время вызова

---

### Сценарий AC-ENTITY-UPSERT-03: Upsert с пустым id

**Given** FS-стор инициализирован
**When** вызван `EntityStore.Upsert(ctx, Entity{ID:"", Type:"service"})`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"MISSING_ID","error":"id is required"}`

---

### Сценарий AC-ENTITY-UPSERT-04: Ошибка чтения существующей записи при upsert

**Given** FS-стор инициализирован, файл `entities/ent-1.json` существует и недоступен для чтения (права 000)
**When** вызван `EntityStore.Upsert(ctx, Entity{ID:"ent-1", Type:"service"})`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"READ_ERROR","error":"failed to read existing record"}`

---

### Сценарий AC-ENTITY-UPSERT-05: Ошибка записи на диск при upsert

**Given** FS-стор инициализирован, директория `entities/` недоступна для записи
**When** вызван `EntityStore.Upsert(ctx, Entity{ID:"ent-new", Type:"service"})`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"WRITE_ERROR","error":"failed to write record"}`

---

### Сценарий AC-ENTITY-UPSERT-06: Temp-файл не остаётся при ошибке записи

**Given** FS-стор инициализирован, операция upsert завершается с ошибкой после создания temp-файла
**When** вызван `EntityStore.Upsert(ctx, Entity{ID:"ent-fail", Type:"service"})`
**Then** возвращает ошибку; в директории `entities/` нет файлов с префиксом `.tmp-`

---

## Блок 3: EntityStore — Get

### Сценарий AC-ENTITY-GET-01: Получение существующей entity

**Given** FS-стор инициализирован, файл `entities/ent-1.json` содержит валидную Entity с `ID="ent-1"`
**When** вызван `EntityStore.Get(ctx, "ent-1")`
**Then** возвращает `(Entity{ID:"ent-1",...}, nil)`; поля объекта соответствуют содержимому файла

---

### Сценарий AC-ENTITY-GET-02: Get несуществующей entity

**Given** FS-стор инициализирован, файл `entities/ent-999.json` не существует
**When** вызван `EntityStore.Get(ctx, "ent-999")`
**Then** возвращает `(Entity{}, err)` где `errors.Is(err, store.ErrNotFound) == true`

---

### Сценарий AC-ENTITY-GET-03: Get с пустым id

**Given** FS-стор инициализирован
**When** вызван `EntityStore.Get(ctx, "")`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"MISSING_ID","error":"id is required"}`

---

### Сценарий AC-ENTITY-GET-04: Ошибка чтения файла при get (не is-not-exist)

**Given** FS-стор инициализирован, файл `entities/ent-1.json` существует и недоступен для чтения (права 000)
**When** вызван `EntityStore.Get(ctx, "ent-1")`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"READ_ERROR","error":"failed to read record"}`

---

## Блок 4: EntityStore — Delete

### Сценарий AC-ENTITY-DELETE-01: Удаление существующей entity

**Given** FS-стор инициализирован, файл `entities/ent-1.json` существует
**When** вызван `EntityStore.Delete(ctx, "ent-1")`
**Then** возвращает `nil`; файл `entities/ent-1.json` не существует на диске

---

### Сценарий AC-ENTITY-DELETE-02: Delete несуществующей entity — не идемпотентный

**Given** FS-стор инициализирован, файл `entities/ent-999.json` не существует
**When** вызван `EntityStore.Delete(ctx, "ent-999")`
**Then** возвращает ошибку; `errors.Is(err, store.ErrNotFound) == true`; при передаче через роутер ответ `{"ok":false,"errorCode":"NOT_FOUND","error":"not found"}`

---

### Сценарий AC-ENTITY-DELETE-03: Delete с пустым id

**Given** FS-стор инициализирован
**When** вызван `EntityStore.Delete(ctx, "")`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"MISSING_ID","error":"id is required"}`

---

### Сценарий AC-ENTITY-DELETE-04: Ошибка удаления файла

**Given** FS-стор инициализирован, файл `entities/ent-1.json` существует, `os.Remove` принудительно возвращает ошибку (через mock FS или временное ограничение прав)
**When** вызван `EntityStore.Delete(ctx, "ent-1")`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"DELETE_ERROR","error":"failed to delete record"}`

---

## Блок 5: EntityStore — List

### Сценарий AC-ENTITY-LIST-01: List возвращает все entity

**Given** FS-стор инициализирован, директория `entities/` содержит файлы `a.json`, `b.json`, `c.json` с валидными Entity
**When** вызван `EntityStore.List(ctx)`
**Then** возвращает `([]Entity, nil)` содержащий ровно 3 элемента соответствующих содержимому файлов; порядок элементов произволен

---

### Сценарий AC-ENTITY-LIST-02: List пустой директории

**Given** FS-стор инициализирован, директория `entities/` пуста
**When** вызван `EntityStore.List(ctx)`
**Then** возвращает `([]Entity{}, nil)` — пустой срез, ошибки нет

---

### Сценарий AC-ENTITY-LIST-03: List при ошибке чтения директории

**Given** FS-стор инициализирован, директория `entities/` недоступна для чтения
**When** вызван `EntityStore.List(ctx)`
**Then** возвращает `(nil, err)` где `err != nil`; при передаче через роутер ответ `{"ok":false,"errorCode":"LIST_ERROR","error":"failed to list records"}`

---

### Сценарий AC-ENTITY-LIST-04: List прекращает обработку при ошибке decode одного файла

**Given** FS-стор инициализирован, директория `entities/` содержит `a.json` (валидный) и `b.json` (невалидный JSON)
**When** вызван `EntityStore.List(ctx)`
**Then** возвращает `(nil, err)` где `err != nil`; ни один элемент не возвращается; при передаче через роутер ответ содержит `{"ok":false,"errorCode":"LIST_ERROR","error":"failed to read record b"}`

---

### Сценарий AC-ENTITY-LIST-05: List игнорирует не-.json файлы

**Given** FS-стор инициализирован, директория `entities/` содержит `a.json` (валидный) и `.tmp-xyz` (temp-файл без расширения .json)
**When** вызван `EntityStore.List(ctx)`
**Then** возвращает `([]Entity, nil)` содержащий ровно 1 элемент; файл `.tmp-xyz` в результат не включён

---

## Блок 6: RelationStore

### Сценарий AC-RELATION-UPSERT-01: Создание новой relation

**Given** FS-стор инициализирован, запись с `id="rel-1"` не существует
**When** вызван `RelationStore.Upsert(ctx, Relation{ID:"rel-1", FromID:"ent-a", ToID:"ent-b", Type:"uses"})`
**Then** возвращает `nil`; файл `relations/rel-1.json` существует на диске; `created_at` и `updated_at` установлены в UTC-время вызова

---

### Сценарий AC-RELATION-UPSERT-02: FS-реализация не проверяет существование from_id/to_id

**Given** FS-стор инициализирован, записи с `id="ent-src"` и `id="ent-dst"` в entities не существуют
**When** вызван `RelationStore.Upsert(ctx, Relation{ID:"rel-1", FromID:"ent-src", ToID:"ent-dst", Type:"uses"})`
**Then** возвращает `nil`; файл `relations/rel-1.json` создан — FS не проверяет FK-целостность

---

### Сценарий AC-RELATION-UPSERT-03: Upsert с пустым id

**Given** FS-стор инициализирован
**When** вызван `RelationStore.Upsert(ctx, Relation{ID:"", FromID:"a", ToID:"b", Type:"uses"})`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"MISSING_ID","error":"id is required"}`

---

### Сценарий AC-RELATION-GET-01: Get существующей relation

**Given** FS-стор инициализирован, файл `relations/rel-1.json` содержит валидную Relation с `ID="rel-1"`
**When** вызван `RelationStore.Get(ctx, "rel-1")`
**Then** возвращает `(Relation{ID:"rel-1",...}, nil)`; поля соответствуют содержимому файла

---

### Сценарий AC-RELATION-GET-02: Get несуществующей relation

**Given** FS-стор инициализирован, файл `relations/rel-999.json` не существует
**When** вызван `RelationStore.Get(ctx, "rel-999")`
**Then** возвращает `(Relation{}, err)` где `errors.Is(err, store.ErrNotFound) == true`

---

### Сценарий AC-RELATION-DELETE-01: Delete существующей relation

**Given** FS-стор инициализирован, файл `relations/rel-1.json` существует
**When** вызван `RelationStore.Delete(ctx, "rel-1")`
**Then** возвращает `nil`; файл `relations/rel-1.json` не существует на диске

---

### Сценарий AC-RELATION-DELETE-02: Delete несуществующей relation — не идемпотентный

**Given** FS-стор инициализирован, файл `relations/rel-999.json` не существует
**When** вызван `RelationStore.Delete(ctx, "rel-999")`
**Then** возвращает ошибку; `errors.Is(err, store.ErrNotFound) == true`

---

### Сценарий AC-RELATION-LIST-01: List возвращает все relations

**Given** FS-стор инициализирован, директория `relations/` содержит 2 файла с валидными Relation
**When** вызван `RelationStore.List(ctx)`
**Then** возвращает `([]Relation, nil)` содержащий ровно 2 элемента; порядок произволен

---

### Сценарий AC-RELATION-LIST-02: List пустой директории

**Given** FS-стор инициализирован, директория `relations/` пуста
**When** вызван `RelationStore.List(ctx)`
**Then** возвращает `([]Relation{}, nil)` — пустой срез, ошибки нет

---

## Блок 7: JobStore

### Сценарий AC-JOB-UPSERT-01: Создание нового job

**Given** FS-стор инициализирован, запись с `id="job-1"` не существует
**When** вызван `JobStore.Upsert(ctx, Job{ID:"job-1", Kind:"index", Worker:"w1", Status:"pending"})`
**Then** возвращает `nil`; файл `jobs/job-1.json` существует на диске; `created_at` и `updated_at` установлены в UTC-время вызова

---

### Сценарий AC-JOB-UPSERT-02: Upsert с пустым id

**Given** FS-стор инициализирован
**When** вызван `JobStore.Upsert(ctx, Job{ID:"", Kind:"index", Worker:"w1", Status:"pending"})`
**Then** возвращает ошибку; при передаче через роутер ответ `{"ok":false,"errorCode":"MISSING_ID","error":"id is required"}`

---

### Сценарий AC-JOB-GET-01: Get существующего job

**Given** FS-стор инициализирован, файл `jobs/job-1.json` содержит валидный Job с `ID="job-1"`
**When** вызван `JobStore.Get(ctx, "job-1")`
**Then** возвращает `(Job{ID:"job-1",...}, nil)`; поля соответствуют содержимому файла

---

### Сценарий AC-JOB-GET-02: Get несуществующего job

**Given** FS-стор инициализирован, файл `jobs/job-999.json` не существует
**When** вызван `JobStore.Get(ctx, "job-999")`
**Then** возвращает `(Job{}, err)` где `errors.Is(err, store.ErrNotFound) == true`

---

### Сценарий AC-JOB-DELETE-01: Delete существующего job

**Given** FS-стор инициализирован, файл `jobs/job-1.json` существует
**When** вызван `JobStore.Delete(ctx, "job-1")`
**Then** возвращает `nil`; файл `jobs/job-1.json` не существует на диске

---

### Сценарий AC-JOB-DELETE-02: Delete несуществующего job — не идемпотентный

**Given** FS-стор инициализирован, файл `jobs/job-999.json` не существует
**When** вызван `JobStore.Delete(ctx, "job-999")`
**Then** возвращает ошибку; `errors.Is(err, store.ErrNotFound) == true`

---

### Сценарий AC-JOB-LIST-01: List возвращает все jobs

**Given** FS-стор инициализирован, директория `jobs/` содержит 3 файла с валидными Job
**When** вызван `JobStore.List(ctx)`
**Then** возвращает `([]Job, nil)` содержащий ровно 3 элемента; порядок произволен

---

### Сценарий AC-JOB-LIST-02: List пустой директории

**Given** FS-стор инициализирован, директория `jobs/` пуста
**When** вызван `JobStore.List(ctx)`
**Then** возвращает `([]Job{}, nil)` — пустой срез, ошибки нет

---

## Блок 8: JSONL-роутер

### Сценарий AC-ROUTER-01: Upsert через роутер — успех

**Given** роутер инициализирован с валидными сторами, соединение установлено
**When** роутер получает строку `{"op":"upsert","type":"entity","data":{"id":"ent-1","type":"service"}}`
**Then** ответ `{"ok":true,"id":"ent-1"}` записан в соединение; соединение не закрыто

---

### Сценарий AC-ROUTER-02: Get через роутер — успех

**Given** роутер инициализирован, запись entity с `id="ent-1"` существует в сторе
**When** роутер получает строку `{"op":"get","type":"entity","id":"ent-1"}`
**Then** ответ `{"ok":true,"data":{...}}` содержит объект с `"id":"ent-1"`; соединение не закрыто

---

### Сценарий AC-ROUTER-03: Delete через роутер — успех

**Given** роутер инициализирован, запись entity с `id="ent-1"` существует в сторе
**When** роутер получает строку `{"op":"delete","type":"entity","id":"ent-1"}`
**Then** ответ `{"ok":true}`; соединение не закрыто

---

### Сценарий AC-ROUTER-04: List через роутер — успех

**Given** роутер инициализирован, entity-стор содержит 2 записи
**When** роутер получает строку `{"op":"list","type":"entity"}`
**Then** ответ `{"ok":true,"data":[..., ...]}` содержит 2 объекта; соединение не закрыто

---

### Сценарий AC-ROUTER-05: Upsert без поля data — соединение сохраняется

**Given** роутер инициализирован, соединение установлено
**When** роутер получает строку `{"op":"upsert","type":"entity"}`
**Then** ответ `{"ok":false,"errorCode":"INVALID_REQUEST","error":"data is required for upsert"}`; соединение не закрыто; следующий запрос обрабатывается нормально

---

### Сценарий AC-ROUTER-06: Upsert с data=null — соединение сохраняется

**Given** роутер инициализирован, соединение установлено
**When** роутер получает строку `{"op":"upsert","type":"entity","data":null}`
**Then** ответ `{"ok":false,"errorCode":"INVALID_REQUEST","error":"data is required for upsert"}`; соединение не закрыто

---

### Сценарий AC-ROUTER-07: Неизвестный op — соединение сохраняется

**Given** роутер инициализирован, соединение установлено
**When** роутер получает строку `{"op":"patch","type":"entity","id":"x"}`
**Then** ответ `{"ok":false,"errorCode":"UNKNOWN_OP","error":"unknown op: patch"}`; соединение не закрыто; следующий запрос обрабатывается нормально

---

### Сценарий AC-ROUTER-08: Неизвестный type — соединение сохраняется

**Given** роутер инициализирован, соединение установлено
**When** роутер получает строку `{"op":"get","type":"table","id":"x"}`
**Then** ответ `{"ok":false,"errorCode":"UNKNOWN_TYPE","error":"unknown type: table"}`; соединение не закрыто; следующий запрос обрабатывается нормально

---

### Сценарий AC-ROUTER-09: Невалидный JSON — поток испорчён, соединение закрывается

**Given** роутер инициализирован, соединение установлено
**When** роутер получает строку `{not-valid-json`
**Then** роутер закрывает соединение; дальнейшие запросы в этом соединении не обрабатываются

---

### Сценарий AC-ROUTER-10: Невалидный JSON в data при upsert — соединение сохраняется

**Given** роутер инициализирован, соединение установлено
**When** роутер получает строку `{"op":"upsert","type":"entity","data":"{broken"}`
**Then** ответ `{"ok":false,"errorCode":"INVALID_REQUEST","error":"invalid data"}`; соединение не закрыто

---

### Сценарий AC-ROUTER-11: id из envelope игнорируется при upsert, используется id из data

**Given** роутер инициализирован, соединение установлено, записей не существует
**When** роутер получает строку `{"op":"upsert","type":"entity","id":"wrong-id","data":{"id":"correct-id","type":"svc"}}`
**Then** ответ `{"ok":true,"id":"correct-id"}`; запись сохранена с id `"correct-id"`; запись с id `"wrong-id"` не создана

---

### Сценарий AC-ROUTER-12: Несколько запросов в одном соединении

**Given** роутер инициализирован, соединение установлено, запись entity `"ent-1"` существует
**When** роутер последовательно получает 3 JSONL-строки: upsert entity `"ent-2"`, get entity `"ent-1"`, list entity
**Then** 3 ответа записаны в соединение в том же порядке; соединение не закрыто после первого или второго ответа

---

### Сценарий AC-ROUTER-13: Корректное завершение при закрытии соединения клиентом

**Given** роутер инициализирован, соединение установлено, несколько запросов успешно обработаны
**When** клиент закрывает соединение (Decode возвращает io.EOF или io.ErrUnexpectedEOF)
**Then** роутер выходит из цикла без ошибки; никакого дополнительного ответа не записывается

---

### Сценарий AC-ROUTER-14: Логирование успешной операции — уровень DEBUG

**Given** роутер инициализирован с логгером, запись entity с `id="ent-1"` существует
**When** роутер успешно обрабатывает `{"op":"get","type":"entity","id":"ent-1"}`
**Then** в лог записана строка уровня DEBUG содержащая поля `op=get`, `type=entity`, `id=ent-1`

---

### Сценарий AC-ROUTER-15: Логирование list — поле count присутствует, поле id отсутствует

**Given** роутер инициализирован с логгером, entity-стор содержит 3 записи
**When** роутер обрабатывает `{"op":"list","type":"entity"}`
**Then** в лог записана строка уровня DEBUG содержащая поля `op=list`, `type=entity`, `count=3`; поле `id` в лог-записи отсутствует

---

### Сценарий AC-ROUTER-16: Логирование ошибки операции — уровень ERROR

**Given** роутер инициализирован с логгером, entity с `id="ent-999"` не существует
**When** роутер обрабатывает `{"op":"get","type":"entity","id":"ent-999"}`
**Then** в лог записана строка уровня ERROR содержащая поля `op=get`, `type=entity`, `id=ent-999`, `error="not found"`

---

## Блок 9: НФТ — атомарность upsert

### Сценарий AC-NFT-ATOMIC-01: Целевой файл не повреждён при прерывании upsert

**Given** FS-стор инициализирован, файл `entities/ent-1.json` содержит версию V1 (валидный JSON)
**When** процесс прерывается (SIGKILL) в момент записи нового содержимого при upsert `ent-1`
**Then** файл `entities/ent-1.json` либо содержит версию V1 (неизменный), либо не существует; файл не содержит частичной или повреждённой записи

---

### Сценарий AC-NFT-ATOMIC-02: UUID v7 используется клиентом для генерации id

**Given** клиент генерирует id перед вызовом Upsert
**When** вызван `uuid.NewV7()` и результат передан как `id` в объект
**Then** функция возвращает непустую строку в формате UUID; `uuid.New()` (UUID v4) в кодовой базе не использован

---

## Блок 10: Конструктор роутера

### Сценарий AC-ROUTER-INIT-01: Роутер принимает интерфейсы, не конкретные типы

**Given** созданы mock-реализации `store.EntityStore`, `store.RelationStore`, `store.JobStore`
**When** вызван `router.New(entityMock, relationMock, jobMock)`
**Then** возвращает `*Router` без ошибки; код компилируется — параметры `router.New` типизированы интерфейсами, не `*fs.Store`

---

## Checklist

- [x] Spec прочитана (v1.4.0, report-004.md)
- [x] Каждый happy path из ФТ имеет хотя бы один сценарий
- [x] Каждый failure mode из ФТ покрыт сценарием
- [x] Граничные условия покрыты (пустой id, пустой список, частичная ошибка в list)
- [x] НФТ которые можно проверить тестом покрыты (атомарность, UUID v7, интерфейсная зависимость)
- [x] Given — только контекст, не содержит действий
- [x] When — ровно одно действие; нет "And" в When
- [x] Then — наблюдаемый результат с точки зрения вызывающей стороны
- [x] Результаты измеримы (конкретные errorCode, точные строки ошибок)
- [x] Сценарии независимы
- [x] Новых failure modes которых нет в spec — не обнаружено

## Решения принятые в процессе

- AC-ENTITY-UPSERT-01: входящие `created_at`/`updated_at` игнорируются — стор выставляет сам
- AC-ENTITY-LIST-04: при ошибке decode одного файла — nil, не частичный результат
- AC-ROUTER-09 vs AC-ROUTER-10: parse error (json.Decoder) закрывает соединение; invalid data в envelope — application-level, соединение сохраняется
- AC-RELATION-UPSERT-02: FS намеренно не проверяет FK — это explicit ограничение из §7

## Открытые вопросы

Нет.
