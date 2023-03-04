package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunCustomSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type TriggerRunCustomRequest struct {
	Request  shared.TriggerRunCustomInput `request:"mediaType=application/json"`
	Security TriggerRunCustomSecurity
}

type TriggerRunCustomResponse struct {
	ContentType                             string
	StatusCode                              int
	RawResponse                             *http.Response
	TriggerRunCustom200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                       *shared.ValidateErrorJSON
}
