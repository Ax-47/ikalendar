package parsehelper

import (
	"errors"
	"strconv"
)

func StrPtr(s string) *string { return &s }
func IntPtr(s string) *int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &i
}

var (
	ErrDuplicateProperty = errors.New("RFC 5545 violation: property MUST NOT occur more than once")
	ErrMissingRequired   = errors.New("RFC 5545 violation: missing REQUIRED property")
	ErrMutuallyExclusive = errors.New("RFC 5545 violation: mutually exclusive properties found")
	ErrNoChildrenAllowed = errors.New("component does not support nested children")
	ErrNoCalendarFound   = errors.New("no VCALENDAR found in stream")
)
