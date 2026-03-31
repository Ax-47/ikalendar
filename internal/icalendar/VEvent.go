package icalendar

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/parse"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

func (ev *VEvent) ProcessProperty(prop parse.Property) error {
	switch prop.Name {

	/* REQUIRED: MUST NOT occur more than once*/
	case "UID":
		if ev.UID != "" {
			return fmt.Errorf("%w: UID", parsehelper.ErrDuplicateProperty)
		}
		ev.UID = prop.Value

	case "DTSTAMP":
		if !ev.DTSTAMP.Time.IsZero() {
			return fmt.Errorf("%w: DTSTAMP", parsehelper.ErrDuplicateProperty)
		}
		ev.DTSTAMP = ParseITIME(prop.Params, prop.Value)

	/* REQUIRED if method isn't specified MUST NOT occur more than once */
	case "DTSTART":
		if ev.DTSTART != nil {
			return fmt.Errorf("%w: DTSTART", parsehelper.ErrDuplicateProperty)
		}
		ev.DTSTART = itimePtr(ParseITIME(prop.Params, prop.Value))

	/* OPTIONAL: MUST NOT occur more than once*/
	case "SUMMARY":
		if ev.SUMMARY != nil {
			return fmt.Errorf("%w: SUMMARY", parsehelper.ErrDuplicateProperty)
		}
		ev.SUMMARY = parsehelper.StrPtr(prop.Value)

	case "DESCRIPTION":
		if ev.DESCRIPTION != nil {
			return fmt.Errorf("%w: DESCRIPTION", parsehelper.ErrDuplicateProperty)
		}
		ev.DESCRIPTION = parsehelper.StrPtr(prop.Value)

	case "LOCATION":
		if ev.LOCATION != nil {
			return fmt.Errorf("%w: LOCATION", parsehelper.ErrDuplicateProperty)
		}
		ev.LOCATION = parsehelper.StrPtr(prop.Value)

	case "GEO":
		if ev.GEO != nil {
			return fmt.Errorf("%w: GEO", parsehelper.ErrDuplicateProperty)
		}
		ev.GEO = parsehelper.StrPtr(prop.Value)

	case "STATUS":
		if ev.STATUS != nil {
			return fmt.Errorf("%w: STATUS", parsehelper.ErrDuplicateProperty)
		}
		ev.STATUS = parsehelper.StrPtr(prop.Value)

	case "CLASS":
		if ev.CLASS != nil {
			return fmt.Errorf("%w: CLASS", parsehelper.ErrDuplicateProperty)
		}
		ev.CLASS = parsehelper.StrPtr(prop.Value)

	case "ORGANIZER":
		if ev.ORGANIZER != nil {
			return fmt.Errorf("%w: ORGANIZER", parsehelper.ErrDuplicateProperty)
		}
		ev.ORGANIZER = parsehelper.StrPtr(prop.Value)

	case "PRIORITY":
		if ev.PRIORITY != nil {
			return fmt.Errorf("%w: PRIORITY", parsehelper.ErrDuplicateProperty)
		}
		pri, err := parsehelper.IntPtr(prop.Value)
		if err != nil {
			return fmt.Errorf("invalid PRIORITY: %w", err)
		}
		ev.PRIORITY = pri

	case "SEQUENCE":
		if ev.SEQUENCE != nil {
			return fmt.Errorf("%w: SEQUENCE", parsehelper.ErrDuplicateProperty)
		}
		seq, err := parsehelper.IntPtr(prop.Value)
		if err != nil {
			return fmt.Errorf("invalid SEQUENCE: %w", err)
		}
		ev.SEQUENCE = seq
	/* OPTIONAL: SHOULD NOT occur more than once*/
	case "RRULE":
		if ev.RRULE != nil {
			return fmt.Errorf("%w: RRULE", parsehelper.ErrDuplicateProperty)
		}
		ev.RRULE = &RECUR{} // unimplemented

	/* OPTIONAL: SHOULD NOT occur in same VEVENT*/
	case "DTEND":
		if ev.DTEND != nil {
			return fmt.Errorf("%w: DTEND", parsehelper.ErrDuplicateProperty)
		}
		ev.DTEND = itimePtr(ParseITIME(prop.Params, prop.Value))

	// case "DURATION":
	// unimplemented

	/* OPTIONAL: MAY occur more than once*/
	case "ATTACH":
		ev.ATTACH = []ATTACH{} // unimplemented
	case "ATTENDEE":
		ev.ATTENDEE = []CALADDRESS{} // unimplemented
	case "CATEGORIES":
		ev.CATEGORIES = append(ev.CATEGORIES, prop.Value)
	case "COMMENT":
		ev.COMMENT = append(ev.COMMENT, prop.Value)
	case "CONTACT":
		ev.CONTACT = append(ev.CONTACT, prop.Value)
	case "EXDATE":
		ev.EXDATE = append(ev.EXDATE, ParseITIME(prop.Params, prop.Value))
	case "REQUESTSTATUS":
		ev.REQUESTSTATUS = append(ev.REQUESTSTATUS, prop.Value)
	case "RELATED":
		ev.RELATED = []RELATED{} // unimplemented
	case "RDATE":
		ev.RDATE = append(ev.RDATE, prop.Value) // unimplemented
	case "RESOURCES":
		ev.RESOURCES = append(ev.RESOURCES, prop.Value)

	}

	return nil
}

