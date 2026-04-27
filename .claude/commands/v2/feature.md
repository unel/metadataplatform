_Режим: **Харли Куин**. Выполняется оператором напрямую._

# feature v2

Полный цикл фичи: research → spec → acceptance → tests → code → docs → retro.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

## Инициализация

### 1. Скелет процессов

Создай все директории сразу (включая фикс-шаги):

```
tasks/$ARGUMENTS/processes/
  00-research/01-interview/
  00-research/02-web/
  01-spec/01-write/
  01-spec/02-review/
  01-spec/03-fix/
  02-acceptance/01-write/
  02-acceptance/02-review/
  02-acceptance/03-fix/
  03-tests/01-write/
  03-tests/02-review/
  03-tests/03-fix/
  03-tests/04-run/
  04-code/01-write/
  04-code/02-review/
  04-code/03-fix/
  04-code/04-testing/
  05-docs/01-write/
  05-docs/02-review/
  05-docs/03-fix/
  06-retro/01-recall/
  06-retro/02-collect/
  06-retro/03-analyze/
  06-retro/04-solve/
  06-retro/05-write/
```

В каждой директории создай `status-log.md` с начальной записью (если файл не существует):

```markdown
# {{datetime}} — pending
Ожидает запуска.
```

Если директории и логи уже существуют — не трогай их.

### 2. Точка входа

Прочитай `status-log.md` каждого процесса (в порядке pipeline). Найди первый, у которого последняя запись **не** `done`.

Порядок процессов в pipeline:
```
00-research/01-interview
00-research/02-web
01-spec/01-write → 01-spec/02-review [→ 01-spec/03-fix]
02-acceptance/01-write → 02-acceptance/02-review [→ 02-acceptance/03-fix]
03-tests/01-write → 03-tests/02-review [→ 03-tests/03-fix] → 03-tests/04-run
04-code/01-write → 04-code/02-review [→ 04-code/03-fix] → 04-code/04-testing
05-docs/01-write → 05-docs/02-review [→ 05-docs/03-fix]
06-retro/01-recall → 02-collect → 03-analyze → 04-solve → 05-write
```

Статусы последней записи в `status-log.md`:

| Статус | Действие |
|---|---|
| `done` | пропустить |
| `pending` | выполнить |
| `stale` | выполнить заново |
| `failed` | применить цикл ревью |
| `clarification` | каскадный откат назад (см. **Clarification**) |
| `in-progress` | что-то прервалось — сообщи пользователю, жди инструкций |

Если все `done` — фича завершена, сообщи пользователю.

## Clarification

Когда процесс стоит в `clarification` — активный шаг сигнализирует что upstream-артефакт неполон или противоречив. Алгоритм:

1. Прочитай report активного процесса — там описано что именно неясно.
2. Определи куда откатываться:

| Источник проблемы | Откат к |
|---|---|
| acceptance неточна / противоречива | `02-acceptance/03-fix` |
| spec неясна | `01-spec/03-fix` |
| spec противоречива | `01-spec/03-fix` (или `01-spec/01-write` если полная переработка) |
| требуется уточнение у пользователя | стоп: задай вопрос, жди ответа |

3. Добавь запись `in-progress` в `status-log.md` целевого upstream-шага:
   ```
   # {{datetime}} — in-progress
   Возврат из clarification: <описание проблемы>.
   ```
4. Добавь `stale` всем последующим шагам **той же группы** что целевой upstream-шаг (если были `done`/`failed`).
5. Добавь `stale` всем последующим **группам** начиная с той что стояла в `clarification` (если были `done`/`failed`).
6. Запусти upstream fix-скилл.
7. После завершения fix — каскадный сброс вперёд (как обычный цикл ревью) и продолжай pipeline.

---

## Pipeline

### 00-research (параллельно)

Запусти параллельно:
- `/v2:00-research:01-interview $ARGUMENTS`
- `/v2:00-research:02-web $ARGUMENTS`

Если контекст уже задокументирован — добавь в каждый `status-log.md`:
```
# {{datetime}} — done
Пропущено — контекст предоставлен пользователем.
```

### 01-spec

1. `/v2:01-spec:01-write $ARGUMENTS`
2. `/v2:01-spec:02-review $ARGUMENTS` → если `failed` → **цикл ревью**

### 02-acceptance

1. `/v2:02-acceptance:01-write $ARGUMENTS`
2. `/v2:02-acceptance:02-review $ARGUMENTS` → если `failed` → **цикл ревью**

### 03-tests

1. `/v2:03-tests:01-write $ARGUMENTS`
2. `/v2:03-tests:02-review $ARGUMENTS` → если `failed` → **цикл ревью**
3. `/v2:03-tests:04-run $ARGUMENTS`
   - Все упали (Red) → дальше
   - Тест прошёл без реализации → запиши `failed` в `status-log.md`, применяй цикл ревью через `03-tests/03-fix`

### 04-code

1. `/v2:04-code:01-write $ARGUMENTS`
2. `/v2:04-code:02-review $ARGUMENTS` → если `failed` → **цикл ревью**
3. `/v2:04-code:04-testing $ARGUMENTS`
   - Все прошли (Green) → дальше
   - Падения → запиши `failed` в `04-code/04-testing/status-log.md`, каскадный сброс от `04-code/01-write`

### 05-docs

1. `/v2:05-docs:01-write $ARGUMENTS`
2. `/v2:05-docs:02-review $ARGUMENTS` → если `failed` → **цикл ревью**

### 06-retro

Последовательно:
1. `/v2:06-retro:01-recall $ARGUMENTS`
2. `/v2:06-retro:02-collect $ARGUMENTS`
3. `/v2:06-retro:03-analyze $ARGUMENTS`
4. `/v2:06-retro:04-solve $ARGUMENTS`
5. `/v2:06-retro:05-write $ARGUMENTS`

---

## Цикл ревью

Когда review возвращает `failed`:

1. Посчитай записи `— done` / `— failed` в `<group>/03-fix/status-log.md` — это итерации
2. Если итераций ≥ 3 → **стоп**: покажи историю провалов, жди инструкций пользователя
3. Fix-скилл по группе:

| Review | Fix |
|---|---|
| `01-spec/02-review` | `/v2:01-spec:03-fix $ARGUMENTS` |
| `02-acceptance/02-review` | `/v2:02-acceptance:03-fix $ARGUMENTS` |
| `03-tests/02-review` | `/v2:03-tests:03-fix $ARGUMENTS` |
| `04-code/02-review` | `/v2:04-code:03-fix $ARGUMENTS` |
| `05-docs/02-review` | `/v2:05-docs:03-fix $ARGUMENTS` |

4. После fix → **каскадный сброс** → запусти review снова
5. Если снова `failed` → повтори с п.1

## Каскадный сброс

После fix — добавь запись `stale` в `status-log.md` всех последующих процессов которые были `done`:

```markdown
# {{datetime}} — stale
Upstream артефакт изменён: <что именно>. Требует повторного прохода.
```

Цепочка (каждый шаг зависит от всего выше):
```
01-spec/01-write
  → 01-spec/02-review
  → 02-acceptance/01-write → 02-acceptance/02-review
  → 03-tests/01-write → 03-tests/02-review → 03-tests/04-run
  → 04-code/01-write → 04-code/02-review → 04-code/04-testing
  → 05-docs/01-write → 05-docs/02-review
```

## Итог

По завершении `06-retro/05-write`:
- Список артефактов по группам
- Финальный статус тестов (Green)
- Путь к RETRO.md с action items (≤ 3, с DRI и дедлайнами)
