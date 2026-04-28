package valarm

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

func (a *VAlarm) SetAction(value string) error {
	return utils.SetOnce(&a.Action, new(value), string(propAction))
}

func (a *VAlarm) SetTrigger(value string) error {
	d, err := share.ParseDURATION(value)
	if err != nil {
		return fmt.Errorf("invalid %s: %w", propTrigger, err)
	}
	return utils.SetOnce(&a.Trigger, new(d), string(propTrigger))
}

func (a *VAlarm) SetDescription(value string) error {
	return utils.SetOnce(&a.Description, new(value), string(propDescription))
}

func (a *VAlarm) SetSummary(value string) error {
	return utils.SetOnce(&a.Summary, new(value), string(propSummary))
}
