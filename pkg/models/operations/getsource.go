// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetSourceSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetSourceRequest struct {
	// The id of the source
	SourceID float64 `pathParam:"style=simple,explode=false,name=sourceId"`
}

type GetSourceResponse struct {
	ContentType string
	// Ok
	Source      *shared.Source
	StatusCode  int
	RawResponse *http.Response
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}
