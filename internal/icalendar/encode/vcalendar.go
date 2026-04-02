package encode

import "github.com/minoplhy/ikalendar/internal/icalendar/models"

func VCalendar(ctx *EncodeContext, cal *models.VCalendar) {
	b := ctx.Builder

	writeProperty(b, "BEGIN", "VCALENDAR")
	writeProperty(b, "VERSION", cal.VERSION)
	writeProperty(b, "PRODID", cal.PRODID)
	writeString(b, "CALSCALE", cal.CALSCALE)
	writeString(b, "METHOD", cal.METHOD)

	for i := range cal.VEVENT {
		VEvent(ctx, &cal.VEVENT[i])
	}

	writeProperty(b, "END", "VCALENDAR")
}
