package shared

import (
	"time"
)

// ModelCustom
// Custom query for sources that doesn't support sql. For example, Airtable.
type ModelCustom struct {
	Query interface{} `json:"query"`
}

// ModelDbt
// Query that is based on a dbt model
type ModelDbt struct {
	CompiledSQL string `json:"compiledSql"`
	Database    string `json:"database"`
	DbtUniqueID string `json:"dbtUniqueId"`
	ModelID     string `json:"modelId"`
	Name        string `json:"name"`
	RawSQL      string `json:"rawSql"`
	Schema      string `json:"schema"`
}

// ModelRaw
// Standard raw SQL query
type ModelRaw struct {
	SQL string `json:"sql"`
}

// ModelTable
// Table-based query that fetches on a table instead of SQL
type ModelTable struct {
	Name string `json:"name"`
}

// ModelVisual
// Visual query, used by audience
type ModelVisual struct {
	Filter         interface{} `json:"filter"`
	ParentID       string      `json:"parentId"`
	PrimaryLabel   string      `json:"primaryLabel"`
	SecondaryLabel string      `json:"secondaryLabel"`
}

// Model
// The SQL query that pulls data from your source to send to your destination.
// We send your SQL query directly to your source so any SQL that is valid for your source (including functions) is valid in Hightouch.
type Model struct {
	CreatedAt   time.Time         `json:"createdAt"`
	Custom      *ModelCustom      `json:"custom,omitempty"`
	Dbt         *ModelDbt         `json:"dbt,omitempty"`
	ID          string            `json:"id"`
	IsSchema    bool              `json:"isSchema"`
	Name        string            `json:"name"`
	PrimaryKey  string            `json:"primaryKey"`
	QueryType   string            `json:"queryType"`
	Raw         *ModelRaw         `json:"raw,omitempty"`
	Slug        string            `json:"slug"`
	SourceID    string            `json:"sourceId"`
	Syncs       []string          `json:"syncs"`
	Table       *ModelTable       `json:"table,omitempty"`
	Tags        map[string]string `json:"tags"`
	UpdatedAt   time.Time         `json:"updatedAt"`
	Visual      *ModelVisual      `json:"visual,omitempty"`
	WorkspaceID string            `json:"workspaceId"`
}
