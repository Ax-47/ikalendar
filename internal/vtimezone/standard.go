package vtimezone

import "github.com/minoplhy/ikalendar/internal/share"

type Standard struct {
	DTSTART      share.ITIME
	TZOFFSETFROM string
	TZOFFSETTO   string

	RRULE  *share.RECUR
	RDATE  []share.ITIME
	TZNAME *string
}
