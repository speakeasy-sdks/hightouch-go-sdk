package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateSyncSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type UpdateSyncRequest struct {
	SyncUpdate shared.SyncUpdate `request:"mediaType=application/json"`
	SyncID     float64           `pathParam:"style=simple,explode=false,name=syncId"`
}

type UpdateSyncResponse struct {
	ContentType                       string
	InternalServerError               *shared.InternalServerErrorEnum
	StatusCode                        int
	RawResponse                       *http.Response
	UpdateSync200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                 *shared.ValidateErrorJSON
}
