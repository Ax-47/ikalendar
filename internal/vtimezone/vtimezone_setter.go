package vtimezone

import (
	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (tz *VTimezone) SetTZID(id string) error {
	return utils.SetOnceValue(&tz.TZID, id, "TZID")
}

// ── Optional ──────────────────────────────────────────────────────────────────

func (tz *VTimezone) SetLastModified(it share.ITIME) error {
	return utils.SetOnce(&tz.LASTMODIFIED, new(it), "LAST-MODIFIED")
}

func (tz *VTimezone) SetTZURL(url string) error {
	return utils.SetOnce(&tz.TZURL, new(url), "TZURL")
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
