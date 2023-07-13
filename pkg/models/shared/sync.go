// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// SyncSchedule - The scheduling configuration. It can be triggerd based on several ways:
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

// Sync - Syncs define how data from models are mapped to destinations. Each time a
// sync runs, Hightouch calculates the rows that have changed since the last
// run, and syncs them to Sync's destination.
type Sync struct {
	// The sync's configuration. This specifies how data is mapped, among other
	// configuration.
	//
	// The schema depends on the destination type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration"`
	// The timestamp when the sync was created
	CreatedAt time.Time `json:"createdAt"`
	// The id of the Destination that sync is connected to
	DestinationID string `json:"destinationId"`
	// Whether the sync has been disabled by the user.
	Disabled bool `json:"disabled"`
	// The sync's id
	ID string `json:"id"`
	// The timestamp of the last sync run
	LastRunAt time.Time `json:"lastRunAt"`
	// The id of the Model that sync is connected to
	ModelID string `json:"modelId"`
	// The primary key that sync uses to identify data from source
	PrimaryKey string `json:"primaryKey"`
	// The reference column that sync depends on to sync data from source
	ReferencedColumns []string `json:"referencedColumns"`
	// The scheduling configuration. It can be triggerd based on several ways:
	//
	// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
	//
	// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
	//
	// Visual: the sync will be trigged based a visual cron configuration on UI
	//
	// DBT-cloud: the sync will be trigged based on a dbt cloud job
	Schedule SyncSchedule `json:"schedule"`
	// The sync's slug
	Slug string `json:"slug"`
	// SyncStatus
	Status SyncStatus `json:"status"`
	// The timestamp when the sync was last updated
	UpdatedAt time.Time `json:"updatedAt"`
	// The id of the workspace that the sync belongs to
	WorkspaceID string `json:"workspaceId"`
}
