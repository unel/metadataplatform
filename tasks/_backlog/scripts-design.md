# Scripts Design: RE-010, RE-011

## Язык и зависимости

**Python 3.12.** Только стандартная библиотека — никаких внешних зависимостей (`pathlib`, `argparse`, `json`, `re`, `sys`).

Frontmatter — YAML-стиль (блок между `---`), парсится вручную. Двоеточие в значении не обрезает — используем `split(':', 1)`:

```python
def parse_frontmatter(text: str) -> dict:
    if not text.startswith('---'):
        return {}
    _, _, rest = text.partition('\n')
    block, sep, _ = rest.partition('\n---')
    if not sep:
        return {}
    meta = {}
    for line in block.splitlines():
        parts = line.split(':', 1)
        if len(parts) == 2 and parts[0].strip():
            meta[parts[0].strip()] = parts[1].strip()
    return meta
```

## Расположение

Все скрипты — в `scripts/` в корне репозитория. Расширение `.py`. Shebang: `#!/usr/bin/env python3`. Права `chmod +x`.

Запуск без аргументов (или с `-h` / `--help`) печатает usage и завершается с `exit 0`.

## Exit codes

| Ситуация | Exit code |
|---|---|
| Успех | 0 |
| Warning (unknown status, active step not found, повторный in-progress) | 0 |
| `--dry-run` | 0 |
| Ошибка (файл не найден, невалидный путь, --to не в цепочке) | 1 |

---

## Frontmatter status-log.md

Frontmatter — **статический**: заполняется один раз при инициализации скелета фичи, никогда не изменяется (см. RE-037).

```yaml
---
next-success-step: 02-acceptance/01-write
next-fail-step: 01-spec/03-fix
previous-step: 01-spec/01-write
---
```

Поля:
- `next-success-step` — следующий шаг при успешном завершении (`done`)
- `next-fail-step` — шаг при провале (`failed`); обычно соответствующий fix-шаг
- `previous-step` — предыдущий шаг в pipeline

Последний шаг pipeline (`05-docs/02-review`): `next-success-step` пустой. `06-retro` в граф не входит — у его шагов `next-success-step` тоже пустой, cascade-stale на них не распространяется. Это ожидаемое поведение.

**Инициализация**: `status-log.md` с frontmatter создаётся scaffold-скриптом при инициализации скелета фичи (RE-037 — система флоу). До инициализации скрипты не вызываются.

**Ограничение**: `status-log.md` не редактируется вручную — только через скрипты. Ручная правка нарушает инварианты на которых строится логика begin-step/end-step/cascade-stale.

`next-success-step` пустая строка или отсутствующий ключ — оба означают конец цепочки (в `parse_frontmatter` возвращается `""` в обоих случаях через `.get('next-success-step', '')`).

Пустой frontmatter (`---\n---\n`) не поддерживается — вернёт пустой dict.

### Номер рана

**Не хранится в frontmatter.** Кодируется в строках статуса: `# <datetime> — <status> run=N`.

Текущий ран = `max(re.findall(r'run=(\d+)', content), default=0)`.

Frontmatter остаётся append-only-safe — обновление не требуется.

---

## Состояния шага

### Нормальные состояния

| Последний статус | brief-N.md | report-N.md | Состояние |
|---|---|---|---|
| `pending` (или нет записей) | нет | нет | `pending` |
| `in-progress run=N` | есть | нет | `in-progress` |
| `done run=N` | есть | есть | `done` |
| `failed run=N` | есть | есть | `failed` |
| `stale` | * | * | `stale` |
| `clarification` | * | * | `clarification` |
| отсутствует `status-log.md` | — | — | `corrupted` |

### Inconsistent-состояния

`detail` — строка, попадает в JSON-вывод `feature-status.py` и в verbose в скобках.

| Последний статус | brief-N.md | report-N.md | detail |
|---|---|---|---|
| `in-progress run=N` | нет | нет | `"brief-{N:03d}.md missing"` |
| `in-progress run=N` | есть | есть | `"report written but end-step not called"` |
| `done/failed run=N` | есть | нет | `"report-{N:03d}.md missing"` |
| `done/failed run=N` | нет | * | `"brief-{N:03d}.md missing"` |
| `pending`/`stale` | любой есть | любой есть | `"stale artifacts from previous run"` |
| max N в статусах ≠ max N в brief/report | — | — | `"run counter mismatch: log={X}, files={Y}"` |

