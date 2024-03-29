# ListDestinationRequest


## Fields

| Field                                                            | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `Limit`                                                          | **float64*                                                       | :heavy_minus_sign:                                               | limit the number of objects returned (default is 100)            |
| `Name`                                                           | **string*                                                        | :heavy_minus_sign:                                               | Filter based on the destination's name                           |
| `Offset`                                                         | **float64*                                                       | :heavy_minus_sign:                                               | set the offset on results (for pagination)                       |
| `OrderBy`                                                        | [*operations.OrderBy](../../../pkg/models/operations/orderby.md) | :heavy_minus_sign:                                               | Order the returned destinations                                  |
| `Slug`                                                           | **string*                                                        | :heavy_minus_sign:                                               | Filter based on destination's slug                               |