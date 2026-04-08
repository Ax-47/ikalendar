package vtimezone

import "github.com/minoplhy/ikalendar/internal/share"

type TimezoneOption func(*VTimezone) error

func WithDaylight(d *Daylight) TimezoneOption {
	return func(tz *VTimezone) error {
		return tz.AddChild(d)
	}
}

func WithStandard(s *Standard) TimezoneOption {
	return func(tz *VTimezone) error {
		return tz.AddChild(s)
	}
}

func WithLastModified(it share.ITIME) TimezoneOption {
	return func(tz *VTimezone) error {
		return tz.SetLastModified(it)
	}
}

func WithTZURL(url string) TimezoneOption {
	return func(tz *VTimezone) error {
		return tz.SetTZURL(url)
	}
}

func NewTimezone(tzid string, opts ...TimezoneOption) (*VTimezone, error) {
	tz := &VTimezone{
		TZID: tzid,
	}

	for _, opt := range opts {
		if err := opt(tz); err != nil {
			return nil, err
		}
	}

	return tz, nil
}
