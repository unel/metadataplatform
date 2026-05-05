# Action: backlog механизм — 04-solve создаёт, 00-research проверяет

## Проблема

Из 8 action items RETRO store/connection → 0 реализовано до старта crud. Action items записывались в RETRO.md и в отдельные файлы `RETRO/actions/`, но не конвертировались в задачи которые проверяются в начале следующей фичи. Разрыв между "описать задачу" и "выполнить задачу" не закрыт механически.

## Решение

**В конце 04-solve:** обязательный шаг — создать/обновить `tasks/_backlog/retro-<feature>.md`:
- чек-лист всех action items с DRI и приоритетом
- ссылки на `RETRO/actions/<file>.md` для деталей

**В начале 00-research следующей фичи:** обязательный шаг — прочитать `tasks/_backlog/` и задокументировать статус каждого пункта:
- выполнен → отметить `[x]` с датой
- не выполнен → перенести в backlog текущей фичи или явно закрыть с причиной

## Шаги

- [ ] Добавить шаг "создать backlog-файл" в `docs/standards/v2/06-retro/04-solve/base-plan.md`
- [ ] Добавить шаг "review backlog предыдущей фичи" в `docs/standards/v2/00-research/01-interview/base-plan.md`
- [ ] Формат backlog-файла: `tasks/_backlog/retro-<feature>.md` (см. `tasks/_backlog/retro-store-crud.md`)

## Источники

Танк, Гримм, Кроули, Харли (ретро store/crud, 2026-05-05)
Рецидив из: RETRO store/connection
