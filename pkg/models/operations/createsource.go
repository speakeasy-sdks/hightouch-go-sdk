// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
	"net/http"
)

type CreateSource200ApplicationJSONType string

const (
	CreateSource200ApplicationJSONTypeSource              CreateSource200ApplicationJSONType = "Source"
	CreateSource200ApplicationJSONTypeValidateErrorJSON   CreateSource200ApplicationJSONType = "ValidateErrorJSON"
	CreateSource200ApplicationJSONTypeInternalServerError CreateSource200ApplicationJSONType = "InternalServerError"
)

type CreateSource200ApplicationJSON struct {
	Source              *shared.Source
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type CreateSource200ApplicationJSONType
}

func CreateCreateSource200ApplicationJSONSource(source shared.Source) CreateSource200ApplicationJSON {
	typ := CreateSource200ApplicationJSONTypeSource

	return CreateSource200ApplicationJSON{
		Source: &source,
		Type:   typ,
	}
}

func CreateCreateSource200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) CreateSource200ApplicationJSON {
	typ := CreateSource200ApplicationJSONTypeValidateErrorJSON

	return CreateSource200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateCreateSource200ApplicationJSONInternalServerError(internalServerError shared.InternalServerError) CreateSource200ApplicationJSON {
	typ := CreateSource200ApplicationJSONTypeInternalServerError

	return CreateSource200ApplicationJSON{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *CreateSource200ApplicationJSON) UnmarshalJSON(data []byte) error {

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = CreateSource200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	source := new(shared.Source)
	if err := utils.UnmarshalJSON(data, &source, "", true, true); err == nil {
		u.Source = source
		u.Type = CreateSource200ApplicationJSONTypeSource
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = internalServerError
		u.Type = CreateSource200ApplicationJSONTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateSource200ApplicationJSON) MarshalJSON() ([]byte, error) {
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

type CreateSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Ok
	CreateSource200ApplicationJSONOneOf *CreateSource200ApplicationJSON
	// Something went wrong
	InternalServerError *shared.InternalServerError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Conflict
	ValidateErrorJSON *shared.ValidateErrorJSON
}

func (o *CreateSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSourceResponse) GetCreateSource200ApplicationJSONOneOf() *CreateSource200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateSource200ApplicationJSONOneOf
}

func (o *CreateSourceResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *CreateSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSourceResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}
