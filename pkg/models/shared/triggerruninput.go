package shared

// TriggerRunInput
// The input of a trigger action to run syncs
type TriggerRunInput struct {
	FullResync *bool `json:"fullResync,omitempty"`
}
