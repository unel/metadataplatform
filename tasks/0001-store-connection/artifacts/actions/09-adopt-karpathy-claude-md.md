# Action: адоптировать практики из CLAUDE.md (andrej-karpathy-skills)

## Источник

https://github.com/forrestchang/andrej-karpathy-skills/blob/main/CLAUDE.md

## Что нашли

Семь механик, три из которых прямо адресуют боли этого ретро:

### 1. Явные критерии верификации на каждом шаге

Шаблон: `[Шаг] → verify: [проверка]`

Закрывает тест-симулякры: если в скилле test-write написано "verify: тест падает без реализации", симулякр невозможен по определению. Добавить в скиллы write-шагов.

### 2. "State your assumptions explicitly"

При неоднозначности в спеке — формулировать допущение явно и поднять через `/ask`, не выбирать молча. Добавить в `code-write.md`: "если спека не покрывает граничный случай — не реализовывать молча, вызвать `/ask`".

### 3. Формула сильных критериев приёмки

Слабо: "система корректно обрабатывает unix socket"  
Сильно: "TestConnectUnixSocket запускается, отправляет пакет, получает ответ без ошибок"

Добавить в `acceptance-write.md`: каждый критерий — конкретный проверяемый assert, не намерение.

### 4. "Every changed line traces to the user's request" (Surgical Changes)

Перед записью файлов — пройтись по diff: каждое изменение должно трассироваться к `acceptance.md`. Если не трассируется — не добавлять. Закрывает drift: добавление `MaxConnections` без флага было бы явным нарушением.

### 5. "No error handling for impossible scenarios"

Не реализовывать сценарии которых нет в спеке или acceptance. Добавить в `docs/standards/go/principles/`: "если сценарий не описан — поднять вопрос, не реализовывать спекулятивно".

### 6. "Don't improve adjacent code"

Изменения только в файлах которые описаны в спеке или явно вытекают из неё. Добавить в `code-write.md`.

### 7. Тест как первичный артефакт воспроизведения бага

"Fix the bug" → "Write a test that reproduces it, then make it pass". Добавить в code-fix: фикс начинается с фэйлящего теста который воспроизводит баг.

## Шаги

- [ ] Обновить `.claude/commands/code-write.md` — assumptions, surgical changes, no adjacent code, no speculative code
- [ ] Обновить `.claude/commands/code-fix.md` — тест как первичный артефакт воспроизведения
- [ ] Обновить `.claude/commands/acceptance-write.md` — формула сильных критериев
- [ ] Написать `docs/standards/go/principles/implementation.md` — surgical changes + no speculative code
- [ ] Добавить `verify:` поля в write-скиллы (spec-write, test-write, code-write)

## Заметки

Документ написан для одиночного LLM-кодера, не для агентного пайплайна. Всё про координацию между агентами — наша собственная разработка и богаче. Адоптируем только механики которые усиливают существующий процесс.

## Источник

Бо (анализ внешнего CLAUDE.md, ретро store/connection, 2026-04-26)
