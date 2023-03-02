package shared

type VisualCronScheduleExpressions struct {
	Days RecordDayBooleanOrUndefined `json:"days"`
	Time string                      `json:"time"`
}

type VisualCronSchedule struct {
	Expressions []VisualCronScheduleExpressions `json:"expressions"`
}
