package vtimezone

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
)

type VTimezone struct {
	/* REQUIRED */
	TZID string

	/* OPTIONAL */
	LASTMODIFIED *share.ITIME
	TZURL        *string

	/* CHILD COMPONENTS */
	STANDARD []Standard
	DAYLIGHT []Daylight
}

func (tz *VTimezone) ProcessProperty(prop componants.Property) error {
	name := share.PropertyName(prop.Name)
	if handler, ok := vtimezoneHandlers[name]; ok {
		return handler(tz, prop)
	}
	// ignore unknown properties per RFC 5545 §3.8.8 (x-prop / iana-prop)
	return nil
}

func (tz *VTimezone) Validate() error {
	if tz.TZID == "" {
		return fmt.Errorf("VTIMEZONE: TZID is required")
	}
	if len(tz.STANDARD) == 0 && len(tz.DAYLIGHT) == 0 {
		return fmt.Errorf("VTIMEZONE: must have at least STANDARD or DAYLIGHT")
	}
	return nil
}

func (tz *VTimezone) AddChild(c componants.Component) error {
	switch child := c.(type) {
	case *Standard:
		return tz.AddStandard(*child)
	case *Daylight:
		return tz.AddDaylight(*child)
	default:
		return fmt.Errorf("VTIMEZONE: unsupported child component")
	}
}
