# SyncSequenceStatusOutput

The status of a given sync sequence run at both a high level and the individual
sync-run level.


## Fields

| Field                                                                                         | Type                                                                                          | Required                                                                                      | Description                                                                                   |
| --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------- |
| `ID`                                                                                          | *string*                                                                                      | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `Status`                                                                                      | [SyncSequenceRequestStatus](../../models/shared/syncsequencerequeststatus.md)                 | :heavy_check_mark:                                                                            | N/A                                                                                           |
| `SyncRuns`                                                                                    | [][SyncSequenceStatusOutputSyncRuns](../../models/shared/syncsequencestatusoutputsyncruns.md) | :heavy_check_mark:                                                                            | N/A                                                                                           |