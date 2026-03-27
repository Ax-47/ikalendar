package icalendar

type EventClass string

const (
	ClassPublic       EventClass = "PUBLIC"
	ClassPrivate      EventClass = "PRIVATE"
	ClassConfidential EventClass = "CONFIDENTIAL"
)
