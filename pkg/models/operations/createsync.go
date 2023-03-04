package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateSyncSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type CreateSyncRequest struct {
	Request  shared.SyncCreate `request:"mediaType=application/json"`
	Security CreateSyncSecurity
}

type CreateSyncResponse struct {
	ContentType                       string
	CreateSync200ApplicationJSONAnyOf interface{}
	InternalServerError               *shared.InternalServerErrorEnum
	StatusCode                        int
	RawResponse                       *http.Response
	ValidateErrorJSON                 *shared.ValidateErrorJSON
}
