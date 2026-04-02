package ikalendar

import (
	"github.com/minoplhy/ikalendar/internal/icalendar/builders"
)

// EventOption configures a VEvent
type EventOption = builders.VeventOption

// Re-export event options
var (
	NewEvent          = builders.NewEvent
	WithSummary       = builders.WithSummary
	WithDescription   = builders.WithDescription
	WithLocation      = builders.WithLocation
	WithStatus        = builders.WithStatus
	WithDtStart       = builders.WithDtStart
	WithDtEnd         = builders.WithDtEnd
	WithDuration      = builders.WithDuration
	WithPriority      = builders.WithPriority
	WithSequence      = builders.WithSequence
	WithCreated       = builders.WithCreated
	WithLastModified  = builders.WithLastModified
	WithClass         = builders.WithClass
	WithOrganizer     = builders.WithOrganizer
	WithTransp        = builders.WithTransp
	WithURL           = builders.WithURL
	WithRRule         = builders.WithRRule
	WithCategory      = builders.WithCategory
	WithAttendee      = builders.WithAttendee
	WithAttach        = builders.WithAttach
	WithComment       = builders.WithComment
	WithContact       = builders.WithContact
	WithExDate        = builders.WithExDate
	WithRDate         = builders.WithRDate
	WithRelated       = builders.WithRelated
	WithResources     = builders.WithResources
	WithRequestStatus = builders.WithRequestStatus
	WithAlarm         = builders.WithAlarm
)
