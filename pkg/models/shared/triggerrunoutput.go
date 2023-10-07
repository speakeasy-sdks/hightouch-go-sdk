// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
)

// TriggerRunOutput - The output of a trigger action to run syncs
type TriggerRunOutput struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The id of the triggered sync run. This can be passed to `/sync/runs` to
	// get the run's status.
	ID string `json:"id"`
}

func (t TriggerRunOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TriggerRunOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *TriggerRunOutput) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *TriggerRunOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}
