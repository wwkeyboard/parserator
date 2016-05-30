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

func TestIsDigit(t *testing.T) {
	tests := []testRune{
		{'1', true},
		{'0', true},
		{'5', true},
		{' ', false},
		{'a', false},
		{'*', false},
	}
	for _, tt := range tests {
		if isDigit(tt.trial) != tt.want {
			t.Errorf("incorrect digit detection of '%v'", tt.trial)
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

func TestScanIdent(t *testing.T) {
	s := NewScanner(strings.NewReader("blar"))
	tok, lit := s.scanIdent()

	if tok != IDENT {
		t.Error("scanIdent didn't detect identifier")
	}
	if lit != "blar" {
		t.Error("scanIdent didn't detect whole identifier")
	}
}
