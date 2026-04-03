package parse

import (
	"bufio"
	"io"

	"github.com/minoplhy/ikalendar/internal/componants"
)

type (
	Validator interface{ Validate() error }
)

func NewParser(r io.Reader) *Parser {
	return &Parser{scanner: bufio.NewScanner(r)}
}

type Parser struct {
	scanner *bufio.Scanner
	peek    string
}
type (
	registryMap map[string]componants.ComponentFactory
	Engine      struct {
		registry registryMap
	}
)
