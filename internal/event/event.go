package event

import "time"

type Event struct {
	LineNumber int
	Raw        string
	Timestamp  time.Time
	Level      string
	Message    string
}