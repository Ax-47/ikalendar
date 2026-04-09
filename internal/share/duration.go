package share

import (
	"fmt"
	"strconv"
	"strings"
)

type DURATION struct {
	Negative bool

	Weeks   int
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func (d *DURATION) FormatDURATION() string {
	var sb strings.Builder

	if d.Negative {
		sb.WriteByte('-')
	}
	sb.WriteByte('P')

	// Weeks are standalone — cannot combine with D/H/M/S per RFC 5545
	if d.Weeks > 0 {
		fmt.Fprintf(&sb, "%dW", d.Weeks)
		return sb.String()
	}

	if d.Days > 0 {
		fmt.Fprintf(&sb, "%dD", d.Days)
	}

	// Time part
	if d.Hours > 0 || d.Minutes > 0 || d.Seconds > 0 {
		sb.WriteByte('T')
		if d.Hours > 0 {
			fmt.Fprintf(&sb, "%dH", d.Hours)
		}
		if d.Minutes > 0 {
			fmt.Fprintf(&sb, "%dM", d.Minutes)
		}
		if d.Seconds > 0 {
			fmt.Fprintf(&sb, "%dS", d.Seconds)
		}
	}

	return sb.String()
}

func ParseDURATION(value string) (DURATION, error) {
	d := DURATION{}
	value = strings.TrimSpace(value)

	if strings.HasPrefix(value, "-") {
		d.Negative = true
		value = value[1:]
	}

	if !strings.HasPrefix(value, "P") {
		return d, fmt.Errorf("invalid DURATION: missing P designator")
	}
	value = value[1:] // strip P

	// Week
	if strings.HasSuffix(value, "W") {
		n, err := strconv.Atoi(strings.TrimSuffix(value, "W"))
		if err != nil {
			return d, fmt.Errorf("invalid DURATION week: %w", err)
		}
		d.Weeks = n
		return d, nil
	}

	// Split date/time on T
	parts := strings.SplitN(value, "T", 2)

	// Date part — days
	if parts[0] != "" {
		if n, err := parseInt(parts[0], "D"); err == nil {
			d.Days = n
		}
	}

	// Time part
	if len(parts) == 2 {
		t := parts[1]
		if h, err := parseInt(t, "H"); err == nil {
			d.Hours = h
			t = t[strings.Index(t, "H")+1:]
		}
		if m, err := parseInt(t, "M"); err == nil {
			d.Minutes = m
			t = t[strings.Index(t, "M")+1:]
		}
		if s, err := parseInt(t, "S"); err == nil {
			d.Seconds = s
		}
	}

	return d, nil
}

func parseInt(s, designator string) (int, error) {
	idx := strings.Index(s, designator)
	if idx < 0 {
		return 0, fmt.Errorf("designator %s not found", designator)
	}
	return strconv.Atoi(s[:idx])
}
