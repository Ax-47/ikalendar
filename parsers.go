package ikalendar

import (
	"errors"
	"io"

	"github.com/minoplhy/ikalendar/internal/icalendar/encode"
	"github.com/minoplhy/ikalendar/internal/icalendar/models"
	"github.com/minoplhy/ikalendar/internal/icalendar/registry"
	"github.com/minoplhy/ikalendar/internal/parse"
)

var Marshal = encode.Marshal

func ParseCalendar(r io.Reader) (*models.VCalendar, error) {
	engine := parse.NewEngine()

	// Automatically register every component supports
	for name, factory := range registry.CalendarComponents {
		engine.Register(name, factory)
	}

	parser := parse.NewParser(r)
	rootComp, err := engine.Run(parser)
	if err != nil {
		return nil, err
	}

	calendar, ok := rootComp.(*models.VCalendar)
	if !ok {
		return nil, errors.New("icalendar: parsed root component was not a VCALENDAR")
	}

	return calendar, nil
}
