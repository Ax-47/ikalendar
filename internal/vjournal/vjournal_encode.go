package vjournal

import (
	"strings"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
)

func (j *VJournal) Encode(ctx *componants.EncodeContext) {
	b := ctx.Builder
	encode.WriteProperty(b, "BEGIN", "VJOURNAL")

	// Required
	encode.WriteProperty(b, "UID", j.UID)
	encode.WriteTimeWithParams(b, "DTSTAMP", &j.DTSTAMP)

	// Optional time
	if j.DTSTART != nil {
		encode.WriteTimeWithParams(b, "DTSTART", j.DTSTART)
	}

	// Optional single
	encode.WriteTimeWithParams(b, "CREATED", j.CREATED)
	encode.WriteTimeWithParams(b, "LAST-MODIFIED", j.LASTMODIFIED)
	encode.WriteString(b, "CLASS", j.CLASS)
	encode.WriteString(b, "DESCRIPTION", j.DESCRIPTION)
	encode.WriteString(b, "ORGANIZER", j.ORGANIZER)
	encode.WriteString(b, "STATUS", j.STATUS)
	encode.WriteString(b, "SUMMARY", j.SUMMARY)
	encode.WriteString(b, "URL", j.URL)
	encode.WriteInt(b, "SEQUENCE", j.SEQUENCE)

	// Recurrence
	if j.RRULE != nil {
		encode.WriteProperty(b, "RRULE", j.RRULE.FormatRECUR())
	}
	for i := range j.EXDATE {
		encode.WriteTimeWithParams(b, "EXDATE", &j.EXDATE[i])
	}
	for i := range j.RDATE {
		encode.WriteTimeWithParams(b, "RDATE", &j.RDATE[i])
	}

	// Multi-value
	if len(j.CATEGORIES) > 0 {
		encode.WriteProperty(b, "CATEGORIES", strings.Join(j.CATEGORIES, ","))
	}
	for _, a := range j.ATTACH {
		if a.URI != nil {
			encode.WriteProperty(b, "ATTACH", *a.URI)
		}
	}
	for _, a := range j.ATTENDEE {
		encode.WriteCalAddress(b, "ATTENDEE", a)
	}
	for _, c := range j.COMMENT {
		encode.WriteProperty(b, "COMMENT", c)
	}
	for _, c := range j.CONTACT {
		encode.WriteProperty(b, "CONTACT", c)
	}
	for _, r := range j.RELATED {
		encode.WriteProperty(b, "RELATED-TO", r.UID)
	}
	for _, rs := range j.REQUESTSTATUS {
		encode.WriteRequestStatus(b, rs)
	}

	encode.WriteProperty(b, "END", "VJOURNAL")
}
