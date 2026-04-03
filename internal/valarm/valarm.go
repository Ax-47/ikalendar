package valarm

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
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
			return fmt.Errorf("%w: ACTION", parsehelper.ErrDuplicateProperty)
		}
		a.Action = parsehelper.Ptr(prop.Value)
	case "TRIGGER":
		if a.Trigger != nil {
			return fmt.Errorf("%w: TRIGGER", parsehelper.ErrDuplicateProperty)
		}
		d, err := share.ParseDURATION(prop.Value)
		if err != nil {
			return fmt.Errorf("invalid TRIGGER: %w", err)
		}
		a.Trigger = parsehelper.Ptr(d)
	case "DESCRIPTION":
		if a.Description != nil {
			return fmt.Errorf("%w: DESCRIPTION", parsehelper.ErrDuplicateProperty)
		}
		a.Description = parsehelper.Ptr(prop.Value)
	case "SUMMARY":
		if a.Summary != nil {
			return fmt.Errorf("%w: SUMMARY", parsehelper.ErrDuplicateProperty)
		}
		a.Summary = parsehelper.Ptr(prop.Value)
	}
	return nil
}

func (a *VAlarm) AddChild(child componants.Component) error {
	return fmt.Errorf("%w: VALARM cannot contain %T",
		parsehelper.ErrInvalidComponent, child)
}

func (a *VAlarm) Validate() error {
	if a.Action == nil {
		return fmt.Errorf("%w: VALARM missing ACTION", parsehelper.ErrMissingRequired)
	}
	if a.Trigger == nil {
		return fmt.Errorf("%w: VALARM missing TRIGGER", parsehelper.ErrMissingRequired)
	}
	return nil
}
