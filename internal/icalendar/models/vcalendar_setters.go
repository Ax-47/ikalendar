package models

import (
	"fmt"

	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

func (cal *VCalendar) SetPRODID(s string) error {
	if cal.PRODID != "" {
		return fmt.Errorf("%w: PRODID", parsehelper.ErrDuplicateProperty)
	}
	cal.PRODID = s
	return nil
}

func (cal *VCalendar) SetVERSION(s string) error {
	if cal.VERSION != "" {
		return fmt.Errorf("%w: VERSION", parsehelper.ErrDuplicateProperty)
	}

	cal.VERSION = s
	return nil
}

func (cal *VCalendar) SetCALSCALE(s string) error {
	return parsehelper.SetOnce(&cal.CALSCALE, parsehelper.Ptr(s), calPropCALSCALE)
}

func (cal *VCalendar) SetMETHOD(s string) error {
	return parsehelper.SetOnce(&cal.METHOD, parsehelper.Ptr(s), calPropMETHOD)
}
