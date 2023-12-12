# Sync

Syncs define how data from models are mapped to destinations. Each time a
sync runs, Hightouch calculates the rows that have changed since the last
run, and syncs them to Sync's destination.


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Configuration`                                                                                                                                                                                                                                                                                                                                                                                      | map[string]*interface{}*                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The sync's configuration. This specifies how data is mapped, among other<br/>configuration.<br/><br/>The schema depends on the destination type.<br/><br/>Consumers should NOT make assumptions on the contents of the<br/>configuration. It may change as Hightouch updates its internal code.                                                                                                      |
| `CreatedAt`                                                                                                                                                                                                                                                                                                                                                                                          | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The timestamp when the sync was created                                                                                                                                                                                                                                                                                                                                                              |
| `DestinationID`                                                                                                                                                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The id of the Destination that sync is connected to                                                                                                                                                                                                                                                                                                                                                  |
| `Disabled`                                                                                                                                                                                                                                                                                                                                                                                           | *bool*                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | Whether the sync has been disabled by the user.                                                                                                                                                                                                                                                                                                                                                      |
| `ID`                                                                                                                                                                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The sync's id                                                                                                                                                                                                                                                                                                                                                                                        |
| `LastRunAt`                                                                                                                                                                                                                                                                                                                                                                                          | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The timestamp of the last sync run                                                                                                                                                                                                                                                                                                                                                                   |
| `ModelID`                                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The id of the Model that sync is connected to                                                                                                                                                                                                                                                                                                                                                        |
| `PrimaryKey`                                                                                                                                                                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The primary key that sync uses to identify data from source                                                                                                                                                                                                                                                                                                                                          |
| `ReferencedColumns`                                                                                                                                                                                                                                                                                                                                                                                  | []*string*                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The reference column that sync depends on to sync data from source                                                                                                                                                                                                                                                                                                                                   |
| `Schedule`                                                                                                                                                                                                                                                                                                                                                                                           | [shared.Schedule](../../../pkg/models/shared/schedule.md)                                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The scheduling configuration. It can be triggerd based on several ways:<br/><br/>Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)<br/><br/>Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.<br/><br/>Visual: the sync will be trigged based a visual cron configuration on UI<br/><br/>DBT-cloud: the sync will be trigged based on a dbt cloud job |
| `Slug`                                                                                                                                                                                                                                                                                                                                                                                               | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The sync's slug                                                                                                                                                                                                                                                                                                                                                                                      |
| `Status`                                                                                                                                                                                                                                                                                                                                                                                             | [shared.SyncStatus](../../../pkg/models/shared/syncstatus.md)                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | SyncStatus                                                                                                                                                                                                                                                                                                                                                                                           |
| `UpdatedAt`                                                                                                                                                                                                                                                                                                                                                                                          | [time.Time](https://pkg.go.dev/time#Time)                                                                                                                                                                                                                                                                                                                                                            | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The timestamp when the sync was last updated                                                                                                                                                                                                                                                                                                                                                         |
| `WorkspaceID`                                                                                                                                                                                                                                                                                                                                                                                        | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The id of the workspace that the sync belongs to                                                                                                                                                                                                                                                                                                                                                     |