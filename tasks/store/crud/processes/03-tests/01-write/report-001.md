---
purpose: Написание тестов для store/crud
process: 03-tests/01-write
run: 1
date: 2026-04-29T18:31:54Z
created: 2026-04-29T18:31:54Z
status: done
agent: Азирафаль + Кроули
checklist: все пункты закрыты
---

## Happy path тесты (Азирафаль)

`tasks/store/crud/tests/happy/`

| Файл | Тесты |
|---|---|
| fs_init_test.go | TestFSInit_NewWithNonexistentBasedir_CreatesStoreAndSubdirs, TestFSInit_SuccessfulInit_LogsInfoWithBasedir |
| entity_upsert_test.go | TestEntity_Upsert_NewRecord_SetsTimestamps, TestEntity_Upsert_NewRecord_PersistsAllFields, TestEntity_Upsert_ExistingRecord_PreservesCreatedAtAndUpdatesUpdatedAt |
| entity_get_test.go | TestEntity_Get_ExistingRecord_ReturnsCorrectData, TestEntity_Get_EmptyID_ReturnsError |
| entity_delete_test.go | TestEntity_Delete_ExistingRecord_RecordNoLongerFound |
| entity_list_test.go | TestEntity_List_MultipleRecords_ReturnsAll, TestEntity_List_EmptyStore_ReturnsEmptySlice, TestEntity_List_TmpFilesPresent_IgnoresTmpFiles |
| relation_test.go | TestRelation_Upsert_NewRecord_SetsTimestamps, TestRelation_Upsert_NonexistentFromToIDs_Succeeds, TestRelation_Get_ExistingRecord_ReturnsCorrectData, TestRelation_Delete_ExistingRecord_RecordNoLongerFound, TestRelation_List_MultipleRecords_ReturnsAll, TestRelation_List_EmptyStore_ReturnsEmptySlice |
| job_test.go | TestJob_Upsert_NewRecord_SetsTimestampsAndPersistsFields, TestJob_Get_ExistingRecord_ReturnsCorrectData, TestJob_Delete_ExistingRecord_RecordNoLongerFound, TestJob_List_MultipleRecords_ReturnsAll, TestJob_List_EmptyStore_ReturnsEmptySlice |
| router_happy_test.go | TestRouter_New_AcceptsStoreInterfaces, TestRouter_Upsert_Entity_ReturnsOkWithID, TestRouter_Get_ExistingEntity_ReturnsOkWithData, TestRouter_Delete_ExistingEntity_ReturnsOk, TestRouter_List_Entities_ReturnsOkWithDataArray |
| router_conn_test.go | TestRouter_Upsert_EnvelopeIDIgnored_TakesIDFromData, TestRouter_MultipleRequestsOnSingleConnection_AllProcessed, TestRouter_ClientClosesConnection_RouterExitsGracefully |
| router_logging_test.go | TestRouter_Upsert_Entity_LogsDebugWithOpTypeID, TestRouter_Get_Entity_LogsDebugWithOpTypeID, TestRouter_List_Entities_LogsDebugWithCount |
| nft_test.go | TestUUID_NewV7_GeneratesVersionSevenUUID, TestUUID_NewV7_SequentialIDs_AreMonotonicallyIncreasing, TestUUID_NewV7_StringRepresentation_CompatibleWithEntityID |

Итого: 36 тестов.

## Adversarial тесты (Кроули)

`tasks/store/crud/tests/adversarial/`

| Файл | Тесты |
|---|---|
| fs_init_error_test.go | TestFSInit_UncreatableBasedir_ReturnsError, TestFSInit_UncreatableBasedir_LogsErrorWithFields, TestFSInit_ValidBasedir_LogsInfoAndNoError |
| entity_upsert_error_test.go | TestEntity_Upsert_EmptyID_ReturnsError, TestEntity_Upsert_ExistingFileUnreadable_ReturnsReadError, TestEntity_Upsert_DirectoryUnwritable_ReturnsWriteError, TestEntity_Upsert_WriteError_NoTempFileLeft |
| entity_get_error_test.go | TestEntity_Get_NonExistentID_ReturnsErrNotFound, TestEntity_Get_FileExistsButUnreadable_ReturnsReadError |
| entity_delete_error_test.go | TestEntity_Delete_NonExistentID_ReturnsErrNotFound, TestEntity_Delete_EmptyID_ReturnsError, TestEntity_Delete_FileExistsButRemoveFails_ReturnsDeleteError |
| entity_list_error_test.go | TestEntity_List_UnreadableDirectory_ReturnsError, TestEntity_List_PartialDecodeFailure_ReturnsNilNotPartial, TestEntity_List_OnError_ReturnsNilSliceNotEmpty |
| relation_error_test.go | TestRelation_Upsert_EmptyID_ReturnsError, TestRelation_Get_NonExistentID_ReturnsErrNotFound, TestRelation_Delete_NonExistentID_ReturnsErrNotFound, TestRelation_Upsert_NonExistentFromToIDs_Succeeds |
| job_error_test.go | TestJob_Upsert_EmptyID_ReturnsError, TestJob_Get_NonExistentID_ReturnsErrNotFound, TestJob_Delete_NonExistentID_ReturnsErrNotFound, TestJob_Upsert_Then_Get_PathIsCorrect |
| router_error_test.go | TestRouter_Upsert_MissingData_ReturnsInvalidRequest_ConnectionAlive, TestRouter_Upsert_DataNull_ReturnsInvalidRequest_ConnectionAlive, TestRouter_UnknownOp_ReturnsUnknownOpError_ConnectionAlive, TestRouter_UnknownType_ReturnsUnknownTypeError_ConnectionAlive, TestRouter_InvalidJSON_ClosesConnection_LogsParseError, TestRouter_InvalidJSONInData_ReturnsInvalidRequest_ConnectionAlive |
| atomic_test.go | TestNFT_Atomic_UpsertInterruptedBySIGKILL_FileNotCorrupted |

