package vevent

import (
	"strings"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
)

func (ev *VEvent) Encode(ctx *componants.EncodeContext) {
	b := ctx.Builder
	encode.WriteProperty(b, "BEGIN", "VEVENT")

	// Required
	encode.WriteProperty(b, "UID", ev.UID)
	encode.WriteTimeWithParams(b, "DTSTAMP", &ev.DTSTAMP)

	// Required conditionally — fallback to DTSTAMP if METHOD is set
	if ev.DTSTART != nil {
		encode.WriteTimeWithParams(b, "DTSTART", ev.DTSTART)
	} else if ctx.Calendar.GetMethod() == nil {
		encode.WriteTimeWithParams(b, "DTSTART", &ev.DTSTAMP)
	}

	// Optional single
	encode.WriteTimeWithParams(b, "DTEND", ev.DTEND)
	if ev.DURATION != nil {
		encode.WriteProperty(b, "DURATION", ev.DURATION.FormatDURATION())
	}
	encode.WriteTimeWithParams(b, "CREATED", ev.CREATED)
	encode.WriteTimeWithParams(b, "LAST-MODIFIED", ev.LASTMODIFIED)
	encode.WriteString(b, "CLASS", ev.CLASS)
	encode.WriteString(b, "DESCRIPTION", ev.DESCRIPTION)
	encode.WriteString(b, "GEO", ev.GEO)
	encode.WriteString(b, "LOCATION", ev.LOCATION)
	encode.WriteString(b, "ORGANIZER", ev.ORGANIZER)
	encode.WriteString(b, "STATUS", ev.STATUS)
	encode.WriteString(b, "SUMMARY", ev.SUMMARY)
	encode.WriteString(b, "TRANSP", ev.TRANSP)
	encode.WriteString(b, "URL", ev.URL)
	encode.WriteInt(b, "PRIORITY", ev.PRIORITY)
	encode.WriteInt(b, "SEQUENCE", ev.SEQUENCE)

	// Recurrence
	if ev.RRULE != nil {
		encode.WriteProperty(b, "RRULE", ev.RRULE.FormatRECUR())
	}
	for i := range ev.EXDATE {
		encode.WriteTimeWithParams(b, "EXDATE", &ev.EXDATE[i])
	}
	for i := range ev.RDATE {
		encode.WriteTimeWithParams(b, "RDATE", &ev.RDATE[i])
	}

	//  Multi-value
	if len(ev.CATEGORIES) > 0 {
		encode.WriteProperty(b, "CATEGORIES", strings.Join(ev.CATEGORIES, ","))
	}
	for _, a := range ev.ATTACH {
		if a.URI != nil {
			encode.WriteProperty(b, "ATTACH", *a.URI)
		}
	}
	for _, a := range ev.ATTENDEE {
		encode.WriteCalAddress(b, "ATTENDEE", a)
	}
	for _, c := range ev.COMMENT {
		encode.WriteProperty(b, "COMMENT", c)
	}
	for _, c := range ev.CONTACT {
		encode.WriteProperty(b, "CONTACT", c)
	}
	for _, r := range ev.RELATED {
		encode.WriteProperty(b, "RELATED-TO", r.UID)
	}
	for _, r := range ev.RESOURCES {
		encode.WriteProperty(b, "RESOURCES", r)
	}
	for _, rs := range ev.REQUESTSTATUS {
		encode.WriteRequestStatus(b, rs)
	}

	// fmt.Println("test")
	// Sub-components
	for i := range ev.VALARM {
		ev.VALARM[i].Encode(ctx)
	}

	encode.WriteProperty(b, "END", "VEVENT")
}
