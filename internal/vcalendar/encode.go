package vcalendar

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
)

func (cal *VCalendar) Encode(ctx *componants.EncodeContext) {
	b := ctx.Builder

	encode.WriteProperty(b, "BEGIN", "VCALENDAR")
	encode.WriteProperty(b, "VERSION", cal.VERSION)
	encode.WriteProperty(b, "PRODID", cal.PRODID)
	encode.WriteString(b, "CALSCALE", cal.CALSCALE)
	encode.WriteString(b, "METHOD", cal.METHOD)

	for i := range cal.VEVENT {
		cal.VEVENT[i].Encode(ctx)
	}

	encode.WriteProperty(b, "END", "VCALENDAR")
}
