package vtodo

import "github.com/minoplhy/ikalendar/internal/share"

type VTodo struct {
	/* REQUIRED */
	UID     string
	DTSTAMP share.ITIME

	/* OPTIONAL (single) */
	CLASS           *string
	COMPLETED       *share.ITIME
	CREATED         *share.ITIME
	DESCRIPTION     *string
	DTSTART         *share.ITIME
	DUE             *share.ITIME
	LASTMODIFIED    *share.ITIME
	LOCATION        *string
	ORGANIZER       *string
	PERCENTCOMPLETE *int
	PRIORITY        *int
	SEQUENCE        *int
	STATUS          *string
	SUMMARY         *string
	URL             *string

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
	RESOURCES     []string
	REQUESTSTATUS []share.RequestStatus
}
