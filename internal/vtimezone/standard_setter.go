package vtimezone

import (
	"fmt"

	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (s *Standard) SetDTSTART(it share.ITIME) error {
	// REQUIRED: MUST NOT occur more than once
	if !s.DTSTART.IsZero() {
		return fmt.Errorf("%w: DTSTART", parsehelper.ErrDuplicateProperty)
	}
	s.DTSTART = it
	return nil
}

func (s *Standard) SetTZOffsetFrom(offset string) error { // TODO: validate
	return parsehelper.SetOnceValue(&s.TZOFFSETFROM, offset, "TZOFFSETFROM")
}

func (s *Standard) SetTZOffsetTo(offset string) error { // TODO: validate
	return parsehelper.SetOnceValue(&s.TZOFFSETTO, offset, "TZOFFSETTO")
}

func (s *Standard) SetRRule(r share.RECUR) error {
	return parsehelper.SetOnce(&s.RRULE, parsehelper.Ptr(r), "RRULE")
}

func (s *Standard) AddRDate(it share.ITIME) error {
	s.RDATE = append(s.RDATE, it)
	return nil
}

func (s *Standard) SetTZName(name string) error {
	return parsehelper.SetOnce(&s.TZNAME, parsehelper.Ptr(name), "TZNAME")
}
