package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type ListModelSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type ListModelOrderByEnum string

const (
	ListModelOrderByEnumID        ListModelOrderByEnum = "id"
	ListModelOrderByEnumName      ListModelOrderByEnum = "name"
	ListModelOrderByEnumSlug      ListModelOrderByEnum = "slug"
	ListModelOrderByEnumCreatedAt ListModelOrderByEnum = "createdAt"
	ListModelOrderByEnumUpdatedAt ListModelOrderByEnum = "updatedAt"
)

type ListModelRequest struct {
	Limit   *float64              `queryParam:"style=form,explode=true,name=limit"`
	Name    *string               `queryParam:"style=form,explode=true,name=name"`
	Offset  *float64              `queryParam:"style=form,explode=true,name=offset"`
	OrderBy *ListModelOrderByEnum `queryParam:"style=form,explode=true,name=orderBy"`
	Slug    *string               `queryParam:"style=form,explode=true,name=slug"`
}

type ListModel200ApplicationJSON struct {
	Data []shared.Model `json:"data"`
}

type ListModelResponse struct {
	ContentType                       string
	ListModel200ApplicationJSONObject *ListModel200ApplicationJSON
	StatusCode                        int
	RawResponse                       *http.Response
	ValidateErrorJSON                 *shared.ValidateErrorJSON
}
