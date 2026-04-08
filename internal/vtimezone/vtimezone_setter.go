package vtimezone

import (
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (tz *VTimezone) SetTZID(id string) error {
	return parsehelper.SetOnceValue(&tz.TZID, id, "TZID")
}

// ── Optional ──────────────────────────────────────────────────────────────────

func (tz *VTimezone) SetLastModified(it share.ITIME) error {
	return parsehelper.SetOnce(&tz.LASTMODIFIED, parsehelper.Ptr(it), "LAST-MODIFIED")
}

func (tz *VTimezone) SetTZURL(url string) error {
	return parsehelper.SetOnce(&tz.TZURL, parsehelper.Ptr(url), "TZURL")
}

// ── Child Components ──────────────────────────────────────────────────────────

func (tz *VTimezone) AddStandard(s Standard) error {
	// RFC: STANDARD สามารถมีหลายตัวได้ (ตาม rule change)
	tz.STANDARD = append(tz.STANDARD, s)
	return nil
}

func (tz *VTimezone) AddDaylight(d Daylight) error {
	// RFC: DAYLIGHT ก็หลายตัวได้
	tz.DAYLIGHT = append(tz.DAYLIGHT, d)
	return nil
}
