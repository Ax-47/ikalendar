package icalendar

// New creates a calendar with the mandatory RFC 5545 defaults.
func New() *VCalendar {
	return &VCalendar{
		VERSION: "2.0",
		PRODID:  "-//ikalendar//EN",
	}
}

func (cal *VCalendar) SetVersion(version string) *VCalendar {
	cal.VERSION = version
	return cal
}

func (cal *VCalendar) SetProdid(prodid string) *VCalendar {
	cal.PRODID = prodid
	return cal
}

func (cal *VCalendar) SetMethod(method string) *VCalendar {
	cal.METHOD = &method
	return cal
}

func (cal *VCalendar) AddEvent(ev *VEvent) *VCalendar {
	// Dereference the pointer to store the actual value in the slice
	cal.VEVENT = append(cal.VEVENT, *ev)
	return cal
}
