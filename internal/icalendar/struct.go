package icalendar

import (
	"time"
)

type VCalendar struct {
	VERSION string
	PRODID  string

	CALSCALE *string
	METHOD   *string

	VEVENT []VEvent
}

/* eventprop */
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
	REQUESTSTATUS []string
	RELATED       []RELATED
	RESOURCES     []string
	RDATE         []string // RRULE, RDATE property name unimplemented
	/* x-prop */
	/* iana-prop */
}

type RECUR struct{}

type DURATION struct{}

type ATTACH struct{}

type CALADDRESS struct{}

type ITIME struct {
	// holds things like TZID or VALUE.
	// {"TZID": "America/New_York"}
	Parameters map[string]string

	Time time.Time

	// IsDateOnly is a helper flag. If true, the property had VALUE=DATE
	// making it all-day event
	IsDateOnly bool
}
type RELATED struct{}
