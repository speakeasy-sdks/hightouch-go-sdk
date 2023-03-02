package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
)

type CreateSourceSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type CreateSourceRequest struct {
	Request  shared.SourceCreate `request:"mediaType=application/json"`
	Security CreateSourceSecurity
}

type CreateSourceResponse struct {
	ContentType                         string
	CreateSource200ApplicationJSONAnyOf interface{}
	InternalServerError                 *shared.InternalServerErrorEnum
	StatusCode                          int
	ValidateErrorJSON                   *shared.ValidateErrorJSON
}
