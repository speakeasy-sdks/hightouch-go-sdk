package shared

// DestinationCreate
// The input for creating a Destination
type DestinationCreate struct {
	Configuration map[string]interface{} `json:"configuration"`
	Name          string                 `json:"name"`
	Slug          string                 `json:"slug"`
	Type          string                 `json:"type"`
}
