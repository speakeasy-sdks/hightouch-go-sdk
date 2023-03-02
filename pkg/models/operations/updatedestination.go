package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

type UpdateDestinationPathParams struct {
	DestinationID float64 `pathParam:"style=simple,explode=false,name=destinationId"`
}

type UpdateDestinationSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
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
	UpdateDestination200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                        *shared.ValidateErrorJSON
}
