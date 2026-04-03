package vevent

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/valarm"
)

type VEvent struct {
	/* REQUIRED: MUST NOT occur more than once*/
	UID     string
	DTSTAMP share.ITIME

	/* REQUIRED if method isn't specified MUST NOT occur more than once */
	DTSTART *share.ITIME

	/* OPTIONAL: MUST NOT occur more than once*/
	CLASS *string

	CREATED      *share.ITIME
	DESCRIPTION  *string
	GEO          *string
	LASTMODIFIED *share.ITIME
	LOCATION     *string
	ORGANIZER    *string
	PRIORITY     *int
	SEQUENCE     *int
	STATUS       *string
	SUMMARY      *string
	TRANSP       *string
	URL          *string

	// child component
	VALARM []valarm.VAlarm

	/* OPTIONAL: SHOULD NOT occur more than once*/
	RRULE *share.RECUR

	/* OPTIONAL: SHOULD NOT occur in same VEVENT*/
	DTEND    *share.ITIME
	DURATION *share.DURATION

	/* OPTIONAL: MAY occur more than once*/
	ATTACH        []share.ATTACH
	ATTENDEE      []share.CALADDRESS
	CATEGORIES    []string
	COMMENT       []string
	CONTACT       []string
	EXDATE        []share.ITIME
	REQUESTSTATUS []share.RequestStatus
	RELATED       []share.RELATED
	RESOURCES     []string
	RDATE         []share.ITIME
}

func (ev *VEvent) ProcessProperty(prop componants.Property) error {
	name := share.PropertyName(prop.Name)
	if handler, ok := veventHandlers[name]; ok {
		return handler(ev, prop)
	}
	// ignore unknown properties per RFC 5545 §3.8.8 (x-prop / iana-prop)
	return nil
}

func (ev *VEvent) AddChild(child componants.Component) error {
	switch c := child.(type) {
	case *valarm.VAlarm:
		ev.VALARM = append(ev.VALARM, *c)
		return nil
	default:
		return fmt.Errorf("%w: VEVENT cannot contain %T",
			parsehelper.ErrInvalidComponent, child)
	}
}

func (ev *VEvent) Validate() error {
	if ev.UID == "" {
		return fmt.Errorf("%w: VEVENT missing UID", parsehelper.ErrMissingRequired)
	}
	if ev.DTSTAMP.Time.IsZero() {
		return fmt.Errorf("%w: VEVENT missing DTSTAMP", parsehelper.ErrMissingRequired)
	}
	if ev.DTEND != nil && ev.DURATION != nil {
		return fmt.Errorf("%w: VEVENT cannot have both DTEND and DURATION",
			parsehelper.ErrMutuallyExclusive)
	}

	return nil
}
