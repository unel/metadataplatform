# Уровни тестирования

**Почему:** E2E тест медленный, хрупкий, и когда падает — непонятно где именно сломалось. Юнит тест быстрый, точный, говорит конкретно что сломалось. Пирамида — это баланс: юниты дают скорость и точность, E2E дают уверенность что система работает как целое.

## Юниты

Тестируют одну функцию/модуль изолированно. Внешние зависимости заменяются заглушками только когда это IO (БД, сеть, файловая система). Бизнес-логика тестируется без моков.

```
test("filterEntities: возвращает только совпадающие по типу"):
  entities = [
    {id: "1", type: "file", subtype: "video"},
    {id: "2", type: "file", subtype: "audio"},
    {id: "3", type: "job",  subtype: "hash"},
  ]
  result = filterEntities(entities, {type: "file"})
  assert result == [{id: "1", ...}, {id: "2", ...}]
  // нет БД, нет сети — только чистая функция
```

## Интеграционные

Тестируют взаимодействие с реальными внешними системами. В этом проекте — store + PostgreSQL через testcontainers. Мок БД здесь недопустим: именно на этом уровне проверяются SQL-запросы, индексы, транзакции.

```
test("store: upsert entity → find entity"):
  // реальный PostgreSQL в контейнере
  db = startTestDB()
  store = Store(db)

  store.upsert({type: "file", meta: {path: "/video.mp4"}})
  result = store.find({type: "file"})

  assert len(result) == 1
  assert result[0].meta.path == "/video.mp4"
```

## E2E

Тестируют критический путь глазами пользователя. Браузер + реальный API + реальная БД. Медленные, хрупкие к UI-изменениям — поэтому только для самых важных сценариев.

```
// Playwright
test("пользователь может найти файл через поисковую строку"):
  page.goto("/")
  page.fill("[data-testid=search]", "video.mp4")
  page.click("[data-testid=search-button]")
  expect(page.locator(".entity-card")).toHaveCount(1)
  expect(page.locator(".entity-card")).toContainText("video.mp4")
```

## Признак перевёрнутой пирамиды

- Юнит тестов мало, E2E тестов много
- CI занимает 20+ минут
- Тесты часто падают по случайным причинам (flaky)
- Сложно понять что именно сломалось по упавшему тесту
