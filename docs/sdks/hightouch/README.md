# Hightouch SDK


## Overview

Hightouch API: Hightouch Public Rest API to access syncs, models, sources and destinations

### Available Operations

* [CreateDestination](#createdestination) - Create Destination
* [CreateModel](#createmodel) - Create Model
* [CreateSource](#createsource) - Create Source
* [CreateSync](#createsync) - Create Sync
* [GetDestination](#getdestination) - Get Destination
* [GetModel](#getmodel) - Get Model
* [GetSource](#getsource) - Get Source
* [GetSync](#getsync) - Get Sync
* [GetSyncSequenceRun](#getsyncsequencerun) - Sync sequence status
* [ListDestination](#listdestination) - List Destinations
* [ListModel](#listmodel) - List Models
* [ListSource](#listsource) - List Sources
* [ListSync](#listsync) - List Syncs
* [ListSyncRuns](#listsyncruns) - List Sync Runs
* [TriggerRun](#triggerrun) - Trigger Sync
* [TriggerRunCustom](#triggerruncustom) - Trigger Sync From ID or Slug
* [TriggerRunIDGraph](#triggerrunidgraph)
* [TriggerSequenceRun](#triggersequencerun) - Trigger Sync sequence
* [UpdateDestination](#updatedestination) - Update Destination
* [UpdateModel](#updatemodel) - Update Model
* [UpdateSource](#updatesource) - Update Source
* [UpdateSync](#updatesync) - Update Sync

## CreateDestination

Create a new destination

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.CreateDestination(ctx, shared.DestinationCreate{
        Configuration: map[string]interface{}{
            "key": "<value>",
        },
        Name: "<value>",
        Slug: "<value>",
        Type: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [shared.DestinationCreate](../../pkg/models/shared/destinationcreate.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.CreateDestinationResponse](../../pkg/models/operations/createdestinationresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 409,422                     | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## CreateModel

Create a new model

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.CreateModel(ctx, operations.CreateModelRequest{
        ModelCreate: shared.ModelCreate{
            IsSchema: false,
            Name: "<value>",
            PrimaryKey: "<value>",
            QueryType: "<value>",
            Slug: "<value>",
            SourceID: 8761.56,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateModelRequest](../../pkg/models/operations/createmodelrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.CreateModelResponse](../../pkg/models/operations/createmodelresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 409,422                     | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## CreateSource

Create a new source

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.CreateSource(ctx, shared.SourceCreate{
        Configuration: map[string]interface{}{
            "key": "<value>",
        },
        Name: "<value>",
        Slug: "<value>",
        Type: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.SourceCreate](../../pkg/models/shared/sourcecreate.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.CreateSourceResponse](../../pkg/models/operations/createsourceresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 409,422                     | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## CreateSync

Create a new sync

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.CreateSync(ctx, shared.SyncCreate{
        Configuration: map[string]interface{}{
            "key": "<value>",
        },
        DestinationID: 8797.77,
        Disabled: false,
        ModelID: 2438.5,
        Schedule: &shared.SyncCreateSchedule{
            Schedule: shared.CreateSyncCreateSchemasScheduleCronSchedule(
                    shared.CronSchedule{
                        Expression: "<value>",
                    },
            ),
            Type: "<value>",
        },
        Slug: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.SyncCreate](../../pkg/models/shared/synccreate.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.CreateSyncResponse](../../pkg/models/operations/createsyncresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 409,422                     | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## GetDestination

Retrieve a destination based on its Hightouch ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetDestination(ctx, operations.GetDestinationRequest{
        DestinationID: 4856.96,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Destination != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetDestinationRequest](../../pkg/models/operations/getdestinationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetDestinationResponse](../../pkg/models/operations/getdestinationresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetModel

Retrieve models from model ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetModel(ctx, operations.GetModelRequest{
        ModelID: 7962.16,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Model != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetModelRequest](../../pkg/models/operations/getmodelrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetModelResponse](../../pkg/models/operations/getmodelresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSource

Retrieve source from source ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetSource(ctx, operations.GetSourceRequest{
        SourceID: 4378.62,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Source != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetSourceRequest](../../pkg/models/operations/getsourcerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetSourceResponse](../../pkg/models/operations/getsourceresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## GetSync

Retrieve sync from sync ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetSync(ctx, operations.GetSyncRequest{
        SyncID: 7319.86,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Sync != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetSyncRequest](../../pkg/models/operations/getsyncrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetSyncResponse](../../pkg/models/operations/getsyncresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetSyncSequenceRun

Get the status of a sync sequence run.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.GetSyncSequenceRun(ctx, operations.GetSyncSequenceRunRequest{
        SyncSequenceRunID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SyncSequenceStatusOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetSyncSequenceRunRequest](../../pkg/models/operations/getsyncsequencerunrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetSyncSequenceRunResponse](../../pkg/models/operations/getsyncsequencerunresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ListDestination

List the destinations in the user's workspace

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ListDestination(ctx, operations.ListDestinationRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListDestinationRequest](../../pkg/models/operations/listdestinationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListDestinationResponse](../../pkg/models/operations/listdestinationresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ListModel

List all the models in the current workspace including parent and related models

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ListModel(ctx, operations.ListModelRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListModelRequest](../../pkg/models/operations/listmodelrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListModelResponse](../../pkg/models/operations/listmodelresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ListSource

List all the sources in the current workspace

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ListSource(ctx, operations.ListSourceRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListSourceRequest](../../pkg/models/operations/listsourcerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListSourceResponse](../../pkg/models/operations/listsourceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListSync

List all the syncs in the current workspace

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ListSync(ctx, operations.ListSyncRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListSyncRequest](../../pkg/models/operations/listsyncrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.ListSyncResponse](../../pkg/models/operations/listsyncresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## ListSyncRuns

List all sync runs under a sync

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.ListSyncRuns(ctx, operations.ListSyncRunsRequest{
        SyncID: 8858.62,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListSyncRunsRequest](../../pkg/models/operations/listsyncrunsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListSyncRunsResponse](../../pkg/models/operations/listsyncrunsresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## TriggerRun

Trigger a new run for the given sync.

If a run is already in progress, this queues a sync run that will get
executed immediately after the current run completes.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.TriggerRun(ctx, operations.TriggerRunRequest{
        SyncID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TriggerRunOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.TriggerRunRequest](../../pkg/models/operations/triggerrunrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.TriggerRunResponse](../../pkg/models/operations/triggerrunresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## TriggerRunCustom

Trigger a new run globally based on sync id or sync slug

If a run is already in progress, this queues a sync run that will get
executed immediately after the current run completes.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.TriggerRunCustom(ctx, shared.TriggerRunCustomInput{})
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.TriggerRunCustomInput](../../pkg/models/shared/triggerruncustominput.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.TriggerRunCustomResponse](../../pkg/models/operations/triggerruncustomresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## TriggerRunIDGraph

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.TriggerRunIDGraph(ctx, operations.TriggerRunIDGraphRequest{
        GraphID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TriggerRunIDGraphOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.TriggerRunIDGraphRequest](../../pkg/models/operations/triggerrunidgraphrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.TriggerRunIDGraphResponse](../../pkg/models/operations/triggerrunidgraphresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## TriggerSequenceRun

Trigger a new run for the given sync sequence.

If a run is already in progress, this queues a sync sequence run that will be
executed immediately after the current run completes.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.TriggerSequenceRun(ctx, operations.TriggerSequenceRunRequest{
        SyncSequenceID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TriggerSequenceRunOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.TriggerSequenceRunRequest](../../pkg/models/operations/triggersequencerunrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.TriggerSequenceRunResponse](../../pkg/models/operations/triggersequencerunresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## UpdateDestination

Update an existing destination

Patch a destination based on its Hightouch ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.UpdateDestination(ctx, operations.UpdateDestinationRequest{
        DestinationUpdate: shared.DestinationUpdate{},
        DestinationID: 8928.88,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpdateDestinationRequest](../../pkg/models/operations/updatedestinationrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpdateDestinationResponse](../../pkg/models/operations/updatedestinationresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## UpdateModel

Update an existing model

Patch a model based on its Hightouch ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.UpdateModel(ctx, operations.UpdateModelRequest{
        ModelUpdate: shared.ModelUpdate{},
        ModelID: 1027.03,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateModelRequest](../../pkg/models/operations/updatemodelrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.UpdateModelResponse](../../pkg/models/operations/updatemodelresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## UpdateSource

Update an existing source

Patch a source based on its Hightouch ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.UpdateSource(ctx, operations.UpdateSourceRequest{
        SourceUpdate: shared.SourceUpdate{},
        SourceID: 6585.68,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateSourceRequest](../../pkg/models/operations/updatesourcerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateSourceResponse](../../pkg/models/operations/updatesourceresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

## UpdateSync

Update an existing sync

Patch a sync based on its Hightouch ID

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"context"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.UpdateSync(ctx, operations.UpdateSyncRequest{
        SyncUpdate: shared.SyncUpdate{},
        SyncID: 5066.51,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateSyncRequest](../../pkg/models/operations/updatesyncrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UpdateSyncResponse](../../pkg/models/operations/updatesyncresponse.md), error**
| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 422                         | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |
