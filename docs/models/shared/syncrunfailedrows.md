# SyncRunFailedRows

The number of rows that we attempted to sync, but were rejected by the
destination.

This does not include rows that weren't attempted due to the sync being
cancelled.


## Fields

| Field                         | Type                          | Required                      | Description                   |
| ----------------------------- | ----------------------------- | ----------------------------- | ----------------------------- |
| `AddedCount`                  | *float64*                     | :heavy_check_mark:            | The number of failed adds.    |
| `ChangedCount`                | *float64*                     | :heavy_check_mark:            | The number of failed changes. |
| `RemovedCount`                | *float64*                     | :heavy_check_mark:            | The number of failed removes. |