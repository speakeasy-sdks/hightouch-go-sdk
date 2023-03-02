package shared

import (
	"time"
)

// SyncRunFailedRows
// The number of rows that we attempted to sync, but were rejected by the
// destination.
//
// This does not include rows that weren't attempted due to the sync being
// cancelled.
type SyncRunFailedRows struct {
	AddedCount   float64 `json:"addedCount"`
	ChangedCount float64 `json:"changedCount"`
	RemovedCount float64 `json:"removedCount"`
}

// SyncRunPlannedRows
// The number of planned rows that this sync run was supposed to execute.
//
// Note that the counts for `successfulRows` and `failedRows` may not add up
// to `plannedRows` if the sync was cancelled.
type SyncRunPlannedRows struct {
	AddedCount   float64 `json:"addedCount"`
	ChangedCount float64 `json:"changedCount"`
	RemovedCount float64 `json:"removedCount"`
}

// SyncRunSuccessfulRows
// The number of rows that were successfully processed by the destination.
type SyncRunSuccessfulRows struct {
	AddedCount   float64 `json:"addedCount"`
	ChangedCount float64 `json:"changedCount"`
	RemovedCount float64 `json:"removedCount"`
}

type SyncRun struct {
	CompletionRatio float64               `json:"completionRatio"`
	CreatedAt       time.Time             `json:"createdAt"`
	Error           *string               `json:"error,omitempty"`
	FailedRows      SyncRunFailedRows     `json:"failedRows"`
	FinishedAt      time.Time             `json:"finishedAt"`
	ID              string                `json:"id"`
	PlannedRows     SyncRunPlannedRows    `json:"plannedRows"`
	QuerySize       float64               `json:"querySize"`
	StartedAt       time.Time             `json:"startedAt"`
	Status          SyncRunStatusEnum     `json:"status"`
	SuccessfulRows  SyncRunSuccessfulRows `json:"successfulRows"`
}
