// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
	"time"
)

type ListSyncSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

// ListSyncOrderByEnum - specify the order
type ListSyncOrderByEnum string

const (
	ListSyncOrderByEnumID        ListSyncOrderByEnum = "id"
	ListSyncOrderByEnumName      ListSyncOrderByEnum = "name"
	ListSyncOrderByEnumSlug      ListSyncOrderByEnum = "slug"
	ListSyncOrderByEnumCreatedAt ListSyncOrderByEnum = "createdAt"
	ListSyncOrderByEnumUpdatedAt ListSyncOrderByEnum = "updatedAt"
)

func (e ListSyncOrderByEnum) ToPointer() *ListSyncOrderByEnum {
	return &e
}

func (e *ListSyncOrderByEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "id":
		fallthrough
	case "name":
		fallthrough
	case "slug":
		fallthrough
	case "createdAt":
		fallthrough
	case "updatedAt":
		*e = ListSyncOrderByEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListSyncOrderByEnum: %v", v)
	}
}

type ListSyncRequest struct {
	// select syncs that were run after given ISO timestamp
	After *time.Time `queryParam:"style=form,explode=true,name=after"`
	// select syncs that were run before given ISO timestamp
	Before *time.Time `queryParam:"style=form,explode=true,name=before"`
	// limit the number of objects returned (default is 100)
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// filter based on modelId
	ModelID *float64 `queryParam:"style=form,explode=true,name=modelId"`
	// set the offset on results (for pagination)
	Offset *float64 `queryParam:"style=form,explode=true,name=offset"`
	// specify the order
	OrderBy *ListSyncOrderByEnum `queryParam:"style=form,explode=true,name=orderBy"`
	// filter based on slug
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

// ListSync200ApplicationJSON - Ok
type ListSync200ApplicationJSON struct {
	Data []shared.Sync `json:"data"`
}

type ListSyncResponse struct {
	ContentType string
	// Ok
	ListSync200ApplicationJSONObject *ListSync200ApplicationJSON
	StatusCode                       int
	RawResponse                      *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
