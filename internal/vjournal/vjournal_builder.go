package vjournal

import (
	"time"

	"github.com/minoplhy/ikalendar/internal/share"
)

type VJournalOption func(*VJournal) error

// NewJournal creates a new NewJournal
//
//	jour, err := ikalendar.NewJournal("uid-123",
//	    ikalendar.WithSummary("Team Standup"),
//	    ikalendar.WithDtStart(time.Now()),
//	)

func NewJournal(uid string, opts ...VJournalOption) (*VJournal, error) {
	jour := &VJournal{}
	if err := jour.SetUID(uid); err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err := opt(jour); err != nil {
			return nil, err
		}
	}
	return jour, nil
}

// ── Single-value ──────────────────────────────────────────────────────────────

func WithSummary(s string) VJournalOption {
	return func(jour *VJournal) error { return jour.SetSummary(s) }
}

func WithDescription(s string) VJournalOption {
	return func(jour *VJournal) error { return jour.SetDescription(s) }
}

func WithStatus(s string) VJournalOption {
	return func(jour *VJournal) error { return jour.SetStatus(s) }
}

func WithClass(s string) VJournalOption {
	return func(jour *VJournal) error { return jour.SetClass(s) }
}

func WithOrganizer(s string) VJournalOption {
	return func(jour *VJournal) error { return jour.SetOrganizer(s) }
}

func WithURL(s string) VJournalOption {
	return func(jour *VJournal) error { return jour.SetURL(s) }
}

func WithSequence(n int) VJournalOption {
	return func(jour *VJournal) error { return jour.SetSequence(n) }
}

func WithRRule(r share.RECUR) VJournalOption {
	return func(jour *VJournal) error { return jour.SetRRule(r) }
}

// ── Time ──────────────────────────────────────────────────────────────────────

func WithDtStart(t time.Time) VJournalOption {
	return func(jour *VJournal) error { return jour.SetDTSTART(share.NewITIME(t.UTC())) }
}

func WithCreated(t time.Time) VJournalOption {
	return func(jour *VJournal) error { return jour.SetCreated(share.NewITIME(t.UTC())) }
}

func WithLastModified(t time.Time) VJournalOption {
	return func(jour *VJournal) error { return jour.SetLastModified(share.NewITIME(t.UTC())) }
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func WithCategory(categories ...string) VJournalOption {
	return func(jour *VJournal) error {
		for _, c := range categories {
			jour.AddCategory(c)
		}
		return nil
	}
}

func WithAttendee(address string, params map[string]string) VJournalOption {
	return func(jour *VJournal) error { return jour.AddAttendee(address, params) }
}

func WithAttach(uri string) VJournalOption {
	return func(jour *VJournal) error { return jour.AddAttach(uri) }
}

func WithComment(comments ...string) VJournalOption {
	return func(jour *VJournal) error {
		for _, c := range comments {
			if err := jour.AddComment(c); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithContact(contacts ...string) VJournalOption {
	return func(jour *VJournal) error {
		for _, c := range contacts {
			if err := jour.AddContact(c); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithExDate(times ...time.Time) VJournalOption {
	return func(jour *VJournal) error {
		for _, t := range times {
			if err := jour.AddExDate(share.NewITIME(t.UTC())); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithRDate(times ...time.Time) VJournalOption {
	return func(jour *VJournal) error {
		for _, t := range times {
			if err := jour.AddRDate(share.NewITIME(t.UTC())); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithRelated(uid string) VJournalOption {
	return func(jour *VJournal) error { return jour.AddRelated(uid) }
}

func WithRequestStatus(code, description string, extra *string) VJournalOption {
	return func(jour *VJournal) error {
		return jour.AddRequestStatus(share.RequestStatus{ // TODO: contruct
			Code:        code,
			Description: description,
			Extra:       extra,
		})
	}
}
