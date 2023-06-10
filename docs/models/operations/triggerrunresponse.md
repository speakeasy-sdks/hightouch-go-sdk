# TriggerRunResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |
| `TriggerRunOutput`                                                    | [*shared.TriggerRunOutput](../../models/shared/triggerrunoutput.md)   | :heavy_minus_sign:                                                    | Ok                                                                    |
| `ValidateErrorJSON`                                                   | [*shared.ValidateErrorJSON](../../models/shared/validateerrorjson.md) | :heavy_minus_sign:                                                    | Validation Failed                                                     |