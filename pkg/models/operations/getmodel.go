package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type GetModelSecurity struct {
	BearerAuth shared.SchemeBearerAuth `security:"scheme,type=http,subtype=bearer"`
}

type GetModelPathParams struct {
	ModelID float64 `pathParam:"style=simple,explode=false,name=modelId"`
}

type GetModelRequest struct {
	PathParams GetModelPathParams
	Security   GetModelSecurity
}

type GetModelResponse struct {
	ContentType string
	Model       *shared.Model
	StatusCode  int
	RawResponse *http.Response
}
