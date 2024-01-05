<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/221538824-87af0e1b-0508-4af5-b3b9-e4b192d8337f.svg" media="(prefers-color-scheme: dark)">
        <img src="https://user-images.githubusercontent.com/6267663/221538828-de1343f2-b249-4ba2-85e3-a2e43cc5f265.svg">
    </picture>
    <h1>Hightouch Go SDK</h1>
   <p>Hightouch exposes a REST API that lets users interact with resources like syncs, models, sources and destinations.</p>
   <a href="https://hightouch.com/docs/api-reference"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/hightouch-go-sdk/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/hightouch-go-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/hightouch-go-sdk/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/hightouch-go-sdk?sort=semver&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/hightouch-go-sdk
```
<!-- End SDK Installation [installation] -->

## Authentication

- Create an [API key](https://app.hightouch.com/settings/api-keys)
- From the API keys tab on the Settings page, select Add API key.
- Enter a descriptive Name for your key.
- Copy your API key and store it in a safe location. The key will only be displayed once.
- Click Create API key.

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Hightouch SDK](docs/sdks/hightouch/README.md)

* [CreateDestination](docs/sdks/hightouch/README.md#createdestination) - Create Destination
* [CreateModel](docs/sdks/hightouch/README.md#createmodel) - Create Model
* [CreateSource](docs/sdks/hightouch/README.md#createsource) - Create Source
* [CreateSync](docs/sdks/hightouch/README.md#createsync) - Create Sync
* [GetDestination](docs/sdks/hightouch/README.md#getdestination) - Get Destination
* [GetModel](docs/sdks/hightouch/README.md#getmodel) - Get Model
* [GetSource](docs/sdks/hightouch/README.md#getsource) - Get Source
* [GetSync](docs/sdks/hightouch/README.md#getsync) - Get Sync
* [GetSyncSequenceRun](docs/sdks/hightouch/README.md#getsyncsequencerun) - Sync sequence status
* [ListDestination](docs/sdks/hightouch/README.md#listdestination) - List Destinations
* [ListModel](docs/sdks/hightouch/README.md#listmodel) - List Models
* [ListSource](docs/sdks/hightouch/README.md#listsource) - List Sources
* [ListSync](docs/sdks/hightouch/README.md#listsync) - List Syncs
* [ListSyncRuns](docs/sdks/hightouch/README.md#listsyncruns) - List Sync Runs
* [TriggerRun](docs/sdks/hightouch/README.md#triggerrun) - Trigger Sync
* [TriggerRunCustom](docs/sdks/hightouch/README.md#triggerruncustom) - Trigger Sync From ID or Slug
* [TriggerRunIDGraph](docs/sdks/hightouch/README.md#triggerrunidgraph)
* [TriggerSequenceRun](docs/sdks/hightouch/README.md#triggersequencerun) - Trigger Sync sequence
* [UpdateDestination](docs/sdks/hightouch/README.md#updatedestination) - Update Destination
* [UpdateModel](docs/sdks/hightouch/README.md#updatemodel) - Update Model
* [UpdateSource](docs/sdks/hightouch/README.md#updatesource) - Update Source
* [UpdateSync](docs/sdks/hightouch/README.md#updatesync) - Update Sync
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| sdkerrors.ValidateErrorJSON | 409,422                     | application/json            |
| sdkerrors.SDKError          | 4xx-5xx                     | */*                         |

### Example

```go
package main

import (
	"context"
	"errors"
	hightouchgosdk "github.com/speakeasy-sdks/hightouch-go-sdk"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/sdkerrors"
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
			"key": "string",
		},
		Name: "string",
		Slug: "string",
		Type: "string",
	})
	if err != nil {

		var e *sdkerrors.ValidateErrorJSON
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.hightouch.com/api/v1` | None |

#### Example

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
		hightouchgosdk.WithServerIndex(0),
		hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		hightouchgosdk.WithServerURL("https://api.hightouch.com/api/v1"),
		hightouchgosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
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
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `BearerAuth` | http         | HTTP Bearer  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
