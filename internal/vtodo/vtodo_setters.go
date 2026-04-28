package vtodo

import (
	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

// ── Required ──────────────────────────────────────────────────────────────────

func (vt *VTodo) SetUID(uid string) error {
	return utils.SetOnceValue(&vt.UID, uid, string(PropUID))
}

func (vt *VTodo) SetDTSTAMP(it share.ITIME) error {
	return utils.SetOnceITIME(&vt.DTSTAMP, it, string(PropDTSTAMP))
}

// ── Time ──────────────────────────────────────────────────────────────────────

func (vt *VTodo) SetDTSTART(it share.ITIME) error {
	return utils.SetOnce(&vt.DTSTART, new(it), string(PropDTSTART))
}

func (vt *VTodo) SetCreated(it share.ITIME) error {
	return utils.SetOnce(&vt.CREATED, new(it), string(PropCREATED))
}

func (vt *VTodo) SetLastModified(it share.ITIME) error {
	return utils.SetOnce(&vt.LASTMODIFIED, new(it), string(PropLASTMODIFIED))
}

func (vt *VTodo) SetRRule(r share.RECUR) error {
	return utils.SetOnce(&vt.RRULE, new(r), string(PropRRULE))
}

// ── Int ───────────────────────────────────────────────────────────────────────

func (vt *VTodo) SetSequence(n int) error {
	return utils.SetOnce(&vt.SEQUENCE, new(n), string(PropSEQUENCE))
}

// ── String ────────────────────────────────────────────────────────────────────

func (vt *VTodo) SetClass(s string) error {
	return utils.SetOnce(&vt.CLASS, new(s), string(PropCLASS))
}

func (vt *VTodo) SetDescription(s string) error {
	return utils.SetOnce(&vt.DESCRIPTION, new(s), string(PropDESCRIPTION))
}

func (vt *VTodo) SetOrganizer(s string) error {
	return utils.SetOnce(&vt.ORGANIZER, new(s), string(PropORGANIZER))
}

func (vt *VTodo) SetStatus(s string) error {
	return utils.SetOnce(&vt.STATUS, new(s), string(PropSTATUS))
}

func (vt *VTodo) SetSummary(s string) error {
	return utils.SetOnce(&vt.SUMMARY, new(s), string(PropSUMMARY))
}

func (vt *VTodo) SetURL(s string) error {
	return utils.SetOnce(&vt.URL, new(s), string(PropURL))
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func (vt *VTodo) AddAttach(uri string) error {
	vt.ATTACH = append(vt.ATTACH, share.ATTACH{URI: new(uri)}) // TODO: construct
	return nil
}

func (vt *VTodo) AddAttendee(address string, params map[string]string) error {
	a := share.CALADDRESS{Address: address} // TODO: constructure
	if cn, ok := params["CN"]; ok {
		a.CN = &cn
	}
	if role, ok := params["ROLE"]; ok {
		a.Role = &role
	}
	vt.ATTENDEE = append(vt.ATTENDEE, a)
	return nil
}

func (vt *VTodo) AddCategory(cat string) {
	vt.CATEGORIES = append(vt.CATEGORIES, cat)
}

func (vt *VTodo) AddComment(s string) error {
	vt.COMMENT = append(vt.COMMENT, s)
	return nil
}

func (vt *VTodo) AddContact(s string) error {
	vt.CONTACT = append(vt.CONTACT, s)
	return nil
}

func (vt *VTodo) AddExDate(it share.ITIME) error {
	vt.EXDATE = append(vt.EXDATE, it)
	return nil
}

func (vt *VTodo) AddRDate(it share.ITIME) error {
	vt.RDATE = append(vt.RDATE, it)
	return nil
}

func (vt *VTodo) AddRelated(uid string) error {
	vt.RELATED = append(vt.RELATED, share.NewRELATED(uid))
	return nil
}

func (vt *VTodo) AddRequestStatus(rs share.RequestStatus) error {
	vt.REQUESTSTATUS = append(vt.REQUESTSTATUS, rs)
	return nil
}
