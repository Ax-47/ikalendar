package encode

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/share"
)

// RFC 5545 §3.1 — contentline max length is 75 octets
const maxLineOctets = 75

func WriteProperty(sb *strings.Builder, name, value string) {
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

func WriteString(b *strings.Builder, name share.PropertyName, val *string) {
	if val != nil {
		WriteProperty(b, string(name), *val)
	}
}

func WriteTime(b *strings.Builder, name share.PropertyName, val *share.ITIME) {
	if val != nil {
		WriteProperty(b, string(name), share.FormatITIME(*val))
	}
}

func WriteTimeWithParams(b *strings.Builder, name share.PropertyName, val *share.ITIME) {
	if val == nil || val.Time.IsZero() {
		return
	}
	prop := string(name)
	if tzid, ok := val.Parameters["TZID"]; ok {
		prop += ";TZID=" + tzid
	} else if val.IsDateOnly {
		prop += ";VALUE=DATE"
	}
	WriteProperty(b, prop, share.FormatITIME(*val))
}

func WriteInt(b *strings.Builder, name share.PropertyName, val *int) {
	if val != nil {
		WriteProperty(b, string(name), fmt.Sprintf("%d", *val))
	}
}

func WriteCalAddress(b *strings.Builder, name share.PropertyName, a share.CALADDRESS) {
	prop := string(name)
	if a.CN != nil {
		prop += ";CN=" + *a.CN
	}
	if a.Role != nil {
		prop += ";ROLE=" + *a.Role
	}
	WriteProperty(b, prop, a.Address)
}

func WriteRequestStatus(b *strings.Builder, rs share.RequestStatus) {
	value := rs.Code + ";" + rs.Description
	if rs.Extra != nil {
		value += ";" + *rs.Extra
	}
	WriteProperty(b, "REQUEST-STATUS", value)
}
