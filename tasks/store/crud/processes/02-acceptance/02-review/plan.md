---
generated: 2026-04-29T12:50:20Z
created: 2026-04-29T12:50:20Z
updated: 2026-04-29T12:50:20Z
base-plan: 02-acceptance-review v1.0.0
---

# План ревью acceptance: store/crud

## 1. Сбор контекста

- Spec: `01-spec/03-fix/report-004.md` (v1.4.0)
- Acceptance: `02-acceptance/01-write/report-001.md`

## 2. Проверка покрытия

Пройтись по каждой секции spec и сверить с acceptance:

### ФТ Happy paths:
- §1 Go-интерфейсы: проверить что интерфейсы покрыты через тест компиляции/конструктора
- §2 Модели: поля Entity, Relation, Job
- §3.1 Upsert: создание, обновление, сохранение created_at, выставление updated_at, игнор входящих timestamps
- §3.2 Get: успешное получение
- §3.3 Delete: успешное удаление
- §3.4 List: непустой список, пустой список, порядок не гарантирован
- §4 FS-реализация: инициализация субдиректорий, atomic write, чтение, удаление, список
- §5 JSONL-роутер: все 12 комбинаций op×type, два запроса в одном соединении, завершение при EOF
- §5.2 Two-pass decode + разрешение конфликта id
- §5.6 Логирование: DEBUG success, ERROR failure, list с count без id

### ФТ Failure modes:
- §3.1: MISSING_ID, READ_ERROR при чтении существующей, WRITE_ERROR
- §3.2: MISSING_ID, NOT_FOUND, READ_ERROR
- §3.3: MISSING_ID, NOT_FOUND, DELETE_ERROR
- §3.4: LIST_ERROR (директория), LIST_ERROR (файл)
- §4.1: ошибка MkdirAll → конструктор возвращает ошибку
- §5.1: невалидный JSON → закрытие соединения
- §5.2: data absent/null для upsert → INVALID_REQUEST
- §5.4: UNKNOWN_OP, UNKNOWN_TYPE, INVALID_REQUEST (json), INVALID_REQUEST (data)

### Граничные условия:
- Пустой id для всех операций (upsert/get/delete)
- Пустой List
- List с невалидным файлом → ни одного элемента, не частичный результат
- List игнорирует не-.json файлы

### НФТ:
- НФТ-Н-1..3: atomic write, fsync, cleanup temp
- НФТ-Н-4: list не возвращает частичный результат
- НФТ-Н-5: failure при инициализации
- НФТ-П-1: комментарий TODO RWMutex — не поведенческий, не тестируется
- НФТ-Б-1: application-level ошибки не разрывают соединение; parse error — разрывает
- НФТ-О-1,2: логирование роутера, инициализации FS

## 3. Проверка качества

- Given — только контекст
- When — ровно одно действие
- Then — наблюдаемый результат с точки зрения вызывающей стороны
- Конкретные значения, конкретные errorCode
- Сценарии независимы
- Нет тривиальных сценариев

## 4. Классификация проблем

По каждой проблеме: пробел / неопределённость в spec / противоречие / качество.

## 5. Вывод

Если нет пробелов — done, следующий 03-tests/01-write.
Если есть пробелы — failed, следующий 02-acceptance/03-fix.
Если spec неполна — clarification, возврат к 01-spec/03-fix.
