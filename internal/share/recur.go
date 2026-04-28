package share

import (
	"fmt"
	"strconv"
	"strings"
)

// RFC 5545 §3.3.10

type RECUR struct {
	Freq       string // SECONDLY, MINUTELY, HOURLY, DAILY, WEEKLY, MONTHLY, YEARLY
	Until      *ITIME
	Count      *int
	Interval   int
	BySecond   []int
	ByMinute   []int
	ByHour     []int
	ByDay      []string // MO, TU, WE, TH, FR, SA, SU — may have prefix e.g. -1MO
	ByMonthDay []int
	ByYearDay  []int
	ByWeekNo   []int
	ByMonth    []int
	BySetPos   []int
	WkSt       string // MO, TU... week start
}

// RFC 5545 §3.3.10 — RRULE part names
const (
	recurFREQ       = "FREQ"
	recurUNTIL      = "UNTIL"
	recurCOUNT      = "COUNT"
	recurINTERVAL   = "INTERVAL"
	recurBYSECOND   = "BYSECOND"
	recurBYMINUTE   = "BYMINUTE"
	recurBYHOUR     = "BYHOUR"
	recurBYDAY      = "BYDAY"
	recurBYMONTHDAY = "BYMONTHDAY"
	recurBYYEARDAY  = "BYYEARDAY"
	recurBYWEEKNO   = "BYWEEKNO"
	recurBYMONTH    = "BYMONTH"
	recurBYSETPOS   = "BYSETPOS"
	recurWKST       = "WKST"
)

// RFC 5545 §3.3.10 — FREQ values
const (
	FreqSecondly = "SECONDLY"
	FreqMinutely = "MINUTELY"
	FreqHourly   = "HOURLY"
	FreqDaily    = "DAILY"
	FreqWeekly   = "WEEKLY"
	FreqMonthly  = "MONTHLY"
	FreqYearly   = "YEARLY"
)

// RFC 5545 §3.3.10 — weekday values
const (
	WeekdayMO = "MO"
	WeekdayTU = "TU"
	WeekdayWE = "WE"
	WeekdayTH = "TH"
	WeekdayFR = "FR"
	WeekdaySA = "SA"
	WeekdaySU = "SU"
)

type recurHandler func(*RECUR, string) error

var recurHandlers = map[string]recurHandler{
	recurFREQ:  func(r *RECUR, val string) error { r.Freq = strings.ToUpper(val); return nil },
	recurWKST:  func(r *RECUR, val string) error { r.WkSt = strings.ToUpper(val); return nil },
	recurBYDAY: func(r *RECUR, val string) error { r.ByDay = strings.Split(val, ","); return nil },
	recurUNTIL: func(r *RECUR, val string) error {
		it := ParseITIME(nil, val)
		r.Until = &it
		return nil
	},
	recurCOUNT: func(r *RECUR, val string) error {
		n, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("invalid COUNT: %w", err)
		}
		r.Count = &n
		return nil
	},
	recurINTERVAL: func(r *RECUR, val string) error {
		n, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("invalid INTERVAL: %w", err)
		}
		r.Interval = n
		return nil
	},
	recurBYSECOND:   intListHandler(recurBYSECOND, func(r *RECUR) *[]int { return &r.BySecond }),
	recurBYMINUTE:   intListHandler(recurBYMINUTE, func(r *RECUR) *[]int { return &r.ByMinute }),
	recurBYHOUR:     intListHandler(recurBYHOUR, func(r *RECUR) *[]int { return &r.ByHour }),
	recurBYMONTHDAY: intListHandler(recurBYMONTHDAY, func(r *RECUR) *[]int { return &r.ByMonthDay }),
	recurBYYEARDAY:  intListHandler(recurBYYEARDAY, func(r *RECUR) *[]int { return &r.ByYearDay }),
	recurBYWEEKNO:   intListHandler(recurBYWEEKNO, func(r *RECUR) *[]int { return &r.ByWeekNo }),
	recurBYMONTH:    intListHandler(recurBYMONTH, func(r *RECUR) *[]int { return &r.ByMonth }),
	recurBYSETPOS:   intListHandler(recurBYSETPOS, func(r *RECUR) *[]int { return &r.BySetPos }),
}

// intListHandler generates a handler for any []int field
func intListHandler(name string, field func(*RECUR) *[]int) recurHandler {
	return func(r *RECUR, val string) error {
		list, err := parseIntList(val)
		if err != nil {
			return fmt.Errorf("invalid %s: %w", name, err)
		}
		*field(r) = list
		return nil
	}
}

func ParseRECUR(value string) (RECUR, error) {
	r := RECUR{}

	for part := range strings.SplitSeq(value, ";") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		eqIdx := strings.Index(part, "=")
		if eqIdx < 0 {
			return r, fmt.Errorf("invalid RRULE part: %q", part)
		}

		key := strings.ToUpper(part[:eqIdx])
		val := part[eqIdx+1:]

		if handler, ok := recurHandlers[key]; ok {
			if err := handler(&r, val); err != nil {
				return r, err
			}
		}
		// unknown keys ignored per RFC 5545
	}

	if r.Freq == "" {
		return r, fmt.Errorf("RRULE missing required %s", recurFREQ)
	}

	return r, nil
}

func (r *RECUR) FormatRECUR() string {
	type part struct {
		key string
		val string
		ok  bool
	}

	parts := []part{
		{recurFREQ, r.Freq, r.Freq != ""},
		{recurUNTIL, func() string {
			if r.Until != nil {
				return FormatITIME(*r.Until)
			}
			return ""
		}(), r.Until != nil},
		{recurCOUNT, func() string {
			if r.Count != nil {
				return fmt.Sprintf("%d", *r.Count)
			}
			return ""
		}(), r.Count != nil},
		{recurINTERVAL, fmt.Sprintf("%d", r.Interval), r.Interval > 0},
		{recurBYDAY, strings.Join(r.ByDay, ","), len(r.ByDay) > 0},
		{recurBYMONTHDAY, formatIntList(r.ByMonthDay), len(r.ByMonthDay) > 0},
		{recurBYYEARDAY, formatIntList(r.ByYearDay), len(r.ByYearDay) > 0},
		{recurBYWEEKNO, formatIntList(r.ByWeekNo), len(r.ByWeekNo) > 0},
		{recurBYMONTH, formatIntList(r.ByMonth), len(r.ByMonth) > 0},
		{recurBYHOUR, formatIntList(r.ByHour), len(r.ByHour) > 0},
		{recurBYMINUTE, formatIntList(r.ByMinute), len(r.ByMinute) > 0},
		{recurBYSECOND, formatIntList(r.BySecond), len(r.BySecond) > 0},
		{recurBYSETPOS, formatIntList(r.BySetPos), len(r.BySetPos) > 0},
		{recurWKST, r.WkSt, r.WkSt != ""},
	}

	var result []string
	for _, p := range parts {
		if p.ok {
			result = append(result, p.key+"="+p.val)
		}
	}
	return strings.Join(result, ";")
}

func parseIntList(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		n, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return result, nil
}

func formatIntList(vals []int) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = strconv.Itoa(v)
	}
	return strings.Join(parts, ",")
}
