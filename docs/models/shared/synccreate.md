# SyncCreate

The input for creating a Sync


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                | Type                                                                                                                                                                                                                                                                                                                                                                                                 | Required                                                                                                                                                                                                                                                                                                                                                                                             | Description                                                                                                                                                                                                                                                                                                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Configuration`                                                                                                                                                                                                                                                                                                                                                                                      | map[string]*interface{}*                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The sync's configuration. This specifies how data is mapped, among other<br/>configuration.<br/><br/>The schema depends on the destination type.<br/><br/>Consumers should NOT make assumptions on the contents of the<br/>configuration. It may change as Hightouch updates its internal code.                                                                                                      |
| `DestinationID`                                                                                                                                                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | Number as a string                                                                                                                                                                                                                                                                                                                                                                                   |
| `Disabled`                                                                                                                                                                                                                                                                                                                                                                                           | *bool*                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | Whether the sync has been disabled by the user.                                                                                                                                                                                                                                                                                                                                                      |
| `ModelID`                                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | Number as a string                                                                                                                                                                                                                                                                                                                                                                                   |
| `Schedule`                                                                                                                                                                                                                                                                                                                                                                                           | [SyncCreateSchedule](../../models/shared/synccreateschedule.md)                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The scheduling configuration. It can be triggerd based on several ways:<br/><br/>Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)<br/><br/>Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.<br/><br/>Visual: the sync will be trigged based a visual cron configuration on UI<br/><br/>DBT-cloud: the sync will be trigged based on a dbt cloud job |
| `Slug`                                                                                                                                                                                                                                                                                                                                                                                               | *string*                                                                                                                                                                                                                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                   | The sync's slug                                                                                                                                                                                                                                                                                                                                                                                      |