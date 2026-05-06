---
purpose: Финальный прогон тестов для store/crud
process: 04-code/04-testing
run: 1
date: 2026-05-04T17:55:00Z
created: 2026-05-04T17:55:00Z
see-also: tasks/0002-store-crud/stages/04-code/02-review/report-002.md
status: done
agent: Азирафаль
checklist: все пункты закрыты
---

## Итог

Прошло: 62 | Упало: 0 | Пропущено: 0

Happy path: 38 тестов, adversarial: 24 теста. Все зелёные. Race detector (-race) не обнаружил гонок.

## Упавшие тесты

Нет.

## Полный вывод

### Happy path (`store/tests/happy/`)

```
=== RUN   TestEntity_Delete_ExistingRecord_RecordNoLongerFound
--- PASS: TestEntity_Delete_ExistingRecord_RecordNoLongerFound (0.00s)
=== RUN   TestEntity_Get_ExistingRecord_ReturnsCorrectData
--- PASS: TestEntity_Get_ExistingRecord_ReturnsCorrectData (0.00s)
=== RUN   TestEntity_List_MultipleRecords_ReturnsAll
--- PASS: TestEntity_List_MultipleRecords_ReturnsAll (0.01s)
=== RUN   TestEntity_List_EmptyStore_ReturnsEmptySlice
--- PASS: TestEntity_List_EmptyStore_ReturnsEmptySlice (0.00s)
=== RUN   TestEntity_List_TmpFilesPresent_IgnoresTmpFiles
--- PASS: TestEntity_List_TmpFilesPresent_IgnoresTmpFiles (0.00s)
=== RUN   TestEntity_Upsert_NewRecord_SetsTimestamps
--- PASS: TestEntity_Upsert_NewRecord_SetsTimestamps (0.00s)
=== RUN   TestEntity_Upsert_NewRecord_PersistsAllFields
--- PASS: TestEntity_Upsert_NewRecord_PersistsAllFields (0.00s)
=== RUN   TestEntity_Upsert_ExistingRecord_PreservesCreatedAtAndUpdatesUpdatedAt
--- PASS: TestEntity_Upsert_ExistingRecord_PreservesCreatedAtAndUpdatesUpdatedAt (0.01s)
=== RUN   TestFSInit_NewWithNonexistentBasedir_CreatesStoreAndSubdirs
--- PASS: TestFSInit_NewWithNonexistentBasedir_CreatesStoreAndSubdirs (0.00s)
=== RUN   TestFSInit_SuccessfulInit_LogsInfoWithBasedir
--- PASS: TestFSInit_SuccessfulInit_LogsInfoWithBasedir (0.00s)
=== RUN   TestJob_Upsert_NewRecord_SetsTimestampsAndPersistsFields
--- PASS: TestJob_Upsert_NewRecord_SetsTimestampsAndPersistsFields (0.00s)
=== RUN   TestJob_Get_ExistingRecord_ReturnsCorrectData
--- PASS: TestJob_Get_ExistingRecord_ReturnsCorrectData (0.00s)
=== RUN   TestJob_Delete_ExistingRecord_RecordNoLongerFound
--- PASS: TestJob_Delete_ExistingRecord_RecordNoLongerFound (0.00s)
=== RUN   TestJob_List_MultipleRecords_ReturnsAll
--- PASS: TestJob_List_MultipleRecords_ReturnsAll (0.00s)
=== RUN   TestJob_List_EmptyStore_ReturnsEmptySlice
--- PASS: TestJob_List_EmptyStore_ReturnsEmptySlice (0.00s)
=== RUN   TestUUID_NewV7_GeneratesVersionSevenUUID
--- PASS: TestUUID_NewV7_GeneratesVersionSevenUUID (0.00s)
=== RUN   TestUUID_NewV7_SequentialIDs_AreMonotonicallyIncreasing
--- PASS: TestUUID_NewV7_SequentialIDs_AreMonotonicallyIncreasing (0.00s)
=== RUN   TestUUID_NewV7_StringRepresentation_CompatibleWithEntityID
--- PASS: TestUUID_NewV7_StringRepresentation_CompatibleWithEntityID (0.00s)
=== RUN   TestRelation_Upsert_NewRecord_SetsTimestamps
--- PASS: TestRelation_Upsert_NewRecord_SetsTimestamps (0.00s)
=== RUN   TestRelation_Upsert_NonexistentFromToIDs_Succeeds
--- PASS: TestRelation_Upsert_NonexistentFromToIDs_Succeeds (0.00s)
=== RUN   TestRelation_Get_ExistingRecord_ReturnsCorrectData
--- PASS: TestRelation_Get_ExistingRecord_ReturnsCorrectData (0.00s)
=== RUN   TestRelation_Delete_ExistingRecord_RecordNoLongerFound
--- PASS: TestRelation_Delete_ExistingRecord_RecordNoLongerFound (0.00s)
=== RUN   TestRelation_List_MultipleRecords_ReturnsAll
--- PASS: TestRelation_List_MultipleRecords_ReturnsAll (0.00s)
=== RUN   TestRelation_List_EmptyStore_ReturnsEmptySlice
--- PASS: TestRelation_List_EmptyStore_ReturnsEmptySlice (0.00s)
=== RUN   TestRouter_Upsert_EnvelopeIDIgnored_TakesIDFromData
--- PASS: TestRouter_Upsert_EnvelopeIDIgnored_TakesIDFromData (0.00s)
=== RUN   TestRouter_MultipleRequestsOnSingleConnection_AllProcessed
--- PASS: TestRouter_MultipleRequestsOnSingleConnection_AllProcessed (0.00s)
=== RUN   TestRouter_ClientClosesConnection_RouterExitsGracefully
--- PASS: TestRouter_ClientClosesConnection_RouterExitsGracefully (0.00s)
=== RUN   TestRouter_New_AcceptsStoreInterfaces
--- PASS: TestRouter_New_AcceptsStoreInterfaces (0.00s)
=== RUN   TestRouter_Upsert_Entity_ReturnsOkWithID
--- PASS: TestRouter_Upsert_Entity_ReturnsOkWithID (0.00s)
=== RUN   TestRouter_Get_ExistingEntity_ReturnsOkWithData
--- PASS: TestRouter_Get_ExistingEntity_ReturnsOkWithData (0.00s)
=== RUN   TestRouter_Delete_ExistingEntity_ReturnsOk
--- PASS: TestRouter_Delete_ExistingEntity_ReturnsOk (0.00s)
=== RUN   TestRouter_List_Entities_ReturnsOkWithDataArray
--- PASS: TestRouter_List_Entities_ReturnsOkWithDataArray (0.00s)
=== RUN   TestRouter_Upsert_Entity_LogsDebugWithOpTypeID
--- PASS: TestRouter_Upsert_Entity_LogsDebugWithOpTypeID (0.00s)
=== RUN   TestRouter_Get_Entity_LogsDebugWithOpTypeID
--- PASS: TestRouter_Get_Entity_LogsDebugWithOpTypeID (0.00s)
=== RUN   TestRouter_List_Entities_LogsDebugWithCount
--- PASS: TestRouter_List_Entities_LogsDebugWithCount (0.00s)
PASS
ok  	github.com/unel/metadataplatform/store/tests/happy	1.088s
```

