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
* [ListDestination](#listdestination) - List Destinations
* [ListModel](#listmodel) - List Models
* [ListSource](#listsource) - List Sources
* [ListSync](#listsync) - List Syncs
* [ListSyncRuns](#listsyncruns) - List Sync Runs
* [TriggerRun](#triggerrun) - Trigger Sync
* [TriggerRunCustom](#triggerruncustom) - Trigger Sync From ID or Slug
* [TriggerRunIDGraph](#triggerrunidgraph)
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
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.CreateDestinationSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.CreateDestination(ctx, shared.DestinationCreate{
        Configuration: map[string]interface{}{
            "ipsa": "delectus",
            "tempora": "suscipit",
            "molestiae": "minus",
            "placeat": "voluptatum",
        },
        Name: "Miriam Huel",
        Slug: "ab",
        Type: "quis",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDestination200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.DestinationCreate](../../models/shared/destinationcreate.md)                         | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.CreateDestinationSecurity](../../models/operations/createdestinationsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.CreateDestinationResponse](../../models/operations/createdestinationresponse.md), error**


## CreateModel

Create a new model

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.CreateModelSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.CreateModel(ctx, shared.ModelCreate{
        Custom: &shared.ModelCreateCustom{
            Query: "veritatis",
        },
        Dbt: &shared.ModelCreateDbt{
            ModelID: "deserunt",
        },
        IsSchema: false,
        Name: "Roberta Sipes",
        PrimaryKey: "odit",
        QueryType: "at",
        Raw: &shared.ModelCreateRaw{
            SQL: "at",
        },
        Slug: "maiores",
        SourceID: "molestiae",
        Table: &shared.ModelCreateTable{
            Name: "Forrest Koepp",
        },
        Visual: &shared.ModelCreateVisual{
            Filter: "dolorum",
            ParentID: "dicta",
            PrimaryLabel: "nam",
            SecondaryLabel: "officia",
        },
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateModel200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.ModelCreate](../../models/shared/modelcreate.md)                         | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.CreateModelSecurity](../../models/operations/createmodelsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.CreateModelResponse](../../models/operations/createmodelresponse.md), error**


## CreateSource

Create a new source

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.CreateSourceSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.CreateSource(ctx, shared.SourceCreate{
        Configuration: map[string]interface{}{
            "fugit": "deleniti",
            "hic": "optio",
            "totam": "beatae",
        },
        Name: "Tanya Gleason",
        Slug: "cum",
        Type: "esse",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSource200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [shared.SourceCreate](../../models/shared/sourcecreate.md)                         | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CreateSourceSecurity](../../models/operations/createsourcesecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.CreateSourceResponse](../../models/operations/createsourceresponse.md), error**


## CreateSync

Create a new sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.CreateSyncSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.CreateSync(ctx, shared.SyncCreate{
        Configuration: map[string]interface{}{
            "excepturi": "aspernatur",
        },
        DestinationID: "perferendis",
        Disabled: false,
        ModelID: "ad",
        Schedule: shared.SyncCreateSchedule{
            Schedule: shared.VisualCronSchedule{
                Expressions: []shared.VisualCronScheduleExpressions{
                    shared.VisualCronScheduleExpressions{
                        Days: shared.RecordDayBooleanOrUndefined{
                            Friday: hightouch.Bool(false),
                            Monday: hightouch.Bool(false),
                            Saturday: hightouch.Bool(false),
                            Sunday: hightouch.Bool(false),
                            Thursday: hightouch.Bool(false),
                            Tuesday: hightouch.Bool(false),
                            Wednesday: hightouch.Bool(false),
                        },
                        Time: "iste",
                    },
                },
            },
            Type: "dolor",
        },
        Slug: "natus",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSync200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.SyncCreate](../../models/shared/synccreate.md)                         | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.CreateSyncSecurity](../../models/operations/createsyncsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.CreateSyncResponse](../../models/operations/createsyncresponse.md), error**


## GetDestination

Retrieve a destination based on its Hightouch ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.GetDestinationSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.GetDestination(ctx, operations.GetDestinationRequest{
        DestinationID: 3864.89,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Destination != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDestinationRequest](../../models/operations/getdestinationrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.GetDestinationSecurity](../../models/operations/getdestinationsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.GetDestinationResponse](../../models/operations/getdestinationresponse.md), error**


## GetModel

Retrieve models from model ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.GetModelSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.GetModel(ctx, operations.GetModelRequest{
        ModelID: 9437.49,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Model != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetModelRequest](../../models/operations/getmodelrequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.GetModelSecurity](../../models/operations/getmodelsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.GetModelResponse](../../models/operations/getmodelresponse.md), error**


## GetSource

Retrieve source from source ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.GetSourceSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.GetSource(ctx, operations.GetSourceRequest{
        SourceID: 9025.99,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Source != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetSourceRequest](../../models/operations/getsourcerequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.GetSourceSecurity](../../models/operations/getsourcesecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.GetSourceResponse](../../models/operations/getsourceresponse.md), error**


## GetSync

Retrieve sync from sync ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.GetSyncSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.GetSync(ctx, operations.GetSyncRequest{
        SyncID: 6818.2,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.Sync != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetSyncRequest](../../models/operations/getsyncrequest.md)   | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [operations.GetSyncSecurity](../../models/operations/getsyncsecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |


### Response

**[*operations.GetSyncResponse](../../models/operations/getsyncresponse.md), error**


## ListDestination

List the destinations in the user's workspace

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.ListDestinationSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.ListDestination(ctx, operations.ListDestinationRequest{
        Limit: hightouch.Float64(4499.5),
        Name: hightouch.String("Sheryl Kertzmann"),
        Offset: hightouch.Float64(992.8),
        OrderBy: operations.ListDestinationOrderByID.ToPointer(),
        Slug: hightouch.String("reiciendis"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDestination200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListDestinationRequest](../../models/operations/listdestinationrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ListDestinationSecurity](../../models/operations/listdestinationsecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ListDestinationResponse](../../models/operations/listdestinationresponse.md), error**


## ListModel

List all the models in the current workspace including parent and related models

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.ListModelSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.ListModel(ctx, operations.ListModelRequest{
        Limit: hightouch.Float64(6667.67),
        Name: hightouch.String("Cameron Dach"),
        Offset: hightouch.Float64(1289.26),
        OrderBy: operations.ListModelOrderByCreatedAt.ToPointer(),
        Slug: hightouch.String("enim"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListModel200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListModelRequest](../../models/operations/listmodelrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.ListModelSecurity](../../models/operations/listmodelsecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.ListModelResponse](../../models/operations/listmodelresponse.md), error**


## ListSource

List all the sources in the current workspace

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.ListSourceSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.ListSource(ctx, operations.ListSourceRequest{
        Limit: hightouch.Float64(6078.31),
        Name: hightouch.String("Ms. Cathy Marks"),
        Offset: hightouch.Float64(9883.74),
        OrderBy: operations.ListSourceOrderByUpdatedAt.ToPointer(),
        Slug: hightouch.String("architecto"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSource200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListSourceRequest](../../models/operations/listsourcerequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.ListSourceSecurity](../../models/operations/listsourcesecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.ListSourceResponse](../../models/operations/listsourceresponse.md), error**


## ListSync

List all the syncs in the current workspace

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/types"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.ListSyncSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.ListSync(ctx, operations.ListSyncRequest{
        After: types.MustTimeFromString("2022-08-01T12:28:44.292Z"),
        Before: types.MustTimeFromString("2022-09-05T05:51:25.673Z"),
        Limit: hightouch.Float64(9953),
        ModelID: hightouch.Float64(6531.08),
        Offset: hightouch.Float64(5818.5),
        OrderBy: operations.ListSyncOrderByName.ToPointer(),
        Slug: hightouch.String("commodi"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSync200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListSyncRequest](../../models/operations/listsyncrequest.md)   | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `security`                                                                 | [operations.ListSyncSecurity](../../models/operations/listsyncsecurity.md) | :heavy_check_mark:                                                         | The security requirements to use for the request.                          |


### Response

**[*operations.ListSyncResponse](../../models/operations/listsyncresponse.md), error**


## ListSyncRuns

List all sync runs under a sync

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/types"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.ListSyncRunsSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.ListSyncRuns(ctx, operations.ListSyncRunsRequest{
        After: types.MustTimeFromString("2022-07-11T17:38:58.953Z"),
        Before: types.MustTimeFromString("2022-05-18T10:03:04.921Z"),
        Limit: hightouch.Float64(1589.69),
        Offset: hightouch.Float64(3380.07),
        OrderBy: operations.ListSyncRunsOrderByID.ToPointer(),
        RunID: hightouch.Float64(6747.52),
        SyncID: 6563.3,
        Within: hightouch.Float64(3172.02),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSyncRuns200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListSyncRunsRequest](../../models/operations/listsyncrunsrequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.ListSyncRunsSecurity](../../models/operations/listsyncrunssecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.ListSyncRunsResponse](../../models/operations/listsyncrunsresponse.md), error**


## TriggerRun

Trigger a new run for the given sync.

If a run is already in progress, this queues a sync run that will get
executed immediately after the current run completes.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.TriggerRunSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.TriggerRun(ctx, operations.TriggerRunRequest{
        TriggerRunInput: &shared.TriggerRunInput{
            FullResync: hightouch.Bool(false),
        },
        SyncID: "odit",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TriggerRunOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.TriggerRunRequest](../../models/operations/triggerrunrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.TriggerRunSecurity](../../models/operations/triggerrunsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.TriggerRunResponse](../../models/operations/triggerrunresponse.md), error**


## TriggerRunCustom

Trigger a new run globally based on sync id or sync slug

If a run is already in progress, this queues a sync run that will get
executed immediately after the current run completes.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.TriggerRunCustomSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.TriggerRunCustom(ctx, shared.TriggerRunCustomInput{
        FullResync: hightouch.Bool(false),
        SyncID: hightouch.String("quo"),
        SyncSlug: hightouch.String("sequi"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TriggerRunCustom200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.TriggerRunCustomInput](../../models/shared/triggerruncustominput.md)               | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.TriggerRunCustomSecurity](../../models/operations/triggerruncustomsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.TriggerRunCustomResponse](../../models/operations/triggerruncustomresponse.md), error**


## TriggerRunIDGraph

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.TriggerRunIDGraphSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.TriggerRunIDGraph(ctx, operations.TriggerRunIDGraphRequest{
        TriggerRunIDGraphInput: &shared.TriggerRunIDGraphInput{
            FullRerun: hightouch.Bool(false),
        },
        GraphID: "tenetur",
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TriggerRunIDGraphOutput != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.TriggerRunIDGraphRequest](../../models/operations/triggerrunidgraphrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.TriggerRunIDGraphSecurity](../../models/operations/triggerrunidgraphsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.TriggerRunIDGraphResponse](../../models/operations/triggerrunidgraphresponse.md), error**


## UpdateDestination

Update an existing destination

Patch a destination based on its Hightouch ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.UpdateDestinationSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.UpdateDestination(ctx, operations.UpdateDestinationRequest{
        DestinationUpdate: shared.DestinationUpdate{
            Configuration: map[string]interface{}{
                "id": "possimus",
                "aut": "quasi",
            },
            Name: hightouch.String("Dr. Jake Pacocha"),
        },
        DestinationID: 8781.94,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDestination200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateDestinationRequest](../../models/operations/updatedestinationrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.UpdateDestinationSecurity](../../models/operations/updatedestinationsecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.UpdateDestinationResponse](../../models/operations/updatedestinationresponse.md), error**


## UpdateModel

Update an existing model

Patch a model based on its Hightouch ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.UpdateModelSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.UpdateModel(ctx, operations.UpdateModelRequest{
        ModelUpdate: shared.ModelUpdate{
            Custom: &shared.ModelUpdateCustom{
                Query: "nihil",
            },
            Dbt: &shared.ModelUpdateDbt{
                ModelID: "praesentium",
            },
            IsSchema: hightouch.Bool(false),
            Name: hightouch.String("Jose Moen"),
            PrimaryKey: hightouch.String("perferendis"),
            Raw: &shared.ModelUpdateRaw{
                SQL: "doloremque",
            },
            Table: &shared.ModelUpdateTable{
                Name: "Mrs. April Wuckert",
            },
            Visual: &shared.ModelUpdateVisual{
                Filter: "iusto",
                ParentID: "dicta",
                PrimaryLabel: "harum",
                SecondaryLabel: "enim",
            },
        },
        ModelID: 8804.76,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateModel200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateModelRequest](../../models/operations/updatemodelrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.UpdateModelSecurity](../../models/operations/updatemodelsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.UpdateModelResponse](../../models/operations/updatemodelresponse.md), error**


## UpdateSource

Update an existing source

Patch a source based on its Hightouch ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.UpdateSourceSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.UpdateSource(ctx, operations.UpdateSourceRequest{
        SourceUpdate: shared.SourceUpdate{
            Configuration: map[string]interface{}{
                "repudiandae": "quae",
                "ipsum": "quidem",
            },
            Name: hightouch.String("Andy Streich"),
        },
        SourceID: 5232.48,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSource200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateSourceRequest](../../models/operations/updatesourcerequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.UpdateSourceSecurity](../../models/operations/updatesourcesecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.UpdateSourceResponse](../../models/operations/updatesourceresponse.md), error**


## UpdateSync

Update an existing sync

Patch a sync based on its Hightouch ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouch.New()
    operationSecurity := operations.UpdateSyncSecurity{
            BearerAuth: "",
        }

    ctx := context.Background()
    res, err := s.Hightouch.UpdateSync(ctx, operations.UpdateSyncRequest{
        SyncUpdate: shared.SyncUpdate{
            Configuration: map[string]interface{}{
                "quasi": "repudiandae",
                "sint": "veritatis",
                "itaque": "incidunt",
                "enim": "consequatur",
            },
            Disabled: hightouch.Bool(false),
            Schedule: &shared.SyncUpdateSchedule{
                Schedule: shared.VisualCronSchedule{
                    Expressions: []shared.VisualCronScheduleExpressions{
                        shared.VisualCronScheduleExpressions{
                            Days: shared.RecordDayBooleanOrUndefined{
                                Friday: hightouch.Bool(false),
                                Monday: hightouch.Bool(false),
                                Saturday: hightouch.Bool(false),
                                Sunday: hightouch.Bool(false),
                                Thursday: hightouch.Bool(false),
                                Tuesday: hightouch.Bool(false),
                                Wednesday: hightouch.Bool(false),
                            },
                            Time: "explicabo",
                        },
                        shared.VisualCronScheduleExpressions{
                            Days: shared.RecordDayBooleanOrUndefined{
                                Friday: hightouch.Bool(false),
                                Monday: hightouch.Bool(false),
                                Saturday: hightouch.Bool(false),
                                Sunday: hightouch.Bool(false),
                                Thursday: hightouch.Bool(false),
                                Tuesday: hightouch.Bool(false),
                                Wednesday: hightouch.Bool(false),
                            },
                            Time: "deserunt",
                        },
                        shared.VisualCronScheduleExpressions{
                            Days: shared.RecordDayBooleanOrUndefined{
                                Friday: hightouch.Bool(false),
                                Monday: hightouch.Bool(false),
                                Saturday: hightouch.Bool(false),
                                Sunday: hightouch.Bool(false),
                                Thursday: hightouch.Bool(false),
                                Tuesday: hightouch.Bool(false),
                                Wednesday: hightouch.Bool(false),
                            },
                            Time: "distinctio",
                        },
                        shared.VisualCronScheduleExpressions{
                            Days: shared.RecordDayBooleanOrUndefined{
                                Friday: hightouch.Bool(false),
                                Monday: hightouch.Bool(false),
                                Saturday: hightouch.Bool(false),
                                Sunday: hightouch.Bool(false),
                                Thursday: hightouch.Bool(false),
                                Tuesday: hightouch.Bool(false),
                                Wednesday: hightouch.Bool(false),
                            },
                            Time: "quibusdam",
                        },
                    },
                },
                Type: "labore",
            },
        },
        SyncID: 2647.3,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSync200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpdateSyncRequest](../../models/operations/updatesyncrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.UpdateSyncSecurity](../../models/operations/updatesyncsecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.UpdateSyncResponse](../../models/operations/updatesyncresponse.md), error**

