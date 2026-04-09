package vcalendar

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
	"github.com/minoplhy/ikalendar/internal/vevent"
)

type VCalendar struct {
	VERSION string
	PRODID  string

	CALSCALE *string
	METHOD   *string

	VEVENT []vevent.VEvent
}

const (
	PropPRODID   share.PropertyName = "PRODID"
	PropVERSION  share.PropertyName = "VERSION"
	PropCALSCALE share.PropertyName = "CALSCALE"
	PropMETHOD   share.PropertyName = "METHOD"
)

func (cal *VCalendar) GetMethod() *string {
	return cal.METHOD
}

func (cal *VCalendar) ProcessProperty(prop componants.Property) error {
	switch share.PropertyName(prop.Name) {
	case PropPRODID:
		return cal.SetPRODID(prop.Value)
	case PropVERSION:
		return cal.SetVERSION(prop.Value)
	case PropCALSCALE:
		return cal.SetCALSCALE(prop.Value)
	case PropMETHOD:
		return cal.SetMETHOD(prop.Value)
	}
	return nil
}

func (cal *VCalendar) AddChild(child componants.Component) error {
	switch c := child.(type) {
	case *vevent.VEvent:
		cal.VEVENT = append(cal.VEVENT, *c)
		return nil
	default:
		return fmt.Errorf("%w: VCALENDAR cannot contain %T",
			utils.ErrInvalidComponent, child)
	}
}

func (cal *VCalendar) Validate() error {
	if cal.PRODID == "" {
		return fmt.Errorf("%w: VCALENDAR missing PRODID", utils.ErrMissingRequired)
	}
	if cal.VERSION == "" {
		return fmt.Errorf("%w: VCALENDAR missing VERSION", utils.ErrMissingRequired)
	}
	return nil
}
