# Hightouch SDK

## Overview

Hightouch Public Rest API to access syncs, models, sources and destinations

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
    }, operations.CreateDestinationSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDestination200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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
    }, operations.CreateModelSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateModel200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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
    }, operations.CreateSourceSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSource200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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
    }, operations.CreateSyncSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateSync200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.GetDestination(ctx, operations.GetDestinationRequest{
        DestinationID: 3864.89,
    }, operations.GetDestinationSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Destination != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.GetModel(ctx, operations.GetModelRequest{
        ModelID: 9437.49,
    }, operations.GetModelSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Model != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.GetSource(ctx, operations.GetSourceRequest{
        SourceID: 9025.99,
    }, operations.GetSourceSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Source != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.GetSync(ctx, operations.GetSyncRequest{
        SyncID: 6818.2,
    }, operations.GetSyncSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Sync != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.ListDestination(ctx, operations.ListDestinationRequest{
        Limit: hightouch.Float64(4499.5),
        Name: hightouch.String("Sheryl Kertzmann"),
        Offset: hightouch.Float64(992.8),
        OrderBy: operations.ListDestinationOrderByEnumID.ToPointer(),
        Slug: hightouch.String("reiciendis"),
    }, operations.ListDestinationSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListDestination200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.ListModel(ctx, operations.ListModelRequest{
        Limit: hightouch.Float64(6667.67),
        Name: hightouch.String("Cameron Dach"),
        Offset: hightouch.Float64(1289.26),
        OrderBy: operations.ListModelOrderByEnumCreatedAt.ToPointer(),
        Slug: hightouch.String("enim"),
    }, operations.ListModelSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListModel200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.ListSource(ctx, operations.ListSourceRequest{
        Limit: hightouch.Float64(6078.31),
        Name: hightouch.String("Ms. Cathy Marks"),
        Offset: hightouch.Float64(9883.74),
        OrderBy: operations.ListSourceOrderByEnumUpdatedAt.ToPointer(),
        Slug: hightouch.String("architecto"),
    }, operations.ListSourceSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSource200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.ListSync(ctx, operations.ListSyncRequest{
        After: types.MustTimeFromString("2022-08-01T12:28:44.292Z"),
        Before: types.MustTimeFromString("2022-09-05T05:51:25.673Z"),
        Limit: hightouch.Float64(9953),
        ModelID: hightouch.Float64(6531.08),
        Offset: hightouch.Float64(5818.5),
        OrderBy: operations.ListSyncOrderByEnumName.ToPointer(),
        Slug: hightouch.String("commodi"),
    }, operations.ListSyncSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSync200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.ListSyncRuns(ctx, operations.ListSyncRunsRequest{
        After: types.MustTimeFromString("2022-07-11T17:38:58.953Z"),
        Before: types.MustTimeFromString("2022-05-18T10:03:04.921Z"),
        Limit: hightouch.Float64(1589.69),
        Offset: hightouch.Float64(3380.07),
        OrderBy: operations.ListSyncRunsOrderByEnumID.ToPointer(),
        RunID: hightouch.Float64(6747.52),
        SyncID: 6563.3,
        Within: hightouch.Float64(3172.02),
    }, operations.ListSyncRunsSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ListSyncRuns200ApplicationJSONObject != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.TriggerRun(ctx, operations.TriggerRunRequest{
        TriggerRunInput: &shared.TriggerRunInput{
            FullResync: hightouch.Bool(false),
        },
        SyncID: "odit",
    }, operations.TriggerRunSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TriggerRunOutput != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.TriggerRunCustom(ctx, shared.TriggerRunCustomInput{
        FullResync: hightouch.Bool(false),
        SyncID: hightouch.String("quo"),
        SyncSlug: hightouch.String("sequi"),
    }, operations.TriggerRunCustomSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TriggerRunCustom200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.UpdateDestination(ctx, operations.UpdateDestinationRequest{
        DestinationUpdate: shared.DestinationUpdate{
            Configuration: map[string]interface{}{
                "ipsam": "id",
                "possimus": "aut",
                "quasi": "error",
                "temporibus": "laborum",
            },
            Name: hightouch.String("Johanna Wolf"),
        },
        DestinationID: 5096.24,
    }, operations.UpdateDestinationSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDestination200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.UpdateModel(ctx, operations.UpdateModelRequest{
        ModelUpdate: shared.ModelUpdate{
            Custom: &shared.ModelUpdateCustom{
                Query: "voluptatibus",
            },
            Dbt: &shared.ModelUpdateDbt{
                ModelID: "ipsa",
            },
            IsSchema: hightouch.Bool(false),
            Name: hightouch.String("Mr. Jared Ritchie"),
            PrimaryKey: hightouch.String("ut"),
            Raw: &shared.ModelUpdateRaw{
                SQL: "maiores",
            },
            Table: &shared.ModelUpdateTable{
                Name: "Stacy Gulgowski MD",
            },
            Visual: &shared.ModelUpdateVisual{
                Filter: "enim",
                ParentID: "accusamus",
                PrimaryLabel: "commodi",
                SecondaryLabel: "repudiandae",
            },
        },
        ModelID: 641.47,
    }, operations.UpdateModelSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateModel200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.UpdateSource(ctx, operations.UpdateSourceRequest{
        SourceUpdate: shared.SourceUpdate{
            Configuration: map[string]interface{}{
                "quidem": "molestias",
            },
            Name: hightouch.String("Ervin Gleason"),
        },
        SourceID: 9167.23,
    }, operations.UpdateSourceSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSource200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```

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

    ctx := context.Background()
    res, err := s.Hightouch.UpdateSync(ctx, operations.UpdateSyncRequest{
        SyncUpdate: shared.SyncUpdate{
            Configuration: map[string]interface{}{
                "repudiandae": "sint",
            },
            Disabled: hightouch.Bool(false),
            Schedule: &shared.SyncUpdateSchedule{
                Schedule: shared.IntervalSchedule{
                    Interval: shared.Interval{
                        Quantity: 9292.97,
                        Unit: shared.IntervalUnitEnumHour,
                    },
                },
                Type: "enim",
            },
        },
        SyncID: 93.56,
    }, operations.UpdateSyncSecurity{
        BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateSync200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```
