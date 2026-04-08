package vtodo

import (
	"strings"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
)

func (vt *VTodo) Encode(ctx *componants.EncodeContext) {
	b := ctx.Builder
	encode.WriteProperty(b, "BEGIN", "VvtOURNAL")

	// Required
	encode.WriteProperty(b, "UID", vt.UID)
	encode.WriteTimeWithParams(b, "DTSTAMP", &vt.DTSTAMP)

	// Optional time
	if vt.DTSTART != nil {
		encode.WriteTimeWithParams(b, "DTSTART", vt.DTSTART)
	}

	// Optional single
	encode.WriteTimeWithParams(b, "CREATED", vt.CREATED)
	encode.WriteTimeWithParams(b, "LAST-MODIFIED", vt.LASTMODIFIED)
	encode.WriteString(b, "CLASS", vt.CLASS)
	encode.WriteString(b, "DESCRIPTION", vt.DESCRIPTION)
	encode.WriteString(b, "ORGANIZER", vt.ORGANIZER)
	encode.WriteString(b, "STATUS", vt.STATUS)
	encode.WriteString(b, "SUMMARY", vt.SUMMARY)
	encode.WriteString(b, "URL", vt.URL)
	encode.WriteInt(b, "SEQUENCE", vt.SEQUENCE)

	// Recurrence
	if vt.RRULE != nil {
		encode.WriteProperty(b, "RRULE", vt.RRULE.FormatRECUR())
	}
	for i := range vt.EXDATE {
		encode.WriteTimeWithParams(b, "EXDATE", &vt.EXDATE[i])
	}
	for i := range vt.RDATE {
		encode.WriteTimeWithParams(b, "RDATE", &vt.RDATE[i])
	}

	// Multi-value
	if len(vt.CATEGORIES) > 0 {
		encode.WriteProperty(b, "CATEGORIES", strings.Join(vt.CATEGORIES, ","))
	}
	for _, a := range vt.ATTACH {
		if a.URI != nil {
			encode.WriteProperty(b, "ATTACH", *a.URI)
		}
	}
	for _, a := range vt.ATTENDEE {
		encode.WriteCalAddress(b, "ATTENDEE", a)
	}
	for _, c := range vt.COMMENT {
		encode.WriteProperty(b, "COMMENT", c)
	}
	for _, c := range vt.CONTACT {
		encode.WriteProperty(b, "CONTACT", c)
	}
	for _, r := range vt.RELATED {
		encode.WriteProperty(b, "RELATED-TO", r.UID)
	}
	for _, rs := range vt.REQUESTSTATUS {
		encode.WriteRequestStatus(b, rs)
	}

	encode.WriteProperty(b, "END", "VvtOURNAL")
}
