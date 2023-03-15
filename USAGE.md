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