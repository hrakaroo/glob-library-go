package glob

import (
	"strings"
	"testing"
)

const pattern = "my%dog%_spot"

func String_testMatch(t *testing.T) {

	matchingEngine, err := Compile(pattern, WithWildcardChar(NullCharacter), WithMatchOneChar(NullCharacter), WithHandleEscapes(true))

	if err != nil {
		t.Error("Compile failed")
	}

	// Verify a basic string compare matches
	if !matchingEngine.Matches(pattern) {
		t.Error("Match failed")
	}
}

func String_testMatchFail(t *testing.T) {

	matchingEngine, err := Compile(pattern, WithWildcardChar(NullCharacter), WithMatchOneChar(NullCharacter), WithHandleEscapes(true))
	if err != nil {
		t.Error("Compile failed")
	}

	// Verify if we swap the '%' for a 'r'
	pattern2 := strings.ReplaceAll(pattern, "%", "r")
	if matchingEngine.Matches(pattern2) {
		t.Error("Match should have failed")
	}
}

func String_testMatchFail2(t *testing.T) {

	matchingEngine, err := Compile(pattern, WithWildcardChar(NullCharacter), WithMatchOneChar(NullCharacter), WithHandleEscapes(true))
	if err != nil {
		t.Error("Compile failed")
	}

	// Verify if we swap the '_' for a 'r'
	pattern2 := strings.ReplaceAll(pattern, "_", "r")
	if matchingEngine.Matches(pattern2) {
		t.Error("Match should have failed")
	}
	// assert (! matchingEngine.matches(pattern.replace('_', 'r')));
}

func String_testMatchFailCase(t *testing.T) {

	matchingEngine, err := Compile(pattern, WithWildcardChar(NullCharacter), WithMatchOneChar(NullCharacter), WithHandleEscapes(true))
	if err != nil {
		t.Error("Compile failed")
	}

	// Upper case the pattern
	pattern2 := strings.ToUpper(pattern)
	if matchingEngine.Matches(pattern2) {
		t.Error("Match should have failed")
	}
}
