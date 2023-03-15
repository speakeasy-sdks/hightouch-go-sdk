package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type ListDestinationSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type ListDestinationOrderByEnum string

const (
	ListDestinationOrderByEnumID        ListDestinationOrderByEnum = "id"
	ListDestinationOrderByEnumName      ListDestinationOrderByEnum = "name"
	ListDestinationOrderByEnumSlug      ListDestinationOrderByEnum = "slug"
	ListDestinationOrderByEnumCreatedAt ListDestinationOrderByEnum = "createdAt"
	ListDestinationOrderByEnumUpdatedAt ListDestinationOrderByEnum = "updatedAt"
)

type ListDestinationRequest struct {
	Limit   *float64                    `queryParam:"style=form,explode=true,name=limit"`
	Name    *string                     `queryParam:"style=form,explode=true,name=name"`
	Offset  *float64                    `queryParam:"style=form,explode=true,name=offset"`
	OrderBy *ListDestinationOrderByEnum `queryParam:"style=form,explode=true,name=orderBy"`
	Slug    *string                     `queryParam:"style=form,explode=true,name=slug"`
}

type ListDestination200ApplicationJSON struct {
	Data []shared.Destination `json:"data"`
}

type ListDestinationResponse struct {
	ContentType                             string
	ListDestination200ApplicationJSONObject *ListDestination200ApplicationJSON
	StatusCode                              int
	RawResponse                             *http.Response
	ValidateErrorJSON                       *shared.ValidateErrorJSON
}
