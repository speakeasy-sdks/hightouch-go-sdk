<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := hightouchgosdk.New(
		hightouchgosdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Hightouch.CreateDestination(ctx, shared.DestinationCreate{
		Configuration: map[string]interface{}{
			"optical": "Sid",
		},
		Name: "Ergonomic mumble blue",
		Slug: "hollow Southeast Mobility",
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