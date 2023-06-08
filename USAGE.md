<!-- Start SDK Example Usage -->
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
    res, err := s.CreateDestination(ctx, shared.DestinationCreate{
        Configuration: map[string]interface{}{
            "provident": "distinctio",
            "quibusdam": "unde",
            "nulla": "corrupti",
        },
        Name: "Ben Mueller",
        Slug: "iure",
        Type: "magnam",
    }, operations.CreateDestinationSecurity{
        BearerAuth: "",
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