package registry

import (
	"github.com/minoplhy/ikalendar/internal/icalendar/models"
	"github.com/minoplhy/ikalendar/internal/parse"
)

var CalendarComponents = map[string]func() parse.Component{
	"VCALENDAR": func() parse.Component { return &models.VCalendar{} },
	"VEVENT":    func() parse.Component { return &models.VEvent{} },
	// "VTODO":     func() parse.IComponent { return &VTodo{} },
	// "VJOURNAL":  func() parse.IComponent { return &VJournal{} },
	// "VALARM": func() parse.IComponent { return &VAlarm{} },
	// "VTIMEZONE": func() parse.IComponent { return &VTimezone{} },
}
