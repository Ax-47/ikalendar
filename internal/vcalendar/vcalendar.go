package vcalendar

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/vevent"
)

type VCalendar struct {
	VERSION string
	PRODID  string

	CALSCALE *string
	METHOD   *string

	VEVENT []vevent.VEvent
}

const (
	calPropPRODID   = "PRODID"
	calPropVERSION  = "VERSION"
	calPropCALSCALE = "CALSCALE"
	calPropMETHOD   = "METHOD"
)

func (c *VCalendar) GetMethod() *string {
	return c.METHOD
}

func (cal *VCalendar) ProcessProperty(prop componants.Property) error {
	switch prop.Name {
	case calPropPRODID:
		return cal.SetPRODID(prop.Value)
	case calPropVERSION:
		return cal.SetVERSION(prop.Value)
	case calPropCALSCALE:
		return cal.SetCALSCALE(prop.Value)
	case calPropMETHOD:
		return cal.SetMETHOD(prop.Value)
	}
	return nil
}

func (cal *VCalendar) AddChild(child componants.Component) error {
	switch c := child.(type) {
	case *vevent.VEvent:
		cal.VEVENT = append(cal.VEVENT, *c)
		return nil
	default:
		return fmt.Errorf("%w: VCALENDAR cannot contain %T",
			parsehelper.ErrInvalidComponent, child)
	}
}

func (cal *VCalendar) Validate() error {
	if cal.PRODID == "" {
		return fmt.Errorf("%w: VCALENDAR missing PRODID", parsehelper.ErrMissingRequired)
	}
	if cal.VERSION == "" {
		return fmt.Errorf("%w: VCALENDAR missing VERSION", parsehelper.ErrMissingRequired)
	}
	return nil
}
