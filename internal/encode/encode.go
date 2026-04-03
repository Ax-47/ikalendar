package encode

import (
	"fmt"
	"strings"

	"github.com/minoplhy/ikalendar/internal/componants"
)

func Marshal(componant componants.CalendarLike) ([]byte, error) {
	if componant == nil {
		return nil, fmt.Errorf("calendar cannot be nil")
	}

	var sb strings.Builder
	ctx := &componants.EncodeContext{
		Builder:  &sb,
		Calendar: componant,
	}

	componant.Encode(ctx)

	return []byte(sb.String()), nil
}
