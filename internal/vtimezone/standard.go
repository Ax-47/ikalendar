package vtimezone

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
)

type Standard struct {
	DTSTART      share.ITIME
	TZOFFSETFROM string
	TZOFFSETTO   string

	RRULE  *share.RECUR
	RDATE  []share.ITIME
	TZNAME *string
}

func (s *Standard) ProcessProperty(prop componants.Property) error {
	name := share.PropertyName(prop.Name)
	if handler, ok := standardHandlers[name]; ok {
		return handler(s, prop)
	}
	// ignore unknown properties per RFC 5545 §3.8.8 (x-prop / iana-prop)
	return nil
}

func (s *Standard) Validate() error {
	if s.DTSTART.IsZero() {
		return fmt.Errorf("DAYLIGHT: DTSTART is required")
	}
	if s.TZOFFSETFROM == "" {
		return fmt.Errorf("DAYLIGHT: TZOFFSETFROM is required")
	}
	if s.TZOFFSETTO == "" {
		return fmt.Errorf("DAYLIGHT: TZOFFSETTO is required")
	}
	return nil
}

func (s *Standard) AddChild(child componants.Component) error {
	return fmt.Errorf("STANDARD: does not support child components")
}
