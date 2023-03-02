package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

type CreateModelSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type CreateModelRequest struct {
	Request  shared.ModelCreate `request:"mediaType=application/json"`
	Security CreateModelSecurity
}

type CreateModelResponse struct {
	ContentType                        string
	CreateModel200ApplicationJSONAnyOf interface{}
	InternalServerError                *shared.InternalServerErrorEnum
	StatusCode                         int
	ValidateErrorJSON                  *shared.ValidateErrorJSON
}
