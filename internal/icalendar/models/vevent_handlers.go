package models

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/parse"
	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
)

type EventHandler func(*VEvent, parse.Property) error

var veventHandlers = map[PropertyName]EventHandler{
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

func handleUID(ev *VEvent, prop parse.Property) error {
	return ev.SetUID(prop.Value)
}

func handleDTSTAMP(ev *VEvent, prop parse.Property) error {
	return ev.SetDTSTAMP(ParseITIME(prop.Params, prop.Value))
}

func handleDTSTART(ev *VEvent, prop parse.Property) error {
	return ev.SetDTSTART(ParseITIME(prop.Params, prop.Value))
}

// ── Optional single ───────────────────────────────────────────────────────────

func handleDTEND(ev *VEvent, prop parse.Property) error {
	return ev.SetDTEND(ParseITIME(prop.Params, prop.Value))
}

func handleDuration(ev *VEvent, prop parse.Property) error {
	d, err := ParseDURATION(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid DURATION: %w", err)
	}
	return ev.SetDuration(d)
}

func handleCreated(ev *VEvent, prop parse.Property) error {
	return ev.SetCreated(ParseITIME(prop.Params, prop.Value))
}

func handleLastModified(ev *VEvent, prop parse.Property) error {
	return ev.SetLastModified(ParseITIME(prop.Params, prop.Value))
}

func handleRRule(ev *VEvent, prop parse.Property) error {
	r, err := ParseRECUR(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid RRULE: %w", err)
	}
	return ev.SetRRule(r)
}

func handlePriority(ev *VEvent, prop parse.Property) error {
	n, err := parsehelper.IntPtr(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid PRIORITY: %w", err)
	}
	return ev.SetPriority(*n)
}

func handleSequence(ev *VEvent, prop parse.Property) error {
	n, err := parsehelper.IntPtr(prop.Value)
	if err != nil {
		return fmt.Errorf("invalid SEQUENCE: %w", err)
	}
	return ev.SetSequence(*n)
}

// ── Pure string fields ────────────────────────────────────────────────────────

func handleClass(ev *VEvent, prop parse.Property) error {
	return ev.SetClass(prop.Value)
}

func handleDescription(ev *VEvent, prop parse.Property) error {
	return ev.SetDescription(prop.Value)
}

func handleGeo(ev *VEvent, prop parse.Property) error {
	return ev.SetGeo(prop.Value)
}

func handleLocation(ev *VEvent, prop parse.Property) error {
	return ev.SetLocation(prop.Value)
}

func handleOrganizer(ev *VEvent, prop parse.Property) error {
	return ev.SetOrganizer(prop.Value)
}

func handleStatus(ev *VEvent, prop parse.Property) error {
	return ev.SetStatus(prop.Value)
}

func handleSummary(ev *VEvent, prop parse.Property) error {
	return ev.SetSummary(prop.Value)
}

func handleTransp(ev *VEvent, prop parse.Property) error {
	return ev.SetTransp(prop.Value)
}

func handleURL(ev *VEvent, prop parse.Property) error {
	return ev.SetURL(prop.Value)
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func handleAttach(ev *VEvent, prop parse.Property) error {
	return ev.AddAttach(prop.Value)
}

func handleAttendee(ev *VEvent, prop parse.Property) error {
	return ev.AddAttendee(prop.Value, prop.Params)
}

func handleCategories(ev *VEvent, prop parse.Property) error {
	for _, cat := range strings.Split(prop.Value, ",") {
		if cat = strings.TrimSpace(cat); cat != "" {
			ev.AddCategory(cat)
		}
	}
	return nil
}

func handleComment(ev *VEvent, prop parse.Property) error {
	return ev.AddComment(prop.Value)
}

func handleContact(ev *VEvent, prop parse.Property) error {
	return ev.AddContact(prop.Value)
}

func handleExDate(ev *VEvent, prop parse.Property) error {
	return ev.AddExDate(ParseITIME(prop.Params, prop.Value))
}

func handleRDate(ev *VEvent, prop parse.Property) error {
	return ev.AddRDate(ParseITIME(prop.Params, prop.Value))
}

func handleRelated(ev *VEvent, prop parse.Property) error {
	return ev.AddRelated(prop.Value)
}

func handleResources(ev *VEvent, prop parse.Property) error {
	return ev.AddResource(prop.Value)
}

func handleRequestStatus(ev *VEvent, prop parse.Property) error {
	parts := strings.SplitN(prop.Value, ";", 3)
	rs := RequestStatus{Code: parts[0]}
	if len(parts) > 1 {
		rs.Description = parts[1]
	}
	if len(parts) > 2 {
		rs.Extra = parsehelper.Ptr(parts[2])
	}
	return ev.AddRequestStatus(rs)
}
