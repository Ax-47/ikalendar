package encode

import (
	"github.com/minoplhy/ikalendar/internal/icalendar/models"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

func VAlarm(ctx *EncodeContext, a *models.VAlarm) {
	b := ctx.Builder

	writeProperty(b, "BEGIN", "VALARM")

	writeString(b, "ACTION", a.Action)
	if a.Trigger != nil {
		writeString(b, "TRIGGER", parsehelper.Ptr(models.FormatDURATION(*a.Trigger)))
	}

	writeString(b, "DESCRIPTION", a.Description)
	writeString(b, "SUMMARY", a.Summary)

	writeProperty(b, "END", "VALARM")
}
