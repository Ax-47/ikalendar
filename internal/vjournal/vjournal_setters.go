package vjournal

import (
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (jour *VJournal) SetUID(uid string) error {
	return parsehelper.SetOnceValue(&jour.UID, uid, string(PropUID))
}

func (jour *VJournal) SetDTSTAMP(it share.ITIME) error {
	return parsehelper.SetOnceITIME(&jour.DTSTAMP, it, string(PropDTSTAMP))
}

// ── Time ──────────────────────────────────────────────────────────────────────

func (jour *VJournal) SetDTSTART(it share.ITIME) error {
	return parsehelper.SetOnce(&jour.DTSTART, parsehelper.Ptr(it), string(PropDTSTART))
}

func (jour *VJournal) SetCreated(it share.ITIME) error {
	return parsehelper.SetOnce(&jour.CREATED, parsehelper.Ptr(it), string(PropCREATED))
}

func (jour *VJournal) SetLastModified(it share.ITIME) error {
	return parsehelper.SetOnce(&jour.LASTMODIFIED, parsehelper.Ptr(it), string(PropLASTMODIFIED))
}

func (jour *VJournal) SetRRule(r share.RECUR) error {
	return parsehelper.SetOnce(&jour.RRULE, parsehelper.Ptr(r), string(PropRRULE))
}

// ── Int ───────────────────────────────────────────────────────────────────────

func (jour *VJournal) SetSequence(n int) error {
	return parsehelper.SetOnce(&jour.SEQUENCE, parsehelper.Ptr(n), string(PropSEQUENCE))
}

// ── String ────────────────────────────────────────────────────────────────────

func (jour *VJournal) SetClass(s string) error {
	return parsehelper.SetOnce(&jour.CLASS, parsehelper.Ptr(s), string(PropCLASS))
}

func (jour *VJournal) SetDescription(s string) error {
	return parsehelper.SetOnce(&jour.DESCRIPTION, parsehelper.Ptr(s), string(PropDESCRIPTION))
}

func (jour *VJournal) SetOrganizer(s string) error {
	return parsehelper.SetOnce(&jour.ORGANIZER, parsehelper.Ptr(s), string(PropORGANIZER))
}

func (jour *VJournal) SetStatus(s string) error {
	return parsehelper.SetOnce(&jour.STATUS, parsehelper.Ptr(s), string(PropSTATUS))
}

func (jour *VJournal) SetSummary(s string) error {
	return parsehelper.SetOnce(&jour.SUMMARY, parsehelper.Ptr(s), string(PropSUMMARY))
}

func (jour *VJournal) SetURL(s string) error {
	return parsehelper.SetOnce(&jour.URL, parsehelper.Ptr(s), string(PropURL))
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func (jour *VJournal) AddAttach(uri string) error {
	jour.ATTACH = append(jour.ATTACH, share.ATTACH{URI: parsehelper.Ptr(uri)}) // TODO: construct
	return nil
}

func (jour *VJournal) AddAttendee(address string, params map[string]string) error {
	a := share.CALADDRESS{Address: address} // TODO: constructure
	if cn, ok := params["CN"]; ok {
		a.CN = &cn
	}
	if role, ok := params["ROLE"]; ok {
		a.Role = &role
	}
	jour.ATTENDEE = append(jour.ATTENDEE, a)
	return nil
}

func (jour *VJournal) AddCategory(cat string) {
	jour.CATEGORIES = append(jour.CATEGORIES, cat)
}

func (jour *VJournal) AddComment(s string) error {
	jour.COMMENT = append(jour.COMMENT, s)
	return nil
}

func (jour *VJournal) AddContact(s string) error {
	jour.CONTACT = append(jour.CONTACT, s)
	return nil
}

func (jour *VJournal) AddExDate(it share.ITIME) error {
	jour.EXDATE = append(jour.EXDATE, it)
	return nil
}

func (jour *VJournal) AddRDate(it share.ITIME) error {
	jour.RDATE = append(jour.RDATE, it)
	return nil
}

func (jour *VJournal) AddRelated(uid string) error {
	jour.RELATED = append(jour.RELATED, share.NewRELATED(uid))
	return nil
}

func (jour *VJournal) AddRequestStatus(rs share.RequestStatus) error {
	jour.REQUESTSTATUS = append(jour.REQUESTSTATUS, rs)
	return nil
}
