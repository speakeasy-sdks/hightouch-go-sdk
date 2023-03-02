package shared

import (
	"time"
)

// SyncSchedule
// The scheduling configuration. It can be triggerd based on several ways:
//
// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
//
// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
//
// Visual: the sync will be trigged based a visual cron configuration on UI
//
// DBT-cloud: the sync will be trigged based on a dbt cloud job
type SyncSchedule struct {
	Schedule interface{} `json:"schedule"`
	Type     string      `json:"type"`
}

// Sync
// Syncs define how data from models are mapped to destinations. Each time a
// sync runs, Hightouch calculates the rows that have changed since the last
// run, and syncs them to Sync's destination.
type Sync struct {
	Configuration     map[string]interface{} `json:"configuration"`
	CreatedAt         time.Time              `json:"createdAt"`
	DestinationID     string                 `json:"destinationId"`
	Disabled          bool                   `json:"disabled"`
	ID                string                 `json:"id"`
	LastRunAt         time.Time              `json:"lastRunAt"`
	ModelID           string                 `json:"modelId"`
	PrimaryKey        string                 `json:"primaryKey"`
	ReferencedColumns []string               `json:"referencedColumns"`
	Schedule          SyncSchedule           `json:"schedule"`
	Slug              string                 `json:"slug"`
	Status            SyncStatusEnum         `json:"status"`
	UpdatedAt         time.Time              `json:"updatedAt"`
	WorkspaceID       string                 `json:"workspaceId"`
}
