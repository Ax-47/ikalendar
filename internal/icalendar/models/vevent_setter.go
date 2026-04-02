package models

import (
	"fmt"

	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (ev *VEvent) SetUID(uid string) error {
	if ev.UID != "" {
		return fmt.Errorf("%w: UID", parsehelper.ErrDuplicateProperty)
	}
	ev.UID = uid
	return nil
}

func (ev *VEvent) SetDTSTAMP(it ITIME) error {
	if !ev.DTSTAMP.Time.IsZero() {
		return fmt.Errorf("%w: DTSTAMP", parsehelper.ErrDuplicateProperty)
	}
	ev.DTSTAMP = it
	return nil
}

// ── Time ──────────────────────────────────────────────────────────────────────

func (ev *VEvent) SetDTSTART(it ITIME) error {
	if ev.DTSTART != nil {
		return fmt.Errorf("%w: DTSTART", parsehelper.ErrDuplicateProperty)
	}
	ev.DTSTART = &it
	return nil
}

func (ev *VEvent) SetDTEND(it ITIME) error {
	if ev.DTEND != nil {
		return fmt.Errorf("%w: DTEND", parsehelper.ErrDuplicateProperty)
	}
	if ev.DURATION != nil {
		return fmt.Errorf("%w: DTEND and DURATION", parsehelper.ErrMutuallyExclusive)
	}
	ev.DTEND = &it
	return nil
}

func (ev *VEvent) SetDuration(d DURATION) error {
	if ev.DURATION != nil {
		return fmt.Errorf("%w: DURATION", parsehelper.ErrDuplicateProperty)
	}
	if ev.DTEND != nil {
		return fmt.Errorf("%w: DURATION and DTEND", parsehelper.ErrMutuallyExclusive)
	}
	ev.DURATION = &d
	return nil
}

func (ev *VEvent) SetCreated(it ITIME) error {
	if ev.CREATED != nil {
		return fmt.Errorf("%w: CREATED", parsehelper.ErrDuplicateProperty)
	}
	ev.CREATED = &it
	return nil
}

func (ev *VEvent) SetLastModified(it ITIME) error {
	if ev.LASTMODIFIED != nil {
		return fmt.Errorf("%w: LAST-MODIFIED", parsehelper.ErrDuplicateProperty)
	}
	ev.LASTMODIFIED = &it
	return nil
}

func (ev *VEvent) SetRRule(r RECUR) error {
	if ev.RRULE != nil {
		return fmt.Errorf("%w: RRULE", parsehelper.ErrDuplicateProperty)
	}
	ev.RRULE = &r
	return nil
}

// ── Int ───────────────────────────────────────────────────────────────────────

func (ev *VEvent) SetPriority(n int) error {
	return parsehelper.SetOnce(&ev.PRIORITY, parsehelper.Ptr(n), "PRIORITY")
}

func (ev *VEvent) SetSequence(n int) error {
	return parsehelper.SetOnce(&ev.SEQUENCE, parsehelper.Ptr(n), "SEQUENCE")
}

// ── String ────────────────────────────────────────────────────────────────────

func (ev *VEvent) SetClass(s string) error {
	return parsehelper.SetOnce(&ev.CLASS, parsehelper.Ptr(s), "CLASS")
}

func (ev *VEvent) SetDescription(s string) error {
	return parsehelper.SetOnce(&ev.DESCRIPTION, parsehelper.Ptr(s), "DESCRIPTION")
}

func (ev *VEvent) SetGeo(s string) error {
	return parsehelper.SetOnce(&ev.GEO, parsehelper.Ptr(s), "GEO")
}

func (ev *VEvent) SetLocation(s string) error {
	return parsehelper.SetOnce(&ev.LOCATION, parsehelper.Ptr(s), "LOCATION")
}

func (ev *VEvent) SetOrganizer(s string) error {
	return parsehelper.SetOnce(&ev.ORGANIZER, parsehelper.Ptr(s), "ORGANIZER")
}

func (ev *VEvent) SetStatus(s string) error {
	return parsehelper.SetOnce(&ev.STATUS, parsehelper.Ptr(s), "STATUS")
}

func (ev *VEvent) SetSummary(s string) error {
	return parsehelper.SetOnce(&ev.SUMMARY, parsehelper.Ptr(s), "SUMMARY")
}

func (ev *VEvent) SetTransp(s string) error {
	return parsehelper.SetOnce(&ev.TRANSP, parsehelper.Ptr(s), "TRANSP")
}

func (ev *VEvent) SetURL(s string) error {
	return parsehelper.SetOnce(&ev.URL, parsehelper.Ptr(s), "URL")
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func (ev *VEvent) AddAttach(uri string) error {
	ev.ATTACH = append(ev.ATTACH, ATTACH{URI: parsehelper.Ptr(uri)})
	return nil
}

func (ev *VEvent) AddAttendee(address string, params map[string]string) error {
	a := CALADDRESS{Address: address}
	if cn, ok := params["CN"]; ok {
		a.CN = &cn
	}
	if role, ok := params["ROLE"]; ok {
		a.Role = &role
	}
	ev.ATTENDEE = append(ev.ATTENDEE, a)
	return nil
}

func (ev *VEvent) AddCategory(cat string) {
	ev.CATEGORIES = append(ev.CATEGORIES, cat)
}

func (ev *VEvent) AddComment(s string) error {
	ev.COMMENT = append(ev.COMMENT, s)
	return nil
}

func (ev *VEvent) AddContact(s string) error {
	ev.CONTACT = append(ev.CONTACT, s)
	return nil
}

func (ev *VEvent) AddExDate(it ITIME) error {
	ev.EXDATE = append(ev.EXDATE, it)
	return nil
}

func (ev *VEvent) AddRDate(it ITIME) error {
	ev.RDATE = append(ev.RDATE, it)
	return nil
}

func (ev *VEvent) AddRelated(uid string) error {
	ev.RELATED = append(ev.RELATED, NewRELATED(uid))
	return nil
}

func (ev *VEvent) AddResource(s string) error {
	ev.RESOURCES = append(ev.RESOURCES, s)
	return nil
}

func (ev *VEvent) AddRequestStatus(rs RequestStatus) error {
	ev.REQUESTSTATUS = append(ev.REQUESTSTATUS, rs)
	return nil
}
