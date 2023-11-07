// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
	"net/http"
)

type CreateModelResponseBodyType string

const (
	CreateModelResponseBodyTypeModel               CreateModelResponseBodyType = "Model"
	CreateModelResponseBodyTypeValidateErrorJSON   CreateModelResponseBodyType = "ValidateErrorJSON"
	CreateModelResponseBodyTypeInternalServerError CreateModelResponseBodyType = "InternalServerError"
)

type CreateModelResponseBody struct {
	Model               *shared.Model
	ValidateErrorJSON   *shared.ValidateErrorJSON
	InternalServerError *shared.InternalServerError

	Type CreateModelResponseBodyType
}

func CreateCreateModelResponseBodyModel(model shared.Model) CreateModelResponseBody {
	typ := CreateModelResponseBodyTypeModel

	return CreateModelResponseBody{
		Model: &model,
		Type:  typ,
	}
}

func CreateCreateModelResponseBodyValidateErrorJSON(validateErrorJSON shared.ValidateErrorJSON) CreateModelResponseBody {
	typ := CreateModelResponseBodyTypeValidateErrorJSON

	return CreateModelResponseBody{
		ValidateErrorJSON: &validateErrorJSON,
		Type:              typ,
	}
}

func CreateCreateModelResponseBodyInternalServerError(internalServerError shared.InternalServerError) CreateModelResponseBody {
	typ := CreateModelResponseBodyTypeInternalServerError

	return CreateModelResponseBody{
		InternalServerError: &internalServerError,
		Type:                typ,
	}
}

func (u *CreateModelResponseBody) UnmarshalJSON(data []byte) error {

	validateErrorJSON := shared.ValidateErrorJSON{}
	if err := utils.UnmarshalJSON(data, &validateErrorJSON, "", true, true); err == nil {
		u.ValidateErrorJSON = &validateErrorJSON
		u.Type = CreateModelResponseBodyTypeValidateErrorJSON
		return nil
	}

	model := shared.Model{}
	if err := utils.UnmarshalJSON(data, &model, "", true, true); err == nil {
		u.Model = &model
		u.Type = CreateModelResponseBodyTypeModel
		return nil
	}

	internalServerError := shared.InternalServerError("")
	if err := utils.UnmarshalJSON(data, &internalServerError, "", true, true); err == nil {
		u.InternalServerError = &internalServerError
		u.Type = CreateModelResponseBodyTypeInternalServerError
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateModelResponseBody) MarshalJSON() ([]byte, error) {
	if u.Model != nil {
		return utils.MarshalJSON(u.Model, "", true)
	}

	if u.ValidateErrorJSON != nil {
		return utils.MarshalJSON(u.ValidateErrorJSON, "", true)
	}

	if u.InternalServerError != nil {
		return utils.MarshalJSON(u.InternalServerError, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type CreateModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	OneOf *CreateModelResponseBody
}

func (o *CreateModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateModelResponse) GetOneOf() *CreateModelResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
