package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateDestinationSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type UpdateDestinationRequest struct {
	DestinationUpdate shared.DestinationUpdate `request:"mediaType=application/json"`
	DestinationID     float64                  `pathParam:"style=simple,explode=false,name=destinationId"`
}

type UpdateDestinationResponse struct {
	ContentType                              string
	InternalServerError                      *shared.InternalServerErrorEnum
	StatusCode                               int
	RawResponse                              *http.Response
	UpdateDestination200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                        *shared.ValidateErrorJSON
}
