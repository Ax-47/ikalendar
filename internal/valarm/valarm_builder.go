package valarm

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

type AlarmOption func(*VAlarm) error

// ── Required ──────────────────────────────────────────────────────────────────

func WithAction(action string) AlarmOption {
	return func(a *VAlarm) error {
		if action == "" {
			return fmt.Errorf("%w: ACTION cannot be empty", utils.ErrMissingRequired)
		}
		a.Action = new(action)
		return nil
	}
}

func WithTrigger(d share.DURATION) AlarmOption {
	return func(a *VAlarm) error {
		if a.Trigger != nil {
			return fmt.Errorf("%w: TRIGGER", utils.ErrDuplicateProperty)
		}
		a.Trigger = new(d)
		return nil
	}
}

// WithTriggerBefore creates a negative duration trigger — fires X hours/minutes before the event
// e.g. WithTriggerBefore(0, 15) → TRIGGER:-PT15M
// e.g. WithTriggerBefore(1, 30) → TRIGGER:-PT1H30M
func WithTriggerBefore(hours, minutes int) AlarmOption {
	return WithTrigger(share.DURATION{ // TODO: construct
		Negative: true,
		Hours:    hours,
		Minutes:  minutes,
	})
}

// WithTriggerAfter creates a positive duration trigger — fires X hours/minutes after the event
// e.g. WithTriggerAfter(0, 5) → TRIGGER:PT5M
func WithTriggerAfter(hours, minutes int) AlarmOption {
	return WithTrigger(share.DURATION{ // TODO: construct
		Negative: false,
		Hours:    hours,
		Minutes:  minutes,
	})
}

// WithTriggerDays — convenience for day-based triggers
// e.g. WithTriggerDays(1, true) → TRIGGER:-P1D (1 day before)
func WithTriggerDays(days int, before bool) AlarmOption {
	return WithTrigger(share.DURATION{ // TODO: construct
		Negative: before,
		Days:     days,
	})
}

// ── Optional ──────────────────────────────────────────────────────────────────

func WithAlarmDescription(desc string) AlarmOption {
	return func(a *VAlarm) error {
		if a.Description != nil {
			return fmt.Errorf("%w: DESCRIPTION", utils.ErrDuplicateProperty)
		}
		a.Description = new(desc)
		return nil
	}
}

func WithAlarmSummary(summary string) AlarmOption {
	return func(a *VAlarm) error {
		if a.Summary != nil {
			return fmt.Errorf("%w: SUMMARY", utils.ErrDuplicateProperty)
		}
		a.Summary = new(summary)
		return nil
	}
}

// ── Constructor ───────────────────────────────────────────────────────────────

func NewAlarm(opts ...AlarmOption) (*VAlarm, error) {
	a := &VAlarm{}
	for _, opt := range opts {
		if err := opt(a); err != nil {
			return nil, err
		}
	}
	if err := validateAlarm(a); err != nil {
		return nil, err
	}
	return a, nil
}

func validateAlarm(a *VAlarm) error {
	if a.Action == nil {
		return fmt.Errorf("%w: VALARM missing ACTION", utils.ErrMissingRequired)
	}
	if a.Trigger == nil {
		return fmt.Errorf("%w: VALARM missing TRIGGER", utils.ErrMissingRequired)
	}
	return nil
}
