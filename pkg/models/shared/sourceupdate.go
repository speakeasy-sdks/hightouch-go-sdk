package shared

// SourceUpdate
// The input for updating a Source
type SourceUpdate struct {
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	Name          *string                `json:"name,omitempty"`
}
