package vtimezone

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (s *Standard) SetDTSTART(it share.ITIME) error {
	// REQUIRED: MUST NOT occur more than once
	if !s.DTSTART.IsZero() {
		return fmt.Errorf("%w: DTSTART", utils.ErrDuplicateProperty)
	}
	s.DTSTART = it
	return nil
}

func (s *Standard) SetTZOffsetFrom(offset string) error { // TODO: validate
	return utils.SetOnceValue(&s.TZOFFSETFROM, offset, "TZOFFSETFROM")
}

func (s *Standard) SetTZOffsetTo(offset string) error { // TODO: validate
	return utils.SetOnceValue(&s.TZOFFSETTO, offset, "TZOFFSETTO")
}

func (s *Standard) SetRRule(r share.RECUR) error {
	return utils.SetOnce(&s.RRULE, utils.Ptr(r), "RRULE")
}

func (s *Standard) AddRDate(it share.ITIME) error {
	s.RDATE = append(s.RDATE, it)
	return nil
}

func (s *Standard) SetTZName(name string) error {
	return utils.SetOnce(&s.TZNAME, utils.Ptr(name), "TZNAME")
}
