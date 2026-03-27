package ikalendar

import (
	"errors"
	"io"

	"github.com/minoplhy/ikalendar/internal/icalendar"
	"github.com/minoplhy/ikalendar/internal/parse"
)

func ParseCalendar(r io.Reader) (*icalendar.VCalendar, error) {
	engine := parse.NewEngine()

	// Automatically register every component supports
	for name, factory := range icalendar.CalendarComponents {
		engine.Register(name, factory)
	}

	parser := parse.NewParser(r)
	rootComp, err := engine.Run(parser)
	if err != nil {
		return nil, err
	}

	calendar, ok := rootComp.(*icalendar.VCalendar)
	if !ok {
		return nil, errors.New("icalendar: parsed root component was not a VCALENDAR")
	}

	return calendar, nil
}
