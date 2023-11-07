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
	res, err := s.CreateDestination(ctx, shared.DestinationCreate{
		Configuration: map[string]interface{}{
			"key": "string",
		},
		Name: "string",
		Slug: "string",
		Type: "string",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.OneOf != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->