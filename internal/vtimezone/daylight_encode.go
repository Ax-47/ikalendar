package vtimezone

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
)

func (d *Daylight) Encode(ctx *componants.EncodeContext) {
	b := ctx.Builder

	encode.WriteProperty(b, "BEGIN", "DAYLIGHT")

	encode.WriteTimeWithParams(b, "DTSTART", &d.DTSTART)
	encode.WriteProperty(b, "TZOFFSETFROM", d.TZOFFSETFROM)
	encode.WriteProperty(b, "TZOFFSETTO", d.TZOFFSETTO)

	if d.TZNAME != nil {
		encode.WriteString(b, "TZNAME", d.TZNAME)
	}

	if d.RRULE != nil {
		encode.WriteProperty(b, "RRULE", d.RRULE.FormatRECUR())
	}

	for i := range d.RDATE {
		encode.WriteTimeWithParams(b, "RDATE", &d.RDATE[i])
	}

	encode.WriteProperty(b, "END", "DAYLIGHT")
}
