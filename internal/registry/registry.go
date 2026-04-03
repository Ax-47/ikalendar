package registry

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/vcalendar"
	"github.com/minoplhy/ikalendar/internal/vevent"
)

var CalendarComponents = map[string]func() componants.Component{
	"VCALENDAR": func() componants.Component { return &vcalendar.VCalendar{} },
	"VEVENT":    func() componants.Component { return &vevent.VEvent{} },
	// "VTODO":     func() parse.IComponent { return &VTodo{} },
	// "VJOURNAL":  func() parse.IComponent { return &VJournal{} },
	// "VALARM": func() parse.IComponent { return &VAlarm{} },
	// "VTIMEZONE": func() parse.IComponent { return &VTimezone{} },
}
