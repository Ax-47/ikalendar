package vcalendar

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/utils"
)

func (cal *VCalendar) SetPRODID(s string) error {
	if cal.PRODID != "" {
		return fmt.Errorf("%w: PRODID", utils.ErrDuplicateProperty)
	}
	cal.PRODID = s
	return nil
}

func (cal *VCalendar) SetVERSION(s string) error {
	if cal.VERSION != "" {
		return fmt.Errorf("%w: VERSION", utils.ErrDuplicateProperty)
	}

	cal.VERSION = s
	return nil
}

func (cal *VCalendar) SetCALSCALE(s string) error {
	return utils.SetOnce(&cal.CALSCALE, utils.Ptr(s), string(PropCALSCALE))
}

func (cal *VCalendar) SetMETHOD(s string) error {
	return utils.SetOnce(&cal.METHOD, utils.Ptr(s), string(PropMETHOD))
}
