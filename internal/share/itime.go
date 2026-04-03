package share

import (
	"strings"
	"time"
)

const (
	iCalUTCFormat      = "20060102T150405Z"
	iCalFloatingFormat = "20060102T150405"
	iCalDateFormat     = "20060102"
)

type ITIME struct {
	// holds things like TZID or VALUE.
	// {"TZID": "America/New_York"}
	Parameters map[string]string

	Time time.Time

	// IsDateOnly is a helper flag. If true, the property had VALUE=DATE
	// making it all-day event
	IsDateOnly bool
}
type ITIMEOption func(*ITIME)

func (it ITIME) IsZero() bool {
	return it.Time.IsZero() && !it.IsDateOnly && len(it.Parameters) == 0
}

func ParseITIME(params map[string]string, value string) ITIME {
	it := ITIME{
		Parameters: params,
	}

	value = strings.TrimSpace(value)
	if value == "" {
		return it
	}

	valParam := ""
	if params != nil {
		valParam = strings.ToUpper(params["VALUE"])
	}

	if valParam == "DATE" || len(value) == 8 {
		it.IsDateOnly = true
		t, err := time.Parse(iCalDateFormat, value[:8])
		if err == nil {
			it.Time = t
		}
		return it
	}

	if strings.HasSuffix(value, "Z") {
		t, err := time.Parse(iCalUTCFormat, value)
		if err == nil {
			it.Time = t
		}
		return it
	}

	t, err := time.Parse(iCalFloatingFormat, value)
	if err == nil {
		it.Time = t
	}
	return it
}

func FormatITIME(it ITIME) string {
	if it.IsDateOnly {
		return it.Time.Format(iCalDateFormat)
	}
	if it.Time.Location().String() == "UTC" {
		return it.Time.Format(iCalUTCFormat)
	}
	return it.Time.Format(iCalFloatingFormat)
}

func NewITIMEUTC(t time.Time) ITIME {
	return ITIME{
		Parameters: make(map[string]string),
		Time:       t.UTC(),
		IsDateOnly: false,
	}
}

func WithTZID(tzid string) ITIMEOption {
	return func(it *ITIME) {
		it.Parameters["TZID"] = tzid
	}
}

func WithDateOnly() ITIMEOption {
	return func(it *ITIME) {
		it.IsDateOnly = true
		it.Parameters["VALUE"] = "DATE"
	}
}

func WithParam(key, value string) ITIMEOption {
	return func(it *ITIME) {
		it.Parameters[key] = value
	}
}

func NewITIME(t time.Time, opts ...ITIMEOption) ITIME {
	it := ITIME{
		Parameters: make(map[string]string),
		Time:       t,
		IsDateOnly: false,
	}
	for _, opt := range opts {
		opt(&it)
	}
	return it
}
