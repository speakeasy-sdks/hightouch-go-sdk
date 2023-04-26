// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type ListSyncRunsSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

// ListSyncRunsOrderByEnum - specify the order
type ListSyncRunsOrderByEnum string

const (
	ListSyncRunsOrderByEnumID         ListSyncRunsOrderByEnum = "id"
	ListSyncRunsOrderByEnumCreatedAt  ListSyncRunsOrderByEnum = "createdAt"
	ListSyncRunsOrderByEnumStartedAt  ListSyncRunsOrderByEnum = "startedAt"
	ListSyncRunsOrderByEnumFinishedAt ListSyncRunsOrderByEnum = "finishedAt"
)

func (e ListSyncRunsOrderByEnum) ToPointer() *ListSyncRunsOrderByEnum {
	return &e
}

func (e *ListSyncRunsOrderByEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "id":
		fallthrough
	case "createdAt":
		fallthrough
	case "startedAt":
		fallthrough
	case "finishedAt":
		*e = ListSyncRunsOrderByEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ListSyncRunsOrderByEnum: %s", s)
	}
}

type ListSyncRunsRequest struct {
	// select sync runs that are started after given ISO timestamp
	After *time.Time `queryParam:"style=form,explode=true,name=after"`
	// select sync runs that are started before certain ISO timestamp
	Before *time.Time `queryParam:"style=form,explode=true,name=before"`
	// limit the number of objects returned (default is 5)
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// set the offset on results (for pagination)
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	// specify the order
	OrderBy *ListSyncRunsOrderByEnum `queryParam:"style=form,explode=true,name=orderBy"`
	// query for specific run id
	RunID  *float64 `queryParam:"style=form,explode=true,name=runId"`
	SyncID float64  `pathParam:"style=simple,explode=false,name=syncId"`
	// select sync runs that are started within last given minutes
	Within *float64 `queryParam:"style=form,explode=true,name=within"`
}

// ListSyncRuns200ApplicationJSON - Ok
type ListSyncRuns200ApplicationJSON struct {
	Data []shared.SyncRun `json:"data"`
}

type ListSyncRunsResponse struct {
	ContentType string
	// Ok
	ListSyncRuns200ApplicationJSONObject *ListSyncRuns200ApplicationJSON
	StatusCode                           int
	RawResponse                          *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
