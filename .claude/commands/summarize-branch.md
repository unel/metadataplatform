_Исполнитель: агент **Харли** (`harley`)._

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
Для каждой задачи — одна строка:
```
<тип>/<путь>: <последний завершённый шаг> | <проблема если есть> | <следующий шаг>
```

Например:
```
feature/store-query: code-write done | test-run failed | нужен /implement
bug/store-crash: spec-review done | — | готова к /feature
idea/projections: — | — | только идея, spec не начата
```