Итого: 27 тестов (+ helpers_test.go, router_helpers_test.go).

## Маппинг: тест → сценарий acceptance

| Сценарий | Тест |
|---|---|
| AC-FS-INIT-01 | TestFSInit_NewWithNonexistentBasedir_CreatesStoreAndSubdirs |
| AC-FS-INIT-02 | TestFSInit_UncreatableBasedir_ReturnsError |
| AC-FS-INIT-03 | TestFSInit_SuccessfulInit_LogsInfoWithBasedir |
| AC-FS-INIT-04 | TestFSInit_UncreatableBasedir_LogsErrorWithFields |
| AC-ENTITY-UPSERT-01 | TestEntity_Upsert_NewRecord_SetsTimestamps, TestEntity_Upsert_NewRecord_PersistsAllFields |
| AC-ENTITY-UPSERT-02 | TestEntity_Upsert_ExistingRecord_PreservesCreatedAtAndUpdatesUpdatedAt |
| AC-ENTITY-UPSERT-03 | TestEntity_Upsert_EmptyID_ReturnsError |
| AC-ENTITY-UPSERT-04 | TestEntity_Upsert_ExistingFileUnreadable_ReturnsReadError |
| AC-ENTITY-UPSERT-05 | TestEntity_Upsert_DirectoryUnwritable_ReturnsWriteError |
| AC-ENTITY-UPSERT-06 | TestEntity_Upsert_WriteError_NoTempFileLeft |
| AC-ENTITY-GET-01 | TestEntity_Get_ExistingRecord_ReturnsCorrectData |
| AC-ENTITY-GET-02 | TestEntity_Get_NonExistentID_ReturnsErrNotFound |
| AC-ENTITY-GET-03 | TestEntity_Get_EmptyID_ReturnsError |
| AC-ENTITY-GET-04 | TestEntity_Get_FileExistsButUnreadable_ReturnsReadError |
| AC-ENTITY-DELETE-01 | TestEntity_Delete_ExistingRecord_RecordNoLongerFound |
| AC-ENTITY-DELETE-02 | TestEntity_Delete_NonExistentID_ReturnsErrNotFound |
| AC-ENTITY-DELETE-03 | TestEntity_Delete_EmptyID_ReturnsError |
| AC-ENTITY-DELETE-04 | TestEntity_Delete_FileExistsButRemoveFails_ReturnsDeleteError |
| AC-ENTITY-LIST-01 | TestEntity_List_MultipleRecords_ReturnsAll |
| AC-ENTITY-LIST-02 | TestEntity_List_EmptyStore_ReturnsEmptySlice |
| AC-ENTITY-LIST-03 | TestEntity_List_UnreadableDirectory_ReturnsError |
| AC-ENTITY-LIST-04 | TestEntity_List_PartialDecodeFailure_ReturnsNilNotPartial |
| AC-ENTITY-LIST-05 | TestEntity_List_TmpFilesPresent_IgnoresTmpFiles, TestEntity_List_IgnoresNonJsonFiles |
| AC-RELATION-UPSERT-01 | TestRelation_Upsert_NewRecord_SetsTimestamps |
| AC-RELATION-UPSERT-02 | TestRelation_Upsert_NonexistentFromToIDs_Succeeds (happy) + TestRelation_Upsert_NonExistentFromToIDs_Succeeds (adversarial) |
| AC-RELATION-UPSERT-03 | TestRelation_Upsert_EmptyID_ReturnsError |
| AC-RELATION-GET-01 | TestRelation_Get_ExistingRecord_ReturnsCorrectData |
| AC-RELATION-GET-02 | TestRelation_Get_NonExistentID_ReturnsErrNotFound |
| AC-RELATION-DELETE-01 | TestRelation_Delete_ExistingRecord_RecordNoLongerFound |
| AC-RELATION-DELETE-02 | TestRelation_Delete_NonExistentID_ReturnsErrNotFound |
| AC-RELATION-LIST-01 | TestRelation_List_MultipleRecords_ReturnsAll |
| AC-RELATION-LIST-02 | TestRelation_List_EmptyStore_ReturnsEmptySlice |
| AC-JOB-UPSERT-01 | TestJob_Upsert_NewRecord_SetsTimestampsAndPersistsFields |
| AC-JOB-UPSERT-02 | TestJob_Upsert_EmptyID_ReturnsError |
| AC-JOB-GET-01 | TestJob_Get_ExistingRecord_ReturnsCorrectData |
| AC-JOB-GET-02 | TestJob_Get_NonExistentID_ReturnsErrNotFound |
| AC-JOB-DELETE-01 | TestJob_Delete_ExistingRecord_RecordNoLongerFound |
| AC-JOB-DELETE-02 | TestJob_Delete_NonExistentID_ReturnsErrNotFound |
| AC-JOB-LIST-01 | TestJob_List_MultipleRecords_ReturnsAll |
| AC-JOB-LIST-02 | TestJob_List_EmptyStore_ReturnsEmptySlice |
| AC-ROUTER-01 | TestRouter_Upsert_Entity_ReturnsOkWithID |
| AC-ROUTER-02 | TestRouter_Get_ExistingEntity_ReturnsOkWithData |
| AC-ROUTER-03 | TestRouter_Delete_ExistingEntity_ReturnsOk |
| AC-ROUTER-04 | TestRouter_List_Entities_ReturnsOkWithDataArray |
| AC-ROUTER-05 | TestRouter_Upsert_MissingData_ReturnsInvalidRequest_ConnectionAlive |
| AC-ROUTER-06 | TestRouter_Upsert_DataNull_ReturnsInvalidRequest_ConnectionAlive |
| AC-ROUTER-07 | TestRouter_UnknownOp_ReturnsUnknownOpError_ConnectionAlive |
| AC-ROUTER-08 | TestRouter_UnknownType_ReturnsUnknownTypeError_ConnectionAlive |
| AC-ROUTER-09 | TestRouter_InvalidJSON_ClosesConnection_LogsParseError |
| AC-ROUTER-10 | TestRouter_InvalidJSONInData_ReturnsInvalidRequest_ConnectionAlive |
| AC-ROUTER-11 | TestRouter_Upsert_EnvelopeIDIgnored_TakesIDFromData |
| AC-ROUTER-12 | TestRouter_MultipleRequestsOnSingleConnection_AllProcessed |
| AC-ROUTER-13 | TestRouter_ClientClosesConnection_RouterExitsGracefully |
| AC-ROUTER-14 | TestRouter_Upsert_Entity_LogsDebugWithOpTypeID |
| AC-ROUTER-15 | TestRouter_List_Entities_LogsDebugWithCount |
| AC-ROUTER-16 | неявно покрыт через TestRouter_UnknownOp/UnknownType — ERROR лог при ошибке операции |
| AC-NFT-ATOMIC-01 | TestNFT_Atomic_UpsertInterruptedBySIGKILL_FileNotCorrupted |
| AC-NFT-ATOMIC-02 | TestUUID_NewV7_* (3 теста) |
| AC-ROUTER-INIT-01 | TestRouter_New_AcceptsStoreInterfaces |

## Непокрытые сценарии

**AC-ROUTER-16** (частично): spec требует лог уровня ERROR при ошибке операции с полями op/type/id/error. Тесты AC-ROUTER-07/-08 в adversarial покрывают ошибочные ответы роутера, но не проверяют явно формат лог-записи ERROR. Подлежит уточнению на ревью.

**AC-ENTITY-UPSERT-04/-05, AC-ENTITY-LIST-03, AC-ENTITY-DELETE-04**: тесты через chmod (права 000). Не работают от root. Технический долг — зафиксирован в коде тестов.

## Архитектурные предположения тестов

Тесты сделаны с предположениями о сигнатурах, которые нужно согласовать при написании кода:
- `fs.New(cfg, logger)` — логгер инжектируется в конструктор
- `router.New(entities, relations, jobs, logger)` — логгер в роутере
- `router.Handle(ctx, conn net.Conn)` — метод обработки соединения
- `fs.Store` реализует все три интерфейса одновременно

## Результаты первого прогона

Тесты не запускались — код ещё не написан. Следующий шаг: 03-tests/02-review.
