package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type CreateDestinationRequest struct {
	Request  shared.DestinationCreate `request:"mediaType=application/json"`
	Security CreateDestinationSecurity
}

type CreateDestinationResponse struct {
	ContentType                              string
	CreateDestination200ApplicationJSONAnyOf interface{}
	InternalServerError                      *shared.InternalServerErrorEnum
	StatusCode                               int
	RawResponse                              *http.Response
	ValidateErrorJSON                        *shared.ValidateErrorJSON
}
