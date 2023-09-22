// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
	"net/http"
)

type UpdateDestinationRequest struct {
	DestinationUpdate shared.DestinationUpdate `request:"mediaType=application/json"`
	// The destination's ID
	DestinationID float64 `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *UpdateDestinationRequest) GetDestinationUpdate() shared.DestinationUpdate {
	if o == nil {
		return shared.DestinationUpdate{}
	}
	return o.DestinationUpdate
}

func (o *UpdateDestinationRequest) GetDestinationID() float64 {
	if o == nil {
		return 0.0
	}
	return o.DestinationID
}

type UpdateDestination200ApplicationJSONType string

const (
	UpdateDestination200ApplicationJSONTypeDestination       UpdateDestination200ApplicationJSONType = "Destination"
	UpdateDestination200ApplicationJSONTypeValidateErrorJSON UpdateDestination200ApplicationJSONType = "ValidateErrorJSON"
	UpdateDestination200ApplicationJSONTypeStr               UpdateDestination200ApplicationJSONType = "str"
)

type UpdateDestination200ApplicationJSON struct {
	Destination       *shared.Destination
	ValidateErrorJSON *shared.ValidateErrorJSON
	Str               *string

	Type UpdateDestination200ApplicationJSONType
}

func CreateUpdateDestination200ApplicationJSONDestination(destination shared.Destination) UpdateDestination200ApplicationJSON {
	typ := UpdateDestination200ApplicationJSONTypeDestination

	return UpdateDestination200ApplicationJSON{
		Destination: &destination,
		Type:        typ,
	}
}

func CreateUpdateDestination200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateDestination200ApplicationJSON {
	typ := UpdateDestination200ApplicationJSONTypeValidateErrorJSON

	return UpdateDestination200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateDestination200ApplicationJSONStr(str string) UpdateDestination200ApplicationJSON {
	typ := UpdateDestination200ApplicationJSONTypeStr

	return UpdateDestination200ApplicationJSON{
		Str:  &str,
		Type: typ,
	}
}

func (u *UpdateDestination200ApplicationJSON) UnmarshalJSON(data []byte) error {

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = UpdateDestination200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	destination := new(shared.Destination)
	if err := utils.UnmarshalJSON(data, &destination, "", true, true); err == nil {
		u.Destination = destination
		u.Type = UpdateDestination200ApplicationJSONTypeDestination
		return nil
	}

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = UpdateDestination200ApplicationJSONTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateDestination200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.Destination != nil {
		return utils.MarshalJSON(u.Destination, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpdateDestinationResponse struct {
	ContentType string
	// Something went wrong
	InternalServerError *string
	StatusCode          int
	RawResponse         *http.Response
	// Ok
	UpdateDestination200ApplicationJSONOneOf *UpdateDestination200ApplicationJSON
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}

func (o *UpdateDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDestinationResponse) GetInternalServerError() *string {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDestinationResponse) GetUpdateDestination200ApplicationJSONOneOf() *UpdateDestination200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDestination200ApplicationJSONOneOf
}

func (o *UpdateDestinationResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}
