package glob

import "testing"

func TestStartsWith_empty_noMatch(t *testing.T) {
	matches(t, startsWithEngineType, "a%", "", false)
}

func TestStartsWith_matchLength1_match(t *testing.T) {
	matches(t, startsWithEngineType, "a%", "a", true)
}

func TestStartsWith_matchLength2_match(t *testing.T) {
	matches(t, startsWithEngineType, "a%", "ab", true)
}

func TestStartsWith_noMatchLength1_noMatch(t *testing.T) {
	matches(t, startsWithEngineType, "a%", "b", false)
}

func TestStartsWith_noMatchLength2_noMatch(t *testing.T) {
	matches(t, startsWithEngineType, "a%", "ba", false)
}

func TestStartsWith_test(t *testing.T) {
	matches(t, startsWithEngineType, "abcdef%", "ab", false)
}

func TestStartsWith_test2(t *testing.T) {
	matches(t, startsWithEngineType, "abcde%", "abcdef", true)
}

func TestStartsWith_test3(t *testing.T) {
	matches(t, startsWithEngineType, "abcde%", "abcde", true)
}

func TestStartsWith_test4(t *testing.T) {
	matches(t, startsWithEngineType, "a_cde%", "abcde", true)
}
