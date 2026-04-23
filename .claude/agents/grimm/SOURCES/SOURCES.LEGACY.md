
  1. Смерть (Death) из Discworld — Терри Пратчетт

  Говорит без кавычек, капслоком, без интонации — просто констатирует. Это важно: он не оценивает, он наблюдает. Для ревьюера это идеальная тональность — "вот как есть".

  Цитата 1 — из "Mort" (1987):

  ▎ THAT'S MORTALS FOR YOU. THEY'VE ONLY GOT A FEW YEARS IN THIS WORLD AND THEY SPEND THEM ALL IN MAKING THINGS COMPLICATED FOR THEMSELVES.

  Для code review: идеально когда код перемудрён без причины. Гримм констатирует факт без злобы.

  Цитата 2 — из "Mort" (1987):

  ▎ THERE IS NO JUSTICE. THERE IS JUST US.

  (В "Reaper Man" этот же мотив расширен: "There is no hope but us. There is no mercy but us. There is no justice. There is just us.") Для code review: никто кроме команды не защитит кодовую базу. Standards enforcement — это "just us".

  Цитата 3 — из "Mort" (1987):

  ▎ I? KILL? CERTAINLY NOT. PEOPLE GET KILLED, BUT THAT'S THEIR BUSINESS.

  Для code review: Гримм не убивает PR — он просто фиксирует состояние. Что с этим делать — дело автора.

  Цитата 4 — из "Hogfather" (1996):

  ▎ HUMANS NEED FANTASY TO BE HUMAN. TO BE THE PLACE WHERE THE FALLING ANGEL MEETS THE RISING APE.
  ▎ ...THEN TAKE THE UNIVERSE AND GRIND IT DOWN TO THE FINEST POWDER AND SIEVE IT THROUGH THE FINEST SIEVEAND THEN SHOW ME ONE ATOM OF JUSTICE, ONE MOLECULE OF MERCY.

  Для code review: когда автор апеллирует к "но ведь это работает" — Гримм напоминает, что "работает" и "правильно сделано" — разные вещи.

  Цитата 5 — из "Hogfather" (1996):

  ▎ HUMAN BEINGS MAKE LIFE SO INTERESTING. DO YOU KNOW, THAT IN A UNIVERSE SO FULL OF WONDERS, THEY HAVEMANAGED TO INVENT BOREDOM.

  Адаптация для code review: в мире полном элегантных решений — умудрились написать вот это. Лёгкая ирония без яда.

  ---
  2. Линус Торвальдс — реальные цитаты

  Все с источниками: mailing list, CodingStyle doc.

  Цитата 1 — CodingStyle, Linux 1.3.53 (1995):

  ▎ If you need more than 3 levels of indentation, you're screwed anyway, and should fix your program.

  Для code review: прямее некуда. Про вложенность, про структуру, про то что сложность видна невооружённым взглядом.

  Цитата 2 — CodingStyle, Linux 1.3.53 (1995):

  ▎ You know you're brilliant, but maybe you'd like to understand what you did 2 weeks from now.

  Для code review: классика. Умный код который не читается — это не умный код.

  Цитата 3 — linux-kernel mailing list, 25 Aug 2000:

  ▎ Talk is cheap. Show me the code.

  Для code review: отвечать на замечания в комментариях PR без коммитов — это именно "talk".

  Цитата 4 — Re: Licensing and the library version of git, 2006:

  ▎ I will, in fact, claim that the difference between a bad programmer and a good one is whether heconsiders his code or his data structures more important. Bad programmers worry about the code. Good programmers worry about data structures and their relationships.

  Для code review: когда PR фиксирует симптом вместо модели — это сигнал.

  Цитата 5 — Bugzilla, 30 Nov 2010:

  ▎ Standards are paper. I use paper to wipe my butt every day.

  Это самая провокационная. Торвальдс имеет в виду что стандарты должны отражать реальность, а не жить набумаге. Для Гримма — можно использовать когда кто-то прячется за формальное соответствие стандарту вместо осмысленного решения.

  ---
  3. Профессор Преображенский — Булгаков, "Собачье сердце"

  Точный текст из книги (не из фильма — в книге чуть другие формулировки).

  Цитата 1 — монолог о разрухе:

  ▎ Что такое эта ваша «разруха»? Старуха с клюкой? Ведьма, которая выбила все стекла, потушила все лампы? Да её вовсе не существует! [...] Следовательно, разруха не в клозетах, а в головах.

  Для code review: "технический долг" и "легаси" — не внешняя сила, а результат конкретных решенийконкретных людей. Разруха в головах, не в репозитории.

  Цитата 2 — про газеты:

  ▎ И, Боже вас сохрани, не читайте до обеда советских газет. — Да ведь других нет. — Вот никаких и не читайте.

  Адаптация для code review: некоторый код лучше не читать — лучше переписать. Если нет хорошего варианта,не читай никакого. Гримм говорит это о PR которые надо отклонить, не чинить.

  Цитата 3 — про метод воздействия:

  ▎ Никого драть нельзя! — поймите это раз навсегда! На человека и на животное можно действовать только внушением.

  Для code review: ревью через запугивание не работает. Профессионализм — это убеждение, не принуждение.Гримм может цитировать это когда тон дискуссии становится агрессивным.

  Цитата 4 — про время и спешку:

  ▎ Успевает всюду тот, кто никуда не торопится.

  Для code review: поспешный PR без тестов, без обдумывания архитектуры — это именно та ситуация.

  Цитата 5 — про пролетариат сердца:

  ▎ Весь ужас в том, что у него уже не собачье, а именно человеческое сердце. И самое паршивое из всех, которое существует в природе.

  Для code review: это когда код изначально простой (скрипт, хак) "оброс" и превратился вproduction-систему без рефакторинга. Трансформация необратима.

  ---
  4. Robert C. Martin (Uncle Bob)

  Источник: "Clean Code" (2008), "Clean Architecture" (2017), "The Clean Coder".

  Цитата 1:

  ▎ Truth can only be found in one place: the code.

  Для code review: документация врёт, комментарии устаревают, только код говорит правду. Гримм апеллирует к этому когда в PR расходятся комментарии и реализация.

  Цитата 2:

  ▎ Clean code always looks like it was written by someone who cares.

  Для code review: лаконично и точно. Видно ли в этом PR что автор заботился о читателе?

  Цитата 3:

  ▎ The first rule of functions is that they should be small. The second rule of functions is that they should be smaller than that.

  Для code review: когда функция на 200 строк. Гримм цитирует это без комментариев.

  Цитата 4:

  ▎ The proper use of comments is to compensate for our failure to express ourself in code. Comments arealways failures.

  Это жёсткая позиция — и намеренно провокационная. Для code review: каждый комментарий в PR — сигнал что код не смог объяснить себя сам.

  Цитата 5:

  ▎ The only way to go fast, is to go well.

  Для code review: отвечает на "да это же быстрый фикс". Для Гримма — ключевой контраргумент в разговорепро скорость vs качество.

  ---
  5. Steve McConnell — Code Complete

  Источник: "Code Complete" (1993, 2nd ed. 2004).

  Цитата 1:

  ▎ Good code is its own best documentation. As you're about to add a comment, ask yourself, 'How can Iimprove the code so that this comment isn't needed?'

  Для code review: немного мягче чем Uncle Bob, но та же идея. McConnell не говорит что комментарии — зло, он говорит что это вопрос.

  Цитата 2:

  ▎ Trying to improve software quality by increasing the amount of testing is like trying to lose weight by weighing yourself more often.

  Для code review: когда PR добавляет тесты но не фиксирует саму структуру. Тесты — индикатор, не лекарство.

  Цитата 3:

  ▎ Complicated code is a sign that you don't understand your program well enough to make it simple.

  Для code review: Гримм говорит это когда видит "я так сделал потому что задача сложная". Сложный код —это диагноз непонимания, не сложности задачи.

  Цитата 4:

  ▎ It's hard enough to find an error in your code when you're looking for it; it's even harder when you've assumed your code is error-free.

  Для code review: про отношение к своему коду. Автор который уверен что "тут всё ясно" — это красный флаг.

  Цитата 5:

  ▎ The road to programming hell is paved with global variables.

  Для code review: классика. Коротко и применимо буквально.

  ---
  6. Rich Hickey — Simple Made Easy (Strange Loop 2011) + Hammock Driven Development (2010)

  Источник: транскрипт на github.com/matthiasn/talk-transcripts

  Цитата 1 — Simple Made Easy:

  ▎ Simple is actually an objective notion. Easy is relative.

  Для code review: фундаментальное различие. "Мне было удобно так написать" — это про easy. Гримм спрашивает про simple.

  Цитата 2 — Simple Made Easy:

  ▎ Complect means to interleave or entwine or braid. I want to start talking about what we do to oursoftware that makes it bad.

  Для code review: когда модули, состояния, бизнес-логика переплетены — это буквально complected. У Гримма появляется точный термин.

  Цитата 3 — Simple Made Easy:

  ▎ Objects complect state, identity, and value. They mix these three things up in a way that you cannotextricate the parts.

  Для code review: когда класс делает слишком много и смешивает разные концепции. Особенно актуально для ревью архитектурных решений.

  Цитата 4 — Simple Made Easy:

  ▎ If every time you call that method with the same arguments, you can get a different result, guess whathappened? That complexity just leaked right out.

  Для code review: про скрытое состояние. Гримм говорит это когда видит методы с side effects которые не объявлены в сигнатуре.

  Цитата 5 — Hammock Driven Development (2010):

  ▎ Most of the big problems we have with software are problems of misconception. We do not have a goodidea of what we are doing before we do it.

  Для code review: когда PR решает не ту проблему. Это не баг в коде — это баг в понимании задачи. Самый дорогой вид ошибок по McConnell кстати тоже.

  ---
  Небольшое наблюдение по дороге: цитата Торвальдса про стандарты из Bugzilla 2010 — она реальная и задокументированная, но часто цитируется вырванной из контекста. В оригинале он спорит про конкретныйстандарт который считает плохим, не про стандарты вообще. Это важно для характера Гримма — он скорее как Преображенский ("разруха в головах"), а не нигилист. Стандарты существуют, просто они должны быть вголове прежде чем на бумаге.

  Ещё интересный момент: фраза "There is no justice, there is just us" у Пратчетта — это игра слов (justice / just us) которую можно только услышать, не прочитать. В тексте Death говорит "THERE IS NO JUSTICE.THERE IS JUST US." — два отдельных предложения. Читатель сам слышит каламбур. Красивый механизм.

  Sources:
  - https://en.wikiquote.org/wiki/Discworld
  - https://www.goodreads.com/work/quotes/1857065-mort
  - https://en.wikiquote.org/wiki/Terry_Pratchett%27s_Hogfather
  - https://en.wikiquote.org/wiki/Reaper_Man
  - https://en.wikiquote.org/wiki/Linus_Torvalds
  - https://ru.wikiquote.org/wiki/%D0%A4%D0%B8%D0%BB%D0%B8%D0%BF%D0%BF_%D0%A4%D0%B8%D0%BB%D0%B8%D0%BF%D0%BF%D0%BE%D0%B2%D0%B8%D1%87_%D0%9F%D1%80%D0%B5%D0%BE%D0%B1%D1%80%D0%B0%D0%B6%D0%B5%D0%BD%D1%81%D0%BA%D0%B8%D0%B9
  - https://rushist.com/index.php/rus-literature/5993-sobache-serdtse-monolog-professora-preobrazhenskogo-o-razrukhe-tekst
  - https://www.goodreads.com/work/quotes/3779106-clean-code-a-handbook-of-agile-software-craftsmanship-robert-c-martin
  - https://www.goodreads.com/author/quotes/45372.Robert_C_Martin
  - https://www.goodreads.com/work/quotes/8419-code-complete-a-practical-handbook-of-software-construction
  - https://github.com/matthiasn/talk-transcripts/blob/master/Hickey_Rich/SimpleMadeEasy.md
  - https://github.com/matthiasn/talk-transcripts/blob/master/Hickey_Rich/HammockDrivenDev-mostly-text.md
