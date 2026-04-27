_Режим: **Харли Куин**. Выполняется оператором напрямую._

Собери сводку по статусам задач в ветке.

Аргументы: `$ARGUMENTS` — имя ветки, например `feature/store-query`

## Шаги

### 1. Найди задачи
Получи список файлов `tasks/*/status.md` из ветки:
```
git ls-tree -r $ARGUMENTS --name-only | grep "^tasks/.*/status.md$"
```

### 2. Прочитай статусы
Для каждого найденного файла:
```
git show $ARGUMENTS:<path>
```

### 3. Верни сводку
Для каждой задачи — одна строка. Формат:
```
<тип>/<путь>: <последний done-шаг> → [<все оставшиеся шаги в порядке цепочки>]
```

Правила:
- В `[...]` перечисли **все** шаги, которые не `done`, в порядке цепочки — включая `needs-recheck` и `pending`
- `needs-recheck` = нужно повторить, отображай как `шаг(recheck)`
- `failed` = нужно переделать, отображай как `шаг(failed)`
- Если все шаги `done` — напиши `✓ завершена`
- Если `status.md` пуст или не существует — напиши `только идея`

Например:
```
feature/store-query: code-write → [test-run(failed), implement]
feature/store-connection: spec-nft → [spec-review, spec(recheck), acceptance-write(recheck), acceptance-review, implement]
bug/store-crash: spec-review → [feature]
idea/projections: — → только идея
feature/done-task: ✓ завершена
```
