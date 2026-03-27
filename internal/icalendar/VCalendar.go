package icalendar

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/parse"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

func (cal *VCalendar) ProcessProperty(prop parse.Property) error {
	switch prop.Name {

	// REQUIRED: MUST NOT occur more than once
	case "PRODID":
		if cal.PRODID != "" {
			return fmt.Errorf("%w: PRODID", parsehelper.ErrDuplicateProperty)
		}
		cal.PRODID = prop.Value

	case "VERSION":
		if cal.VERSION != "" {
			return fmt.Errorf("%w: VERSION", parsehelper.ErrDuplicateProperty)
		}
		cal.VERSION = prop.Value

	// OPTIONAL: MUST NOT occur more than once
	case "CALSCALE":
		if cal.CALSCALE != nil {
			return fmt.Errorf("%w: CALSCALE", parsehelper.ErrDuplicateProperty)
		}
		cal.CALSCALE = parsehelper.StrPtr(prop.Value)

	case "METHOD":
		if cal.METHOD != nil {
			return fmt.Errorf("%w: METHOD", parsehelper.ErrDuplicateProperty)
		}
		cal.METHOD = parsehelper.StrPtr(prop.Value)
	}

	return nil
}

func (cal *VCalendar) Encode(ctx *EncodeContext) {
	WriteProperty(ctx.Builder, "BEGIN", "VCALENDAR")
	WriteProperty(ctx.Builder, "VERSION", cal.VERSION)
	WriteProperty(ctx.Builder, "PRODID", cal.PRODID)

	if cal.METHOD != nil {
		WriteProperty(ctx.Builder, "METHOD", *cal.METHOD)
	}

	for _, ev := range cal.VEVENT {
		ev.Encode(ctx)
	}

	WriteProperty(ctx.Builder, "END", "VCALENDAR")
}

// AddChild is called by the Engine when a nested BEGIN...END block finishes parsing
func (cal *VCalendar) AddChild(child parse.IComponent) error {
	switch v := child.(type) {
	case *VEvent:
		cal.VEVENT = append(cal.VEVENT, *v)
	default:
		// silently ignore unsupported children
	}
	return nil
}

// ensures the VCalendar block is legally constructed
func (cal *VCalendar) Validate() error {
	if cal.PRODID == "" {
		return fmt.Errorf("%w: VCALENDAR missing REQUIRED property PRODID", parsehelper.ErrMissingRequired)
	}
	if cal.VERSION == "" {
		return fmt.Errorf("%w: VCALENDAR missing REQUIRED property VERSION", parsehelper.ErrMissingRequired)
	}
	return nil
}
