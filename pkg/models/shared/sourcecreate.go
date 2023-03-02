package shared

// SourceCreate
// The input for creating a Source
type SourceCreate struct {
	Configuration map[string]interface{} `json:"configuration"`
	Name          string                 `json:"name"`
	Slug          string                 `json:"slug"`
	Type          string                 `json:"type"`
}
