package shared

// ModelUpdateCustom
// Custom query for sources that doesn't support sql. For example, Airtable.
type ModelUpdateCustom struct {
	Query interface{} `json:"query"`
}

type ModelUpdateDbt struct {
	ModelID string `json:"modelId"`
}

// ModelUpdateRaw
// Standard raw SQL query
type ModelUpdateRaw struct {
	SQL string `json:"sql"`
}

// ModelUpdateTable
// Table-based query that fetches on a table instead of SQL
type ModelUpdateTable struct {
	Name string `json:"name"`
}

// ModelUpdateVisual
// Visual query, used by audience
type ModelUpdateVisual struct {
	Filter         interface{} `json:"filter"`
	ParentID       string      `json:"parentId"`
	PrimaryLabel   string      `json:"primaryLabel"`
	SecondaryLabel string      `json:"secondaryLabel"`
}

// ModelUpdate
// The input for updating a Model
type ModelUpdate struct {
	Custom     *ModelUpdateCustom `json:"custom,omitempty"`
	Dbt        *ModelUpdateDbt    `json:"dbt,omitempty"`
	IsSchema   *bool              `json:"isSchema,omitempty"`
	Name       *string            `json:"name,omitempty"`
	PrimaryKey *string            `json:"primaryKey,omitempty"`
	Raw        *ModelUpdateRaw    `json:"raw,omitempty"`
	Table      *ModelUpdateTable  `json:"table,omitempty"`
	Visual     *ModelUpdateVisual `json:"visual,omitempty"`
}
