package icalendar

import (
	"strings"
	"time"
)

const (
	iCalUTCFormat      = "20060102T150405Z"
	iCalFloatingFormat = "20060102T150405"
	iCalDateFormat     = "20060102"
)

func itimePtr(it ITIME) *ITIME { return &it }

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

func NewITIMELocal(t time.Time) ITIME {
	params := make(map[string]string)
	loc := t.Location()

	if loc != nil && loc.String() != "UTC" && loc.String() != "Local" {
		params["TZID"] = loc.String()
	}

	return ITIME{
		Parameters: params,
		Time:       t,
		IsDateOnly: false,
	}
}

func NewITIMEDate(t time.Time) ITIME {
	return ITIME{
		Parameters: map[string]string{"VALUE": "DATE"},
		Time:       t,
		IsDateOnly: true,
	}
}
