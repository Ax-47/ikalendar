package ikalendar

import "github.com/minoplhy/ikalendar/internal/icalendar/builders"

type AlarmOption = builders.AlarmOption

var (
	NewAlarm             = builders.NewAlarm
	WithAction           = builders.WithAction
	WithTrigger          = builders.WithTrigger
	WithTriggerBefore    = builders.WithTriggerBefore
	WithTriggerAfter     = builders.WithTriggerAfter
	WithTriggerDays      = builders.WithTriggerDays
	WithAlarmDescription = builders.WithAlarmDescription
	WithAlarmSummary     = builders.WithAlarmSummary
)
