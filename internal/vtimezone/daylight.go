package vtimezone

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
)

type Daylight struct {
	DTSTART      share.ITIME
	TZOFFSETFROM string
	TZOFFSETTO   string

	RRULE  *share.RECUR
	RDATE  []share.ITIME
	TZNAME *string
}

func (d *Daylight) ProcessProperty(prop componants.Property) error {
	name := share.PropertyName(prop.Name)
	if handler, ok := daylightHandlers[name]; ok {
		return handler(d, prop)
	}
	// ignore unknown properties per RFC 5545 §3.8.8 (x-prop / iana-prop)
	return nil
}

func (d *Daylight) Validate() error {
	if d.DTSTART.IsZero() {
		return fmt.Errorf("DAYLIGHT: DTSTART is required")
	}
	if d.TZOFFSETFROM == "" {
		return fmt.Errorf("DAYLIGHT: TZOFFSETFROM is required")
	}
	if d.TZOFFSETTO == "" {
		return fmt.Errorf("DAYLIGHT: TZOFFSETTO is required")
	}
	return nil
}

func (d *Daylight) AddChild(child componants.Component) error {
	return fmt.Errorf("DAYLIGHT: does not support child components")
}
