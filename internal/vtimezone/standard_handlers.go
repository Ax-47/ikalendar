package vtimezone

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
)

type StandardHandler func(*Standard, componants.Property) error

var standardHandlers = map[share.PropertyName]StandardHandler{
	PropDTSTART:      handleStandardDTSTART,
	PropTZOFFSETFROM: handleStandardTZOffsetFrom,
	PropTZOFFSETTO:   handleStandardTZOffsetTo,
	PropTZNAME:       handleStandardTZName,
	PropRRULE:        handleStandardRRule,
	PropRDATE:        handleStandardRDate,
}

// ── Required ──────────────────────────────────────────────────────────────────

func handleStandardDTSTART(s *Standard, prop componants.Property) error {
	return s.SetDTSTART(share.ParseITIME(prop.Params, prop.Value))
}

func handleStandardTZOffsetFrom(s *Standard, prop componants.Property) error {
	return s.SetTZOffsetFrom(prop.Value)
}

func handleStandardTZOffsetTo(s *Standard, prop componants.Property) error {
	return s.SetTZOffsetTo(prop.Value)
}

func handleStandardTZName(s *Standard, prop componants.Property) error {
	return s.SetTZName(prop.Value)
}

func handleStandardRRule(s *Standard, prop componants.Property) error {
	r, err := share.ParseRECUR(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid RRULE: %w", err)
	}
	return s.SetRRule(r)
}

func handleStandardRDate(s *Standard, prop componants.Property) error {
	return s.AddRDate(share.ParseITIME(prop.Params, prop.Value))
}
