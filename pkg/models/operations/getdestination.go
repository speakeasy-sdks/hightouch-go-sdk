package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetDestinationSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type GetDestinationPathParams struct {
	DestinationID float64 `pathParam:"style=simple,explode=false,name=destinationId"`
}

type GetDestinationRequest struct {
	PathParams GetDestinationPathParams
	Security   GetDestinationSecurity
}

type GetDestinationResponse struct {
	ContentType string
	Destination *shared.Destination
	StatusCode  int
	RawResponse *http.Response
}
