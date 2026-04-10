_Исполнитель: агент **Харли** (`harley`) — оркестрирует **Аду**, **Гримма**, **Кроули**, **Азирафаля**._

Реализуй фичу от code-write до build — оркестрирует шаги и обрабатывает сбои.

Аргументы: `$ARGUMENTS` — путь к задаче, например `feature/store-query`

## Подготовка
Прочитай `.claude/project.yaml` — возьми `max_retries` для шагов `code-write`, `build`, `test-run`, `code-review`.

Счётчик `implement-iterations` в `tasks/$ARGUMENTS/status.md` отражает общее количество возвратов к code-write. Лимит — `code-write.max_retries`.

## Шаги

### 1. code-write
Вызови `/code-write $ARGUMENTS`.

**Если `code-write: failed`:**
- Если причина — `needs-rfc`: остановись, сообщи пользователю, жди решения
- Если причина — неясная спека: остановись, предложи уточнить spec через `/spec $ARGUMENTS`
- Иначе: покажи проблемы пользователю, жди инструкций

### 2. build
Вызови `/build $ARGUMENTS`.

**Если `build: failed`:**
- Проверь счётчик — если `implement-iterations ≥ build.max_retries`: остановись, покажи историю сбоев, жди инструкций
- Иначе — инкрементируй счётчик, передай ошибки сборки как контекст в `/code-write $ARGUMENTS`, повтори с шага 2

### 3. test-run
Вызови `/test-run $ARGUMENTS`.

**Если `test-run: failed`:**
- Проверь счётчик — если `implement-iterations ≥ test-run.max_retries`: остановись, покажи историю сбоев, жди инструкций
- Иначе — инкрементируй счётчик, передай упавшие тесты как контекст в `/code-write $ARGUMENTS`, повтори с шага 2

### 4. code-review
Вызови `/code-review $ARGUMENTS`.

**Если `code-review: failed`:**
- Проверь счётчик — если `implement-iterations ≥ code-review.max_retries`: остановись, покажи историю сбоев, жди инструкций
- Иначе — инкрементируй счётчик, передай критичные замечания как контекст в `/code-write $ARGUMENTS`, повтори с шага 2

### 5. Финальная проверка
Вызови `/test-run $ARGUMENTS`. Если прошло — вызови coverage (`go test -cover ./...` для Go, `bun test --coverage` для frontend).

**Если что-то упало:**
- Проверь счётчик — если `implement-iterations ≥ code-write.max_retries`: остановись, покажи историю сбоев, жди инструкций
- Иначе — инкрементируй счётчик, передай контекст ошибок в `/code-write $ARGUMENTS`, повтори с шага 2

### Успех
Финальные тесты и coverage прошли — сообщи вызывающему что реализация завершена.
