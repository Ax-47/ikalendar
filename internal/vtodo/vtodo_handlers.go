package vtodo

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/componants"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
	"github.com/minoplhy/ikalendar/internal/share"
)

type vtodoHandler func(*VTodo, componants.Property) error

const (
	// Required
	PropUID     share.PropertyName = "UID"
	PropDTSTAMP share.PropertyName = "DTSTAMP"
	PropDTSTART share.PropertyName = "DTSTART"

	// Optional single
	PropDTEND        share.PropertyName = "DTEND"
	PropDURATION     share.PropertyName = "DURATION"
	PropCLASS        share.PropertyName = "CLASS"
	PropCREATED      share.PropertyName = "CREATED"
	PropDESCRIPTION  share.PropertyName = "DESCRIPTION"
	PropGEO          share.PropertyName = "GEO"
	PropLASTMODIFIED share.PropertyName = "LAST-MODIFIED"
	PropLOCATION     share.PropertyName = "LOCATION"
	PropORGANIZER    share.PropertyName = "ORGANIZER"
	PropPRIORITY     share.PropertyName = "PRIORITY"
	PropSEQUENCE     share.PropertyName = "SEQUENCE"
	PropSTATUS       share.PropertyName = "STATUS"
	PropSUMMARY      share.PropertyName = "SUMMARY"
	PropTRANSP       share.PropertyName = "TRANSP"
	PropURL          share.PropertyName = "URL"
	PropRRULE        share.PropertyName = "RRULE"

	// Optional multi
	PropATTACH        share.PropertyName = "ATTACH"
	PropATTENDEE      share.PropertyName = "ATTENDEE"
	PropCATEGORIES    share.PropertyName = "CATEGORIES"
	PropCOMMENT       share.PropertyName = "COMMENT"
	PropCONTACT       share.PropertyName = "CONTACT"
	PropEXDATE        share.PropertyName = "EXDATE"
	PropREQUESTSTATUS share.PropertyName = "REQUEST-STATUS"
	PropRELATED       share.PropertyName = "RELATED-TO"
	PropRDATE         share.PropertyName = "RDATE"
	PropRESOURCES     share.PropertyName = "RESOURCES"
) // RFC 5545 §3.3.6
var vtodoHandlers = map[share.PropertyName]vtodoHandler{
	PropUID:          handleUID,
	PropDTSTAMP:      handleDTSTAMP,
	PropDTSTART:      handleDTSTART,
	PropCLASS:        handleClass,
	PropCREATED:      handleCreated,
	PropDESCRIPTION:  handleDescription,
	PropLASTMODIFIED: handleLastModified,
	PropORGANIZER:    handleOrganizer,
	PropSEQUENCE:     handleSequence,
	PropSTATUS:       handleStatus,
	PropSUMMARY:      handleSummary,
	PropURL:          handleURL,
	PropRRULE:        handleRRule,

	// multi
	PropATTACH:        handleAttach,
	PropATTENDEE:      handleAttendee,
	PropCATEGORIES:    handleCategories,
	PropCOMMENT:       handleComment,
	PropCONTACT:       handleContact,
	PropEXDATE:        handleExDate,
	PropREQUESTSTATUS: handleRequestStatus,
	PropRELATED:       handleRelated,
	PropRDATE:         handleRDate,
}

// ── Required ──────────────────────────────────────────────────────────────────

func handleUID(vt *VTodo, prop componants.Property) error {
	return vt.SetUID(prop.Value)
}

func handleDTSTAMP(vt *VTodo, prop componants.Property) error {
	return vt.SetDTSTAMP(share.ParseITIME(prop.Params, prop.Value))
}

func handleDTSTART(vt *VTodo, prop componants.Property) error {
	return vt.SetDTSTART(share.ParseITIME(prop.Params, prop.Value))
}

// ── Optional single ───────────────────────────────────────────────────────────

func handleCreated(vt *VTodo, prop componants.Property) error {
	return vt.SetCreated(share.ParseITIME(prop.Params, prop.Value))
}

func handleLastModified(vt *VTodo, prop componants.Property) error {
	return vt.SetLastModified(share.ParseITIME(prop.Params, prop.Value))
}

func handleRRule(vt *VTodo, prop componants.Property) error {
	r, err := share.ParseRECUR(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid RRULE: %w", err)
	}
	return vt.SetRRule(r)
}

func handleSequence(vt *VTodo, prop componants.Property) error {
	n, err := parsehelper.IntPtr(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid SEQUENCE: %w", err)
	}
	return vt.SetSequence(*n)
}

// ── Pure string fields ────────────────────────────────────────────────────────

func handleClass(vt *VTodo, prop componants.Property) error {
	return vt.SetClass(prop.Value)
}

func handleDescription(vt *VTodo, prop componants.Property) error {
	return vt.SetDescription(prop.Value)
}

func handleOrganizer(vt *VTodo, prop componants.Property) error {
	return vt.SetOrganizer(prop.Value)
}

func handleStatus(vt *VTodo, prop componants.Property) error {
	return vt.SetStatus(prop.Value)
}

func handleSummary(vt *VTodo, prop componants.Property) error {
	return vt.SetSummary(prop.Value)
}

func handleURL(vt *VTodo, prop componants.Property) error {
	return vt.SetURL(prop.Value)
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func handleAttach(vt *VTodo, prop componants.Property) error {
	return vt.AddAttach(prop.Value)
}

func handleAttendee(vt *VTodo, prop componants.Property) error {
	return vt.AddAttendee(prop.Value, prop.Params)
}

func handleCategories(vt *VTodo, prop componants.Property) error {
	for _, cat := range strings.Split(prop.Value, ",") {
		if cat = strings.TrimSpace(cat); cat != "" {
			vt.AddCategory(cat)
		}
	}
	return nil
}

func handleComment(vt *VTodo, prop componants.Property) error {
	return vt.AddComment(prop.Value)
}

func handleContact(vt *VTodo, prop componants.Property) error {
	return vt.AddContact(prop.Value)
}

func handleExDate(vt *VTodo, prop componants.Property) error {
	return vt.AddExDate(share.ParseITIME(prop.Params, prop.Value))
}

func handleRDate(vt *VTodo, prop componants.Property) error {
	return vt.AddRDate(share.ParseITIME(prop.Params, prop.Value))
}

func handleRelated(vt *VTodo, prop componants.Property) error {
	return vt.AddRelated(prop.Value)
}

func handleRequestStatus(vt *VTodo, prop componants.Property) error {
	parts := strings.SplitN(prop.Value, ";", 3)
	rs := share.RequestStatus{Code: parts[0]} // TODO: construct
	if len(parts) > 1 {
		rs.Description = parts[1]
	}
	if len(parts) > 2 {
		rs.Extra = parsehelper.Ptr(parts[2])
	}
	return vt.AddRequestStatus(rs)
}
