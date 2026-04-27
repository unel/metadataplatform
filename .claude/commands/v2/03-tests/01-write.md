_Исполнитель: **параллельно Кроули и Азирафаль**._

# test-write

Написание тестов по acceptance. Два агента работают параллельно.

## Инициализация

Скопируй `README.md` из `docs/standards/v2/03-tests/01-write/README.md` в `tasks/$ARGUMENTS/processes/03-tests/01-write/README.md` и добавь в метаданные: `feature: $ARGUMENTS`, `generated: {{datetime}}`, `source: docs/standards/v2/03-tests/01-write/README.md`.

## Исполнение

Запусти **параллельно** двух агентов:

### Азирафаль — happy path

Прочитай `docs/standards/v2/03-tests/01-write/base-plan.md` и `base-checklist.md`.

Найди финальный acceptance: последний `tasks/$ARGUMENTS/processes/02-acceptance/03-fix/report-*.md`, если нет — `tasks/$ARGUMENTS/processes/02-acceptance/01-write/report-*.md`.

Пиши happy path и contract тесты в `tasks/$ARGUMENTS/tests/happy/`.

Держи каждый тест-файл ≤ 150 строк (без комментариев).

### Кроули — adversarial

Прочитай `docs/standards/v2/03-tests/01-write/base-plan.md` и `base-checklist.md`.

Найди финальный acceptance (аналогично Азирафалю).

Пиши adversarial тесты: edge cases, failure modes, граничные условия — в `tasks/$ARGUMENTS/tests/adversarial/`.

Держи каждый тест-файл ≤ 150 строк (без комментариев).

## Отчёт

После завершения обоих агентов — запиши `tasks/$ARGUMENTS/processes/03-tests/01-write/report-NNN.md`:

```markdown
---
purpose: Написание тестов для $ARGUMENTS
process: 03-tests/01-write
run: {{N}}
date: {{datetime}}
created: {{datetime}}
see-also:
status: done | failed | clarification
agent: Азирафаль + Кроули
checklist: все пункты закрыты | открытые: {{список}}
---

## Happy path тесты (Азирафаль)

{{список файлов и тестов}}

## Adversarial тесты (Кроули)

{{список файлов и тестов}}

## Маппинг: тест → сценарий acceptance

{{таблица}}

## Непокрытые сценарии

{{если есть — перечислить с объяснением}}

## Результаты первого прогона

{{pass/fail}}
```

Обнови `tasks/$ARGUMENTS/processes/03-tests/01-write/README.md`: добавь `updated: {{datetime}}`.
