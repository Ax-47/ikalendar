package parse

import "errors"

func (e *Engine) Register(name string, registrant func() IComponent) {
	e.registry[name] = registrant
}

func NewEngine() *Engine {
	return &Engine{
		registry: make(map[string]func() IComponent),
	}
}

var ErrNoRootComponent = errors.New("parse: no root component found in stream")

// Run takes input Parser and builds the Component tree
func (e *Engine) Run(parser *Parser) (IComponent, error) {
	var root IComponent
	var stack []IComponent
	var ignoreDepth int // Tracks nested levels of unsupported components

	for {
		prop, hasMore, err := parser.Next()
		if err != nil {
			return nil, err
		}
		if !hasMore {
			break // EOF
		}

		if prop.Name == "" {
			continue // Skip malformed/empty lines
		}

		// ignore anything until END tag for unsupported method
		if ignoreDepth > 0 {
			switch prop.Name {
			case "BEGIN":
				ignoreDepth++ // Handle nested unsupported blocks gracefully
			case "END":
				ignoreDepth--
			}
			continue
		}

		switch prop.Name {
		case "BEGIN":
			factory, exists := e.registry[prop.Value]
			if !exists {
				// invalid or We don't support this component
				// Tell the parser to ignore everything until the END tag
				ignoreDepth = 1
				continue
			}

			stack = append(stack, factory())

		case "END":
			if len(stack) == 0 {
				continue // Ignore mismatched END tags
			}

			finished := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // Pop from stack

			if err := finished.Validate(); err != nil {
				return nil, err
			}

			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				if err := parent.AddChild(finished); err != nil {
					return nil, err
				}
			} else {
				// root node
				root = finished
			}

		default:
			if len(stack) > 0 {
				activeComponent := stack[len(stack)-1]
				if err := activeComponent.ProcessProperty(prop); err != nil {
					return nil, err
				}
			}
		}
	}

	if root == nil {
		return nil, ErrNoRootComponent
	}
	return root, nil
}
