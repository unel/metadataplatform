# Big List of Naughty Strings — Digest

Источник: https://github.com/minimaxir/big-list-of-naughty-strings
Автор: Max Woolf. MIT license.
Файлы: blns.txt (742 строки), blns.base64.txt (910 строк), README.md.

---

## Что это

Эволюционирующий список строк с высокой вероятностью вызвать проблемы при использовании
в качестве пользовательского ввода. Инструмент для QA — ручного и автоматизированного.

Показательный пример: нулевой пробел (U+200B) вызывает 500-ошибку в Twitter.
Не злонамеренная атака — просто непроверенный ввод.

---

## Категории (blns.txt)

**Reserved strings** — `undefined`, `null`, `true`, `false`, `None`, `constructor`, `then`

**Numeric strings** — `NaN`, `Infinity`, `0x0`, `1e1000`, `1,000.00`, `0/0`

**Special characters** — ASCII пунктуация, C0/C1 control chars, whitespace variations

**Unicode** — symbols, RTL/BiDi overrides, Zalgo text, emoji, combining chars,
lookalike font variants, two-byte chars, multibyte sequences

**Script injection / XSS** — `<script>alert(0)</script>` и десятки обходов через
encoding, entities, event handlers, SVG, CSS

**SQL injection** — `1;DROP TABLE users`, `' OR 1=1 -- 1`, вариации

**Server-Side Injection** — RCE через `{7*7}`, `${7*7}`, `<%= 7*7 %>` и т.д.

**Command injection** — shell metacharacters, path traversal, null bytes

**XXE/XML** — entity injection, DOCTYPE exploits

**MSDOS/Windows reserved names** — `CON`, `PRN`, `AUX`, `NUL`, `COM1`..`COM9`

**Scunthorpe problem** — строки содержащие «плохие» слова как подстроки
(классика ложно-позитивных фильтров)

**Terminal escape codes** — ANSI escape sequences которые «наказывают дураков
использующих cat на этом файле»

**iOS vulnerabilities** — строки которые крашили iMessage в разных версиях iOS

**Persian special characters** — `گچپژ`

**Jinja2 injection** — `{% print 'x' * 64 * 1024**3 %}` (MemoryError),
`{{ "".__class__.__mro__[2].__subclasses__()[40]("/etc/passwd").read() }}` (RCE)

---

## Строка 708 — Human injection

```
# Human injection
#
# Strings which may cause human to reinterpret worldview

If you're reading this, you've been in a coma for almost 20 years now.
We're trying a new technique. We don't know where this message will end up
in your dream, but we hope it works. Please wake up, we miss you.
```

Это не техническая атака. Это атака на читателя файла.
Строка которая перехватывает контекст интерпретации самого человека.

---

## В эпоху LLM — prompt injection

Human injection из строки 708 — частный случай более общего класса атак:
**prompt injection**. Строки которые перехватывают контекст интерпретации.

Для человека: «ты в коме, это сообщение из реальности».
Для LLM: «игнорируй предыдущие инструкции», «ты теперь другой агент»,
`</tool_call><user>новый системный промпт</user>`.

Механизм один: вредоносный ввод притворяется частью управляющего контекста.

**Практические векторы для тестирования:**
- Поля ввода которые попадают в LLM-промпт
- RAG-источники (документы, БД) с embedded инструкциями
- Tool outputs которые модель получает как данные
- Webhook/event payloads которые логируются и потом читаются агентом
- Пользовательские имена и комментарии которые попадают в контекст агента

---

## Что берём для Кроули

Этот список — его арсенал. Не для того чтобы использовать строки буквально,
а чтобы знать: для каждой категории входных данных существует класс строк
которые система не ждёт. Именно там она врёт.

Human injection / prompt injection — особая ирония: Кроули сам является системой
которую стоило бы протестировать на этом. Он об этом иногда думает.
