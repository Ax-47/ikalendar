package builders

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/icalendar/models"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

type AlarmOption func(*models.VAlarm) error

// ── Required ──────────────────────────────────────────────────────────────────

func WithAction(action string) AlarmOption {
	return func(a *models.VAlarm) error {
		if action == "" {
			return fmt.Errorf("%w: ACTION cannot be empty", parsehelper.ErrMissingRequired)
		}
		a.Action = parsehelper.Ptr(action)
		return nil
	}
}

func WithTrigger(d models.DURATION) AlarmOption {
	return func(a *models.VAlarm) error {
		if a.Trigger != nil {
			return fmt.Errorf("%w: TRIGGER", parsehelper.ErrDuplicateProperty)
		}
		a.Trigger = parsehelper.Ptr(d)
		return nil
	}
}

// WithTriggerBefore creates a negative duration trigger — fires X hours/minutes before the event
// e.g. WithTriggerBefore(0, 15) → TRIGGER:-PT15M
// e.g. WithTriggerBefore(1, 30) → TRIGGER:-PT1H30M
func WithTriggerBefore(hours, minutes int) AlarmOption {
	return WithTrigger(models.DURATION{
		Negative: true,
		Hours:    hours,
		Minutes:  minutes,
	})
}

// WithTriggerAfter creates a positive duration trigger — fires X hours/minutes after the event
// e.g. WithTriggerAfter(0, 5) → TRIGGER:PT5M
func WithTriggerAfter(hours, minutes int) AlarmOption {
	return WithTrigger(models.DURATION{
		Negative: false,
		Hours:    hours,
		Minutes:  minutes,
	})
}

// WithTriggerDays — convenience for day-based triggers
// e.g. WithTriggerDays(1, true) → TRIGGER:-P1D (1 day before)
func WithTriggerDays(days int, before bool) AlarmOption {
	return WithTrigger(models.DURATION{
		Negative: before,
		Days:     days,
	})
}

// ── Optional ──────────────────────────────────────────────────────────────────

func WithAlarmDescription(desc string) AlarmOption {
	return func(a *models.VAlarm) error {
		if a.Description != nil {
			return fmt.Errorf("%w: DESCRIPTION", parsehelper.ErrDuplicateProperty)
		}
		a.Description = parsehelper.Ptr(desc)
		return nil
	}
}

func WithAlarmSummary(summary string) AlarmOption {
	return func(a *models.VAlarm) error {
		if a.Summary != nil {
			return fmt.Errorf("%w: SUMMARY", parsehelper.ErrDuplicateProperty)
		}
		a.Summary = parsehelper.Ptr(summary)
		return nil
	}
}

// ── Constructor ───────────────────────────────────────────────────────────────

func NewAlarm(opts ...AlarmOption) (*models.VAlarm, error) {
	a := &models.VAlarm{}
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

func validateAlarm(a *models.VAlarm) error {
	if a.Action == nil {
		return fmt.Errorf("%w: VALARM missing ACTION", parsehelper.ErrMissingRequired)
	}
	if a.Trigger == nil {
		return fmt.Errorf("%w: VALARM missing TRIGGER", parsehelper.ErrMissingRequired)
	}
	return nil
}
