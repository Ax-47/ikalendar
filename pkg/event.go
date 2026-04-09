package ikalendar

import (
	"github.com/minoplhy/ikalendar/internal/vevent"
)

// EventOption configures a VEvent
type EventOption = vevent.VeventOption

// Re-export event options
var (
	NewEvent          = vevent.NewEvent
	WithSummary       = vevent.WithSummary
	WithDescription   = vevent.WithDescription
	WithLocation      = vevent.WithLocation
	WithStatus        = vevent.WithStatus
	WithDtStart       = vevent.WithDtStart
	WithDtEnd         = vevent.WithDtEnd
	WithDuration      = vevent.WithDuration
	WithPriority      = vevent.WithPriority
	WithSequence      = vevent.WithSequence
	WithCreated       = vevent.WithCreated
	WithLastModified  = vevent.WithLastModified
	WithClass         = vevent.WithClass
	WithOrganizer     = vevent.WithOrganizer
	WithTransp        = vevent.WithTransp
	WithURL           = vevent.WithURL
	WithRRule         = vevent.WithRRule
	WithCategory      = vevent.WithCategory
	WithAttendee      = vevent.WithAttendee
	WithAttach        = vevent.WithAttach
	WithComment       = vevent.WithComment
	WithContact       = vevent.WithContact
	WithExDate        = vevent.WithExDate
	WithRDate         = vevent.WithRDate
	WithRelated       = vevent.WithRelated
	WithResources     = vevent.WithResources
	WithRequestStatus = vevent.WithRequestStatus
	WithAlarm         = vevent.WithAlarm
)
