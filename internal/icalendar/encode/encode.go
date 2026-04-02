package encode

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/icalendar/models"
)

// RFC 5545 §3.1 — contentline max length is 75 octets
const maxLineOctets = 75

type EncodeContext struct {
	Builder  *strings.Builder
	Calendar *models.VCalendar
}

func Marshal(cal *models.VCalendar) ([]byte, error) {
	if cal == nil {
		return nil, fmt.Errorf("calendar cannot be nil")
	}

	var sb strings.Builder
	ctx := &EncodeContext{
		Builder:  &sb,
		Calendar: cal,
	}

	VCalendar(ctx, cal)

	return []byte(sb.String()), nil
}

func writeProperty(sb *strings.Builder, name, value string) {
	line := fmt.Sprintf("%s:%s", name, value)
	for len(line) > 0 {
		if len(line) <= maxLineOctets {
			sb.WriteString(line + "\r\n")
			break
		}
		sb.WriteString(line[:maxLineOctets] + "\r\n ")
		line = line[maxLineOctets:]
	}
}

func writeString(b *strings.Builder, name models.PropertyName, val *string) {
	if val != nil {
		writeProperty(b, string(name), *val)
	}
}

func writeTime(b *strings.Builder, name models.PropertyName, val *models.ITIME) {
	if val != nil {
		writeProperty(b, string(name), models.FormatITIME(*val))
	}
}

func writeTimeWithParams(b *strings.Builder, name models.PropertyName, val *models.ITIME) {
	if val == nil || val.Time.IsZero() {
		return
	}
	prop := string(name)
	if tzid, ok := val.Parameters["TZID"]; ok {
		prop += ";TZID=" + tzid
	} else if val.IsDateOnly {
		prop += ";VALUE=DATE"
	}
	writeProperty(b, prop, models.FormatITIME(*val))
}

func writeInt(b *strings.Builder, name models.PropertyName, val *int) {
	if val != nil {
		writeProperty(b, string(name), fmt.Sprintf("%d", *val))
	}
}

func writeCalAddress(b *strings.Builder, name models.PropertyName, a models.CALADDRESS) {
	prop := string(name)
	if a.CN != nil {
		prop += ";CN=" + *a.CN
	}
	if a.Role != nil {
		prop += ";ROLE=" + *a.Role
	}
	writeProperty(b, prop, a.Address)
}

func writeRequestStatus(b *strings.Builder, rs models.RequestStatus) {
	value := rs.Code + ";" + rs.Description
	if rs.Extra != nil {
		value += ";" + *rs.Extra
	}
	writeProperty(b, "REQUEST-STATUS", value)
}
