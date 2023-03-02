package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

type UpdateSyncPathParams struct {
	SyncID float64 `pathParam:"style=simple,explode=false,name=syncId"`
}

type UpdateSyncSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type UpdateSyncRequest struct {
	PathParams UpdateSyncPathParams
	Request    shared.SyncUpdate `request:"mediaType=application/json"`
	Security   UpdateSyncSecurity
}

type UpdateSyncResponse struct {
	ContentType                       string
	InternalServerError               *shared.InternalServerErrorEnum
	StatusCode                        int
	UpdateSync200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                 *shared.ValidateErrorJSON
}
