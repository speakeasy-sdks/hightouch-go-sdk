package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type TriggerRunPathParams struct {
	SyncID string `pathParam:"style=simple,explode=false,name=syncId"`
}

type TriggerRunRequest struct {
	PathParams TriggerRunPathParams
	Request    *shared.TriggerRunInput `request:"mediaType=application/json"`
	Security   TriggerRunSecurity
}

type TriggerRunResponse struct {
	ContentType       string
	StatusCode        int
	RawResponse       *http.Response
	TriggerRunOutput  *shared.TriggerRunOutput
	ValidateErrorJSON *shared.ValidateErrorJSON
}
