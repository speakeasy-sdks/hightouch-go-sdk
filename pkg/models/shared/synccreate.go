package shared

// SyncCreateSchedule
// The scheduling configuration. It can be triggerd based on several ways:
//
// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
//
// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
//
// Visual: the sync will be trigged based a visual cron configuration on UI
//
// DBT-cloud: the sync will be trigged based on a dbt cloud job
type SyncCreateSchedule struct {
	Schedule interface{} `json:"schedule"`
	Type     string      `json:"type"`
}

// SyncCreate
// The input for creating a Sync
type SyncCreate struct {
	Configuration map[string]interface{} `json:"configuration"`
	DestinationID string                 `json:"destinationId"`
	Disabled      bool                   `json:"disabled"`
	ModelID       string                 `json:"modelId"`
	Schedule      SyncCreateSchedule     `json:"schedule"`
	Slug          string                 `json:"slug"`
}
