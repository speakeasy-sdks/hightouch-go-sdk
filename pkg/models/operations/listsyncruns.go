package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type ListSyncRunsSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type ListSyncRunsPathParams struct {
	SyncID float64 `pathParam:"style=simple,explode=false,name=syncId"`
}

type ListSyncRunsOrderByEnum string

const (
	ListSyncRunsOrderByEnumID         ListSyncRunsOrderByEnum = "id"
	ListSyncRunsOrderByEnumCreatedAt  ListSyncRunsOrderByEnum = "createdAt"
	ListSyncRunsOrderByEnumStartedAt  ListSyncRunsOrderByEnum = "startedAt"
	ListSyncRunsOrderByEnumFinishedAt ListSyncRunsOrderByEnum = "finishedAt"
)

type ListSyncRunsQueryParams struct {
	After   *time.Time               `queryParam:"style=form,explode=true,name=after"`
	Before  *time.Time               `queryParam:"style=form,explode=true,name=before"`
	Limit   *float64                 `queryParam:"style=form,explode=true,name=limit"`
	Offset  *float64                 `queryParam:"style=form,explode=true,name=offset"`
	OrderBy *ListSyncRunsOrderByEnum `queryParam:"style=form,explode=true,name=orderBy"`
	RunID   *float64                 `queryParam:"style=form,explode=true,name=runId"`
	Within  *float64                 `queryParam:"style=form,explode=true,name=within"`
}

type ListSyncRunsRequest struct {
	PathParams  ListSyncRunsPathParams
	QueryParams ListSyncRunsQueryParams
	Security    ListSyncRunsSecurity
}

type ListSyncRuns200ApplicationJSON struct {
	Data []shared.SyncRun `json:"data"`
}

type ListSyncRunsResponse struct {
	ContentType                          string
	ListSyncRuns200ApplicationJSONObject *ListSyncRuns200ApplicationJSON
	StatusCode                           int
	RawResponse                          *http.Response
	ValidateErrorJSON                    *shared.ValidateErrorJSON
}
