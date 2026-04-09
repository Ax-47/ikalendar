package vtodo

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
)

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

func (vt *VTodo) ProcessProperty(prop componants.Property) error {
	name := share.PropertyName(prop.Name)
	if handler, ok := vtodoHandlers[name]; ok {
		return handler(vt, prop)
	}
	// ignore unknown properties per RFC 5545 §3.8.8 (x-prop / iana-prop)
	return nil
}

func (vt *VTodo) Validate() error {
	if vt.UID == "" {
		return fmt.Errorf("VJOURNAL: UID is required")
	}
	if vt.DTSTAMP.IsZero() {
		return fmt.Errorf("VJOURNAL: DTSTAMP is required")
	}

	// Recurrence rule constraint
	if vt.RRULE != nil && vt.DTSTART == nil {
		return fmt.Errorf("VJOURNAL: DTSTART is required when RRULE is present")
	}

	return nil
}

func (vt *VTodo) AddChild(child componants.Component) error {
	return fmt.Errorf("JOURNAL: does not support child components")
}
