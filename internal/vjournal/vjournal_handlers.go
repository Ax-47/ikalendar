package vjournal

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

type JournalHandler func(*VJournal, componants.Property) error

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
var vjournalHandlers = map[share.PropertyName]JournalHandler{
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

func handleUID(jour *VJournal, prop componants.Property) error {
	return jour.SetUID(prop.Value)
}

func handleDTSTAMP(jour *VJournal, prop componants.Property) error {
	return jour.SetDTSTAMP(share.ParseITIME(prop.Params, prop.Value))
}

func handleDTSTART(jour *VJournal, prop componants.Property) error {
	return jour.SetDTSTART(share.ParseITIME(prop.Params, prop.Value))
}

// ── Optional single ───────────────────────────────────────────────────────────

func handleCreated(jour *VJournal, prop componants.Property) error {
	return jour.SetCreated(share.ParseITIME(prop.Params, prop.Value))
}

func handleLastModified(jour *VJournal, prop componants.Property) error {
	return jour.SetLastModified(share.ParseITIME(prop.Params, prop.Value))
}

func handleRRule(jour *VJournal, prop componants.Property) error {
	r, err := share.ParseRECUR(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid RRULE: %w", err)
	}
	return jour.SetRRule(r)
}

func handleSequence(jour *VJournal, prop componants.Property) error {
	n, err := utils.IntPtr(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid SEQUENCE: %w", err)
	}
	return jour.SetSequence(*n)
}

// ── Pure string fields ────────────────────────────────────────────────────────

func handleClass(jour *VJournal, prop componants.Property) error {
	return jour.SetClass(prop.Value)
}

func handleDescription(jour *VJournal, prop componants.Property) error {
	return jour.SetDescription(prop.Value)
}

func handleOrganizer(jour *VJournal, prop componants.Property) error {
	return jour.SetOrganizer(prop.Value)
}

func handleStatus(jour *VJournal, prop componants.Property) error {
	return jour.SetStatus(prop.Value)
}

func handleSummary(jour *VJournal, prop componants.Property) error {
	return jour.SetSummary(prop.Value)
}

func handleURL(jour *VJournal, prop componants.Property) error {
	return jour.SetURL(prop.Value)
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func handleAttach(jour *VJournal, prop componants.Property) error {
	return jour.AddAttach(prop.Value)
}

func handleAttendee(jour *VJournal, prop componants.Property) error {
	return jour.AddAttendee(prop.Value, prop.Params)
}

func handleCategories(jour *VJournal, prop componants.Property) error {
	for _, cat := range strings.Split(prop.Value, ",") {
		if cat = strings.TrimSpace(cat); cat != "" {
			jour.AddCategory(cat)
		}
	}
	return nil
}

func handleComment(jour *VJournal, prop componants.Property) error {
	return jour.AddComment(prop.Value)
}

func handleContact(jour *VJournal, prop componants.Property) error {
	return jour.AddContact(prop.Value)
}

func handleExDate(jour *VJournal, prop componants.Property) error {
	return jour.AddExDate(share.ParseITIME(prop.Params, prop.Value))
}

func handleRDate(jour *VJournal, prop componants.Property) error {
	return jour.AddRDate(share.ParseITIME(prop.Params, prop.Value))
}

func handleRelated(jour *VJournal, prop componants.Property) error {
	return jour.AddRelated(prop.Value)
}

func handleRequestStatus(jour *VJournal, prop componants.Property) error {
	parts := strings.SplitN(prop.Value, ";", 3)
	rs := share.RequestStatus{Code: parts[0]} // TODO: construct
	if len(parts) > 1 {
		rs.Description = parts[1]
	}
	if len(parts) > 2 {
		rs.Extra = new(parts[2])
	}
	return jour.AddRequestStatus(rs)
}
