package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateSourceSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type UpdateSourcePathParams struct {
	SourceID float64 `pathParam:"style=simple,explode=false,name=sourceId"`
}

type UpdateSourceRequest struct {
	PathParams UpdateSourcePathParams
	Request    shared.SourceUpdate `request:"mediaType=application/json"`
	Security   UpdateSourceSecurity
}

type UpdateSourceResponse struct {
	ContentType                         string
	InternalServerError                 *shared.InternalServerErrorEnum
	StatusCode                          int
	RawResponse                         *http.Response
	UpdateSource200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                   *shared.ValidateErrorJSON
}
