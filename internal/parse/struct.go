package parse

import (
	"bufio"
	"io"
)

type Component interface {
	// AddProperty applies a parsed line (like SUMMARY or DTSTART) to struct
	AddProperty(name string, params map[string]string, value string) error

	// AddChild adds a nested component (like a VEVENT inside a VCALENDAR)
	AddChild(child Component) error
}
type IComponent interface {
	ProcessProperty(prop Property) error
	AddChild(child IComponent) error
	Validate() error
}

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

type Engine struct {
	registry map[string]func() IComponent
}
