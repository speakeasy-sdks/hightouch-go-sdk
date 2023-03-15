package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateModelSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type CreateModelResponse struct {
	ContentType                        string
	CreateModel200ApplicationJSONAnyOf interface{}
	InternalServerError                *shared.InternalServerErrorEnum
	StatusCode                         int
	RawResponse                        *http.Response
	ValidateErrorJSON                  *shared.ValidateErrorJSON
}
