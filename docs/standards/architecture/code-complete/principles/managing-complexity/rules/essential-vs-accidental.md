# Essential vs Accidental Complexity

**Почему:** essential complexity нельзя устранить — она в природе задачи. Но accidental complexity можно и нужно. Разработчик, не различающий эти два вида, тратит силы не там: борется с проблемой вместо того чтобы упрощать своё решение.

## Границы применимости

Любое архитектурное и тактическое решение. Перед добавлением абстракции, уровня косвенности, нового паттерна — спроси: это essential или accidental?

## Если соблюдать рьяно

Стремление к нулевой accidental complexity иногда ведёт к over-engineering — сложная архитектура ради «правильности» сама становится источником accidental complexity.

## Если игнорировать

Кодовая база зарастает случайной сложностью: god objects, глубокие цепочки вызовов, абстракции ради абстракций. Onboarding новых разработчиков занимает недели вместо дней.

## Когда можно отступить

Некоторая accidental complexity оправдана performance-требованиями или ограничениями платформы. Важно осознавать что это выбор, а не случайность.

## Примеры accidental complexity

```
// лишний уровень косвенности без причины
class EntityServiceFacadeAdapter:
  EntityServiceFacade facade

  function find(id):
    return facade.findEntity(id)  // просто пробрасывает — зачем?

// запутанный поток управления вместо early return
function validate(entity):
  result = null
  if entity != null:
    if entity.type != "":
      if entity.type in allowedTypes:
        result = Ok
      else:
        result = Error("unknown type")
    else:
      result = Error("type required")
  else:
    result = Error("entity required")
  return result

// то же самое без accidental complexity
function validate(entity):
  if entity == null:    return Error("entity required")
  if entity.type == "": return Error("type required")
  if entity.type not in allowedTypes: return Error("unknown type")
  return Ok
```

## Метрики как сигналы

McConnell предлагает конкретные пороги — не как жёсткие правила, а как сигналы что стоит пересмотреть:

- **Цикломатическая сложность > 10** — функция содержит слишком много независимых путей
- **Вложенность > 4** — читатель теряет контекст к тому моменту как добирается до тела
- **Длина файла > ~400 строк** — вероятно модуль делает слишком много
