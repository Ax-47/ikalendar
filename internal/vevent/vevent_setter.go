package vevent

import (
	"fmt"

	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (ev *VEvent) SetUID(uid string) error {
	return parsehelper.SetOnceValue(&ev.UID, uid, string(PropUID))
}

func (ev *VEvent) SetDTSTAMP(it share.ITIME) error {
	return parsehelper.SetOnceITIME(&ev.DTSTAMP, it, string(PropDTSTAMP))
}

// ── Time ──────────────────────────────────────────────────────────────────────

func (ev *VEvent) SetDTSTART(it share.ITIME) error {
	return parsehelper.SetOnce(&ev.DTSTART, parsehelper.Ptr(it), string(PropDTSTART))
}

func (ev *VEvent) SetDTEND(it share.ITIME) error {
	if ev.DURATION != nil {
		return fmt.Errorf("%w: DTEND and DURATION", parsehelper.ErrMutuallyExclusive)
	}
	return parsehelper.SetOnce(&ev.DTEND, parsehelper.Ptr(it), string(PropDTEND))
}

func (ev *VEvent) SetDuration(d share.DURATION) error {
	if ev.DTEND != nil {
		return fmt.Errorf("%w: DURATION and DTEND", parsehelper.ErrMutuallyExclusive)
	}
	return parsehelper.SetOnce(&ev.DURATION, parsehelper.Ptr(d), string(PropDURATION))
}

func (ev *VEvent) SetCreated(it share.ITIME) error {
	if ev.CREATED != nil {
		return fmt.Errorf("%w: CREATED", parsehelper.ErrDuplicateProperty)
	}
	return parsehelper.SetOnce(&ev.CREATED, parsehelper.Ptr(it), string(PropCREATED))
}

func (ev *VEvent) SetLastModified(it share.ITIME) error {
	return parsehelper.SetOnce(&ev.LASTMODIFIED, parsehelper.Ptr(it), string(PropLASTMODIFIED))
}

func (ev *VEvent) SetRRule(r share.RECUR) error {
	return parsehelper.SetOnce(&ev.RRULE, parsehelper.Ptr(r), string(PropRRULE))
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
	ev.ATTACH = append(ev.ATTACH, share.ATTACH{URI: parsehelper.Ptr(uri)}) // TODO: construct
	return nil
}

func (ev *VEvent) AddAttendee(address string, params map[string]string) error {
	a := share.CALADDRESS{Address: address} // TODO: constructure
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

func (ev *VEvent) AddExDate(it share.ITIME) error {
	ev.EXDATE = append(ev.EXDATE, it)
	return nil
}

func (ev *VEvent) AddRDate(it share.ITIME) error {
	ev.RDATE = append(ev.RDATE, it)
	return nil
}

func (ev *VEvent) AddRelated(uid string) error {
	ev.RELATED = append(ev.RELATED, share.NewRELATED(uid))
	return nil
}

func (ev *VEvent) AddResource(s string) error {
	ev.RESOURCES = append(ev.RESOURCES, s)
	return nil
}

func (ev *VEvent) AddRequestStatus(rs share.RequestStatus) error {
	ev.REQUESTSTATUS = append(ev.REQUESTSTATUS, rs)
	return nil
}
