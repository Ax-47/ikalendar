package encode

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/icalendar/models"
)

func FormatRECUR(r *models.RECUR) string {
	var parts []string

	parts = append(parts, "FREQ="+r.Freq)

	if r.Until != nil {
		parts = append(parts, "UNTIL="+models.FormatITIME(*r.Until))
	}
	if r.Count != nil {
		parts = append(parts, fmt.Sprintf("COUNT=%d", *r.Count))
	}
	if r.Interval > 0 {
		parts = append(parts, fmt.Sprintf("INTERVAL=%d", r.Interval))
	}
	if len(r.ByDay) > 0 {
		parts = append(parts, "BYDAY="+strings.Join(r.ByDay, ","))
	}
	if len(r.ByMonthDay) > 0 {
		parts = append(parts, "BYMONTHDAY="+joinInts(r.ByMonthDay))
	}
	if len(r.ByMonth) > 0 {
		parts = append(parts, "BYMONTH="+joinInts(r.ByMonth))
	}

	return strings.Join(parts, ";")
}

func joinInts(vals []int) string {
	s := make([]string, len(vals))
	for i, v := range vals {
		s[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(s, ",")
}
