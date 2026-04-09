package ikalendar

import "github.com/minoplhy/ikalendar/internal/valarm"

type AlarmOption = valarm.AlarmOption

var (
	NewAlarm             = valarm.NewAlarm
	WithAction           = valarm.WithAction
	WithTrigger          = valarm.WithTrigger
	WithTriggerBefore    = valarm.WithTriggerBefore
	WithTriggerAfter     = valarm.WithTriggerAfter
	WithTriggerDays      = valarm.WithTriggerDays
	WithAlarmDescription = valarm.WithAlarmDescription
	WithAlarmSummary     = valarm.WithAlarmSummary
)
