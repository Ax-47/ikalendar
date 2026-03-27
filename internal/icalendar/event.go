// File: internal/icalendar/event.go
package icalendar

import (
	"time"

	"github.com/minoplhy/ikalendar/internal/parse"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

// NewEvent initializes an empty event and automatically sets the creation time (DTSTAMP).
func NewEvent() *VEvent {
	return &VEvent{
		DTSTAMP: ITIME{
			Time:       time.Now().UTC(),
			IsDateOnly: false,
		},
	}
}

func (ev *VEvent) SetUID(uid string) *VEvent {
	ev.UID = uid
	return ev
}

func (ev *VEvent) SetSummary(summary string) *VEvent {
	ev.ProcessProperty(parse.Property{
		Name:  "SUMMARY",
		Value: summary,
	})
	//ev.SUMMARY = parsehelper.StrPtr(summary)
	return ev
}

func (ev *VEvent) SetDescription(desc string) *VEvent {
	ev.DESCRIPTION = parsehelper.StrPtr(desc)
	return ev
}

func (ev *VEvent) SetLocation(location string) *VEvent {
	ev.LOCATION = parsehelper.StrPtr(location)
	return ev
}

func (ev *VEvent) SetStatus(status string) *VEvent {
	ev.STATUS = parsehelper.StrPtr(status) // e.g., "TENTATIVE", "CONFIRMED", "CANCELLED"
	return ev
}

func (ev *VEvent) AddCategory(category string) *VEvent {
	ev.CATEGORIES = append(ev.CATEGORIES, category)
	return ev
}

// SetDtStart converts a native Go time.Time into your internal ITIME format.
func (ev *VEvent) SetDtStart(t time.Time) *VEvent {
	ev.DTSTART = &ITIME{
		Time:       t,
		IsDateOnly: false,
	}
	return ev
}

// SetDtEnd handles the end time.
func (ev *VEvent) SetDtEnd(t time.Time) *VEvent {
	ev.DTEND = &ITIME{
		Time:       t,
		IsDateOnly: false,
	}
	return ev
}
