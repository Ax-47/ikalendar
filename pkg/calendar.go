package ikalendar

import (
	"github.com/minoplhy/ikalendar/internal/vcalendar"
)

// CalendarOption configures a Calendar
type CalendarOption = vcalendar.VcalendarOption

// New creates a new Calendar with RFC 5545 defaults
//
//	cal, err := ikalendar.New(
//	    ikalendar.WithProdID("-//MyApp//EN"),
//	    ikalendar.WithEvent(ev),
//	)

// Re-export calendar options
var (
	NewCalendar  = vcalendar.NewCalendar
	WithProdID   = vcalendar.WithProdID
	WithVersion  = vcalendar.WithVersion
	WithMethod   = vcalendar.WithMethod
	WithCalScale = vcalendar.WithCalScale
	WithEvent    = vcalendar.WithEvent
)
