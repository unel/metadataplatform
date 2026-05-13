---
version: 1.0.0
updated: 2026-04-26T12:00
---

# Базовый план: 02-acceptance-fix

## 1. Сбор контекста

- Прочитай актуальную spec (последний report из `01-spec/`)
- Прочитай актуальный acceptance: последний `02-acceptance/03-fix/report-*.md`, если нет — `02-acceptance/01-write/report-*.md`
- Прочитай замечания: последний `02-acceptance/02-review/report-*.md`

## 2. Составление плана исправлений

Для каждой проблемы из ревью:
- Определи тип: пробел в acceptance, качество сценария
- Определи что именно изменить или добавить
- Проблемы типа "неопределённость/противоречие в spec" — не исправляй в acceptance, эскалируй

## 3. Согласование плана

- Покажи план исправлений пользователю
- Жди подтверждения или правок

## 4. Внесение исправлений

- Исправляй только то что указано в замечаниях
- Не переписывай acceptance целиком

## 5. Формирование отчёта

В frontmatter `report-NNN.md` обязательно укажи:

```yaml
in-response-to: tasks/<feature>/stages/02-acceptance/02-review/report-NNN.md
```

Содержимое отчёта:
- Полный обновлённый текст acceptance
- Changelog: по каждой проблеме — что изменено или почему пропущено

## 6. Обновление статусов

- `02-acceptance/02-review` → `pending`
- `03-tests`, `04-code`, `05-docs` → `stale` (если уже выполнены)

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

**Done когда:** все замечания типа "пробел/качество" закрыты, report написан
**Следующий:** 02-acceptance/02-review (повторный проход)
**Передаём:** report-NNN.md (содержит обновлённый acceptance + changelog)

**Возврат назад:**
- замечание типа "неопределённость/противоречие в spec" → `clarification`, возврат к 01-spec/03-fix (через оркестратора)
