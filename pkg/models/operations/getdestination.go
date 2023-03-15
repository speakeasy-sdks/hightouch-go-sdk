package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetDestinationSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetDestinationRequest struct {
	DestinationID float64 `pathParam:"style=simple,explode=false,name=destinationId"`
}

type GetDestinationResponse struct {
	ContentType string
	Destination *shared.Destination
	StatusCode  int
	RawResponse *http.Response
}
