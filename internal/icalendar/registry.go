package icalendar

import "github.com/minoplhy/ikalendar/internal/parse"

var CalendarComponents = map[string]func() parse.IComponent{
	"VCALENDAR": func() parse.IComponent { return &VCalendar{} },
	"VEVENT":    func() parse.IComponent { return &VEvent{} },
	// "VTODO":     func() parse.IComponent { return &VTodo{} },
	// "VJOURNAL":  func() parse.IComponent { return &VJournal{} },
	// "VALARM":    func() parse.IComponent { return &VAlarm{} },
	// "VTIMEZONE": func() parse.IComponent { return &VTimezone{} },
}
