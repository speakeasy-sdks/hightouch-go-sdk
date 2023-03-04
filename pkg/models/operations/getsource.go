package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetSourceSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type GetSourcePathParams struct {
	SourceID float64 `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceRequest struct {
	PathParams GetSourcePathParams
	Security   GetSourceSecurity
}

type GetSourceResponse struct {
	ContentType       string
	Source            *shared.Source
	StatusCode        int
	RawResponse       *http.Response
	ValidateErrorJSON *shared.ValidateErrorJSON
}
