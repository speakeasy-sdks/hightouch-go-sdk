package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type ListSourceSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type ListSourceOrderByEnum string

const (
	ListSourceOrderByEnumID        ListSourceOrderByEnum = "id"
	ListSourceOrderByEnumName      ListSourceOrderByEnum = "name"
	ListSourceOrderByEnumSlug      ListSourceOrderByEnum = "slug"
	ListSourceOrderByEnumCreatedAt ListSourceOrderByEnum = "createdAt"
	ListSourceOrderByEnumUpdatedAt ListSourceOrderByEnum = "updatedAt"
)

type ListSourceQueryParams struct {
	Limit   *float64               `queryParam:"style=form,explode=true,name=limit"`
	Name    *string                `queryParam:"style=form,explode=true,name=name"`
	Offset  *float64               `queryParam:"style=form,explode=true,name=offset"`
	OrderBy *ListSourceOrderByEnum `queryParam:"style=form,explode=true,name=orderBy"`
	Slug    *string                `queryParam:"style=form,explode=true,name=slug"`
}

type ListSourceRequest struct {
	QueryParams ListSourceQueryParams
	Security    ListSourceSecurity
}

type ListSource200ApplicationJSON struct {
	Data []shared.Source `json:"data"`
}

type ListSourceResponse struct {
	ContentType                        string
	ListSource200ApplicationJSONObject *ListSource200ApplicationJSON
	StatusCode                         int
	RawResponse                        *http.Response
}
