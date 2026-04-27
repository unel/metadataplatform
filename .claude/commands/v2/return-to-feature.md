_Режим: **Харли Куин**. Выполняется оператором напрямую._

# return-to-feature v2

Восстанови контекст и продолжи работу над фичей. Используется после перерыва, после прерванной сессии, или когда нужно понять где мы.

Аргументы: `$ARGUMENTS` — путь к фиче, например `store/query`

---

## 1. Контекст

Прочитай параллельно:
- `tasks/$ARGUMENTS/spec.md` (если есть) — цель и суть фичи
- `tasks/$ARGUMENTS/acceptance.md` (если есть) — критерии приёмки
- Последние коммиты через `git log --oneline -10`

---

## 2. Git-ветка

Проверь через `git branch --list "feature/$ARGUMENTS" "bug/$ARGUMENTS"`:
- Если ветка есть — переключись: `git checkout feature/$ARGUMENTS`
- Если ветки нет и нужна — отбранчуйся от `master`/`main`
- Если уже на нужной ветке — продолжай

---

## 3. Сканирование статусов

Прочитай `status-log.md` каждого процесса в порядке pipeline. Читай **последнюю запись** каждого файла.

Порядок:
```
00-research/01-interview
00-research/02-web
01-spec/01-write → 01-spec/02-review → 01-spec/03-fix
02-acceptance/01-write → 02-acceptance/02-review → 02-acceptance/03-fix
03-tests/01-write → 03-tests/02-review → 03-tests/03-fix → 03-tests/04-run
04-code/01-write → 04-code/02-review → 04-code/03-fix → 04-code/04-testing
05-docs/01-write → 05-docs/02-review → 05-docs/03-fix
06-retro/01-recall → 06-retro/02-collect → 06-retro/03-analyze → 06-retro/04-solve → 06-retro/05-write
```

Покажи пользователю сводку:

```
Статус фичи: store/query
─────────────────────────────────────────
  00-research/01-interview   done
  00-research/02-web         done
  01-spec/01-write           done
  01-spec/02-review          done
  02-acceptance/01-write     done
  02-acceptance/02-review    failed  ← точка входа
  02-acceptance/03-fix       pending
  ...
```

Если скаффолд не существует (нет директории `tasks/$ARGUMENTS/processes/`) — запусти `/v2:feature $ARGUMENTS` и остановись.

---

## 4. Точка входа

Найди первый процесс с последней записью **не** `done`.

| Статус | Действие |
|---|---|
| `pending` | запустить скилл этого процесса |
| `stale` | запустить скилл заново (артефакт изменился выше) |
| `failed` | см. **Цикл ревью** |
| `clarification` | см. **Clarification** |
| `in-progress` | стоп — показать контекст, спросить инструкций |

Если все `done` — фича завершена. Сообщи пользователю, предложи `/v2:06-retro:01-recall $ARGUMENTS` если ретро ещё не делали.

---

## 5. Продолжение

### Обычный запуск (`pending` / `stale`)

Запусти скилл процесса:

| Процесс | Скилл |
|---|---|
| `00-research/01-interview` | `/v2:00-research:01-interview $ARGUMENTS` |
| `00-research/02-web` | `/v2:00-research:02-web $ARGUMENTS` |
| `01-spec/01-write` | `/v2:01-spec:01-write $ARGUMENTS` |
| `01-spec/02-review` | `/v2:01-spec:02-review $ARGUMENTS` |
| `01-spec/03-fix` | `/v2:01-spec:03-fix $ARGUMENTS` |
| `02-acceptance/01-write` | `/v2:02-acceptance:01-write $ARGUMENTS` |
| `02-acceptance/02-review` | `/v2:02-acceptance:02-review $ARGUMENTS` |
| `02-acceptance/03-fix` | `/v2:02-acceptance:03-fix $ARGUMENTS` |
| `03-tests/01-write` | `/v2:03-tests:01-write $ARGUMENTS` |
| `03-tests/02-review` | `/v2:03-tests:02-review $ARGUMENTS` |
| `03-tests/03-fix` | `/v2:03-tests:03-fix $ARGUMENTS` |
| `03-tests/04-run` | `/v2:03-tests:04-run $ARGUMENTS` |
| `04-code/01-write` | `/v2:04-code:01-write $ARGUMENTS` |
| `04-code/02-review` | `/v2:04-code:02-review $ARGUMENTS` |
| `04-code/03-fix` | `/v2:04-code:03-fix $ARGUMENTS` |
| `04-code/04-testing` | `/v2:04-code:04-testing $ARGUMENTS` |
| `05-docs/01-write` | `/v2:05-docs:01-write $ARGUMENTS` |
| `05-docs/02-review` | `/v2:05-docs:02-review $ARGUMENTS` |
| `05-docs/03-fix` | `/v2:05-docs:03-fix $ARGUMENTS` |
| `06-retro/01-recall` | `/v2:06-retro:01-recall $ARGUMENTS` |
| `06-retro/02-collect` | `/v2:06-retro:02-collect $ARGUMENTS` |
| `06-retro/03-analyze` | `/v2:06-retro:03-analyze $ARGUMENTS` |
| `06-retro/04-solve` | `/v2:06-retro:04-solve $ARGUMENTS` |
| `06-retro/05-write` | `/v2:06-retro:05-write $ARGUMENTS` |

После завершения скилла — продолжи pipeline через `/v2:feature $ARGUMENTS`.

### Цикл ревью (`failed`)

1. Посчитай записи `— done` и `— failed` в `<group>/03-fix/status-log.md` — это итерации
2. Если итераций ≥ 3 → **стоп**: показать историю провалов, ждать инструкций пользователя
3. Иначе — запусти fix-скилл:

| Review с failed | Fix |
|---|---|
| `01-spec/02-review` | `/v2:01-spec:03-fix $ARGUMENTS` |
| `02-acceptance/02-review` | `/v2:02-acceptance:03-fix $ARGUMENTS` |
| `03-tests/02-review` | `/v2:03-tests:03-fix $ARGUMENTS` |
| `04-code/02-review` | `/v2:04-code:03-fix $ARGUMENTS` |
| `05-docs/02-review` | `/v2:05-docs:03-fix $ARGUMENTS` |

4. После fix → каскадный сброс (`stale` всем последующим `done`-процессам) → продолжи через `/v2:feature $ARGUMENTS`

### Clarification

Процесс стоит в `clarification` — upstream-артефакт неполон или противоречив.

1. Прочитай report процесса с `clarification` — там описано что неясно.
2. Определи откат:
   - Проблема в acceptance → `02-acceptance/03-fix`
   - Проблема в spec → `01-spec/03-fix`
   - Нужно уточнение у пользователя → **стоп**, задай вопрос

3. Если откат определён:
   - Добавь `in-progress` в `status-log.md` целевого upstream-шага с объяснением возврата
   - Добавь `stale` последующим шагам **той же группы** (если были `done`/`failed`)
   - Добавь `stale` всем группам начиная с группы в `clarification` (если были `done`/`failed`)
   - Запусти upstream fix-скилл
   - После fix → продолжи через `/v2:feature $ARGUMENTS`

---

## 6. Доклад перед запуском

Прежде чем запускать скилл, скажи пользователю:
- Где находимся (процесс, статус)
- Что делаешь прямо сейчас
- Если `clarification` или ≥3 итераций — объясни ситуацию и жди ответа

Для `pending`/`stale` — запускай немедленно, без лишних вопросов.
