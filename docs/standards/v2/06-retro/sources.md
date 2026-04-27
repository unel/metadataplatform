---
purpose: Источники по практикам ретроспектив
updated: 2026-04-27
---

# Источники: 06-retro

## Foundational

**Derby & Larsen — Agile Retrospectives: Making Good Teams Great** (Pragmatic Programmers, 2nd ed.)
- Пять фаз: Set the Stage → Gather Data → Generate Insights → Decide What to Do → Close
- Каждая фаза начинается с 3–5 мин молчаливого письма прежде чем обсуждать — равный голос для всех
- https://pragprog.com/titles/dlret2/agile-retrospectives-second-edition/

**Norm Kerth — Prime Directive**
- "Regardless of what we discover, we understand and truly believe that everyone did the best job they could, given what they knew at the time, their skills and abilities, the resources available, and the situation at hand."
- Читается вслух в начале каждой сессии — ритуал, не информация
- https://retrospectivewiki.org/index.php?title=The_Prime_Directive

## Форматы и техники

**Retrospective Wiki — форматы**
- Start/Stop/Continue, 4Ls, Starfish, Mad/Sad/Glad, Sailboat, Timeline Retro
- Выбор формата по: зрелость команды, цель сессии, эмоциональный климат
- https://retrospectivewiki.org

**5 Whys в ретро**
- Применять в фазе Generate Insights, не в Gather Data
- Останавливаться на системной причине — не доходить до личной вины
- Без фасилитации превращается в петлю обвинений

## Action items

**Scrum.org — Sprint Retrospective**
- Sprint Retrospective может добавлять items в Sprint Backlog — если item не в backlog, он не существует
- https://www.scrum.org/resources/blog/sprint-retrospective-dysfunctions-and-how-overcome-them
- https://www.scrum.org/resources/blog/ditch-unfinished-action-items

**EasyAgile / Echometer — практики action items**
- Max 3 items за ретро: completion rate при 1–3 кратно выше чем при 10+
- DRI (Directly Responsible Individual): "команда сделает X" не работает — нужен конкретный владелец
- SMART в контексте ретро: Specific + Measurable + Achievable + Relevant + Time-bound
- "15% Solutions": находи что команда может сделать прямо сейчас, пусть частично
- https://www.easyagile.com/blog/improve-sprint-retrospective-action-items
- https://echometerapp.com/en/retrospective-action-items-tips-examples/

**Age of Product — Unfinished Action Items**
- "Action Item Backlog" (список переносимых невыполненных items) — антипаттерн
- Устаревшие items: не переносить автоматически; переоценить или явно списать
- https://age-of-product.com/action-items-retrospectives/

## Антипаттерны

**Retrium — Retrospective Anti-Patterns**
- Ретро без цели, ретро без outcome, 25 action items, вентиляция вместо решений, доминирующий голос
- https://www.retrium.com/ultimate-guide-to-agile-retrospectives/retrospective-anti-patterns

**TeamRetro — Anti-Patterns 2026**
- Менеджеры на ретро разрушают психологическую безопасность — документированный антипаттерн
- https://www.teamretro.com/avoid-these-retrospective-anti-patterns-in-2025

## Удалённые / асинхронные ретро

**Retrium — Remote Retrospectives**
- "Если один участник удалённый — все удалённые": гибрид создаёт коммуникационное неравенство
- Async-first: наблюдения вносятся за 24–48 ч до синхронной сессии; синхрон — только для обсуждения инсайтов
- https://www.retrium.com/ultimate-guide-to-agile-retrospectives/retrospectives-with-remote-and-distributed-teams

## Применимость к агентному контексту

Агентное ретро — естественно async-first: агенты не имеют состояния между сессиями.

| Derby & Larsen фаза | Наш процесс |
|---|---|
| Set the Stage | `01-recall` (Prime Directive, review предыдущего RETRO.md) |
| Gather Data | `01-recall` (notes + complaints — молчаливые "стикеры") |
| Generate Insights | `02-collect` + `03-analyze` (выжимки → группировка → 5 Whys) |
| Decide What to Do | `04-solve` (решения + DRI + дедлайн, ≤3 items) |
| Close | `05-write` (RETRO.md) |
