# Action: стандарт `// inv:` для инвариантов в коде

## Проблема

CR-7 был косметическим замечанием: "cancel() дублируется в каждой ветке select — можно вынести". Ада вынесла `cancel()` за select. CR-9: deadlock — `cancel()` вызывается после блокирующего `<-done` при живых горутинах. Инвариант "cancel() должен выполняться до любого блокирующего ожидания в ветке таймаута" нигде не был зафиксирован. Ни ревьювер, ни реализатор не увидели что косметика меняет семантику.

## Решение

Ввести стандарт `// inv:` комментариев для блоков с lifecycle, concurrency, shutdown, signal handling, idempotency, error classification.

Правила:
1. Автор при первом написании критичного блока обязан расставить `// inv:` с формулировкой инварианта
2. В шаблоне code-review для блоков с `// inv:` — обязательное поле "Как фикс сохраняет инвариант". Без этого поля замечание не выпускается
3. В lifecycle/concurrency блоках нет категории "minor по стилю" — любое замечание про структуру или порядок операций → автоматически critical

## Пример

```go
// inv: cancel() must run before any blocking <-done in shutdown branches
// to ensure context propagation reaches active goroutines
select {
case <-done:
    cancel()
    return nil
case <-time.After(timeout):
    cancel() // must precede <-activeConns
    <-activeConns
    return ErrShutdownTimeout
}
```

## Шаги

- [ ] Написать `docs/standards/go/invariants.md` с правилом и примерами
- [ ] Добавить поле "Инвариант / Как фикс его сохраняет" в шаблон code-review для критичных блоков
- [ ] Обновить `docs/skills-guide/` или чек-лист Гримма

## Источник

Ада (ретро store/connection, 2026-04-26)
