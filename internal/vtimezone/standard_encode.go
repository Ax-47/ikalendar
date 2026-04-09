package vtimezone

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
)

func (s *Standard) Encode(ctx *componants.EncodeContext) {
	b := ctx.Builder

	encode.WriteProperty(b, "BEGIN", "STANDARD")

	encode.WriteTimeWithParams(b, "DTSTART", &s.DTSTART)
	encode.WriteProperty(b, "TZOFFSETFROM", s.TZOFFSETFROM)
	encode.WriteProperty(b, "TZOFFSETTO", s.TZOFFSETTO)

	if s.TZNAME != nil {
		encode.WriteString(b, "TZNAME", s.TZNAME)
	}

	if s.RRULE != nil {
		encode.WriteProperty(b, "RRULE", s.RRULE.FormatRECUR())
	}

	for i := range s.RDATE {
		encode.WriteTimeWithParams(b, "RDATE", &s.RDATE[i])
	}

	encode.WriteProperty(b, "END", "STANDARD")
}
