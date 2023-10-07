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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.CreateDestination(ctx, shared.DestinationCreate{
        AdditionalProperties: map[string]interface{}{
            "optical": "Sid",
        },
        Configuration: map[string]interface{}{
            "kilogram": "Southwest",
        },
        Name: "blue wilt",
        Slug: "Southeast Mobility",
        Type: "Jaguar Ford",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDestination200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.DestinationCreate](../../models/shared/destinationcreate.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.CreateModel(ctx, shared.ModelCreate{
        AdditionalProperties: map[string]interface{}{
            "zesty": "veritatis",
        },
        Custom: &shared.ModelCreateCustom{
            Query: "Specialist",
        },
        Dbt: &shared.ModelCreateDbt{
            ModelID: "violet digital Data",
        },
        IsSchema: false,
        Name: "nor North",
        PrimaryKey: "solid success Rochester",
        QueryType: "if East",
        Raw: &shared.ModelCreateRaw{
            SQL: "vortals Vanadium Division",
        },
        Slug: "application unloosen yahoo",
        SourceID: "Practical Bike",
        Table: &shared.ModelCreateTable{
            Name: "Kentucky invoice South",
        },
        Visual: &shared.ModelCreateVisual{
            Filter: "precedent",
            ParentID: "Keebler instead",
            PrimaryLabel: "Saint Brookline Bacon",
            SecondaryLabel: "owlishly Principal Music",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateModel200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [shared.ModelCreate](../../models/shared/modelcreate.md) | :heavy_check_mark:                                       | The request object to use for the request.               |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.CreateSource(ctx, shared.SourceCreate{
        AdditionalProperties: map[string]interface{}{
            "North": "tightly",
        },
        Configuration: map[string]interface{}{
            "Van": "deposit",
        },
        Name: "matrix York ouch",
        Slug: "Architect",
        Type: "gold Baby Interactions",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSource200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.SourceCreate](../../models/shared/sourcecreate.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.CreateSync(ctx, shared.SyncCreate{
        AdditionalProperties: map[string]interface{}{
            "euphoric": "experiences",
        },
        Configuration: map[string]interface{}{
            "North": "synergies",
        },
        DestinationID: "ranch",
        Disabled: false,
        ModelID: "International",
        Schedule: &shared.SyncCreateSchedule{
            Schedule: shared.CreateSyncCreateScheduleScheduleIntervalSchedule(
                    shared.IntervalSchedule{
                        AdditionalProperties: map[string]interface{}{
                            "mobile": "copying",
                        },
                        Interval: shared.Interval{
                            AdditionalProperties: map[string]interface{}{
                                "database": "meanwhile",
                            },
                            Quantity: 5827.46,
                            Unit: shared.IntervalUnitWeek,
                        },
                    },
            ),
            Type: "SDD",
        },
        Slug: "Fremont",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSync200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.SyncCreate](../../models/shared/synccreate.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.GetDestination(ctx, operations.GetDestinationRequest{
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetDestinationRequest](../../models/operations/getdestinationrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.GetModel(ctx, operations.GetModelRequest{
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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetModelRequest](../../models/operations/getmodelrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.GetSource(ctx, operations.GetSourceRequest{
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetSourceRequest](../../models/operations/getsourcerequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.GetSync(ctx, operations.GetSyncRequest{
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

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetSyncRequest](../../models/operations/getsyncrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.ListDestination(ctx, operations.ListDestinationRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDestination200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListDestinationRequest](../../models/operations/listdestinationrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.ListModel(ctx, operations.ListModelRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ListModel200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.ListModelRequest](../../models/operations/listmodelrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.ListSource(ctx, operations.ListSourceRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSource200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListSourceRequest](../../models/operations/listsourcerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.ListSync(ctx, operations.ListSyncRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSync200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ListSyncRequest](../../models/operations/listsyncrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.ListSyncRuns(ctx, operations.ListSyncRunsRequest{
        SyncID: 8858.62,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSyncRuns200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListSyncRunsRequest](../../models/operations/listsyncrunsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.TriggerRun(ctx, operations.TriggerRunRequest{
        TriggerRunInput: &shared.TriggerRunInput{},
        SyncID: "Digitized withdrawal Midland",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.TriggerRunRequest](../../models/operations/triggerrunrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.TriggerRunCustom(ctx, shared.TriggerRunCustomInput{})
    if err != nil {
        log.Fatal(err)
    }

    if res.TriggerRunCustom200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.TriggerRunCustomInput](../../models/shared/triggerruncustominput.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.TriggerRunCustomResponse](../../models/operations/triggerruncustomresponse.md), error**


## TriggerRunIDGraph

### Example Usage

```go
package main

import(
	"context"
	"log"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.TriggerRunIDGraph(ctx, operations.TriggerRunIDGraphRequest{
        TriggerRunIDGraphInput: &shared.TriggerRunIDGraphInput{},
        GraphID: "hop Bentley AI",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.TriggerRunIDGraphRequest](../../models/operations/triggerrunidgraphrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.UpdateDestination(ctx, operations.UpdateDestinationRequest{
        DestinationUpdate: shared.DestinationUpdate{
            AdditionalProperties: map[string]interface{}{
                "nervously": "enable",
            },
            Configuration: map[string]interface{}{
                "Northwest": "woot",
            },
        },
        DestinationID: 8585.6,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDestination200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateDestinationRequest](../../models/operations/updatedestinationrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.UpdateModel(ctx, operations.UpdateModelRequest{
        ModelUpdate: shared.ModelUpdate{
            AdditionalProperties: map[string]interface{}{
                "Northwest": "gold",
            },
            Custom: &shared.ModelUpdateCustom{
                Query: "administration",
            },
            Dbt: &shared.ModelUpdateDbt{
                ModelID: "Xenogender East flexibility",
            },
            Raw: &shared.ModelUpdateRaw{
                SQL: "content Chair reinvent",
            },
            Table: &shared.ModelUpdateTable{
                Name: "minus Gasoline Rhodium",
            },
            Visual: &shared.ModelUpdateVisual{
                Filter: "mint",
                ParentID: "Madagascar",
                PrimaryLabel: "stimulating alarm",
                SecondaryLabel: "Mazda",
            },
        },
        ModelID: 8905.93,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateModel200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpdateModelRequest](../../models/operations/updatemodelrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.UpdateSource(ctx, operations.UpdateSourceRequest{
        SourceUpdate: shared.SourceUpdate{
            AdditionalProperties: map[string]interface{}{
                "Internal": "esse",
            },
            Configuration: map[string]interface{}{
                "selfish": "female",
            },
        },
        SourceID: 1942.93,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSource200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateSourceRequest](../../models/operations/updatesourcerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


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
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := hightouchgosdk.New(
        hightouchgosdk.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Hightouch.UpdateSync(ctx, operations.UpdateSyncRequest{
        SyncUpdate: shared.SyncUpdate{
            AdditionalProperties: map[string]interface{}{
                "hacking": "Keyboard",
            },
            Configuration: map[string]interface{}{
                "deposit": "Bulgarian",
            },
            Schedule: &shared.SyncUpdateSchedule{
                Schedule: shared.CreateSyncUpdateScheduleScheduleVisualCronSchedule(
                        shared.VisualCronSchedule{
                            AdditionalProperties: map[string]interface{}{
                                "York": "intensely",
                            },
                            Expressions: []shared.VisualCronScheduleExpressions{
                                shared.VisualCronScheduleExpressions{
                                    Days: shared.RecordDayBooleanOrUndefined{},
                                    Time: "Buckinghamshire",
                                },
                            },
                        },
                ),
                Type: "Convertible",
            },
        },
        SyncID: 8847.7,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSync200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpdateSyncRequest](../../models/operations/updatesyncrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UpdateSyncResponse](../../models/operations/updatesyncresponse.md), error**

