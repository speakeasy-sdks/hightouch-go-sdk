package shared

type ValidateErrorJSONMessageEnum string

const (
	ValidateErrorJSONMessageEnumValidationFailed ValidateErrorJSONMessageEnum = "Validation failed"
)

type ValidateErrorJSON struct {
	Details map[string]interface{}       `json:"details"`
	Message ValidateErrorJSONMessageEnum `json:"message"`
}
