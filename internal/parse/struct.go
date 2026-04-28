package parse

import (
	"bufio"
	"io"

	"github.com/minoplhy/ikalendar/internal/componants"
)

func NewParser(r io.Reader) *Parser {
	return &Parser{scanner: bufio.NewScanner(r)}
}

type Parser struct {
	scanner *bufio.Scanner
	peek    string
}
type (
	Engine struct {
		registry componants.RegistryMap
	}
)
