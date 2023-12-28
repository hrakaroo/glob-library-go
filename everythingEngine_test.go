package glob

import "testing"

func TestEverything_empty_match(t *testing.T) {
	matches(t, everythingEngineType, "%", "", true)
}

func TestEverything_length1_match(t *testing.T) {
	matches(t, everythingEngineType, "%", "a", true)
}

func TestEverything_length2_match(t *testing.T) {
	matches(t, everythingEngineType, "%", "ab", true)
}
