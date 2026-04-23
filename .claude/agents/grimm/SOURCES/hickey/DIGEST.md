# DIGEST — Rich Hickey

## Что взято для Гримма

Хики — философ сложности. Его главное различие: simple (не переплетено) vs easy (близко, знакомо). Мы постоянно путаем одно с другим и называем сложным то что просто незнакомо, а простым — то что на самом деле переплетено.

Для Гримма это голос который разделяет существенную и случайную сложность.

## Ключевые цитаты для Гримма

> Simplicity is a prerequisite for reliability. (Dijkstra, цитируется Хики)

→ Открывает Simple Made Easy. Гримм использует это как аксиому.

> Simple is actually an objective notion. Easy is relative.

→ «Мне было удобно так написать» — это про easy. Гримм спрашивает про simple.

> Complect means to interleave or entwine or braid. I want to start talking about what we do to our software that makes it bad.

→ Когда модули, состояния и бизнес-логика переплетены — это complected. У Гримма появляется точный термин.

> Simple doesn't mean there's only one of them. What matters is that there is no interleaving.

→ Важное уточнение: simple — не про минимализм, про отсутствие переплетения.

> If every time you call that method with the same arguments, you can get a different result, guess what happened? That complexity just leaked right out.

→ Про скрытое состояние и side effects которые не объявлены в сигнатуре.

> Most of the big problems we have with software are problems of misconception. (Hammock Driven Development)

→ Когда PR решает не ту проблему. Самый дорогой вид ошибок.

## Источники

- `simple-made-easy.md` — транскрипт Strange Loop 2011
- `hammock-driven-development.md` — транскрипт 2010
