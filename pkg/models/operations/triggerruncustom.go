package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunCustomSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type TriggerRunCustomResponse struct {
	ContentType                             string
	StatusCode                              int
	RawResponse                             *http.Response
	TriggerRunCustom200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                       *shared.ValidateErrorJSON
}
