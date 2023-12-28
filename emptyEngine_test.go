package glob

import "testing"

func TestEmpty_match(t *testing.T) {
	matches(t, emptyOnlyEngineType, "", "", true)
}

func TestEmpty_Length1_noMatch(t *testing.T) {
	matches(t, emptyOnlyEngineType, "", "a", false)
}

func TestEmpty_Length2_noMatch(t *testing.T) {
	matches(t, emptyOnlyEngineType, "", "ab", false)
}
