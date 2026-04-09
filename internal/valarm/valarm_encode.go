package valarm

import (
	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/encode"
	"github.com/minoplhy/ikalendar/internal/utils"
)

func (a *VAlarm) Encode(ctx *componants.EncodeContext) {
	b := ctx.Builder

	encode.WriteProperty(b, "BEGIN", "VALARM")

	encode.WriteString(b, "ACTION", a.Action)
	if a.Trigger != nil {
		encode.WriteString(b, "TRIGGER", utils.Ptr(a.Trigger.FormatDURATION()))
	}

	encode.WriteString(b, "DESCRIPTION", a.Description)
	encode.WriteString(b, "SUMMARY", a.Summary)

	encode.WriteProperty(b, "END", "VALARM")
}
