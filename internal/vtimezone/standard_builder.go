package vtimezone

import "github.com/minoplhy/ikalendar/internal/share"

type StandardOption func(*Standard) error

func WithStandardDTSTART(it share.ITIME) StandardOption {
	return func(s *Standard) error {
		return s.SetDTSTART(it)
	}
}

func WithStandardTZOffsetFrom(offset string) StandardOption {
	return func(s *Standard) error {
		return s.SetTZOffsetFrom(offset)
	}
}

func WithStandardTZOffsetTo(offset string) StandardOption {
	return func(s *Standard) error {
		return s.SetTZOffsetTo(offset)
	}
}

func WithStandardTZName(name string) StandardOption {
	return func(s *Standard) error {
		return s.SetTZName(name)
	}
}

func WithStandardRRule(r share.RECUR) StandardOption {
	return func(s *Standard) error {
		return s.SetRRule(r)
	}
}

func WithStandardRDate(it share.ITIME) StandardOption {
	return func(s *Standard) error {
		return s.AddRDate(it)
	}
}

func NewStandard(opts ...StandardOption) (*Standard, error) {
	s := &Standard{}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	if err := s.Validate(); err != nil {
		return nil, err
	}

	return s, nil
}
