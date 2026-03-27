package icalendar

import (
	"fmt"
	"strings"
)

type EncodeContext struct {
	Builder  *strings.Builder
	Calendar *VCalendar
}

type IEncodable interface {
	Encode(ctx *EncodeContext)
}

func WriteProperty(sb *strings.Builder, name, value string) {
	line := fmt.Sprintf("%s:%s", name, value)
	const maxLen = 75

	for len(line) > 0 {
		if len(line) <= maxLen {
			sb.WriteString(line + "\r\n")
			break
		}
		sb.WriteString(line[:maxLen] + "\r\n ") // fold with a space
		line = line[maxLen:]
	}
}
