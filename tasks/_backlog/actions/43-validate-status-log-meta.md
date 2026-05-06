# Action: scripts/validate-status-log-meta.sh

## Проблема

Frontmatter status-log.md содержит машиночитаемые поля (`run`, `next-step`, `previous-step`). Нет способа быстро проверить что они присутствуют и консистентны (например, `next-step` одного шага совпадает с `previous-step` следующего).

## Решение

Написать `scripts/validate-status-log-meta.sh <feature>`:
- Проверяет наличие frontmatter во всех status-log.md фичи
- Проверяет наличие обязательных полей: `run`, `next-step`, `previous-step`
- Проверяет консистентность цепочки next/previous между шагами
- Выводит список проблем или "OK"

## Источники

Гримм (ревью scripts-design.md, 2026-05-06)
Пользователь (обсуждение cascade-stale и frontmatter, 2026-05-06)
