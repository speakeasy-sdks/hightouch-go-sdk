# CreateSourceResponse


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ContentType`                                                             | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `CreateSource200ApplicationJSONAnyOf`                                     | *interface{}*                                                             | :heavy_minus_sign:                                                        | Ok                                                                        |
| `InternalServerError`                                                     | [*shared.InternalServerError](../../models/shared/internalservererror.md) | :heavy_minus_sign:                                                        | Something went wrong                                                      |
| `StatusCode`                                                              | *int*                                                                     | :heavy_check_mark:                                                        | N/A                                                                       |
| `RawResponse`                                                             | [*http.Response](https://pkg.go.dev/net/http#Response)                    | :heavy_minus_sign:                                                        | N/A                                                                       |
| `ValidateErrorJSON`                                                       | [*shared.ValidateErrorJSON](../../models/shared/validateerrorjson.md)     | :heavy_minus_sign:                                                        | Conflict                                                                  |