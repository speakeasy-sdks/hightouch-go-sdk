// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetModelSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type GetModelRequest struct {
	// The id of the model
	ModelID float64 `pathParam:"style=simple,explode=false,name=modelId"`
}

type GetModelResponse struct {
	ContentType string
	// Ok
	Model       *shared.Model
	StatusCode  int
	RawResponse *http.Response
}
