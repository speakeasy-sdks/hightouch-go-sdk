// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunIDGraphRequest struct {
	TriggerRunIDGraphInput *shared.TriggerRunIDGraphInput `request:"mediaType=application/json"`
	GraphID                string                         `pathParam:"style=simple,explode=false,name=graphId"`
}

func (o *TriggerRunIDGraphRequest) GetTriggerRunIDGraphInput() *shared.TriggerRunIDGraphInput {
	if o == nil {
		return nil
	}
	return o.TriggerRunIDGraphInput
}

func (o *TriggerRunIDGraphRequest) GetGraphID() string {
	if o == nil {
		return ""
	}
	return o.GraphID
}

type TriggerRunIDGraphResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	TriggerRunIDGraphOutput *shared.TriggerRunIDGraphOutput
}

func (o *TriggerRunIDGraphResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TriggerRunIDGraphResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TriggerRunIDGraphResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TriggerRunIDGraphResponse) GetTriggerRunIDGraphOutput() *shared.TriggerRunIDGraphOutput {
	if o == nil {
		return nil
	}
	return o.TriggerRunIDGraphOutput
}
