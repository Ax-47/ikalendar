package encode

import (
	"strings"

	"github.com/minoplhy/ikalendar/internal/icalendar/models"
)

func VEvent(ctx *EncodeContext, ev *models.VEvent) {
	b := ctx.Builder

	writeProperty(b, "BEGIN", "VEVENT")

	// Required
	writeProperty(b, "UID", ev.UID)
	writeTimeWithParams(b, "DTSTAMP", &ev.DTSTAMP)

	// Required conditionally — fallback to DTSTAMP if METHOD is set
	if ev.DTSTART != nil {
		writeTimeWithParams(b, "DTSTART", ev.DTSTART)
	} else if ctx.Calendar.METHOD == nil {
		writeTimeWithParams(b, "DTSTART", &ev.DTSTAMP)
	}

	// Optional single
	writeTimeWithParams(b, "DTEND", ev.DTEND)
	if ev.DURATION != nil {
		writeProperty(b, "DURATION", models.FormatDURATION(*ev.DURATION))
	}
	writeTimeWithParams(b, "CREATED", ev.CREATED)
	writeTimeWithParams(b, "LAST-MODIFIED", ev.LASTMODIFIED)
	writeString(b, "CLASS", ev.CLASS)
	writeString(b, "DESCRIPTION", ev.DESCRIPTION)
	writeString(b, "GEO", ev.GEO)
	writeString(b, "LOCATION", ev.LOCATION)
	writeString(b, "ORGANIZER", ev.ORGANIZER)
	writeString(b, "STATUS", ev.STATUS)
	writeString(b, "SUMMARY", ev.SUMMARY)
	writeString(b, "TRANSP", ev.TRANSP)
	writeString(b, "URL", ev.URL)
	writeInt(b, "PRIORITY", ev.PRIORITY)
	writeInt(b, "SEQUENCE", ev.SEQUENCE)

	// Recurrence
	if ev.RRULE != nil {
		writeProperty(b, "RRULE", FormatRECUR(ev.RRULE))
	}
	for i := range ev.EXDATE {
		writeTimeWithParams(b, "EXDATE", &ev.EXDATE[i])
	}
	for i := range ev.RDATE {
		writeTimeWithParams(b, "RDATE", &ev.RDATE[i])
	}

	// Multi-value
	if len(ev.CATEGORIES) > 0 {
		writeProperty(b, "CATEGORIES", strings.Join(ev.CATEGORIES, ","))
	}
	for _, a := range ev.ATTACH {
		if a.URI != nil {
			writeProperty(b, "ATTACH", *a.URI)
		}
	}
	for _, a := range ev.ATTENDEE {
		writeCalAddress(b, "ATTENDEE", a)
	}
	for _, c := range ev.COMMENT {
		writeProperty(b, "COMMENT", c)
	}
	for _, c := range ev.CONTACT {
		writeProperty(b, "CONTACT", c)
	}
	for _, r := range ev.RELATED {
		writeProperty(b, "RELATED-TO", r.UID)
	}
	for _, r := range ev.RESOURCES {
		writeProperty(b, "RESOURCES", r)
	}
	for _, rs := range ev.REQUESTSTATUS {
		writeRequestStatus(b, rs)
	}

	// Sub-components
	for i := range ev.VALARM {
		VAlarm(ctx, &ev.VALARM[i])
	}

	writeProperty(b, "END", "VEVENT")
}
