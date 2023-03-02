package shared

type SyncStatusEnum string

const (
	SyncStatusEnumDisabled    SyncStatusEnum = "disabled"
	SyncStatusEnumPending     SyncStatusEnum = "pending"
	SyncStatusEnumCancelled   SyncStatusEnum = "cancelled"
	SyncStatusEnumFailed      SyncStatusEnum = "failed"
	SyncStatusEnumQueued      SyncStatusEnum = "queued"
	SyncStatusEnumSuccess     SyncStatusEnum = "success"
	SyncStatusEnumWarning     SyncStatusEnum = "warning"
	SyncStatusEnumQuerying    SyncStatusEnum = "querying"
	SyncStatusEnumProcessing  SyncStatusEnum = "processing"
	SyncStatusEnumReporting   SyncStatusEnum = "reporting"
	SyncStatusEnumInterrupted SyncStatusEnum = "interrupted"
)
