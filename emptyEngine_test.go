package glob

import "testing"

func TestEmpty_match(t *testing.T) {
	matches(t, emptyOnlyEngineType, "", "", true)
}

func TestLength1_noMatch(t *testing.T) {
	matches(t, emptyOnlyEngineType, "", "a", false)
}

func TestLength2_noMatch(t *testing.T) {
	matches(t, emptyOnlyEngineType, "", "ab", false)
}
