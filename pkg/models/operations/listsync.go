package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type ListSyncSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type ListSyncOrderByEnum string

const (
	ListSyncOrderByEnumID        ListSyncOrderByEnum = "id"
	ListSyncOrderByEnumName      ListSyncOrderByEnum = "name"
	ListSyncOrderByEnumSlug      ListSyncOrderByEnum = "slug"
	ListSyncOrderByEnumCreatedAt ListSyncOrderByEnum = "createdAt"
	ListSyncOrderByEnumUpdatedAt ListSyncOrderByEnum = "updatedAt"
)

type ListSyncRequest struct {
	After   *time.Time           `queryParam:"style=form,explode=true,name=after"`
	Before  *time.Time           `queryParam:"style=form,explode=true,name=before"`
	Limit   *float64             `queryParam:"style=form,explode=true,name=limit"`
	ModelID *float64             `queryParam:"style=form,explode=true,name=modelId"`
	Offset  *float64             `queryParam:"style=form,explode=true,name=offset"`
	OrderBy *ListSyncOrderByEnum `queryParam:"style=form,explode=true,name=orderBy"`
	Slug    *string              `queryParam:"style=form,explode=true,name=slug"`
}

type ListSync200ApplicationJSON struct {
	Data []shared.Sync `json:"data"`
}

type ListSyncResponse struct {
	ContentType                      string
	ListSync200ApplicationJSONObject *ListSync200ApplicationJSON
	StatusCode                       int
	RawResponse                      *http.Response
	ValidateErrorJSON                *shared.ValidateErrorJSON
}
