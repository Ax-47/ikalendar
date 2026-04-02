package models

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/parse"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

type VEvent struct {
	/* REQUIRED: MUST NOT occur more than once*/
	UID     string
	DTSTAMP ITIME

	/* REQUIRED if method isn't specified MUST NOT occur more than once */
	DTSTART *ITIME

	/* OPTIONAL: MUST NOT occur more than once*/
	CLASS        *string
	CREATED      *ITIME
	DESCRIPTION  *string
	GEO          *string
	LASTMODIFIED *ITIME
	LOCATION     *string
	ORGANIZER    *string
	PRIORITY     *int
	SEQUENCE     *int
	STATUS       *string
	SUMMARY      *string
	TRANSP       *string
	URL          *string

	// child component
	VALARM []VAlarm

	/* OPTIONAL: SHOULD NOT occur more than once*/
	RRULE *RECUR

	/* OPTIONAL: SHOULD NOT occur in same VEVENT*/
	DTEND    *ITIME
	DURATION *DURATION

	/* OPTIONAL: MAY occur more than once*/
	ATTACH        []ATTACH
	ATTENDEE      []CALADDRESS
	CATEGORIES    []string
	COMMENT       []string
	CONTACT       []string
	EXDATE        []ITIME
	REQUESTSTATUS []RequestStatus
	RELATED       []RELATED
	RESOURCES     []string
	RDATE         []ITIME
}

func (ev *VEvent) ProcessProperty(prop parse.Property) error {
	name := PropertyName(prop.Name)
	if handler, ok := veventHandlers[name]; ok {
		return handler(ev, prop)
	}
	// ignore unknown properties per RFC 5545 §3.8.8 (x-prop / iana-prop)
	return nil
}

func (ev *VEvent) AddChild(child parse.Component) error {
	switch c := child.(type) {
	case *VAlarm:
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
