package shared

import (
	"time"
)

// Source
// The database or warehouse where your data is stored. The starting point for
// a Hightouch data pipeline.
type Source struct {
	Configuration map[string]interface{} `json:"configuration"`
	CreatedAt     time.Time              `json:"createdAt"`
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Slug          string                 `json:"slug"`
	Type          string                 `json:"type"`
	UpdatedAt     time.Time              `json:"updatedAt"`
	WorkspaceID   string                 `json:"workspaceId"`
}