---

## feature-status.py (RE-010)

### Вызов

```bash
scripts/feature-status.py [--feature <pattern>] [--all] [--short] [--json] [--dry-run]
                           [--max-matches <N>] [--unique]
```

Примеры:
```bash
scripts/feature-status.py                           # auto: CWD → ищем tasks/<feat>/ вверх
scripts/feature-status.py --feature store/crud      # glob: tasks/*store*crud*/
scripts/feature-status.py --feature store --max-matches 3
scripts/feature-status.py --feature store --unique  # alias для --max-matches 1
scripts/feature-status.py --all
scripts/feature-status.py --short
scripts/feature-status.py --json
```

### Резолюция фичи (порядок приоритетов)

1. `--all` → все `tasks/*/` кроме `_backlog`
2. `--feature <pattern>` → glob `tasks/*<pattern>*/`; `--max-matches N` — ошибка если найдено > N (exit 1); `--unique` = `--max-matches 1`
3. нет флага → CWD: идём вверх, ищем `tasks/<something>/` в пути
4. ничего → stderr + exit 1

### Парсинг status-log.md

Текущий статус — последняя строка вида `# <datetime> — <status>` (может содержать `run=N`).

Статус не ограничен фиксированным списком. Неизвестные → иконка `*`.

Порядок шагов — filesystem order: `sorted((task_dir / 'stages').rglob('status-log.md'))`, где `task_dir` — директория задачи найденная при резолюции фичи (не CWD). Директории используют числовые префиксы (стандарт: `docs/standards/v2/stages.md`).

### Определение активного шага

| Приоритет | Условие | Смысл |
|---|---|---|
| 1 | первый `in-progress` | выполняется прямо сейчас |
| 2 | первый `failed` или `clarification` | застряли, нужен fix |
| 3 | первый `stale` | upstream изменился, повторить |
| 4 | первый `pending` | следующий в очереди |
| 5 | все `done` | фича завершена |

### Inconsistent-детект

Для каждого шага перед отображением — проверка по таблице состояний выше. При несоответствии — статус заменяется на `inconsistent` с деталью в скобках.

### Форматы вывода

**Verbose (default):**
```
0002-store-crud
  01-spec/01-write       ✓ done         2026-04-27T09:12:44Z
→ 04-code/03-fix         ~ in-progress  2026-05-04T17:30:56Z
  04-code/04-testing     · pending
  05-docs/01-write       ? corrupted
  05-docs/02-review      % inconsistent (report deleted)
  06-retro/01-recall     * unknownstat  2026-05-01T08:00:00Z
```

**Short (`--short`):**
```
0002-store-crud  →  04-code/03-fix  [in-progress]  2026-05-04T17:30:56Z
```

**JSON (`--json`):**
```json
{
  "feature": "0002-store-crud",
  "active_step": "04-code/03-fix",
  "steps": [
    { "step": "01-spec/01-write", "status": "done", "timestamp": "2026-04-27T09:12:44Z", "active": false },
    { "step": "04-code/03-fix",   "status": "in-progress", "timestamp": "2026-05-04T17:30:56Z", "active": true },
    { "step": "05-docs/02-review","status": "inconsistent", "detail": "report deleted", "timestamp": null, "active": false }
  ]
}
```

При `--all` verbose — пустая строка между фичами; JSON — массив объектов.

### Иконки

| Статус | Иконка |
|---|---|
| done | `✓` |
| failed | `✗` |
| in-progress | `~` |
| pending | `·` |
| stale | `↻` |
| clarification | `!` |
| corrupted | `?` |
| inconsistent | `%` |
| неизвестный | `*` |

### Тестовые сценарии (проверено 2026-05-11)

