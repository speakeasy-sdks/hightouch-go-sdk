# SyncRunPlannedRows

The number of planned rows that this sync run was supposed to execute.

Note that the counts for `successfulRows` and `failedRows` may not add up
to `plannedRows` if the sync was cancelled.


## Fields

| Field                       | Type                        | Required                    | Description                 |
| --------------------------- | --------------------------- | --------------------------- | --------------------------- |
| `AddedCount`                | *float64*                   | :heavy_check_mark:          | The number of added rows.   |
| `ChangedCount`              | *float64*                   | :heavy_check_mark:          | The number of changed rows. |
| `RemovedCount`              | *float64*                   | :heavy_check_mark:          | The number of removed rows. |