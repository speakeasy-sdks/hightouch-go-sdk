// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
	"net/http"
)

// ListSourceQueryParamOrderBy - specify the order
type ListSourceQueryParamOrderBy string

const (
	ListSourceQueryParamOrderByID        ListSourceQueryParamOrderBy = "id"
	ListSourceQueryParamOrderByName      ListSourceQueryParamOrderBy = "name"
	ListSourceQueryParamOrderBySlug      ListSourceQueryParamOrderBy = "slug"
	ListSourceQueryParamOrderByCreatedAt ListSourceQueryParamOrderBy = "createdAt"
	ListSourceQueryParamOrderByUpdatedAt ListSourceQueryParamOrderBy = "updatedAt"
)

func (e ListSourceQueryParamOrderBy) ToPointer() *ListSourceQueryParamOrderBy {
	return &e
}

func (e *ListSourceQueryParamOrderBy) UnmarshalJSON(data []byte) error {
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
		*e = ListSourceQueryParamOrderBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListSourceQueryParamOrderBy: %v", v)
	}
}

type ListSourceRequest struct {
	// limit the number of objects returned (default is 100)
	Limit *float64 `default:"100" queryParam:"style=form,explode=true,name=limit"`
	// filter based on name
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// set the offset on results (for pagination)
	Offset *float64 `default:"0" queryParam:"style=form,explode=true,name=offset"`
	// specify the order
	OrderBy *ListSourceQueryParamOrderBy `default:"id" queryParam:"style=form,explode=true,name=orderBy"`
	// filter based on slug
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (l ListSourceRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListSourceRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListSourceRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListSourceRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListSourceRequest) GetOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListSourceRequest) GetOrderBy() *ListSourceQueryParamOrderBy {
	if o == nil {
		return nil
	}
	return o.OrderBy
}

func (o *ListSourceRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

// ListSourceResponseBody - Ok
type ListSourceResponseBody struct {
	Data []shared.Source `json:"data"`
}

func (o *ListSourceResponseBody) GetData() []shared.Source {
	if o == nil {
		return []shared.Source{}
	}
	return o.Data
}

type ListSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	Object *ListSourceResponseBody
}

func (o *ListSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListSourceResponse) GetObject() *ListSourceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
