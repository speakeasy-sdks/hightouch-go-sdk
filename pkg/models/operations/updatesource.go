// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
	"net/http"
)

type UpdateSourceRequest struct {
	SourceUpdate shared.SourceUpdate `request:"mediaType=application/json"`
	// The source's ID
	SourceID float64 `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *UpdateSourceRequest) GetSourceUpdate() shared.SourceUpdate {
	if o == nil {
		return shared.SourceUpdate{}
	}
	return o.SourceUpdate
}

func (o *UpdateSourceRequest) GetSourceID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SourceID
}

type UpdateSource200ApplicationJSONType string

const (
	UpdateSource200ApplicationJSONTypeSource              UpdateSource200ApplicationJSONType = "Source"
	UpdateSource200ApplicationJSONTypeValidateErrorJSON   UpdateSource200ApplicationJSONType = "ValidateErrorJSON"
	UpdateSource200ApplicationJSONTypeInternalServerError UpdateSource200ApplicationJSONType = "InternalServerError"
)

type UpdateSource200ApplicationJSON struct {
	Source              *shared.Source
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateSource200ApplicationJSONType
}

func CreateUpdateSource200ApplicationJSONSource(source shared.Source) UpdateSource200ApplicationJSON {
	typ := UpdateSource200ApplicationJSONTypeSource

	return UpdateSource200ApplicationJSON{
		Source: &source,
		Type:   typ,
	}
}

func CreateUpdateSource200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateSource200ApplicationJSON {
	typ := UpdateSource200ApplicationJSONTypeValidateErrorJSON

	return UpdateSource200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateSource200ApplicationJSONInternalServerError(internalServerError shared.InternalServerError) UpdateSource200ApplicationJSON {
	typ := UpdateSource200ApplicationJSONTypeInternalServerError

	return UpdateSource200ApplicationJSON{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateSource200ApplicationJSON) UnmarshalJSON(data []byte) error {

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = UpdateSource200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	source := new(shared.Source)
	if err := utils.UnmarshalJSON(data, &source, "", true, true); err == nil {
		u.Source = source
		u.Type = UpdateSource200ApplicationJSONTypeSource
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = internalServerError
		u.Type = UpdateSource200ApplicationJSONTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateSource200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.Source != nil {
		return utils.MarshalJSON(u.Source, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	if u.InternalServerError != nil {
		return utils.MarshalJSON(u.InternalServerError, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpdateSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Something went wrong
	InternalServerError *shared.InternalServerError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	UpdateSource200ApplicationJSONOneOf *UpdateSource200ApplicationJSON
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}

func (o *UpdateSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSourceResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSourceResponse) GetUpdateSource200ApplicationJSONOneOf() *UpdateSource200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateSource200ApplicationJSONOneOf
}

func (o *UpdateSourceResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}
