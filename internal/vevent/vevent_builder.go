package vevent

import (
	"time"

	"github.com/minoplhy/ikalendar/internal/share"
	"github.com/minoplhy/ikalendar/internal/valarm"
)

type VeventOption func(*VEvent) error

// NewEvent creates a new VEvent
//
//	ev, err := ikalendar.NewEvent("uid-123",
//	    ikalendar.WithSummary("Team Standup"),
//	    ikalendar.WithDtStart(time.Now()),
//	)

func NewEvent(uid string, opts ...VeventOption) (*VEvent, error) {
	ev := &VEvent{}
	if err := ev.SetUID(uid); err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err := opt(ev); err != nil {
			return nil, err
		}
	}
	return ev, nil
}

// ── Single-value ──────────────────────────────────────────────────────────────

func WithSummary(s string) VeventOption {
	return func(ev *VEvent) error { return ev.SetSummary(s) }
}

func WithDescription(s string) VeventOption {
	return func(ev *VEvent) error { return ev.SetDescription(s) }
}

func WithLocation(s string) VeventOption {
	return func(ev *VEvent) error { return ev.SetLocation(s) }
}

func WithStatus(s string) VeventOption {
	return func(ev *VEvent) error { return ev.SetStatus(s) }
}

func WithClass(s string) VeventOption {
	return func(ev *VEvent) error { return ev.SetClass(s) }
}

func WithOrganizer(s string) VeventOption {
	return func(ev *VEvent) error { return ev.SetOrganizer(s) }
}

func WithTransp(s string) VeventOption {
	return func(ev *VEvent) error { return ev.SetTransp(s) }
}

func WithURL(s string) VeventOption {
	return func(ev *VEvent) error { return ev.SetURL(s) }
}

func WithPriority(p int) VeventOption {
	return func(ev *VEvent) error { return ev.SetPriority(p) }
}

func WithSequence(n int) VeventOption {
	return func(ev *VEvent) error { return ev.SetSequence(n) }
}

func WithRRule(r share.RECUR) VeventOption {
	return func(ev *VEvent) error { return ev.SetRRule(r) }
}

// ── Time ──────────────────────────────────────────────────────────────────────

func WithDtStamp(t time.Time) VeventOption {
	return func(ev *VEvent) error { return ev.SetDTSTAMP(share.NewITIME(t.UTC())) }
}

func WithDtStart(t time.Time) VeventOption {
	return func(ev *VEvent) error { return ev.SetDTSTART(share.NewITIME(t.UTC())) }
}

func WithDtEnd(t time.Time) VeventOption {
	return func(ev *VEvent) error { return ev.SetDTEND(share.NewITIME(t.UTC())) }
}

func WithCreated(t time.Time) VeventOption {
	return func(ev *VEvent) error { return ev.SetCreated(share.NewITIME(t.UTC())) }
}

func WithLastModified(t time.Time) VeventOption {
	return func(ev *VEvent) error { return ev.SetLastModified(share.NewITIME(t.UTC())) }
}

func WithDuration(d share.DURATION) VeventOption {
	return func(ev *VEvent) error { return ev.SetDuration(d) }
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func WithCategory(categories ...string) VeventOption {
	return func(ev *VEvent) error {
		for _, c := range categories {
			ev.AddCategory(c)
		}
		return nil
	}
}

func WithAttendee(address string, params map[string]string) VeventOption {
	return func(ev *VEvent) error { return ev.AddAttendee(address, params) }
}

func WithAlarm(alarm *valarm.VAlarm) VeventOption {
	return func(ev *VEvent) error { return ev.AddChild(alarm) }
}

func WithAttach(uri string) VeventOption {
	return func(ev *VEvent) error { return ev.AddAttach(uri) }
}

func WithComment(comments ...string) VeventOption {
	return func(ev *VEvent) error {
		for _, c := range comments {
			if err := ev.AddComment(c); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithContact(contacts ...string) VeventOption {
	return func(ev *VEvent) error {
		for _, c := range contacts {
			if err := ev.AddContact(c); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithResources(resources ...string) VeventOption {
	return func(ev *VEvent) error {
		for _, r := range resources {
			if err := ev.AddResource(r); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithExDate(times ...time.Time) VeventOption {
	return func(ev *VEvent) error {
		for _, t := range times {
			if err := ev.AddExDate(share.NewITIME(t.UTC())); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithRDate(times ...time.Time) VeventOption {
	return func(ev *VEvent) error {
		for _, t := range times {
			if err := ev.AddRDate(share.NewITIME(t.UTC())); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithRelated(uid string) VeventOption {
	return func(ev *VEvent) error { return ev.AddRelated(uid) }
}

func WithRequestStatus(code, description string, extra *string) VeventOption {
	return func(ev *VEvent) error {
		return ev.AddRequestStatus(share.RequestStatus{ // TODO: contruct
			Code:        code,
			Description: description,
			Extra:       extra,
		})
	}
}
