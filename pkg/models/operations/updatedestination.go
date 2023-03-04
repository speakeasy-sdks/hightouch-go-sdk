package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateDestinationSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type UpdateDestinationPathParams struct {
	DestinationID float64 `pathParam:"style=simple,explode=false,name=destinationId"`
}

type UpdateDestinationRequest struct {
	PathParams UpdateDestinationPathParams
	Request    shared.DestinationUpdate `request:"mediaType=application/json"`
	Security   UpdateDestinationSecurity
}

type UpdateDestinationResponse struct {
	ContentType                              string
	InternalServerError                      *shared.InternalServerErrorEnum
	StatusCode                               int
	RawResponse                              *http.Response
	UpdateDestination200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                        *shared.ValidateErrorJSON
}
