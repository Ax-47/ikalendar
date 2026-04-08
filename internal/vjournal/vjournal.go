package vjournal

import "github.com/minoplhy/ikalendar/internal/share"

type VJournal struct {
	/* REQUIRED */
	UID     string
	DTSTAMP share.ITIME

	/* OPTIONAL */
	CLASS        *string
	CREATED      *share.ITIME
	DTSTART      *share.ITIME
	LASTMODIFIED *share.ITIME
	ORGANIZER    *string
	SEQUENCE     *int
	STATUS       *string
	SUMMARY      *string
	URL          *string
	DESCRIPTION  *string

	/* recurrence */
	RRULE  *share.RECUR
	RDATE  []share.ITIME
	EXDATE []share.ITIME

	/* multi */
	ATTACH        []share.ATTACH
	ATTENDEE      []share.CALADDRESS
	CATEGORIES    []string
	COMMENT       []string
	CONTACT       []string
	RELATED       []share.RELATED
	REQUESTSTATUS []share.RequestStatus
}
