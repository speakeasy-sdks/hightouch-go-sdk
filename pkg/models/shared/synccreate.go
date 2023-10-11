// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/hightouch-go-sdk/pkg/utils"
)

type SyncCreateScheduleScheduleType string

const (
	SyncCreateScheduleScheduleTypeIntervalSchedule   SyncCreateScheduleScheduleType = "IntervalSchedule"
	SyncCreateScheduleScheduleTypeCronSchedule       SyncCreateScheduleScheduleType = "CronSchedule"
	SyncCreateScheduleScheduleTypeVisualCronSchedule SyncCreateScheduleScheduleType = "VisualCronSchedule"
	SyncCreateScheduleScheduleTypeDBTSchedule        SyncCreateScheduleScheduleType = "DBTSchedule"
)

type SyncCreateScheduleSchedule struct {
	IntervalSchedule   *IntervalSchedule
	CronSchedule       *CronSchedule
	VisualCronSchedule *VisualCronSchedule
	DBTSchedule        *DBTSchedule

	Type SyncCreateScheduleScheduleType
}

func CreateSyncCreateScheduleScheduleIntervalSchedule(intervalSchedule IntervalSchedule) SyncCreateScheduleSchedule {
	typ := SyncCreateScheduleScheduleTypeIntervalSchedule

	return SyncCreateScheduleSchedule{
		IntervalSchedule: &intervalSchedule,
		Type:             typ,
	}
}

func CreateSyncCreateScheduleScheduleCronSchedule(cronSchedule CronSchedule) SyncCreateScheduleSchedule {
	typ := SyncCreateScheduleScheduleTypeCronSchedule

	return SyncCreateScheduleSchedule{
		CronSchedule: &cronSchedule,
		Type:         typ,
	}
}

func CreateSyncCreateScheduleScheduleVisualCronSchedule(visualCronSchedule VisualCronSchedule) SyncCreateScheduleSchedule {
	typ := SyncCreateScheduleScheduleTypeVisualCronSchedule

	return SyncCreateScheduleSchedule{
		VisualCronSchedule: &visualCronSchedule,
		Type:               typ,
	}
}

func CreateSyncCreateScheduleScheduleDBTSchedule(dbtSchedule DBTSchedule) SyncCreateScheduleSchedule {
	typ := SyncCreateScheduleScheduleTypeDBTSchedule

	return SyncCreateScheduleSchedule{
		DBTSchedule: &dbtSchedule,
		Type:        typ,
	}
}

func (u *SyncCreateScheduleSchedule) UnmarshalJSON(data []byte) error {

	intervalSchedule := new(IntervalSchedule)
	if err := utils.UnmarshalJSON(data, &intervalSchedule, "", true, true); err == nil {
		u.IntervalSchedule = intervalSchedule
		u.Type = SyncCreateScheduleScheduleTypeIntervalSchedule
		return nil
	}

	cronSchedule := new(CronSchedule)
	if err := utils.UnmarshalJSON(data, &cronSchedule, "", true, true); err == nil {
		u.CronSchedule = cronSchedule
		u.Type = SyncCreateScheduleScheduleTypeCronSchedule
		return nil
	}

	visualCronSchedule := new(VisualCronSchedule)
	if err := utils.UnmarshalJSON(data, &visualCronSchedule, "", true, true); err == nil {
		u.VisualCronSchedule = visualCronSchedule
		u.Type = SyncCreateScheduleScheduleTypeVisualCronSchedule
		return nil
	}

	dbtSchedule := new(DBTSchedule)
	if err := utils.UnmarshalJSON(data, &dbtSchedule, "", true, true); err == nil {
		u.DBTSchedule = dbtSchedule
		u.Type = SyncCreateScheduleScheduleTypeDBTSchedule
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SyncCreateScheduleSchedule) MarshalJSON() ([]byte, error) {
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

// SyncCreateSchedule - The scheduling configuration. It can be triggerd based on several ways:
//
// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
//
// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
//
// Visual: the sync will be trigged based a visual cron configuration on UI
//
// DBT-cloud: the sync will be trigged based on a dbt cloud job
type SyncCreateSchedule struct {
	Schedule SyncCreateScheduleSchedule `json:"schedule"`
	Type     string                     `json:"type"`
}

func (o *SyncCreateSchedule) GetSchedule() SyncCreateScheduleSchedule {
	if o == nil {
		return SyncCreateScheduleSchedule{}
	}
	return o.Schedule
}

func (o *SyncCreateSchedule) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

// SyncCreate - The input for creating a Sync
type SyncCreate struct {
	// The sync's configuration. This specifies how data is mapped, among other
	// configuration.
	//
	// The schema depends on the destination type.
	//
	// Consumers should NOT make assumptions on the contents of the
	// configuration. It may change as Hightouch updates its internal code.
	Configuration map[string]interface{} `json:"configuration"`
	// The id of the Destination that sync is connected to
	DestinationID string `json:"destinationId"`
	// Whether the sync has been disabled by the user.
	Disabled bool `json:"disabled"`
	// The id of the Model that sync is connected to
	ModelID string `json:"modelId"`
	// The scheduling configuration. It can be triggerd based on several ways:
	//
	// Interval: the sync will be trigged based on certain interval(minutes/hours/days/weeks)
	//
	// Cron: the sync will be trigged based on cron expression https://en.wikipedia.org/wiki/Cron.
	//
	// Visual: the sync will be trigged based a visual cron configuration on UI
	//
	// DBT-cloud: the sync will be trigged based on a dbt cloud job
	Schedule *SyncCreateSchedule `json:"schedule"`
	// The sync's slug
	Slug string `json:"slug"`
}

func (o *SyncCreate) GetConfiguration() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Configuration
}

func (o *SyncCreate) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *SyncCreate) GetDisabled() bool {
	if o == nil {
		return false
	}
	return o.Disabled
}

func (o *SyncCreate) GetModelID() string {
	if o == nil {
		return ""
	}
	return o.ModelID
}

func (o *SyncCreate) GetSchedule() *SyncCreateSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *SyncCreate) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}
