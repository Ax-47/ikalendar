package ikalendar

import (
	"errors"
	"io"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
	"github.com/minoplhy/ikalendar/internal/parse"
	"github.com/minoplhy/ikalendar/internal/registry"
	"github.com/minoplhy/ikalendar/internal/vcalendar"
)

var Marshal = encode.Marshal

func ParseCalendar(r io.Reader) (*vcalendar.VCalendar, error) {
	engine := parse.NewEngine()

	// Automatically register every component supports
	for name, factory := range registry.CalendarComponents {
		engine.Register(componants.ComponentName(name), factory)
	}

	parser := parse.NewParser(r)
	rootComp, err := engine.Run(parser)
	if err != nil {
		return nil, err
	}

	calendar, ok := rootComp.(*vcalendar.VCalendar)
	if !ok {
		return nil, errors.New("icalendar: parsed root component was not a VCALENDAR")
	}

	return calendar, nil
}
