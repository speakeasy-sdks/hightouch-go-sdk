package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetSyncSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type GetSyncPathParams struct {
	SyncID float64 `pathParam:"style=simple,explode=false,name=syncId"`
}

type GetSyncRequest struct {
	PathParams GetSyncPathParams
	Security   GetSyncSecurity
}

type GetSyncResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Sync        *shared.Sync
}
