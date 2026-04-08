package vjournal

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
)

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

func (j *VJournal) ProcessProperty(prop componants.Property) error {
	name := share.PropertyName(prop.Name)
	if handler, ok := vjournalHandlers[name]; ok {
		return handler(j, prop)
	}
	// ignore unknown properties per RFC 5545 §3.8.8 (x-prop / iana-prop)
	return nil
}

func (j *VJournal) Validate() error {
	if j.UID == "" {
		return fmt.Errorf("VJOURNAL: UID is required")
	}
	if j.DTSTAMP.IsZero() {
		return fmt.Errorf("VJOURNAL: DTSTAMP is required")
	}

	// Recurrence rule constraint
	if j.RRULE != nil && j.DTSTART == nil {
		return fmt.Errorf("VJOURNAL: DTSTART is required when RRULE is present")
	}

	return nil
}

func (j *VJournal) AddChild(child componants.Component) error {
	return fmt.Errorf("JOURNAL: does not support child components")
}
