package componants

import "strings"

type CalendarLike interface {
	Encode(*EncodeContext)
	GetMethod() *string
}

type EncodeContext struct {
	Builder  *strings.Builder
	Calendar CalendarLike
}
