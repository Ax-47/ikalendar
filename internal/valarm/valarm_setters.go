package valarm

import (
	"fmt"

	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
)

func (a *VAlarm) SetAction(value string) error {
	return parsehelper.SetOnce(&a.Action, parsehelper.Ptr(value), string(propAction))
}

func (a *VAlarm) SetTrigger(value string) error {
	d, err := share.ParseDURATION(value)
	if err != nil {
		return fmt.Errorf("invalid %s: %w", propTrigger, err)
	}
	return parsehelper.SetOnce(&a.Trigger, parsehelper.Ptr(d), string(propTrigger))
}

func (a *VAlarm) SetDescription(value string) error {
	return parsehelper.SetOnce(&a.Description, parsehelper.Ptr(value), string(propDescription))
}

func (a *VAlarm) SetSummary(value string) error {
	return parsehelper.SetOnce(&a.Summary, parsehelper.Ptr(value), string(propSummary))
}
