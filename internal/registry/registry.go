package registry

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/valarm"
	"github.com/minoplhy/ikalendar/internal/vcalendar"
	"github.com/minoplhy/ikalendar/internal/vevent"
)

var CalendarComponents = map[string]func() componants.Component{
	string(componants.ComponentVCalendar): func() componants.Component { return &vcalendar.VCalendar{} },
	string(componants.ComponentVEvent):    func() componants.Component { return &vevent.VEvent{} },
	string(componants.ComponentVAlarm):    func() componants.Component { return &valarm.VAlarm{} },
	// "VTODO":     func() componants.Component { return &VTodo{} },
	// "VJOURNAL":  func() componants.Component { return &VJournal{} },
	// "VTIMEZONE": func() componants.Component { return &VTimezone{} },
}