```bash
# Verbose по имени фичи
scripts/feature-status.py --feature 0002-store-crud
# → все шаги done; 02-acceptance/02-review показывает * pass (нестандартный статус — корректно)

# Short — одна строка
scripts/feature-status.py --feature 0002-store-crud --short
# → "0002-store-crud  ✓  (all done)"

# Все фичи коротко (включая старые без stages/ → error-строки)
scripts/feature-status.py --all --short

# JSON для одной фичи
scripts/feature-status.py --feature 0002-store-crud --json
# → {"feature": "0002-store-crud", "active_step": null, "steps": [...]}

# JSON массив (--all); фичи без stages/ → {"feature": "...", "error": "no stages/ directory"}
scripts/feature-status.py --all --json

# CWD-резолюция изнутри папки шага
cd tasks/0002-store-crud/stages/04-code/03-fix && scripts/feature-status.py --short
# → резолюция работает, возвращает статус 0002-store-crud

# --unique с неоднозначным паттерном → exit 1 с именами совпадений
scripts/feature-status.py --feature store --unique
# → exit 1: "matched 3 tasks (max 1): 0001-store-connection, 0002-store-crud, store"

# Без параметров из корня репо → usage + exit 0
scripts/feature-status.py
```

### Design-решения принятые при реализации

**Толерантность к старому формату без `run=N`**: если в `status-log.md` нет ни одной строки `run=N`, inconsistent-проверки на brief/report-N пропускаются. Задачи написанные до введения системы run=N отображаются корректно без false-positive `inconsistent`.

**Фичи без `stages/`** (старый формат v1): в verbose/short выводятся как `error — no stages/ directory`; в JSON — `{"feature": "...", "error": "..."}`.

---

## begin-step.py

### Вызов

```bash
scripts/begin-step.py [--step <stage>/<step>] [--dry-run]
```

### Алгоритм

1. Определяет целевой шаг (CWD или `--step`); валидирует что путь существует внутри `stages/` (exit 1 если нет)
2. Определяет состояние рана из status-log.md:
   - Последний статус (последнее вхождение паттерна `# <datetime> — <status>`) = `in-progress run=M` → **recovery**: текущий открытый ран = M
   - Последний статус = `done run=*` → stderr warning: `warning: step is already done, opening new run`; продолжает
   - Иначе: `last_run` = max N по всем записям вида `run=N` (default 0); следующий ран = `last_run + 1`
3. Проверяет наличие `brief-{run:03d}.md` (run = M для recovery, last_run+1 для нового) — exit 1 если нет. Ответственность за создание brief до вызова begin-step лежит на агенте (см. RE-045)
4. Recovery-путь: перезаписывает datetime в строке последнего вхождения `# <datetime> — in-progress run=M` (единственное нарушение append-only; реализуется через чтение всего файла, замену строки, запись обратно) + stderr warning: `warning: step already in-progress (run {M}), updating timestamp`; exit 0
5. Если `--dry-run` — печатает что сделал бы; exit 0
6. Пишет в status-log (append):
   ```
   # <datetime> — in-progress run=N+1
   → brief-{N+1:03d}.md

   ```

### Резолюция целевой папки

1. `--step <stage>/<step>` → `tasks/<feat>/stages/<stage>/<step>/`; фича из CWD
2. CWD внутри шага (`…/stages/<stage>/<step>/`) → пишем туда
3. ничего → stderr + exit 1

---

## end-step.py

### Вызов

```bash
scripts/end-step.py --status <status> --message "<комментарий>"
                    [--step <stage>/<step>] [--dry-run] [--amend]
```

`--status` и `--message` обязательны (кроме `--amend` — см. ниже).

### Алгоритм

1. Определяет целевой шаг; валидирует путь
2. Ищет последнее вхождение паттерна `# <datetime> — in-progress run=M` в status-log.md — exit 1 если нет (begin-step не был вызван для текущего рана)
3. Проверяет наличие `brief-{M:03d}.md` — exit 1 если нет
4. Проверяет наличие `report-{M:03d}.md` — exit 1 если нет
5. Если статус не из известных (`done`, `failed`, `stale`, `clarification`) → stderr warning: `warning: unknown status "<status>"` (exit code остаётся 0)
6. Если `--dry-run` — печатает что сделал бы (включая cascade-stale `--dry-run`); exit 0
7. Пишет в status-log (append):
   ```
   # <datetime> — <status> run=M
   <комментарий> → report-{M:03d}.md

   ```
8. Если статус `done` → вызывает `cascade-stale.py --from <текущий шаг> --message "<комментарий>"` (с `--dry-run` если был)
9. Если статус не `done` → cascade-stale не вызывается

### Флаг --amend

Перезаписывает последнюю запись в status-log вместо добавления. `--status` и `--message` опциональны (не переданное берётся из текущей последней записи). Не меняет счётчик рана.

`--message` — однострочный. Многострочный комментарий не поддерживается.

