# scripts/

Инструменты для управления workflow v2. Все скрипты запускаются из корня репозитория.

---

## set-step-status.py

Меняет статус шага в `status-log.md` с валидацией предусловий. **Единственный способ менять статус** — агенты не пишут в `status-log.md` напрямую.

```bash
python3 scripts/set-step-status.py <feature> <stage/step> <status> [--comment TEXT] [--dry-run]
```

**Статусы и их валидация:**

| Статус | Что проверяет |
|---|---|
| `in-progress` | `brief-NNN.md` существует; predecessor шаг terminal |
| `done` / `failed` | текущий статус `in-progress`; `report-NNN.md` существует; для `03-fix`: `in-response-to:` в frontmatter report |
| `stale` | `--comment` обязателен |
| `clarification` | `--comment` обязателен |
| `pending` | без проверок |

Run-номер определяется автоматически. Predecessor берётся из frontmatter `status-log.md` (проставляется скаффолдером).

**Порядок работы агента:**
1. Написать `brief-NNN.md` (входные данные, план прогона)
2. `set-step-status.py ... in-progress --comment "..."`
3. Выполнить работу
4. Написать `report-NNN.md`
5. `set-step-status.py ... done --comment "..."`

**Примеры:**
```bash
python3 scripts/set-step-status.py store-crud 01-spec/01-write in-progress \
  --comment "Читаю PROJECT.md и research-репорты."

python3 scripts/set-step-status.py store-crud 01-spec/01-write done \
  --comment "Spec v1 написана и согласована."

python3 scripts/set-step-status.py store-crud 02-acceptance/02-review failed \
  --comment "AC-3, AC-7 — пробелы в сценариях."

python3 scripts/set-step-status.py store-crud 01-spec/02-review stale \
  --comment "Spec обновлена в 03-fix run 2."

# Проверить без записи:
python3 scripts/set-step-status.py store-crud 04-code/03-fix done --dry-run
```

---

## feature-status.py

Показывает статус всех шагов фичи. Оркестратор использует для навигации по pipeline.

```bash
python3 scripts/feature-status.py [--feature PATTERN] [--all] [--short] [--json]
```

**Примеры:**
```bash
# Статус текущей фичи (по CWD):
python3 scripts/feature-status.py

# Конкретная фича:
python3 scripts/feature-status.py --feature store-crud

# Все фичи коротко:
python3 scripts/feature-status.py --all --short

# JSON для скриптов:
python3 scripts/feature-status.py --feature store-crud --json
```

**Иконки:** `✓` done · `✗` failed · `~` in-progress · `·` pending · `↻` stale · `!` clarification

---

## scaffold-feature.py

Создаёт директорию `tasks/<slug>/stages/` со всеми шагами по флоу. Записывает frontmatter (`previous-step`, `next-success-step`, `next-fail-step`) в каждый `status-log.md`.

```bash
python3 scripts/scaffold-feature.py <feature-slug> [--flow <name>] [--dry-run]
```

Флоу-файлы: `docs/standards/v2/flows/`. По умолчанию — `default`.

**Примеры:**
```bash
python3 scripts/scaffold-feature.py 0003-store-protocol
python3 scripts/scaffold-feature.py 0003-store-protocol --flow minimal
python3 scripts/scaffold-feature.py 0003-store-protocol --dry-run
```

---

## log-note.py

Добавляет наблюдение в `notes-<agent>.md` текущего шага.

```bash
python3 scripts/log-note.py --agent <agent> --message "<текст>" [--step <stage/step>]
```

**Примеры:**
```bash
python3 scripts/log-note.py --agent ада \
  --message "[propose] Вынести валидацию ID в отдельную функцию"

python3 scripts/log-note.py --user \
  --message "[miss] Не описали поведение при пустом ID"

# Явно указать шаг (если не определяется по CWD):
python3 scripts/log-note.py --agent гримм --step 04-code/02-review \
  --message "[friction] Список файлов не был передан в промпт"
```

**Теги:** `[propose]` `[friction]` `[miss]` `[rework]` `[doc]` `[whatever]`

---

## log-complaint.py

Добавляет жалобу в `complaints-<agent>.md`. Формат свободный, без тегов.

```bash
python3 scripts/log-complaint.py --agent <agent> --message "<текст>" [--step <stage/step>]
```

**Примеры:**
```bash
python3 scripts/log-complaint.py --agent кроули \
  --message "Тесты писались к несуществующему API, пришлось угадывать сигнатуры"

python3 scripts/log-complaint.py --user \
  --message "Три итерации ревью на то что можно было поймать в acceptance"
```

---

## validate-write-path.py

PreToolUse-хук для Claude Code. Блокирует запись агента в пути вне его разрешений из `AGENT.md`.

Вызывается автоматически через hooks — не запускается вручную.
