// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateSyncSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *CreateSyncSecurity) GetBearerAuth() string {
	if o == nil {
		return ""
	}
	return o.BearerAuth
}

type CreateSyncResponse struct {
	ContentType string
	// Ok
	CreateSync200ApplicationJSONAnyOf interface{}
	// Something went wrong
	InternalServerError *shared.InternalServerError
	StatusCode          int
	RawResponse         *http.Response
	// Conflict
	ValidateErrorJSON *shared.ValidateErrorJSON
}

func (o *CreateSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSyncResponse) GetCreateSync200ApplicationJSONAnyOf() interface{} {
	if o == nil {
		return nil
	}
	return o.CreateSync200ApplicationJSONAnyOf
}

func (o *CreateSyncResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSyncResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}
