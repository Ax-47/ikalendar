package vtimezone

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
)

type VTimezoneHandler func(*VTimezone, componants.Property) error

const (
	// Required
	PropTZID         share.PropertyName = "TZID"
	PropLASTMODIFIED share.PropertyName = "LASTMODIFIED"
	PropTZURL        share.PropertyName = "TZURL"
) // RFC 5545 §3.3.6
var vtimezoneHandlers = map[share.PropertyName]VTimezoneHandler{
	PropTZID:         handleTZID,
	PropLASTMODIFIED: handleLastModified,
	PropTZURL:        handleTZURL,
}

func handleTZID(tz *VTimezone, prop componants.Property) error {
	return tz.SetTZID(prop.Value)
}

func handleLastModified(tz *VTimezone, prop componants.Property) error {
	return tz.SetLastModified(share.ParseITIME(prop.Params, prop.Value))
}

func handleTZURL(tz *VTimezone, prop componants.Property) error {
	return tz.SetTZURL(prop.Value)
}
