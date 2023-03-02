package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
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
	ValidateErrorJSON                 *shared.ValidateErrorJSON
}
