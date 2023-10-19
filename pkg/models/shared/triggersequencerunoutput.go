// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TriggerSequenceRunOutput - The output of a trigger action to run sync sequences
type TriggerSequenceRunOutput struct {
	// The id of the triggered sync sequence run.
	ID string `json:"id"`
}

func (o *TriggerSequenceRunOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}