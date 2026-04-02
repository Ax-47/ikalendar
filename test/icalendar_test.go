package test

import (
	"strings"
	"testing"
	"time"

	"github.com/minoplhy/ikalendar"
)

func TestParseCalendar(t *testing.T) {
	input := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//test//EN
BEGIN:VEVENT
UID:uid1@example.com
DTSTAMP:19970701T100000Z
SUMMARY:Bastille Day Party
GEO:48.85299;2.36885
END:VEVENT
END:VCALENDAR`

	cal, err := ikalendar.ParseCalendar(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	if len(cal.VEVENT) != 1 {
		t.Fatalf("expected 1 event, got %d", len(cal.VEVENT))
	}

	ev := cal.VEVENT[0]

	if ev.UID != "uid1@example.com" {
		t.Fatalf("unexpected UID: %s", ev.UID)
	}
	if ev.SUMMARY == nil || *ev.SUMMARY != "Bastille Day Party" {
		t.Fatalf("unexpected SUMMARY: %v", ev.SUMMARY)
	}
	if ev.GEO == nil {
		t.Fatal("GEO should not be nil")
	}
}

func TestParseMultipleCases(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		wantUID string
	}{
		{
			name: "basic",
			input: `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//test//EN
BEGIN:VEVENT
UID:abc
DTSTAMP:20240101T000000Z
END:VEVENT
END:VCALENDAR`,
			wantUID: "abc",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cal, err := ikalendar.ParseCalendar(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}
			if cal.VEVENT[0].UID != tc.wantUID {
				t.Fatalf("got %s, want %s", cal.VEVENT[0].UID, tc.wantUID)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	input := `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//test//EN
BEGIN:VEVENT
UID:1
DTSTAMP:20240101T000000Z
SUMMARY:Test Event
END:VEVENT
END:VCALENDAR`

	cal, err := ikalendar.ParseCalendar(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	out, err := ikalendar.Marshal(cal)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(out), "SUMMARY:Test Event") {
		t.Fatal("round-trip lost SUMMARY")
	}
}

func TestMarshalCalendar(t *testing.T) {
	ev, err := ikalendar.NewEvent("uid-test-"+time.Now().Format("20060102T150405Z"),
		ikalendar.WithSummary("Follow-up Meeting"),
		ikalendar.WithLocation("My Home"),
		ikalendar.WithStatus("TENTATIVE"),
		ikalendar.WithDtStart(time.Now()),
	)
	if err != nil {
		t.Fatal(err)
	}

	cal, err := ikalendar.NewCalendar(
		ikalendar.WithEvent(ev),
	)
	if err != nil {
		t.Fatal(err)
	}

	out, err := ikalendar.Marshal(cal)
	if err != nil {
		t.Fatal(err)
	}

	s := string(out)
	if !strings.Contains(s, "SUMMARY:Follow-up Meeting") {
		t.Fatal("missing SUMMARY")
	}
	if !strings.Contains(s, "LOCATION:My Home") {
		t.Fatal("missing LOCATION")
	}
	if !strings.Contains(s, "STATUS:TENTATIVE") {
		t.Fatal("missing STATUS")
	}
}

func TestMarshalWithAlarm(t *testing.T) {
	alarm, err := ikalendar.NewAlarm(
		ikalendar.WithAction("DISPLAY"),
		ikalendar.WithTriggerBefore(0, 15),
		ikalendar.WithAlarmDescription("Reminder"),
	)
	if err != nil {
		t.Fatal(err)
	}

	ev, err := ikalendar.NewEvent("uid-alarm-test",
		ikalendar.WithSummary("Meeting with Alarm"),
		ikalendar.WithDtStart(time.Now()),
		ikalendar.WithAlarm(alarm),
	)
	if err != nil {
		t.Fatal(err)
	}

	cal, err := ikalendar.NewCalendar(
		ikalendar.WithEvent(ev),
	)
	if err != nil {
		t.Fatal(err)
	}

	out, err := ikalendar.Marshal(cal)
	if err != nil {
		t.Fatal(err)
	}

	s := string(out)
	if !strings.Contains(s, "BEGIN:VALARM") {
		t.Fatal("missing VALARM")
	}
	if !strings.Contains(s, "ACTION:DISPLAY") {
		t.Fatal("missing ACTION")
	}
	if !strings.Contains(s, "TRIGGER:-PT15M") {
		t.Fatal("missing TRIGGER")
	}
}
