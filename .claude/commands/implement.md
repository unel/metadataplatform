_Режим: **Харли Куин**. Выполняется оператором напрямую._

Реализуй фичу от code-write до build — оркестрирует шаги и обрабатывает сбои.

Аргументы: `$ARGUMENTS` — путь к задаче, например `feature/store-query`

## Подготовка
Прочитай `.claude/project.yaml` — возьми `max_retries` для шагов `code-write`, `build`, `test-run`, `code-review`, а также `retry_escalation` для тонального контекста.

Счётчик `implement-iterations` в `tasks/$ARGUMENTS/status.md` отражает общее количество возвратов к code-write. Лимит — `code-write.max_retries`.

## Паттерн цикла ревью

Применяется к каждому шагу у которого есть `max_retries` в конфиге. При провале:
1. Проверь счётчик (`implement-iterations` или специфичный для шага) vs `max_retries` шага
2. Если лимит исчерпан — **стоп**, показать историю сбоев, ждать инструкций
3. Иначе — инкрементируй счётчик, возьми тон из `retry_escalation[iterations]`, передай проблемы как контекст в шаг переделки
4. После переделки — **обязательно** повтори ревью-шаг (не пропускай его, переходя сразу дальше)

## Шаги

### 1. test-write
Вызови `/test-write $ARGUMENTS`.

### 1b. test-review
Вызови `/test-review $ARGUMENTS`.

**Если `test-review: failed`:**
- Проверь `test-iterations` в `status.md` vs `test-review.max_retries` из `.claude/project.yaml`
- Если лимит исчерпан — стоп, показать историю ревью, ждать инструкций
- Иначе — инкрементируй `test-iterations`, вызови `/fix-test $ARGUMENTS`, повтори с шага 1b

### 2. code-write
Вызови `/code-write $ARGUMENTS`.

**Если `code-write: failed`:**
- Если причина — `needs-rfc`: остановись, сообщи пользователю, жди решения
- Если причина — неясная спека: остановись, предложи уточнить spec через `/spec $ARGUMENTS`
- Иначе: покажи проблемы пользователю, жди инструкций

### 3. build
Вызови `/build $ARGUMENTS`.

**Если `build: failed`:**
- Проверь счётчик — если `implement-iterations ≥ build.max_retries`: остановись, покажи историю сбоев, жди инструкций
- Иначе — инкрементируй счётчик, передай ошибки сборки как контекст в `/code-write $ARGUMENTS`, повтори с шага 2

### 4. test-run
Вызови `/test-run $ARGUMENTS`.

**Если `test-run: failed`:**
- Проверь счётчик — если `implement-iterations ≥ test-run.max_retries`: остановись, покажи историю сбоев, жди инструкций
- Иначе — инкрементируй счётчик, передай упавшие тесты как контекст в `/code-write $ARGUMENTS`, повтори с шага 2

### 5. code-review
Вызови `/code-review $ARGUMENTS`.

**Если `code-review: failed`:**
- Проверь счётчик — если `implement-iterations ≥ code-review.max_retries`: остановись, покажи историю сбоев, жди инструкций
- Иначе — инкрементируй счётчик, вызови `/fix-code $ARGUMENTS`, повтори с шага 5

### 6. Финальная проверка
Вызови `/test-run $ARGUMENTS`. Если прошло — вызови coverage (`go test -cover ./...` для Go, `bun test --coverage` для frontend).

**Если что-то упало:**
- Проверь счётчик — если `implement-iterations ≥ code-write.max_retries`: остановись, покажи историю сбоев, жди инструкций
- Иначе — инкрементируй счётчик, передай контекст ошибок в `/code-write $ARGUMENTS`, повтори с шага 2

### Успех
Финальные тесты и coverage прошли — сообщи вызывающему что реализация завершена.
