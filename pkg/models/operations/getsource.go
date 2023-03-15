package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetSourceSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetSourceRequest struct {
	SourceID float64 `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceResponse struct {
	ContentType       string
	Source            *shared.Source
	StatusCode        int
	RawResponse       *http.Response
	ValidateErrorJSON *shared.ValidateErrorJSON
}
