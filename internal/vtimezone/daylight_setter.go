package vtimezone

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (d *Daylight) SetDTSTART(it share.ITIME) error {
	// REQUIRED: MUST NOT occur more than once
	if !d.DTSTART.IsZero() {
		return fmt.Errorf("%w: DTSTART", utils.ErrDuplicateProperty)
	}
	d.DTSTART = it
	return nil
}

func (d *Daylight) SetTZOffsetFrom(offset string) error { // TODO: validate
	return utils.SetOnceValue(&d.TZOFFSETFROM, offset, "TZOFFSETFROM")
}

func (d *Daylight) SetTZOffsetTo(offset string) error { // TODO: validate
	return utils.SetOnceValue(&d.TZOFFSETTO, offset, "TZOFFSETTO")
}

func (d *Daylight) SetRRule(r share.RECUR) error {
	return utils.SetOnce(&d.RRULE, utils.Ptr(r), "RRULE")
}

func (d *Daylight) AddRDate(it share.ITIME) error {
	d.RDATE = append(d.RDATE, it)
	return nil
}

func (d *Daylight) SetTZName(name string) error {
	return utils.SetOnce(&d.TZNAME, utils.Ptr(name), "TZNAME")
}
