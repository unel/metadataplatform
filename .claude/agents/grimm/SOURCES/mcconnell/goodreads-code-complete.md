# Goodreads - Code Complete quotes

URL: https://www.goodreads.com/work/quotes/8419-code-complete-a-practical-handbook-of-software-construction
Fetched: 2026-04-22
Format: Extracted Goodreads quoteText blocks from page HTML

## Content

```text
---
"The big optimizations come from refining the high-level design, not the individual routines."

  ---

    Steve McConnell,

      Code Complete: A Practical Handbook of Software Construction

---
"Good code is its own best documentation. As you're about to add a comment, ask yourself, 'How can I improve the code so that this comment isn't needed?' Improve the code and then document it to make it even clearer."

  ---

    Steve McConnell,

      Code Complete: A Practical Handbook of Software Construction

---
"the road to programming hell is paved with global variables,"

  ---

    Steve McConnell,

      Code Complete

---
"Programmers working with high-level languages achieve better productivity and quality than those working with lower-level languages. Languages such as C++, Java, Smalltalk, and Visual Basic have been credited with improving productivity, reliability, simplicity, and comprehensibility by factors of 5 to 15 over low-level languages such as assembly and C (Brooks 1987, Jones 1998, Boehm 2000). You save time when you don't need to have an awards ceremony every time a C statement does what it's supposed to."

  ---

    Steve McConnell,

      Code Complete: A Practical Handbook of Software Construction

---
"Trying to improve software quality by increasing the amount of testing is like trying to lose weight by weighing yourself more often. What you eat before you step onto the scale determines how much you will weigh, and the software-development techniques you use determine how many errors testing will find."

  ---

    Steve McConnell,

      Code Complete: A Practical Handbook of Software Construction

---
"once gotos are introduced, they spread through the code like termites through a rotting house."

  ---

    Steve McConnell,

      Code Complete

---
"Heuristic is an algorithm in a clown suit. It’s less predictable, it’s more fun, and it comes without a 30-day, money-back guarantee."

  ---

    Steve McConnell,

      Code Complete

---
"Managers of programming projects aren’t always aware that certain programming
issues are matters of religion. If you’re a manager and you try to require compliance
with certain programming practices, you’re inviting your programmers’ ire. Here’s a
list of religious issues:
■ Programming language
■ Indentation style
■ Placing of braces
■ Choice of IDE
■ Commenting style
■ Efficiency vs. readability tradeoffs
■ Choice of methodology—for example, Scrum vs. Extreme Programming vs. evolutionary
delivery
■ Programming utilities
■ Naming conventions
■ Use of gotos
■ Use of global variables
■ Measurements, especially productivity measures such as lines of code per day"

  ---

    Steve McConnell,

      Code Complete: A Practical Handbook of Software Construction

---
"Copy and paste is a design error"

  ---

    Steve McConnell,

      Code Complete

---
"usually more time is spent in making good-looking presentation slides than in improving the quality of the software."

  ---

    Steve McConnell,

      Code Complete

---
"complicated code is a sign that you don't understand your program well enough to make it simple."

  ---

    Steve McConnell,

      Code Complete

---
"One of the paradoxes of defensive programming is that during development, you'd like an error to be noticeable—you'd rather have it be obnoxious than risk overlooking it. But during production, you'd rather have the error be as unobtrusive as possible, to have the program recover or fail gracefully."

  ---

    Steve McConnell,

      Code Complete

---
"Reduce complexity. The single most important reason to create a routine is to reduce a program's complexity. Create a routine to hide information so that you won't need to think about it."

  ---

    Steve McConnell,

      Code Complete

---
"Inheritance adds complexity to a program, and, as such, it's a dangerous technique. As Java guru Joshua Bloch says, "Design and document for inheritance, or prohibit it." If a class isn't designed to be inherited from, make its members non-virtual in C++, final in Java, or non-overridable in Microsoft Visual Basic so that you can't inherit from it."

  ---

    Steve McConnell,

      Code Complete

---
"Try to create modules that depend little on other modules. Make them detached, as business associates are, rather than attached, as Siamese twins are."

  ---

    Steve McConnell,

      Code Complete

---
"The goal is to minimize the amount of a program you have to think about at any one time. You might think of this as mental juggling—the more mental balls the program requires you to keep in the air at once, the more likely you'll drop one of the balls, leading to a design or coding error."

  ---

    Steve McConnell,

      Code Complete

---
"Managing complexity is the most important technical topic in software development. In my view, it's so important that Software's Primary Technical Imperative has to be managing complexity. Complexity is not a new feature of software development."

  ---

    Steve McConnell,

      Code Complete

---
"Don't differentiate routine names solely by number. One developer wrote all his code in one big function. Then he took every 15 lines and created functions named Part1, Part2, and so on. After that, he created one high-level function that called each part."

  ---

    Steve McConnell,

      Code Complete

---
"People who are effective at developing high-quality software have spent years accumulating dozens of techniques, tricks, and magic incantations. The techniques are not rules; they are analytical tools."

  ---

    Steve McConnell,

      Code Complete

---
"Spend your time on the 20 percent of the refactorings that provide 80 percent of the benefit."

  ---

    Steve McConnell,

      Code Complete

---
"developers insert an average of 1 to 3 defects per hour into their designs and 5 to 8 defects per hour into code"

  ---

    Steve McConnell,

      Code Complete

---
"few people can understand more than three levels of nested ifs"

  ---

    Steve McConnell,

      Code Complete

---
"You can do anything with stacks and iteration that you can do with recursion."

  ---

    Steve McConnell,

      Code Complete

---
"Building software implies various stages of planning, preparation and execution that vary in kind and degree depending on what's being built. [...]
Building a four-foot tower requires a steady hand, a level surface, and 10 undamaged beer cans. Building a tower 100 times that size doesn't merely require 100 times as many beer cans."

  ---

    Steve McConnell,

      Code Complete: A Practical Handbook of Software Construction

---
"Because successful programming depends on minimizing complexity, a skilled programmer will build in as much flexibility as needed to meet the software's requirements but will not add flexibility—and related complexity—beyond what's required."

  ---

    Steve McConnell,

      Code Complete

---
"people who believed the old theory thought the new theory was just as ridiculous then as you think the old theory is now."

  ---

    Steve McConnell,

      Code Complete

---
"To experiment effectively, you must be willing to change your beliefs based on the results of the experiment"

  ---

    Steve McConnell,

      Code Complete

---
"It's tempting to trivialize the power of metaphors. To each of the earlier examples, the natural response is to say, "Well, of course the right metaphor is more useful. The other metaphor was wrong!" Though that's a natural reaction, it's simplistic. The history of science isn't a series of switches from the "wrong" metaphor to the "right" one. It's a series of changes from "worse" metaphors to "better" ones, from less inclusive to more inclusive, from suggestive in one area to suggestive in another."

  ---

    Steve McConnell,

      Code Complete

---
"Object-Oriented Design Heuristics (1996), Arthur Riel"

  ---

    Steve McConnell,

      Code Complete

---
"From time to time, a complex algorithm will lead to a longer routine, and in those circumstances, the routine should be allowed to grow organically up to 100–200 lines. (A line is a noncomment, nonblank line of source code.) Decades of evidence say that routines of such length are no more error prone than shorter routines. Let issues such as the routine's cohesion, depth of nesting, number of variables, number of decision points, number of comments needed to explain the routine, and other complexity-related considerations dictate the length of the routine rather than imposing a length restriction per se."

  ---

    Steve McConnell,

      Code Complete


```
