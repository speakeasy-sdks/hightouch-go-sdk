package shared

// SyncUpdateSchedule
// The scheduling configuration. It can be triggerd based on several ways:
//
// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
//
// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
//
// Visual: the sync will be trigged based a visual cron configuration on UI
//
// DBT-cloud: the sync will be trigged based on a dbt cloud job
type SyncUpdateSchedule struct {
	Schedule interface{} `json:"schedule"`
	Type     string      `json:"type"`
}

// SyncUpdate
// The input for updating a Sync
type SyncUpdate struct {
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	Disabled      *bool                  `json:"disabled,omitempty"`
	Schedule      *SyncUpdateSchedule    `json:"schedule,omitempty"`
}
