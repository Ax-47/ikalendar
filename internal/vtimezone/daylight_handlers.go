package vtimezone

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
)

type DaylightHandler func(*Daylight, componants.Property) error

const (
	PropDTSTART      share.PropertyName = "DTSTART"
	PropTZOFFSETFROM share.PropertyName = "TZOFFSETFROM"
	PropTZOFFSETTO   share.PropertyName = "TZOFFSETTO"

	// Optional
	PropTZNAME share.PropertyName = "TZNAME"

	// Recurrence
	PropRRULE share.PropertyName = "RRULE"
	PropRDATE share.PropertyName = "RDATE"
)

var daylightHandlers = map[share.PropertyName]DaylightHandler{
	PropDTSTART:      handleDaylightDTSTART,
	PropTZOFFSETFROM: handleDaylightTZOffsetFrom,
	PropTZOFFSETTO:   handleDaylightTZOffsetTo,
	PropTZNAME:       handleDaylightTZName,
	PropRRULE:        handleDaylightRRule,
	PropRDATE:        handleDaylightRDate,
}

// ── Required ──────────────────────────────────────────────────────────────────

func handleDaylightDTSTART(d *Daylight, prop componants.Property) error {
	return d.SetDTSTART(share.ParseITIME(prop.Params, prop.Value))
}

func handleDaylightTZOffsetFrom(d *Daylight, prop componants.Property) error {
	return d.SetTZOffsetFrom(prop.Value)
}

func handleDaylightTZOffsetTo(d *Daylight, prop componants.Property) error {
	return d.SetTZOffsetTo(prop.Value)
}

func handleDaylightTZName(d *Daylight, prop componants.Property) error {
	return d.SetTZName(prop.Value)
}

func handleDaylightRRule(d *Daylight, prop componants.Property) error {
	r, err := share.ParseRECUR(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid RRULE: %w", err)
	}
	return d.SetRRule(r)
}

func handleDaylightRDate(d *Daylight, prop componants.Property) error {
	return d.AddRDate(share.ParseITIME(prop.Params, prop.Value))
}
