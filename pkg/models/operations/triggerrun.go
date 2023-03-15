package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type TriggerRunRequest struct {
	TriggerRunInput *shared.TriggerRunInput `request:"mediaType=application/json"`
	SyncID          string                  `pathParam:"style=simple,explode=false,name=syncId"`
}

type TriggerRunResponse struct {
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
	TriggerRunOutput  *shared.TriggerRunOutput
	ValidateErrorJSON *shared.ValidateErrorJSON
}
