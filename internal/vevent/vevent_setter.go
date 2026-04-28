package vevent

import (
	"fmt"

	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

// ── Required ──────────────────────────────────────────────────────────────────
func (ev *VEvent) SetUID(uid string) error {
	return utils.SetOnceValue(&ev.UID, uid, string(PropUID))
}

func (ev *VEvent) SetDTSTAMP(it share.ITIME) error {
	return utils.SetOnceITIME(&ev.DTSTAMP, it, string(PropDTSTAMP))
}

// ── Time ──────────────────────────────────────────────────────────────────────

func (ev *VEvent) SetDTSTART(it share.ITIME) error {
	return utils.SetOnce(&ev.DTSTART, new(it), string(PropDTSTART))
}

func (ev *VEvent) SetDTEND(it share.ITIME) error {
	if ev.DURATION != nil {
		return fmt.Errorf("%w: DTEND and DURATION", utils.ErrMutuallyExclusive)
	}
	return utils.SetOnce(&ev.DTEND, new(it), string(PropDTEND))
}

func (ev *VEvent) SetDuration(d share.DURATION) error {
	if ev.DTEND != nil {
		return fmt.Errorf("%w: DURATION and DTEND", utils.ErrMutuallyExclusive)
	}
	return utils.SetOnce(&ev.DURATION, new(d), string(PropDURATION))
}

func (ev *VEvent) SetCreated(it share.ITIME) error {
	if ev.CREATED != nil {
		return fmt.Errorf("%w: CREATED", utils.ErrDuplicateProperty)
	}
	return utils.SetOnce(&ev.CREATED, new(it), string(PropCREATED))
}

func (ev *VEvent) SetLastModified(it share.ITIME) error {
	return utils.SetOnce(&ev.LASTMODIFIED, new(it), string(PropLASTMODIFIED))
}

func (ev *VEvent) SetRRule(r share.RECUR) error {
	return utils.SetOnce(&ev.RRULE, new(r), string(PropRRULE))
}

// ── Int ───────────────────────────────────────────────────────────────────────

func (ev *VEvent) SetPriority(n int) error {
	return utils.SetOnce(&ev.PRIORITY, new(n), "PRIORITY")
}

func (ev *VEvent) SetSequence(n int) error {
	return utils.SetOnce(&ev.SEQUENCE, new(n), "SEQUENCE")
}

// ── String ────────────────────────────────────────────────────────────────────

func (ev *VEvent) SetClass(s string) error {
	return utils.SetOnce(&ev.CLASS, new(s), "CLASS")
}

func (ev *VEvent) SetDescription(s string) error {
	return utils.SetOnce(&ev.DESCRIPTION, new(s), "DESCRIPTION")
}

func (ev *VEvent) SetGeo(s string) error {
	return utils.SetOnce(&ev.GEO, new(s), "GEO")
}

func (ev *VEvent) SetLocation(s string) error {
	return utils.SetOnce(&ev.LOCATION, new(s), "LOCATION")
}

func (ev *VEvent) SetOrganizer(s string) error {
	return utils.SetOnce(&ev.ORGANIZER, new(s), "ORGANIZER")
}

func (ev *VEvent) SetStatus(s string) error {
	return utils.SetOnce(&ev.STATUS, new(s), "STATUS")
}

func (ev *VEvent) SetSummary(s string) error {
	return utils.SetOnce(&ev.SUMMARY, new(s), "SUMMARY")
}

func (ev *VEvent) SetTransp(s string) error {
	return utils.SetOnce(&ev.TRANSP, new(s), "TRANSP")
}

func (ev *VEvent) SetURL(s string) error {
	return utils.SetOnce(&ev.URL, new(s), "URL")
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func (ev *VEvent) AddAttach(uri string) error {
	ev.ATTACH = append(ev.ATTACH, share.ATTACH{URI: new(uri)}) // TODO: construct
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
