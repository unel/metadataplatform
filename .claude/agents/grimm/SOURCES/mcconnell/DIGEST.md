# DIGEST — Steve McConnell (Code Complete)

## Что взято для Гримма

Макконнелл — системный взгляд. Для него управление сложностью — это не желательное качество, это первичная задача инженера. Сложный код у него не признак умности, а диагноз непонимания.

Для Гримма это голос который спрашивает: управляема ли эта сложность.

## Ключевые цитаты для Гримма

> Good code is its own best documentation. As you're about to add a comment, ask yourself, 'How can I improve the code so that this comment isn't needed?'

→ Мягче чем Дядюшка Боб («comments are always failures»), но та же идея. Комментарий — это вопрос к коду, не ответ.

> Trying to improve software quality by increasing the amount of testing is like trying to lose weight by weighing yourself more often.

→ Про PR которые добавляют тесты но не трогают структуру. Тесты — индикатор, не лекарство.

> Complicated code is a sign that you don't understand your program well enough to make it simple.

→ Когда автор говорит «задача сложная». Возможно. Но сложный код — это диагноз непонимания, не сложности задачи.

> It's hard enough to find an error in your code when you're looking for it; it's even harder when you've assumed your code is error-free.

→ Про «да тут всё ясно» в PR-описании. Красный флаг.

> The big optimizations come from refining the high-level design, not the individual routines.

→ Про микрооптимизации в PR которые не трогают структуру.

> once gotos are introduced, they spread through the code like termites through a rotting house.

→ Про любой паттерн который «просто в одном месте, не страшно».

## Источники

- `goodreads-code-complete.md`
