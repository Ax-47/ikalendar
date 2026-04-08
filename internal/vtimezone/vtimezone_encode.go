package vtimezone

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
)

func (tz *VTimezone) Encode(ctx *componants.EncodeContext) {
	b := ctx.Builder

	encode.WriteProperty(b, "BEGIN", "VTIMEZONE")
	encode.WriteProperty(b, "TZID", tz.TZID)

	encode.WriteTimeWithParams(b, "LAST-MODIFIED", tz.LASTMODIFIED)
	encode.WriteString(b, "TZURL", tz.TZURL)

	// STANDARD
	for i := range tz.STANDARD {
		tz.STANDARD[i].Encode(ctx)
	}

	// DAYLIGHT
	for i := range tz.DAYLIGHT {
		tz.DAYLIGHT[i].Encode(ctx)
	}

	encode.WriteProperty(b, "END", "VTIMEZONE")
}
