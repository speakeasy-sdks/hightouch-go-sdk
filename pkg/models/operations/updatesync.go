// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
	"net/http"
)

type UpdateSyncRequest struct {
	SyncUpdate shared.SyncUpdate `request:"mediaType=application/json"`
	// The sync's ID
	SyncID float64 `pathParam:"style=simple,explode=false,name=syncId"`
}

func (o *UpdateSyncRequest) GetSyncUpdate() shared.SyncUpdate {
	if o == nil {
		return shared.SyncUpdate{}
	}
	return o.SyncUpdate
}

func (o *UpdateSyncRequest) GetSyncID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SyncID
}

type UpdateSync200ApplicationJSONType string

const (
	UpdateSync200ApplicationJSONTypeSync                UpdateSync200ApplicationJSONType = "Sync"
	UpdateSync200ApplicationJSONTypeValidateErrorJSON   UpdateSync200ApplicationJSONType = "ValidateErrorJSON"
	UpdateSync200ApplicationJSONTypeInternalServerError UpdateSync200ApplicationJSONType = "InternalServerError"
)

type UpdateSync200ApplicationJSON struct {
	Sync                *shared.Sync
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type UpdateSync200ApplicationJSONType
}

func CreateUpdateSync200ApplicationJSONSync(sync shared.Sync) UpdateSync200ApplicationJSON {
	typ := UpdateSync200ApplicationJSONTypeSync

	return UpdateSync200ApplicationJSON{
		Sync: &sync,
		Type: typ,
	}
}

func CreateUpdateSync200ApplicationJSONValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) UpdateSync200ApplicationJSON {
	typ := UpdateSync200ApplicationJSONTypeValidateErrorJSON

	return UpdateSync200ApplicationJSON{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateUpdateSync200ApplicationJSONInternalServerError(internalServerError shared.InternalServerError) UpdateSync200ApplicationJSON {
	typ := UpdateSync200ApplicationJSONTypeInternalServerError

	return UpdateSync200ApplicationJSON{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *UpdateSync200ApplicationJSON) UnmarshalJSON(data []byte) error {

	validateErrorJSON := new(shared.ValidateErrorJSON)
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = validateErrorJSON
		u.Type = UpdateSync200ApplicationJSONTypeValidateErrorJSON
		return nil
	}

	sync := new(shared.Sync)
	if err := utils.UnmarshalJSON(data, &sync, "", true, true); err == nil {
		u.Sync = sync
		u.Type = UpdateSync200ApplicationJSONTypeSync
		return nil
	}

	internalServerError := new(shared.InternalServerError)
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = internalServerError
		u.Type = UpdateSync200ApplicationJSONTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateSync200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.Sync != nil {
		return utils.MarshalJSON(u.Sync, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	if u.InternalServerError != nil {
		return utils.MarshalJSON(u.InternalServerError, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UpdateSyncResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Something went wrong
	InternalServerError *shared.InternalServerError
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	UpdateSync200ApplicationJSONOneOf *UpdateSync200ApplicationJSON
	// Validation Failed
	ValidateErrorJSON *shared.ValidateErrorJSON
}

func (o *UpdateSyncResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSyncResponse) GetInternalServerError() *shared.InternalServerError {
	if o == nil {
		return nil
	}
	return o.InternalServerError
}

func (o *UpdateSyncResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSyncResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSyncResponse) GetUpdateSync200ApplicationJSONOneOf() *UpdateSync200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateSync200ApplicationJSONOneOf
}

func (o *UpdateSyncResponse) GetValidateErrorJSON() *shared.ValidateErrorJSON {
	if o == nil {
		return nil
	}
	return o.ValidateErrorJSON
}
