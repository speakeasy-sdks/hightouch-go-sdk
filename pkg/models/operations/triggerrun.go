package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

type TriggerRunPathParams struct {
	SyncID string `pathParam:"style=simple,explode=false,name=syncId"`
}

type TriggerRunSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type TriggerRunRequest struct {
	PathParams TriggerRunPathParams
	Request    *shared.TriggerRunInput `request:"mediaType=application/json"`
	Security   TriggerRunSecurity
}

type TriggerRunResponse struct {
	ContentType       string
	StatusCode        int
	TriggerRunOutput  *shared.TriggerRunOutput
	ValidateErrorJSON *shared.ValidateErrorJSON
}
