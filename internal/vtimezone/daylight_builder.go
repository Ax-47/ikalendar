package vtimezone

import "github.com/minoplhy/ikalendar/internal/share"

type DaylightOption func(*Daylight) error

func WithDaylightDTSTART(it share.ITIME) DaylightOption {
	return func(d *Daylight) error {
		return d.SetDTSTART(it)
	}
}

func WithDaylightTZOffsetFrom(offset string) DaylightOption {
	return func(d *Daylight) error {
		return d.SetTZOffsetFrom(offset)
	}
}

func WithDaylightTZOffsetTo(offset string) DaylightOption {
	return func(d *Daylight) error {
		return d.SetTZOffsetTo(offset)
	}
}

func WithDaylightTZName(name string) DaylightOption {
	return func(d *Daylight) error {
		return d.SetTZName(name)
	}
}

func WithDaylightRRule(r share.RECUR) DaylightOption {
	return func(d *Daylight) error {
		return d.SetRRule(r)
	}
}

func WithDaylightRDate(it share.ITIME) DaylightOption {
	return func(d *Daylight) error {
		return d.AddRDate(it)
	}
}

func NewDaylight(opts ...DaylightOption) (*Daylight, error) {
	d := &Daylight{}

	for _, opt := range opts {
		if err := opt(d); err != nil {
			return nil, err
		}
	}

	if err := d.Validate(); err != nil {
		return nil, err
	}

	return d, nil
}
