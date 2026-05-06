---
process: 02-acceptance/02-review
run: 2
date: 2026-04-29T18:16:41Z
created: 2026-04-29T18:16:41Z
see-also: tasks/0002-store-crud/stages/02-acceptance/02-review/report-001.md
status: failed
agent: Гримм
checklist: открытые: AR-5
---

## Результат

AR-1—AR-4 закрыты корректно. Обнаружено 1 новое замечание: пробел в AC-ROUTER-09 — сценарий не проверяет поле `op=unknown` в лог-записи, заданное spec §5.1.

## Закрытые замечания (run 1)

**AR-1** — закрыт. AC-ROUTER-09 дополнен проверкой лог-записи ERROR с полем `parse_error`. Given добавлен логгер. Соответствует spec §5.1.

**AR-2** — закрыт. Добавлены AC-FS-INIT-03 и AC-FS-INIT-04. Given явный путь + логгер. Then соответствует НФТ-О-2 дословно.

**AR-3** — закрыт. AC-NFT-ATOMIC-02 переписан как поведенческий: uuid.NewV7() вызван, результат передан в Upsert, операция успешна. Статическая проверка кодовой базы убрана.

**AR-4** — закрыт. AC-ENTITY-UPSERT-01: Given содержит явные `CreatedAt=time.Date(2000,1,1,0,0,0,0,time.UTC)` и `UpdatedAt=time.Date(2000,1,1,0,0,0,0,time.UTC)`. Then проверяет что сохранённые поля не равны этим значениям и установлены в UTC-время вызова.

## Замечания

### AR-5 [пробел в acceptance]

**Сценарий**: AC-ROUTER-09
**Spec**: §5.1

Spec §5.1 задаёт конкретный формат лог-записи при parse error: `ERROR op=unknown parse_error="<text>"`. Then AC-ROUTER-09 проверяет наличие поля `parse_error` с текстом ошибки — но не проверяет наличие поля `op=unknown` в той же записи. Поле `op=unknown` — не опциональное украшение: это единственный способ отличить parse-error запись от любой другой ERROR-записи с полем `parse_error` в будущих реализациях.

Then должен явно указывать: в лог-записи присутствуют поля `op=unknown` и `parse_error=<text>`.
