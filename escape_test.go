package glob

import (
	"testing"
)

func TestCompile(t *testing.T) {
	// glob.Compile()
}

func TestUnicode(t *testing.T) {
	pattern := "foo\\u0010bar"

	matchingEngine, err := Compile(pattern)

	if err != nil {
		t.Errorf("failed")
	}

	if !matchingEngine.Matches("foo\u0010bar") {
		t.Errorf("failed to match")
	}
}

func TestBadUnicode(t *testing.T) {
	pattern := "foo\\u001zbar"

	_, err := Compile(pattern)

	if err == nil {
		t.Errorf("Bad unicode not detected")
	}
}

func TestBadUnicodeAtEnd(t *testing.T) {
	pattern := "foo\\u001"

	_, err := Compile(pattern)

	if err == nil {
		t.Errorf("Bad unicode not detected")
	}
}

func TestBadEscape(t *testing.T) {
	pattern := "foo\\e"

	_, err := Compile(pattern)

	if err == nil {
		t.Errorf("Bad escape not detected")
	}
}
