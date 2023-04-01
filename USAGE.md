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
    s := hightouch.New()

    req := shared.DestinationCreate{
        Configuration: map[string]interface{}{
            "provident": "distinctio",
            "quibusdam": "unde",
            "nulla": "corrupti",
        },
        Name: "illum",
        Slug: "vel",
        Type: "error",
    }

    ctx := context.Background()
    res, err := s.CreateDestination(ctx, req, operations.CreateDestinationSecurity{
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
<!-- End SDK Example Usage -->