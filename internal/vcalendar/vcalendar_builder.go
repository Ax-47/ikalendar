package vcalendar

import (
	"github.com/minoplhy/ikalendar/internal/vevent"
)

type VcalendarOption func(*VCalendar) error

// ── Required ──────────────────────────────────────────────────────────────────

func WithVersion(version string) VcalendarOption {
	return func(cal *VCalendar) error {
		return cal.SetVERSION(version)
	}
}

func WithProdID(prodid string) VcalendarOption {
	return func(cal *VCalendar) error {
		return cal.SetPRODID(prodid)
	}
}

func WithMethod(method string) VcalendarOption {
	return func(cal *VCalendar) error {
		return cal.SetMETHOD(method)
	}
}

func WithCalScale(scale string) VcalendarOption {
	return func(cal *VCalendar) error {
		return cal.SetCALSCALE(scale)
	}
}

// ── Optional ──────────────────────────────────────────────────────────────────

func WithEvent(ev *vevent.VEvent) VcalendarOption {
	return func(cal *VCalendar) error {
		cal.AddChild(ev)
		return nil
	}
}

// ── Constructor ───────────────────────────────────────────────────────────────

// New creates a new Calendar with RFC 5545 defaults
//
//	cal, err := ikalendar.NewCalendar(
//	    ikalendar.WithProdID("-//MyApp//EN"),
//	    ikalendar.WithEvent(ev),
//	)
func NewCalendar(opts ...VcalendarOption) (*VCalendar, error) {
	cal := &VCalendar{
		VERSION: "2.0",
		PRODID:  "-//ikalendar//EN",
	}
	for _, opt := range opts {
		if err := opt(cal); err != nil {
			return nil, err
		}
	}
	return cal, nil
}
