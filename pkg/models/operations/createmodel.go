// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type CreateModelSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type CreateModelResponse struct {
	ContentType string
	// Ok
	CreateModel200ApplicationJSONAnyOf interface{}
	// Something went wrong
	InternalServerError *shared.InternalServerErrorEnum
	StatusCode          int
	RawResponse         *http.Response
	// Conflict
	ValidateErrorJSON *shared.ValidateErrorJSON
}
