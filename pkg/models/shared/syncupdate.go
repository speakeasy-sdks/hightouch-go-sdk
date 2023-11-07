// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
)

type SyncUpdateSchemasScheduleType string

const (
	SyncUpdateSchemasScheduleTypeIntervalSchedule   SyncUpdateSchemasScheduleType = "IntervalSchedule"
	SyncUpdateSchemasScheduleTypeCronSchedule       SyncUpdateSchemasScheduleType = "CronSchedule"
	SyncUpdateSchemasScheduleTypeVisualCronSchedule SyncUpdateSchemasScheduleType = "VisualCronSchedule"
	SyncUpdateSchemasScheduleTypeDBTSchedule        SyncUpdateSchemasScheduleType = "DBTSchedule"
)

type SyncUpdateSchemasSchedule struct {
	IntervalSchedule   *IntervalSchedule
	CronSchedule       *CronSchedule
	VisualCronSchedule *VisualCronSchedule
	DBTSchedule        *DBTSchedule

	Type SyncUpdateSchemasScheduleType
}

func CreateSyncUpdateSchemasScheduleIntervalSchedule(intervalSchedule IntervalSchedule) SyncUpdateSchemasSchedule {
	typ := SyncUpdateSchemasScheduleTypeIntervalSchedule

	return SyncUpdateSchemasSchedule{
		IntervalSchedule: &intervalSchedule,
		Type:             typ,
	}
}

func CreateSyncUpdateSchemasScheduleCronSchedule(cronSchedule CronSchedule) SyncUpdateSchemasSchedule {
	typ := SyncUpdateSchemasScheduleTypeCronSchedule

	return SyncUpdateSchemasSchedule{
		CronSchedule: &cronSchedule,
		Type:         typ,
	}
}

func CreateSyncUpdateSchemasScheduleVisualCronSchedule(visualCronSchedule VisualCronSchedule) SyncUpdateSchemasSchedule {
	typ := SyncUpdateSchemasScheduleTypeVisualCronSchedule

	return SyncUpdateSchemasSchedule{
		VisualCronSchedule: &visualCronSchedule,
		Type:               typ,
	}
}

func CreateSyncUpdateSchemasScheduleDBTSchedule(dbtSchedule DBTSchedule) SyncUpdateSchemasSchedule {
	typ := SyncUpdateSchemasScheduleTypeDBTSchedule

	return SyncUpdateSchemasSchedule{
		DBTSchedule: &dbtSchedule,
		Type:        typ,
	}
}

func (u *SyncUpdateSchemasSchedule) UnmarshalJSON(data []byte) error {

	intervalSchedule := IntervalSchedule{}
	if err := utils.UnmarshalJSON(data, &intervalSchedule, "", true, true); err == nil {
		u.IntervalSchedule = &intervalSchedule
		u.Type = SyncUpdateSchemasScheduleTypeIntervalSchedule
		return nil
	}

	cronSchedule := CronSchedule{}
	if err := utils.UnmarshalJSON(data, &cronSchedule, "", true, true); err == nil {
		u.CronSchedule = &cronSchedule
		u.Type = SyncUpdateSchemasScheduleTypeCronSchedule
		return nil
	}

	visualCronSchedule := VisualCronSchedule{}
	if err := utils.UnmarshalJSON(data, &visualCronSchedule, "", true, true); err == nil {
		u.VisualCronSchedule = &visualCronSchedule
		u.Type = SyncUpdateSchemasScheduleTypeVisualCronSchedule
		return nil
	}

	dbtSchedule := DBTSchedule{}
	if err := utils.UnmarshalJSON(data, &dbtSchedule, "", true, true); err == nil {
		u.DBTSchedule = &dbtSchedule
		u.Type = SyncUpdateSchemasScheduleTypeDBTSchedule
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SyncUpdateSchemasSchedule) MarshalJSON() ([]byte, error) {
	if u.IntervalSchedule != nil {
		return utils.MarshalJSON(u.IntervalSchedule, "", true)
	}

	if u.CronSchedule != nil {
		return utils.MarshalJSON(u.CronSchedule, "", true)
	}

	if u.VisualCronSchedule != nil {
		return utils.MarshalJSON(u.VisualCronSchedule, "", true)
	}

	if u.DBTSchedule != nil {
		return utils.MarshalJSON(u.DBTSchedule, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// SyncUpdateSchedule - The scheduling configuration. It can be triggerd based on several ways:
//
// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
//
// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
//
// Visual: the sync will be trigged based a visual cron configuration on UI
//
// DBT-cloud: the sync will be trigged based on a dbt cloud job
type SyncUpdateSchedule struct {
	Schedule SyncUpdateSchemasSchedule `json:"schedule"`
	Type     string                    `json:"type"`
}

func (o *SyncUpdateSchedule) GetSchedule() SyncUpdateSchemasSchedule {
	if o == nil {
		return SyncUpdateSchemasSchedule{}
	}
	return o.Schedule
}

func (o *SyncUpdateSchedule) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

// SyncUpdate - The input for updating a Sync
type SyncUpdate struct {
	// The sync's configuration. This specifies how data is mapped, among other
	// configuration.
	//
	// The schema depends on the destination type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	// Whether the sync has been disabled by the user.
	Disabled *bool `json:"disabled,omitempty"`
	// The scheduling configuration. It can be triggerd based on several ways:
	//
	// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
	//
	// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
	//
	// Visual: the sync will be trigged based a visual cron configuration on UI
	//
	// DBT-cloud: the sync will be trigged based on a dbt cloud job
	Schedule *SyncUpdateSchedule `json:"schedule,omitempty"`
}

func (o *SyncUpdate) GetConfiguration() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Configuration
}

func (o *SyncUpdate) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *SyncUpdate) GetSchedule() *SyncUpdateSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}
