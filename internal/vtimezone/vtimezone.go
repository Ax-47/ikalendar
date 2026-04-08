package vtimezone

import "github.com/minoplhy/ikalendar/internal/share"

type VTimezone struct {
	/* REQUIRED */
	TZID string

	/* OPTIONAL */
	LASTMODIFIED *share.ITIME
	TZURL        *string

	/* CHILD COMPONENTS */
	STANDARD []Standard
	DAYLIGHT []Daylight
}
