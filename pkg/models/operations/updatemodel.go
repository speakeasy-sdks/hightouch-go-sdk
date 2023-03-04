package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type UpdateModelSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type UpdateModelPathParams struct {
	ModelID float64 `pathParam:"style=simple,explode=false,name=modelId"`
}

type UpdateModelRequest struct {
	PathParams UpdateModelPathParams
	Request    shared.ModelUpdate `request:"mediaType=application/json"`
	Security   UpdateModelSecurity
}

type UpdateModelResponse struct {
	ContentType                        string
	InternalServerError                *shared.InternalServerErrorEnum
	StatusCode                         int
	RawResponse                        *http.Response
	UpdateModel200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                  *shared.ValidateErrorJSON
}
