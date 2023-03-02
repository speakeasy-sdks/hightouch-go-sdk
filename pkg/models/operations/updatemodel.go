package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

type UpdateModelPathParams struct {
	ModelID float64 `pathParam:"style=simple,explode=false,name=modelId"`
}

type UpdateModelSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
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
	UpdateModel200ApplicationJSONAnyOf interface{}
	ValidateErrorJSON                  *shared.ValidateErrorJSON
}
