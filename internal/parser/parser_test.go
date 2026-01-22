package parser

import (
	"testing"

	"github.com/filipe/loglens/internal/event"
)

func TestParseLineBasic(t *testing.T) {
	e := event.Event{
		Raw: "ERROR failed to connect",
	}

	ParseLine(&e)

	if e.Level != "ERROR" {
		t.Errorf("expected Level=ERROR, got %s", e.Level)
	}

	if e.Message != "failed to connect" {
		t.Errorf("expected Message='failed to connect', got %s", e.Message)
	}
}

func TestParseLineEmpty(t *testing.T) {
	e := event.Event{
		Raw: "",
	}

	ParseLine(&e)

	if e.Level != "" {
		t.Errorf("expected empty Level, got %s", e.Level)
	}

	if e.Message != "" {
		t.Errorf("expected empty Message, got %s", e.Message)
	}
}