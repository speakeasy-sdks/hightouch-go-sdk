# TriggerRunCustomInput

The input of a trigger action to run syncs based on sync ID, slug or other filters


## Fields

| Field                                                                                   | Type                                                                                    | Required                                                                                | Description                                                                             |
| --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| `FullResync`                                                                            | **bool*                                                                                 | :heavy_minus_sign:                                                                      | Whether to resync all the rows in the query (i.e. ignoring previously<br/>synced rows). |
| `ResetCDC`                                                                              | **bool*                                                                                 | :heavy_minus_sign:                                                                      | Whether to sync all the rows in the query without executing changes on the destination. |
| `SyncID`                                                                                | **string*                                                                               | :heavy_minus_sign:                                                                      | Trigger run based on sync id                                                            |
| `SyncSlug`                                                                              | **string*                                                                               | :heavy_minus_sign:                                                                      | Trigger run based on sync slug                                                          |