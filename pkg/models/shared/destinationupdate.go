package shared

// DestinationUpdate
// The input for updating a Destination
type DestinationUpdate struct {
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	Name          *string                `json:"name,omitempty"`
}
