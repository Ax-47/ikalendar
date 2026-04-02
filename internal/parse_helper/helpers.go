package parsehelper

import (
	"errors"
	"fmt"
	"strconv"
)

//nolint:gocritic
func Ptr[T any](t T) *T { return &t }

func IntPtr(s string) (*int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func SetOnce[T any](field **T, value *T, name string) error {
	if *field != nil {
		return fmt.Errorf("%w: %s", ErrDuplicateProperty, name)
	}
	*field = value
	return nil
}

var (
	ErrDuplicateProperty = errors.New("RFC 5545 violation: property MUST NOT occur more than once")
	ErrMissingRequired   = errors.New("RFC 5545 violation: missing REQUIRED property")
	ErrMutuallyExclusive = errors.New("RFC 5545 violation: mutually exclusive properties found")
	ErrNoChildrenAllowed = errors.New("component does not support nested children")
	ErrNoCalendarFound   = errors.New("no VCALENDAR found in stream")
	ErrInvalidComponent  = errors.New("invalid component")
	ErrInvalidProperty   = errors.New("invalid component")
)
