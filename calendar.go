package ikalendar

import (
	"github.com/minoplhy/ikalendar/internal/icalendar/builders"
)

// CalendarOption configures a Calendar
type CalendarOption = builders.VcalendarOption

// New creates a new Calendar with RFC 5545 defaults
//
//	cal, err := ikalendar.New(
//	    ikalendar.WithProdID("-//MyApp//EN"),
//	    ikalendar.WithEvent(ev),
//	)

// Re-export calendar options
var (
	NewCalendar  = builders.NewCalendar
	WithProdID   = builders.WithProdID
	WithVersion  = builders.WithVersion
	WithMethod   = builders.WithMethod
	WithCalScale = builders.WithCalScale
	WithEvent    = builders.WithEvent
)
