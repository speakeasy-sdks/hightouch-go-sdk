package shared

type DBTScheduleAccount struct {
	ID string `json:"id"`
}

type DBTScheduleJob struct {
	ID string `json:"id"`
}

type DBTSchedule struct {
	Account         DBTScheduleAccount `json:"account"`
	DbtCredentialID string             `json:"dbtCredentialId"`
	Job             DBTScheduleJob     `json:"job"`
}
