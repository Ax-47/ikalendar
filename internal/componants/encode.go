package componants

import "strings"

type CalendarLike interface {
	Encodable
	GetMethod() *string
}

type EncodeContext struct {
	Builder  *strings.Builder
	Calendar CalendarLike
}
type (
	Encodable interface{ Encode(*EncodeContext) }
)
