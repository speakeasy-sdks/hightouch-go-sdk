package shared

type SyncRunStatusEnum string

const (
	SyncRunStatusEnumCancelled   SyncRunStatusEnum = "cancelled"
	SyncRunStatusEnumFailed      SyncRunStatusEnum = "failed"
	SyncRunStatusEnumQueued      SyncRunStatusEnum = "queued"
	SyncRunStatusEnumSuccess     SyncRunStatusEnum = "success"
	SyncRunStatusEnumWarning     SyncRunStatusEnum = "warning"
	SyncRunStatusEnumQuerying    SyncRunStatusEnum = "querying"
	SyncRunStatusEnumProcessing  SyncRunStatusEnum = "processing"
	SyncRunStatusEnumReporting   SyncRunStatusEnum = "reporting"
	SyncRunStatusEnumInterrupted SyncRunStatusEnum = "interrupted"
)
