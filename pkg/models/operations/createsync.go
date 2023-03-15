package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateSyncSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type CreateSyncResponse struct {
	ContentType                       string
	CreateSync200ApplicationJSONAnyOf interface{}
	InternalServerError               *shared.InternalServerErrorEnum
	StatusCode                        int
	RawResponse                       *http.Response
	ValidateErrorJSON                 *shared.ValidateErrorJSON
}
