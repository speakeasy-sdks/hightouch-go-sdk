// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"net/http"
)

type TriggerRunCustomSecurity struct {
	BearerAuth string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

func (o *TriggerRunCustomSecurity) GetBearerAuth() string {
	if o == nil {
		return ""
	}
	return o.BearerAuth
}

type TriggerRunCustomResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Ok
	TriggerRunCustom200ApplicationJSONAnyOf interface{}
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}

func (o *TriggerRunCustomResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *TriggerRunCustomResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *TriggerRunCustomResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *TriggerRunCustomResponse) GetTriggerRunCustom200ApplicationJSONAnyOf() interface{} {
	if o == nil {
		return nil
	}
	return o.TriggerRunCustom200ApplicationJSONAnyOf
}

func (o *TriggerRunCustomResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}
