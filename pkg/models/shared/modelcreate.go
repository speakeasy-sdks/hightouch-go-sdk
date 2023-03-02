package shared

// ModelCreateCustom
// Custom query for sources that doesn't support sql. For example, Airtable.
type ModelCreateCustom struct {
	Query interface{} `json:"query"`
}

type ModelCreateDbt struct {
	ModelID string `json:"modelId"`
}

// ModelCreateRaw
// Standard raw SQL query
type ModelCreateRaw struct {
	SQL string `json:"sql"`
}

// ModelCreateTable
// Table-based query that fetches on a table instead of SQL
type ModelCreateTable struct {
	Name string `json:"name"`
}

// ModelCreateVisual
// Visual query, used by audience
type ModelCreateVisual struct {
	Filter         interface{} `json:"filter"`
	ParentID       string      `json:"parentId"`
	PrimaryLabel   string      `json:"primaryLabel"`
	SecondaryLabel string      `json:"secondaryLabel"`
}

// ModelCreate
// The input for creating a Model
type ModelCreate struct {
	Custom     *ModelCreateCustom `json:"custom,omitempty"`
	Dbt        *ModelCreateDbt    `json:"dbt,omitempty"`
	IsSchema   bool               `json:"isSchema"`
	Name       string             `json:"name"`
	PrimaryKey string             `json:"primaryKey"`
	QueryType  string             `json:"queryType"`
	Raw        *ModelCreateRaw    `json:"raw,omitempty"`
	Slug       string             `json:"slug"`
	SourceID   string             `json:"sourceId"`
	Table      *ModelCreateTable  `json:"table,omitempty"`
	Visual     *ModelCreateVisual `json:"visual,omitempty"`
}
