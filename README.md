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

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/hightouch-go-sdk
```
<!-- End SDK Installation -->

## Authentication

- Create an [API key](https://app.hightouch.com/settings/api-keys)
- From the API keys tab on the Settings page, select Add API key.
- Enter a descriptive Name for your key.
- Copy your API key and store it in a safe location. The key will only be displayed once.
- Click Create API key.

## SDK Example Usage
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
        BearerAuth: "YOUR_BEARER_TOKEN_HERE",
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [Hightouch SDK](docs/hightouch/README.md)

* [CreateDestination](docs/hightouch/README.md#createdestination) - Create Destination
* [CreateModel](docs/hightouch/README.md#createmodel) - Create Model
* [CreateSource](docs/hightouch/README.md#createsource) - Create Source
* [CreateSync](docs/hightouch/README.md#createsync) - Create Sync
* [GetDestination](docs/hightouch/README.md#getdestination) - Get Destination
* [GetModel](docs/hightouch/README.md#getmodel) - Get Model
* [GetSource](docs/hightouch/README.md#getsource) - Get Source
* [GetSync](docs/hightouch/README.md#getsync) - Get Sync
* [ListDestination](docs/hightouch/README.md#listdestination) - List Destinations
* [ListModel](docs/hightouch/README.md#listmodel) - List Models
* [ListSource](docs/hightouch/README.md#listsource) - List Sources
* [ListSync](docs/hightouch/README.md#listsync) - List Syncs
* [ListSyncRuns](docs/hightouch/README.md#listsyncruns) - List Sync Runs
* [TriggerRun](docs/hightouch/README.md#triggerrun) - Trigger Sync
* [TriggerRunCustom](docs/hightouch/README.md#triggerruncustom) - Trigger Sync From ID or Slug
* [UpdateDestination](docs/hightouch/README.md#updatedestination) - Update Destination
* [UpdateModel](docs/hightouch/README.md#updatemodel) - Update Model
* [UpdateSource](docs/hightouch/README.md#updatesource) - Update Source
* [UpdateSync](docs/hightouch/README.md#updatesync) - Update Sync
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
