package glob

import (
	"testing"
)

func TestLength3_match(t *testing.T) {
	matches(t, containsEngineType, "%a%", "bac", true)
}

func TestEscape_match(t *testing.T) {
	matches(t, containsEngineType, "%a\\%%", "ba%dfo", true)
}

func TestLength4a_match(t *testing.T) {
	matches(t, containsEngineType, "%a%", "bacd", true)
}

func TestLength4b_match(t *testing.T) {
	matches(t, containsEngineType, "%a%", "dbac", true)
}

func TestEmpty_noMatch(t *testing.T) {
	matches(t, containsEngineType, "%a%", "", false)
}

func TestLength3_noMatch(t *testing.T) {
	matches(t, containsEngineType, "%a%", "bdc", false)
}

func TestLength7_match(t *testing.T) {
	matches(t, containsEngineType, "%abcd%", "abababcde", true)
}

func Test1(t *testing.T) {
	matches(t, containsEngineType, "%a_c%", "aabcc", true)
}
