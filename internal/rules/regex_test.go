package rules

import (
	"regexp"
	"testing"
)

func TestRegexRuleMatch(t *testing.T) {
	pattern := "(failed to connect|timeout)"
	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatalf("failed to compile regex: %v", err)
	}

	tests := []struct {
		msg    string
		expect bool
	}{
		{"failed to connect to db", true},
		{"timeout while waiting", true},
		{"connection ok", false},
	}

	for _, tt := range tests {
		if re.MatchString(tt.msg) != tt.expect {
			t.Errorf("regex match failed for msg: %s", tt.msg)
		}
	}
}