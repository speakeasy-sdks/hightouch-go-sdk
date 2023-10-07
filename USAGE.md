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
<!-- End SDK Example Usage -->