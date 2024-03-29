// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type TriggerSequenceRunRequest struct {
	// The id of the sync sequence to trigger a run
	SyncSequenceID string `pathParam:"style=simple,explode=false,name=syncSequenceId"`
}

func (o *TriggerSequenceRunRequest) GetSyncSequenceID() string {
	if o == nil {
		return ""
	}
	return o.SyncSequenceID
}

type TriggerSequenceRunResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	TriggerSequenceRunOutput *shared.TriggerSequenceRunOutput
}

func (o *TriggerSequenceRunResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TriggerSequenceRunResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TriggerSequenceRunResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TriggerSequenceRunResponse) GetTriggerSequenceRunOutput() *shared.TriggerSequenceRunOutput {
	if o == nil {
		return nil
	}
	return o.TriggerSequenceRunOutput
}
