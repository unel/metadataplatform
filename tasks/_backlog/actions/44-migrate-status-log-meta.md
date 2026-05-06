# Action: scripts/add-status-log-meta.sh

## Проблема

Существующие status-log.md файлы не имеют frontmatter. Новая механика `begin-step`/`end-step` и `cascade-stale` опираются на поля `run`, `next-step`, `previous-step` в frontmatter. Нужен способ добавить их к существующим файлам.

## Решение

Написать `scripts/add-status-log-meta.sh <feature>`:
- Для каждого status-log.md в фиче определяет текущий номер рана (по количеству `in-progress` записей в файле)
- Определяет `next-step` / `previous-step` по порядку шагов в выбранном флоу фичи
- Добавляет frontmatter в начало файла
- Не перезаписывает frontmatter если уже есть (или только добавляет отсутствующие поля)

Зависит от BL-037 (система флоу — нужно знать порядок шагов).

## Источники

Гримм (ревью scripts-design.md, 2026-05-06)
Пользователь (2026-05-06)
