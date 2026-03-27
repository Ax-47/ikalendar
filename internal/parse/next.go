package parse

import (
	"strings"
)

// parsePropertyLine does the string splitting (Name:Value;Param=X)
func (p *Parser) parsePropertyLine(line string) Property {
	leftPart, rawValue, found := strings.Cut(line, ":")
	if !found {
		return Property{}
	}

	namePart, paramStr, hasParams := strings.Cut(leftPart, ";")
	prop := Property{
		Name:   strings.ToUpper(strings.TrimSpace(namePart)),
		Value:  Unescape(rawValue), // Your unescape helper
		Params: make(map[string]string),
	}

	if hasParams {
		for _, param := range strings.Split(paramStr, ";") {
			k, v, _ := strings.Cut(param, "=")
			prop.Params[strings.ToUpper(k)] = strings.Trim(v, `"`)
		}
	}
	return prop
}

// Next reads the input and yields the next complete
// Returns false when the input is fully consumed.
func (p *Parser) Next() (Property, bool, error) {
	var currentLine strings.Builder

	if p.peek != "" {
		currentLine.WriteString(p.peek)
		p.peek = ""
	} else if p.scanner.Scan() {
		currentLine.WriteString(p.scanner.Text())
	} else {
		return Property{}, false, p.scanner.Err() // EOF
	}

	for p.scanner.Scan() {
		line := p.scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == ' ' || line[0] == '\t' {
			currentLine.WriteString(line[1:])
		} else {
			p.peek = line
			break
		}
	}

	prop := p.parsePropertyLine(currentLine.String())
	return prop, true, nil
}
