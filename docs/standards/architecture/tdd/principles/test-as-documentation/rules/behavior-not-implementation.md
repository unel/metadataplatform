# Поведение, не реализация

**Почему:** тест привязанный к реализации падает при любом рефакторинге — даже когда поведение не изменилось. Такой тест создаёт ложную работу и снижает доверие к suite. Тест привязанный к поведению — живёт долго и защищает от регрессий.

## Границы применимости

Все тесты публичных интерфейсов. Тестировать приватные методы напрямую — антипаттерн: это детали реализации.

## Если соблюдать рьяно

Иногда для сложного алгоритма полезно протестировать промежуточные шаги. Если промежуточный шаг выделен в отдельную функцию с понятной ответственностью — тестировать её допустимо.

## Если игнорировать

Тесты падают при рефакторинге который не ломает поведение. Команда боится рефакторить — тесты мешают вместо того чтобы помогать.

## Когда можно отступить

Performance тесты проверяют не только поведение но и характеристики реализации (время, память). Это другой вид тестов с другими правилами.

## Плохо — тест знает про реализацию

```
test("processJob вызывает validate затем save"):
  job = makeJob()
  spy_validate = spy(job, "validate")
  spy_save = spy(job, "save")

  processJob(job)

  assert spy_validate.calledOnce     // проверяем что вызвали validate
  assert spy_save.calledAfter(spy_validate)  // проверяем порядок вызовов
  // тест упадёт при любом рефакторинге внутренностей processJob
```

## Хорошо — тест описывает поведение

```
test("processJob: валидный job переходит в статус done"):
  job = {id: "1", status: "pending", payload: {worker: "hash"}}
  processJob(job)
  assert job.status == "done"

test("processJob: невалидный job переходит в статус failed с сообщением"):
  job = {id: "2", status: "pending", payload: {}}  // нет worker
  processJob(job)
  assert job.status == "failed"
  assert job.error contains "worker required"

// рефакторинг processJob не сломает эти тесты — они проверяют результат
```

## Структура теста: Arrange-Act-Assert

```
test("queryEntities: фильтр по типу возвращает только совпадающие"):
  // Arrange — подготовка
  store.create({type: "file", subtype: "video"})
  store.create({type: "file", subtype: "audio"})
  store.create({type: "job"})

  // Act — действие
  result = store.query({type: "file"})

  // Assert — проверка
  assert len(result) == 2
  assert all(e.type == "file" for e in result)
```
