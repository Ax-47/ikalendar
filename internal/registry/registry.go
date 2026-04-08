package registry

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/valarm"
	"github.com/minoplhy/ikalendar/internal/vcalendar"
	"github.com/minoplhy/ikalendar/internal/vevent"
	"github.com/minoplhy/ikalendar/internal/vjournal"
	"github.com/minoplhy/ikalendar/internal/vtimezone"
	"github.com/minoplhy/ikalendar/internal/vtodo"
)

var CalendarComponents = componants.RegistryMap{
	componants.ComponentVCalendar: func() componants.Component { return &vcalendar.VCalendar{} },
	componants.ComponentVEvent:    func() componants.Component { return &vevent.VEvent{} },
	componants.ComponentVAlarm:    func() componants.Component { return &valarm.VAlarm{} },
	componants.ComponentVTodo:     func() componants.Component { return &vtodo.VTodo{} },
	componants.ComponentVJournal:  func() componants.Component { return &vjournal.VJournal{} },
	componants.ComponentVTimezone: func() componants.Component { return &vtimezone.VTimezone{} },
}
