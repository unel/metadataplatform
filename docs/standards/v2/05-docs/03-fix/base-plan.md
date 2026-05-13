---
version: 1.0.0
updated: 2026-04-26T16:45
---

# Базовый план: 05-docs-fix

## 1. Сбор контекста

- Прочитай `05-docs/02-review/report-*.md` — замечания с классификацией
- Прочитай затронутые файлы документации
- Прочитай код если нужно проверить точность исправления

## 2. Согласование плана

- Покажи план фикса пользователю
- Жди подтверждения

## 3. Исправление

- Исправляй только то что указано в review
- Сначала critical, потом warning, потом Nit:

## 4. Формирование отчёта

В frontmatter `report-NNN.md` обязательно укажи:

```yaml
in-response-to: tasks/<feature>/stages/05-docs/02-review/report-NNN.md
```

Содержимое отчёта:
- Список исправленных замечаний
- Список неисправленных с объяснением

## Скрипты

```bash
# Начало: написать brief-NNN.md, затем
python3 scripts/set-step-status.py <feature> <stage/step> in-progress --comment "..."
# Конец: написать report-NNN.md, затем
python3 scripts/set-step-status.py <feature> <stage/step> done --comment "..."
```

## Финальный шаг: notes/complaints

```bash
python3 scripts/log-note.py --agent <агент> --message "[propose] текст"
python3 scripts/log-complaint.py --agent <агент> --message "текст"
```

Теги: `[doc]`, `[propose]`, `[friction]`, `[miss]`, `[rework]`, `[whatever]`
Минимум — одна строка.

## Движение вперёд и назад

**Done когда:** исправимые замечания закрыты, report написан
**Следующий:** 05-docs/02-review (повторное ревью)
