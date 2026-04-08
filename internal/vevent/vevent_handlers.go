package vevent

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/componants"
	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/utils"
)

type EventHandler func(*VEvent, componants.Property) error

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
var veventHandlers = map[share.PropertyName]EventHandler{
	PropUID:           handleUID,
	PropDTSTAMP:       handleDTSTAMP,
	PropDTSTART:       handleDTSTART,
	PropDTEND:         handleDTEND,
	PropDURATION:      handleDuration,
	PropCLASS:         handleClass,
	PropCREATED:       handleCreated,
	PropDESCRIPTION:   handleDescription,
	PropGEO:           handleGeo,
	PropLASTMODIFIED:  handleLastModified,
	PropLOCATION:      handleLocation,
	PropORGANIZER:     handleOrganizer,
	PropPRIORITY:      handlePriority,
	PropSEQUENCE:      handleSequence,
	PropSTATUS:        handleStatus,
	PropSUMMARY:       handleSummary,
	PropTRANSP:        handleTransp,
	PropURL:           handleURL,
	PropRRULE:         handleRRule,
	PropATTACH:        handleAttach,
	PropATTENDEE:      handleAttendee,
	PropCATEGORIES:    handleCategories,
	PropCOMMENT:       handleComment,
	PropCONTACT:       handleContact,
	PropEXDATE:        handleExDate,
	PropREQUESTSTATUS: handleRequestStatus,
	PropRELATED:       handleRelated,
	PropRDATE:         handleRDate,
	PropRESOURCES:     handleResources,
}

// ── Required ──────────────────────────────────────────────────────────────────

func handleUID(ev *VEvent, prop componants.Property) error {
	return ev.SetUID(prop.Value)
}

func handleDTSTAMP(ev *VEvent, prop componants.Property) error {
	return ev.SetDTSTAMP(share.ParseITIME(prop.Params, prop.Value))
}

func handleDTSTART(ev *VEvent, prop componants.Property) error {
	return ev.SetDTSTART(share.ParseITIME(prop.Params, prop.Value))
}

// ── Optional single ───────────────────────────────────────────────────────────

func handleDTEND(ev *VEvent, prop componants.Property) error {
	return ev.SetDTEND(share.ParseITIME(prop.Params, prop.Value))
}

func handleDuration(ev *VEvent, prop componants.Property) error {
	d, err := share.ParseDURATION(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid DURATION: %w", err)
	}
	return ev.SetDuration(d)
}

func handleCreated(ev *VEvent, prop componants.Property) error {
	return ev.SetCreated(share.ParseITIME(prop.Params, prop.Value))
}

func handleLastModified(ev *VEvent, prop componants.Property) error {
	return ev.SetLastModified(share.ParseITIME(prop.Params, prop.Value))
}

func handleRRule(ev *VEvent, prop componants.Property) error {
	r, err := share.ParseRECUR(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid RRULE: %w", err)
	}
	return ev.SetRRule(r)
}

func handlePriority(ev *VEvent, prop componants.Property) error {
	n, err := utils.IntPtr(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid PRIORITY: %w", err)
	}
	return ev.SetPriority(*n)
}

func handleSequence(ev *VEvent, prop componants.Property) error {
	n, err := utils.IntPtr(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid SEQUENCE: %w", err)
	}
	return ev.SetSequence(*n)
}

// ── Pure string fields ────────────────────────────────────────────────────────

func handleClass(ev *VEvent, prop componants.Property) error {
	return ev.SetClass(prop.Value)
}

func handleDescription(ev *VEvent, prop componants.Property) error {
	return ev.SetDescription(prop.Value)
}

func handleGeo(ev *VEvent, prop componants.Property) error {
	return ev.SetGeo(prop.Value)
}

func handleLocation(ev *VEvent, prop componants.Property) error {
	return ev.SetLocation(prop.Value)
}

func handleOrganizer(ev *VEvent, prop componants.Property) error {
	return ev.SetOrganizer(prop.Value)
}

func handleStatus(ev *VEvent, prop componants.Property) error {
	return ev.SetStatus(prop.Value)
}

func handleSummary(ev *VEvent, prop componants.Property) error {
	return ev.SetSummary(prop.Value)
}

func handleTransp(ev *VEvent, prop componants.Property) error {
	return ev.SetTransp(prop.Value)
}

func handleURL(ev *VEvent, prop componants.Property) error {
	return ev.SetURL(prop.Value)
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func handleAttach(ev *VEvent, prop componants.Property) error {
	return ev.AddAttach(prop.Value)
}

func handleAttendee(ev *VEvent, prop componants.Property) error {
	return ev.AddAttendee(prop.Value, prop.Params)
}

func handleCategories(ev *VEvent, prop componants.Property) error {
	for _, cat := range strings.Split(prop.Value, ",") {
		if cat = strings.TrimSpace(cat); cat != "" {
			ev.AddCategory(cat)
		}
	}
	return nil
}

func handleComment(ev *VEvent, prop componants.Property) error {
	return ev.AddComment(prop.Value)
}

func handleContact(ev *VEvent, prop componants.Property) error {
	return ev.AddContact(prop.Value)
}

func handleExDate(ev *VEvent, prop componants.Property) error {
	return ev.AddExDate(share.ParseITIME(prop.Params, prop.Value))
}

func handleRDate(ev *VEvent, prop componants.Property) error {
	return ev.AddRDate(share.ParseITIME(prop.Params, prop.Value))
}

func handleRelated(ev *VEvent, prop componants.Property) error {
	return ev.AddRelated(prop.Value)
}

func handleResources(ev *VEvent, prop componants.Property) error {
	return ev.AddResource(prop.Value)
}

func handleRequestStatus(ev *VEvent, prop componants.Property) error {
	parts := strings.SplitN(prop.Value, ";", 3)
	rs := share.RequestStatus{Code: parts[0]} // TODO: construct
	if len(parts) > 1 {
		rs.Description = parts[1]
	}
	if len(parts) > 2 {
		rs.Extra = utils.Ptr(parts[2])
	}
	return ev.AddRequestStatus(rs)
}
