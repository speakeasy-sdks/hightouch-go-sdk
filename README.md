# github.com/speakeasy-sdks/hightouch-go-sdk

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/hightouch-go-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/hightouch-go-sdk"
    "github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/operations"
)

func main() {
    s := sdk.New()
    
    req := operations.CreateDestinationRequest{
        Security: operations.CreateDestinationSecurity{
            BearerAuth: shared.SchemeBearerAuth{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        },
        Request: shared.DestinationCreate{
            Configuration: map[string]interface{}{
                "deserunt": "porro",
                "nulla": "id",
                "vero": "perspiciatis",
            },
            Name: "nulla",
            Slug: "nihil",
            Type: "fuga",
        },
    }

    ctx := context.Background()
    res, err := s.CreateDestination(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDestination200ApplicationJSONAnyOf != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations

### SDK SDK

* `CreateDestination` - Create Destination
* `CreateModel` - Create Model
* `CreateSource` - Create Source
* `CreateSync` - Create Sync
* `GetDestination` - Get Destination
* `GetModel` - Get Model
* `GetSource` - Get Source
* `GetSync` - Get Sync
* `ListDestination` - List Destinations
* `ListModel` - List Models
* `ListSource` - List Sources
* `ListSync` - List Syncs
* `ListSyncRuns` - List Sync Runs
* `TriggerRun` - Trigger Sync
* `TriggerRunCustom` - Trigger Sync From ID or Slug
* `UpdateDestination` - Update Destination
* `UpdateModel` - Update Model
* `UpdateSource` - Update Source
* `UpdateSync` - Update Sync
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)