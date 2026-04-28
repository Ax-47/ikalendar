package vtodo

import (
	"time"

	"github.com/minoplhy/ikalendar/internal/share"
)

type VTodoOption func(*VTodo) error

// NewVTodo creates a new VTodo
//
//	vt, err := ikalendar.NewVTodo("uid-123",
//	    ikalendar.WithSummary("Team Standup"),
//	    ikalendar.WithDtStart(time.Now()),
//	)

func NewVTodo(uid string, opts ...VTodoOption) (*VTodo, error) {
	vt := &VTodo{}
	if err := vt.SetUID(uid); err != nil {
		return nil, err
	}
	for _, opt := range opts {
		if err := opt(vt); err != nil {
			return nil, err
		}
	}
	return vt, nil
}

// ── Single-value ──────────────────────────────────────────────────────────────

func WithSummary(s string) VTodoOption {
	return func(vt *VTodo) error { return vt.SetSummary(s) }
}

func WithDescription(s string) VTodoOption {
	return func(vt *VTodo) error { return vt.SetDescription(s) }
}

func WithStatus(s string) VTodoOption {
	return func(vt *VTodo) error { return vt.SetStatus(s) }
}

func WithClass(s string) VTodoOption {
	return func(vt *VTodo) error { return vt.SetClass(s) }
}

func WithOrganizer(s string) VTodoOption {
	return func(vt *VTodo) error { return vt.SetOrganizer(s) }
}

func WithURL(s string) VTodoOption {
	return func(vt *VTodo) error { return vt.SetURL(s) }
}

func WithSequence(n int) VTodoOption {
	return func(vt *VTodo) error { return vt.SetSequence(n) }
}

func WithRRule(r share.RECUR) VTodoOption {
	return func(vt *VTodo) error { return vt.SetRRule(r) }
}

// ── Time ──────────────────────────────────────────────────────────────────────

func WithDtStart(t time.Time) VTodoOption {
	return func(vt *VTodo) error { return vt.SetDTSTART(share.NewITIME(t.UTC())) }
}

func WithCreated(t time.Time) VTodoOption {
	return func(vt *VTodo) error { return vt.SetCreated(share.NewITIME(t.UTC())) }
}

func WithLastModified(t time.Time) VTodoOption {
	return func(vt *VTodo) error { return vt.SetLastModified(share.NewITIME(t.UTC())) }
}

// ── Multi-value ───────────────────────────────────────────────────────────────

func WithCategory(categories ...string) VTodoOption {
	return func(vt *VTodo) error {
		for _, c := range categories {
			vt.AddCategory(c)
		}
		return nil
	}
}

func WithAttendee(address string, params map[string]string) VTodoOption {
	return func(vt *VTodo) error { return vt.AddAttendee(address, params) }
}

func WithAttach(uri string) VTodoOption {
	return func(vt *VTodo) error { return vt.AddAttach(uri) }
}

func WithComment(comments ...string) VTodoOption {
	return func(vt *VTodo) error {
		for _, c := range comments {
			if err := vt.AddComment(c); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithContact(contacts ...string) VTodoOption {
	return func(vt *VTodo) error {
		for _, c := range contacts {
			if err := vt.AddContact(c); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithExDate(times ...time.Time) VTodoOption {
	return func(vt *VTodo) error {
		for _, t := range times {
			if err := vt.AddExDate(share.NewITIME(t.UTC())); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithRDate(times ...time.Time) VTodoOption {
	return func(vt *VTodo) error {
		for _, t := range times {
			if err := vt.AddRDate(share.NewITIME(t.UTC())); err != nil {
				return err
			}
		}
		return nil
	}
}

func WithRelated(uid string) VTodoOption {
	return func(vt *VTodo) error { return vt.AddRelated(uid) }
}

func WithRequestStatus(code, description string, extra *string) VTodoOption {
	return func(vt *VTodo) error {
		return vt.AddRequestStatus(share.RequestStatus{ // TODO: contruct
			Code:        code,
			Description: description,
			Extra:       extra,
		})
	}
}
