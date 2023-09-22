<!-- Start SDK Example Usage -->


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
    res, err := s.CreateDestination(ctx, shared.DestinationCreate{
        Configuration: map[string]interface{}{
            "corrupti": "provident",
        },
        Name: "Ellis Mitchell",
        Slug: "illum",
        Type: "vel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDestination200ApplicationJSONOneOf != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->