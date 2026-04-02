package main

//
// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"strings"
// 	"time"
//
// 	"github.com/minoplhy/ikalendar"
// 	"github.com/minoplhy/ikalendar/internal/icalendar"
// 	parsehelper "github.com/minoplhy/ikalendar/internal/parse_helper"
// )
//
// func main() {
// 	test := `BEGIN:VCALENDAR
// VERSION:2.0
// PRODID:-//hacksw/handcal//NONSGML v1.0//EN
// BEGIN:VEVENT
// UID:uid1@example.com
// ORGANIZER;CN=John Doe:MAILTO:john.doe@example.com
// DTSTAMP:19970701T100000Z
// DTSTART:19970714T170000Z
// DTEND:19970715T040000Z
// SUMMARY:Bastille Day Party
// GEO:48.85299;2.36885
// END:VEVENT
// END:VCALENDAR`
//
// 	buf := new(strings.Builder)
//
// 	r := strings.NewReader(test)
//
// 	_, err := io.Copy(buf, r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// ParseCalendar expects io.Reader, pass a reader
// 	cal, err := ikalendar.ParseCalendar(strings.NewReader(buf.String()))
// 	if err != nil {
// 		fmt.Println("error:", err)
// 		return
// 	}
//
// 	fmt.Printf("%#v\n", *cal.VEVENT[0].GEO)
//
// 	newEvent := ikalendar.NewEvent().
// 		SetSummary("Follow-up Meeting").
// 		SetLocation("My Home").SetStatus("TENTATIVE").SetSummary("Follow-up Meeting").SetSummary("67")
// 	newEvent.SetSummary("67")
// 	cal.AddEvent(newEvent)
//
// 	output, err := ikalendar.Marshal(cal)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Println(string(output))
//
// 	s := ikalendar.New()
// 	s.AddEvent(newEvent)
//
// 	newCalendar := &icalendar.VCalendar{
// 		VERSION: "2.0",
// 		PRODID:  "-//ikarl//My Awesome App//EN",
// 		METHOD:  parsehelper.StrPtr("PUBLISH"),
//
// 		VEVENT: []icalendar.VEvent{
// 			{
// 				UID:     "event-001@my-app.com",
// 				SUMMARY: parsehelper.StrPtr("Project Launch"),
// 				STATUS:  parsehelper.StrPtr("CONFIRMED"),
// 				DTSTAMP: icalendar.NewITIMEUTC(time.Now()),
// 			},
// 			{
// 				UID:     "event-002@my-app.com",
// 				SUMMARY: parsehelper.StrPtr("Team Dinner"),
// 				DTSTAMP: icalendar.NewITIMEUTC(time.Now()),
// 			},
// 		},
// 	}
// 	newCalendar.AddEvent(newEvent)
//
// 	outputCalendar, _ := ikalendar.Marshal(newCalendar)
//
// 	fmt.Println(string(outputCalendar))
// }