### Adversarial (`store/tests/adversarial/`)

```
=== RUN   TestNFT_Atomic_UpsertInterruptedBySIGKILL_FileNotCorrupted
--- PASS: TestNFT_Atomic_UpsertInterruptedBySIGKILL_FileNotCorrupted (0.01s)
=== RUN   TestEntity_Delete_NonExistentID_ReturnsErrNotFound
--- PASS: TestEntity_Delete_NonExistentID_ReturnsErrNotFound (0.00s)
=== RUN   TestEntity_Delete_EmptyID_ReturnsError
--- PASS: TestEntity_Delete_EmptyID_ReturnsError (0.00s)
=== RUN   TestEntity_Delete_FileExistsButRemoveFails_ReturnsDeleteError
--- PASS: TestEntity_Delete_FileExistsButRemoveFails_ReturnsDeleteError (0.00s)
=== RUN   TestEntity_Get_EmptyID_ReturnsError
--- PASS: TestEntity_Get_EmptyID_ReturnsError (0.00s)
=== RUN   TestEntity_Get_NonExistentID_ReturnsErrNotFound
--- PASS: TestEntity_Get_NonExistentID_ReturnsErrNotFound (0.00s)
=== RUN   TestEntity_Get_FileExistsButUnreadable_ReturnsReadError
--- PASS: TestEntity_Get_FileExistsButUnreadable_ReturnsReadError (0.00s)
=== RUN   TestEntity_List_UnreadableDirectory_ReturnsError
--- PASS: TestEntity_List_UnreadableDirectory_ReturnsError (0.00s)
=== RUN   TestEntity_List_PartialDecodeFailure_ReturnsNilNotPartial
--- PASS: TestEntity_List_PartialDecodeFailure_ReturnsNilNotPartial (0.00s)
=== RUN   TestEntity_List_IgnoresNonJsonFiles
--- PASS: TestEntity_List_IgnoresNonJsonFiles (0.00s)
=== RUN   TestEntity_List_OnError_ReturnsNilSliceNotEmpty
--- PASS: TestEntity_List_OnError_ReturnsNilSliceNotEmpty (0.00s)
=== RUN   TestEntity_Upsert_EmptyID_ReturnsError
--- PASS: TestEntity_Upsert_EmptyID_ReturnsError (0.00s)
=== RUN   TestEntity_Upsert_ExistingFileUnreadable_ReturnsReadError
--- PASS: TestEntity_Upsert_ExistingFileUnreadable_ReturnsReadError (0.00s)
=== RUN   TestEntity_Upsert_DirectoryUnwritable_ReturnsWriteError
--- PASS: TestEntity_Upsert_DirectoryUnwritable_ReturnsWriteError (0.00s)
=== RUN   TestEntity_Upsert_WriteError_NoTempFileLeft
--- PASS: TestEntity_Upsert_WriteError_NoTempFileLeft (0.00s)
=== RUN   TestFSInit_UncreatableBasedir_ReturnsError
--- PASS: TestFSInit_UncreatableBasedir_ReturnsError (0.00s)
=== RUN   TestFSInit_UncreatableBasedir_LogsErrorWithFields
--- PASS: TestFSInit_UncreatableBasedir_LogsErrorWithFields (0.00s)
=== RUN   TestFSInit_ValidBasedir_LogsInfoAndNoError
--- PASS: TestFSInit_ValidBasedir_LogsInfoAndNoError (0.00s)
=== RUN   TestJob_Upsert_EmptyID_ReturnsError
--- PASS: TestJob_Upsert_EmptyID_ReturnsError (0.00s)
=== RUN   TestJob_Get_NonExistentID_ReturnsErrNotFound
--- PASS: TestJob_Get_NonExistentID_ReturnsErrNotFound (0.00s)
=== RUN   TestJob_Delete_NonExistentID_ReturnsErrNotFound
--- PASS: TestJob_Delete_NonExistentID_ReturnsErrNotFound (0.00s)
=== RUN   TestJob_Upsert_Then_Get_PathIsCorrect
--- PASS: TestJob_Upsert_Then_Get_PathIsCorrect (0.00s)
=== RUN   TestRelation_Upsert_EmptyID_ReturnsError
--- PASS: TestRelation_Upsert_EmptyID_ReturnsError (0.00s)
=== RUN   TestRelation_Get_NonExistentID_ReturnsErrNotFound
--- PASS: TestRelation_Get_NonExistentID_ReturnsErrNotFound (0.00s)
=== RUN   TestRelation_Delete_NonExistentID_ReturnsErrNotFound
--- PASS: TestRelation_Delete_NonExistentID_ReturnsErrNotFound (0.00s)
=== RUN   TestRelation_Upsert_NonExistentFromToIDs_Succeeds
--- PASS: TestRelation_Upsert_NonExistentFromToIDs_Succeeds (0.00s)
=== RUN   TestRouter_Upsert_MissingData_ReturnsInvalidRequest_ConnectionAlive
--- PASS: TestRouter_Upsert_MissingData_ReturnsInvalidRequest_ConnectionAlive (0.00s)
=== RUN   TestRouter_Upsert_DataNull_ReturnsInvalidRequest_ConnectionAlive
--- PASS: TestRouter_Upsert_DataNull_ReturnsInvalidRequest_ConnectionAlive (0.00s)
=== RUN   TestRouter_UnknownOp_ReturnsUnknownOpError_ConnectionAlive
--- PASS: TestRouter_UnknownOp_ReturnsUnknownOpError_ConnectionAlive (0.00s)
=== RUN   TestRouter_UnknownType_ReturnsUnknownTypeError_ConnectionAlive
--- PASS: TestRouter_UnknownType_ReturnsUnknownTypeError_ConnectionAlive (0.00s)
=== RUN   TestRouter_InvalidJSON_ClosesConnection_LogsParseError
--- PASS: TestRouter_InvalidJSON_ClosesConnection_LogsParseError (0.00s)
=== RUN   TestRouter_InvalidJSONInData_ReturnsInvalidRequest_ConnectionAlive
--- PASS: TestRouter_InvalidJSONInData_ReturnsInvalidRequest_ConnectionAlive (0.00s)
=== RUN   TestRouter_GetNonexistentEntity_LogsError
--- PASS: TestRouter_GetNonexistentEntity_LogsError (0.00s)
PASS
ok  	github.com/unel/metadataplatform/store/tests/adversarial	1.051s
```
