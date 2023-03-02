package shared

// TriggerRunCustomInput
// The input of a trigger action to run syncs based on sync ID, slug or other filters
type TriggerRunCustomInput struct {
	FullResync *bool   `json:"fullResync,omitempty"`
	SyncID     *string `json:"syncId,omitempty"`
	SyncSlug   *string `json:"syncSlug,omitempty"`
}
