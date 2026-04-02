package models

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/parse"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

type VCalendar struct {
	VERSION string
	PRODID  string

	CALSCALE *string
	METHOD   *string

	VEVENT []VEvent
}

const (
	calPropPRODID   = "PRODID"
	calPropVERSION  = "VERSION"
	calPropCALSCALE = "CALSCALE"
	calPropMETHOD   = "METHOD"
)

func (cal *VCalendar) ProcessProperty(prop parse.Property) error {
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

func (cal *VCalendar) AddChild(child parse.Component) error {
	switch c := child.(type) {
	case *VEvent:
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
