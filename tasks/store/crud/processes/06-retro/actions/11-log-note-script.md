# Action: scripts/log-note.sh и scripts/log-complaint.sh

## Проблема

Дозапись notes и complaints требует вручную найти нужный файл, открыть, дописать в правильном формате. Барьер достаточно высок чтобы агенты не делали этого по ходу работы.

## Решение

```bash
scripts/log-note.sh store/crud ада "[propose] Лучше использовать интерфейс вместо конкретного типа"
scripts/log-complaint.sh store/crud кроули "[friction] Тесты писались к несуществующему API"
```

Скрипт находит нужный файл по feature + агент, дозаписывает строку с тегом. Создаёт файл если не существует.

## Шаги

- [ ] Написать `scripts/log-note.sh <feature> <агент> "<тег> текст"`
- [ ] Написать `scripts/log-complaint.sh <feature> <агент> "<тег> текст"`
- [ ] Опционально: интеграция с feature-status.sh — показывать количество notes/complaints на шаг

## Источники

Пользователь (ретро store/crud, 2026-05-05)