**Парсинг последней записи**: от конца файла до предыдущего заголовка `# <datetime> ...` включительно. Первая строка этого блока — заголовок (`# <datetime> — <status> run=M`), следующая строка — комментарий (до ` → report-NNN.md`). При записи обратно — `→ report-{M:03d}.md` добавляется ровно один раз (если уже присутствует в комментарии — не дублируется).

- Если статус меняется `* → done` → вызывает cascade-stale
- Если статус меняется `done → *` → stderr warning: `warning: step was done, downstream steps may have been cascaded; validate their statuses manually`; cascade-stale не вызывается

### Резолюция целевой папки

Аналогично `begin-step.py`.

---

## cascade-stale.py

### Вызов

```bash
scripts/cascade-stale.py --message "<причина>"
                         [--step <stage>/<step>] [--from <stage>/<step>] [--to <stage>/<step>]
                         [--dry-run]
```

`--message` обязателен.

Обычно вызывается автоматически из `end-step.py`. Можно вызвать вручную.

### Алгоритм

1. Определяет стартовый шаг (`--from` или текущий из CWD / `--step`)
2. **Предварительно** обходит всю цепочку `next-success-step` от старта до конца (или `--to`), строит упорядоченный список шагов
3. Если `--to` указан но не найден в цепочке → stderr + exit 1 (до изменения файлов)
4. Если `--dry-run` → печатает все шаги из списка с их действием (`mark stale` / `skip: pending` / `skip: already stale`); exit 0
5. Для каждого шага в списке (эксклюзивно от стартового):
   - `stale` → пропускает (уже сброшен)
   - `pending` — поведение зависит от наличия `--to`:
     - `--to` **не указан** → останавливает обход (последующие шаги тоже `pending`)
     - `--to` **указан** → помечает `stale` (явный диапазон override; агент сам ответственен за корректность)
   - иначе → помечает `stale`
   - иначе → пишет в status-log (append), используя run=N из предыдущей записи шага (`max run=N` в его status-log.md):
     ```
     # <datetime> — stale run=N
     <причина>

     ```

`06-retro` в каскад не входит — `05-docs/02-review` имеет пустой `next-success-step`.

### Резолюция целевой папки

1. `--from` → явный стартовый шаг (эксклюзивно)
2. `--step <stage>/<step>` → стартовый шаг из явного аргумента
3. CWD внутри шага → стартовый шаг из CWD
4. ничего → stderr + exit 1

---

## log-note.py / log-complaint.py (RE-011)

### Вызов

```bash
scripts/log-note.py --agent <агент> --message "<тег> текст" [--step <stage>/<step>] [--dry-run]
scripts/log-complaint.py --agent <агент> --message "<тег> текст" [--step <stage>/<step>] [--dry-run]
```

### Резолюция целевой папки

1. `--step <stage>/<step>` → `tasks/<feat>/stages/<stage>/<step>/`; фича из CWD
2. CWD внутри шага (`…/stages/<stage>/<step>/`) → пишем туда
3. CWD внутри фичи, не в шаге → запрашиваем активный шаг через `feature-status.py --json` (subprocess без `--all`), парсим JSON-объект (не массив), берём `active_step`
4. Активный шаг не определён (все `done`, фича завершена) → пишем в `stages/06-retro/01-recall/`; сообщение предваряем `[from-unknown-step]`; stderr warning: `warning: active step not found, writing to retro/01-recall`; exit 0
5. ничего → stderr + exit 1

### Целевой файл

```
notes-<агент>.md      (log-note.py)
complaints-<агент>.md (log-complaint.py)
```

Создаётся если не существует. Дозаписываем строку в конец без timestamp, без разделителя между записями — каждая запись ровно одна строка. `--message` передаётся as-is:

```
[propose] Лучше использовать интерфейс вместо конкретного типа
[friction] Тесты писались к несуществующему API
[rework] Пришлось переписать половину из-за изменившегося контракта
```

### Конкурентный доступ

Параллельные агенты (Кроули + Азирафаль) пишут в **разные шаги** — конфликта нет по дизайну. За `status-log.md` в `06-retro/*` отвечает только Харли — один писатель, конфликтов нет.

### Резолюция фичи (для всех скриптов)

Из CWD: идём вверх, ищем `tasks/<feat>/` в пути. Если не нашли → stderr + exit 1.
