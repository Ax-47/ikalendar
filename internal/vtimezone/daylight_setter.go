package vtimezone

import (
	"fmt"

	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (d *Daylight) SetDTSTART(it share.ITIME) error {
	// REQUIRED: MUST NOT occur more than once
	if !d.DTSTART.IsZero() {
		return fmt.Errorf("%w: DTSTART", parsehelper.ErrDuplicateProperty)
	}
	d.DTSTART = it
	return nil
}

func (d *Daylight) SetTZOffsetFrom(offset string) error { // TODO: validate
	return parsehelper.SetOnceValue(&d.TZOFFSETFROM, offset, "TZOFFSETFROM")
}

func (d *Daylight) SetTZOffsetTo(offset string) error { // TODO: validate
	return parsehelper.SetOnceValue(&d.TZOFFSETTO, offset, "TZOFFSETTO")
}

func (d *Daylight) SetRRule(r share.RECUR) error {
	return parsehelper.SetOnce(&d.RRULE, parsehelper.Ptr(r), "RRULE")
}

func (d *Daylight) AddRDate(it share.ITIME) error {
	d.RDATE = append(d.RDATE, it)
	return nil
}

func (d *Daylight) SetTZName(name string) error {
	return parsehelper.SetOnce(&d.TZNAME, parsehelper.Ptr(name), "TZNAME")
}
