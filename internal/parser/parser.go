package parser

import (
	"strings"

	"github.com/filipe/loglens/internal/event"
)

func ParseLine(e *event.Event) {
	parts := strings.Fields(e.Raw)
	if len(parts) == 0 {
		return
	}

	e.Level = parts[0]

	if len(parts) > 1 {
		e.Message = strings.Join(parts[1:], " ")
	}
}