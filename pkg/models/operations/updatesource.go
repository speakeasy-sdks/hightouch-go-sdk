package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateSourceSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type UpdateSourceRequest struct {
	SourceUpdate shared.SourceUpdate `request:"mediaType=application/json"`
	SourceID     float64             `pathParam:"style=simple,explode=false,name=sourceId"`
}

type UpdateSourceResponse struct {
	ContentType                         string
	InternalServerError                 *shared.InternalServerErrorEnum
	StatusCode                          int
	RawResponse                         *http.Response
	UpdateSource200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                   *shared.ValidateErrorJSON
}
