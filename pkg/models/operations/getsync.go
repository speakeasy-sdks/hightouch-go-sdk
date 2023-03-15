package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetSyncSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetSyncRequest struct {
	SyncID float64 `pathParam:"style=simple,explode=false,name=syncId"`
}

type GetSyncResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	Sync        *shared.Sync
}