func (ev *VEvent) Encode(ctx *EncodeContext) {
	WriteProperty(ctx.Builder, "BEGIN", "VEVENT")
	WriteProperty(ctx.Builder, "UID", ev.UID)

	// DTSTAMP is required
	WriteProperty(ctx.Builder, "DTSTAMP", FormatITIME(ev.DTSTAMP))

	if ev.DTSTART != nil {
		WriteProperty(ctx.Builder, "DTSTART", FormatITIME(*ev.DTSTART))
	} else if ev.DTSTART == nil && ctx.Calendar.METHOD == nil {
		WriteProperty(ctx.Builder, "DTSTART", FormatITIME(ev.DTSTAMP)) // NON STANDARD- we do use DSTAMP VALUE if DTSTART NOT PROVIDED
	}

	if ev.DTEND != nil {
		WriteProperty(ctx.Builder, "DTEND", FormatITIME(*ev.DTEND))
	}
	if ev.SUMMARY != nil {
		WriteProperty(ctx.Builder, "SUMMARY", *ev.SUMMARY)
	}
	if ev.DESCRIPTION != nil {
		WriteProperty(ctx.Builder, "DESCRIPTION", *ev.DESCRIPTION)
	}
	if ev.GEO != nil {
		WriteProperty(ctx.Builder, "GEO", *ev.GEO)
	}
	if ev.LOCATION != nil {
		WriteProperty(ctx.Builder, "LOCATION", *ev.LOCATION)
	}
	if ev.STATUS != nil {
		WriteProperty(ctx.Builder, "STATUS", *ev.STATUS)
	}

	if len(ev.CATEGORIES) > 0 {
		WriteProperty(ctx.Builder, "CATEGORIES", strings.Join(ev.CATEGORIES, ","))
	}

	WriteProperty(ctx.Builder, "END", "VEVENT")
}

func (ev *VEvent) AddChild(child parse.IComponent) error {
	// Implementation here for future expansion
	return fmt.Errorf("%w: VEVENT", parsehelper.ErrNoChildrenAllowed)
}

func (ev *VEvent) Validate() error {
	if ev.UID == "" {
		return fmt.Errorf("%w: VEVENT missing UID", parsehelper.ErrMissingRequired)
	}
	if ev.DTSTAMP.Time.IsZero() {
		return fmt.Errorf("%w: VEVENT missing DTSTAMP", parsehelper.ErrMissingRequired)
	}

	if ev.DTEND != nil && ev.DURATION != nil {
		return fmt.Errorf("%w: VEVENT cannot have both DTEND and DURATION", parsehelper.ErrMutuallyExclusive)
	}

	return nil
}
