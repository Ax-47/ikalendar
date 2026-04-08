package componants

type ComponentName string

const (
	ComponentVCalendar ComponentName = "VCALENDAR"
	ComponentVEvent    ComponentName = "VEVENT"
	ComponentVAlarm    ComponentName = "VALARM"
	ComponentVJournal  ComponentName = "VJOURNAL"
	ComponentVTodo     ComponentName = "VTODO"
)

type (
	Validalbe interface{ Validate() error }
)

type Component interface {
	// AddProperty applies a parsed line (like SUMMARY or DTSTART) to struct

	// AddProperty(prop Property) error
	ProcessProperty(prop Property) error
	// AddChild adds a nested component (like a VEVENT inside a VCALENDAR)
	AddChild(child Component) error
	Validalbe
}

type Property struct {
	Name   string
	Params map[string]string
	Value  string
}
type (
	ComponentFactory func() Component
)
