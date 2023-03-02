package shared

// RecordDayBooleanOrUndefined
// Construct a type with a set of properties K of type T
type RecordDayBooleanOrUndefined struct {
	Friday    *bool `json:"friday,omitempty"`
	Monday    *bool `json:"monday,omitempty"`
	Saturday  *bool `json:"saturday,omitempty"`
	Sunday    *bool `json:"sunday,omitempty"`
	Thursday  *bool `json:"thursday,omitempty"`
	Tuesday   *bool `json:"tuesday,omitempty"`
	Wednesday *bool `json:"wednesday,omitempty"`
}
