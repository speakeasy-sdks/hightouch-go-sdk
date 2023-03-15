package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateModelSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type UpdateModelRequest struct {
	ModelUpdate shared.ModelUpdate `request:"mediaType=application/json"`
	ModelID     float64            `pathParam:"style=simple,explode=false,name=modelId"`
}

type UpdateModelResponse struct {
	ContentType                        string
	InternalServerError                *shared.InternalServerErrorEnum
	StatusCode                         int
	RawResponse                        *http.Response
	UpdateModel200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                  *shared.ValidateErrorJSON
}
