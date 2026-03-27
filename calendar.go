package ikalendar

import (
	"strings"

	"github.com/minoplhy/ikalendar/internal/icalendar"
)

func New() *icalendar.VCalendar {
	return icalendar.New()
}

func NewEvent() *icalendar.VEvent {
	return icalendar.NewEvent()
}

func Marshal(cal *icalendar.VCalendar) ([]byte, error) {
	var sb strings.Builder

	ctx := &icalendar.EncodeContext{
		Builder:  &sb,
		Calendar: cal,
	}

	cal.Encode(ctx)

	return []byte(sb.String()), nil
}
