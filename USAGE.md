<!-- Start SDK Example Usage [usage] -->
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
		hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.CreateDestination(ctx, shared.DestinationCreate{
		Configuration: map[string]interface{}{
			"key": "<value>",
		},
		Name: "<value>",
		Slug: "<value>",
		Type: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.OneOf != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->