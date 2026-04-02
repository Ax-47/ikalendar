package parse

import (
	"bufio"
	"io"
)

type Component interface {
	// AddProperty applies a parsed line (like SUMMARY or DTSTART) to struct
	// AddProperty(prop Property) error
	ProcessProperty(prop Property) error
	// AddChild adds a nested component (like a VEVENT inside a VCALENDAR)
	AddChild(child Component) error

	Validate() error
}

type (
	ComponentFactory func() Component
)

type (
	Validator interface{ Validate() error }
)

func NewParser(r io.Reader) *Parser {
	return &Parser{scanner: bufio.NewScanner(r)}
}

type Property struct {
	Name   string
	Params map[string]string
	Value  string
}

type Parser struct {
	scanner *bufio.Scanner
	peek    string
}
type (
	registryMap map[string]ComponentFactory
	Engine      struct {
		registry registryMap
	}
)
