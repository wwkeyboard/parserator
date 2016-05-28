package parserator

import (
	"strings"
	"testing"
)

type testRune struct {
	trial rune
	want  bool
}

func TestIsWhitespace(t *testing.T) {
	tests := []testRune{
		{' ', true},
		{'\n', true},
		{'\t', true},
		{'a', false},
		{'9', false},
	}
	for _, tt := range tests {
		if isWhitespace(tt.trial) != tt.want {
			t.Errorf("incorrect whitespace detection of '%s'", tt.trial)
		}
	}
}

func TestIsLetter(t *testing.T) {
	tests := []testRune{
		{'a', true},
		{'z', true},
		{'B', true},
		{' ', false},
		{'9', false},
		{'*', false},
	}
	for _, tt := range tests {
		if isLetter(tt.trial) != tt.want {
			t.Errorf("incorrect whitespace detection of '%s'", tt.trial)
		}
	}
}

func TestScanWhitespace(t *testing.T) {
	s := NewScanner(strings.NewReader("    foo"))
	tok, lit := s.scanWhitespace()

	if tok != WS {
		t.Error("scanWhitespace returned incorrect token")
	}

	if lit != "    " {
		t.Error("scanWhitespace didn't slurp all of the whitespace")
	}
}
