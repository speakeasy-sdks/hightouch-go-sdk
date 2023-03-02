package shared

import (
	"time"
)

// Destination
// The service receiving your data (e.g. Salesforce, Hubspot, Customer.io, or a
// SFTP server)
type Destination struct {
	Configuration map[string]interface{} `json:"configuration"`
	CreatedAt     time.Time              `json:"createdAt"`
	ID            string                 `json:"id"`
	Name          string                 `json:"name"`
	Slug          string                 `json:"slug"`
	Syncs         []string               `json:"syncs"`
	Type          string                 `json:"type"`
	UpdatedAt     time.Time              `json:"updatedAt"`
	WorkspaceID   string                 `json:"workspaceId"`
}
