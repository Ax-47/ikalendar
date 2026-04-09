package valarm

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

const (
	propAction      share.PropertyName = "ACTION"
	propTrigger     share.PropertyName = "TRIGGER"
	propDescription share.PropertyName = "DESCRIPTION"
	propSummary     share.PropertyName = "SUMMARY"
)

type VAlarm struct {
	Action      *string
	Trigger     *share.DURATION
	Description *string
	Summary     *string
}

func (a *VAlarm) ProcessProperty(prop componants.Property) error {
	switch prop.Name {
	case "ACTION":
		if a.Action != nil {
			return fmt.Errorf("%w: ACTION", utils.ErrDuplicateProperty)
		}
		a.Action = utils.Ptr(prop.Value)
	case "TRIGGER":
		if a.Trigger != nil {
			return fmt.Errorf("%w: TRIGGER", utils.ErrDuplicateProperty)
		}
		d, err := share.ParseDURATION(prop.Value)
		if err != nil {
			return fmt.Errorf("invalid TRIGGER: %w", err)
		}
		a.Trigger = utils.Ptr(d)
	case "DESCRIPTION":
		if a.Description != nil {
			return fmt.Errorf("%w: DESCRIPTION", utils.ErrDuplicateProperty)
		}
		a.Description = utils.Ptr(prop.Value)
	case "SUMMARY":
		if a.Summary != nil {
			return fmt.Errorf("%w: SUMMARY", utils.ErrDuplicateProperty)
		}
		a.Summary = utils.Ptr(prop.Value)
	}
	return nil
}

func (a *VAlarm) AddChild(child componants.Component) error {
	return fmt.Errorf("%w: VALARM cannot contain %T",
		utils.ErrInvalidComponent, child)
}

func (a *VAlarm) Validate() error {
	if a.Action == nil {
		return fmt.Errorf("%w: VALARM missing ACTION", utils.ErrMissingRequired)
	}
	if a.Trigger == nil {
		return fmt.Errorf("%w: VALARM missing TRIGGER", utils.ErrMissingRequired)
	}
	return nil
}
