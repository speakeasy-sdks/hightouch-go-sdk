# SyncSequenceStatusOutput

The status of a given sync sequence run at both a high level and the individual
sync-run level.


## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ID`                                                                                        | *string*                                                                                    | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `Status`                                                                                    | [shared.SyncSequenceRequestStatus](../../../pkg/models/shared/syncsequencerequeststatus.md) | :heavy_check_mark:                                                                          | N/A                                                                                         |
| `SyncRuns`                                                                                  | [][shared.SyncRuns](../../../pkg/models/shared/syncruns.md)                                 | :heavy_check_mark:                                                                          | N/A                                                                                         |