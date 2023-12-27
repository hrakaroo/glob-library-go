package glob

import (
	"reflect"
	"testing"
)

var containsEngineType = reflect.TypeOf(&containsEngine{})

func matches(t *testing.T, typ reflect.Type, pattern string, match string, result bool) {
	m, err := Compile(pattern, WithWildcardChar('%'), WithMatchOneChar('_'), WithCaseInsensitive(true))
	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}

	mTyp := reflect.TypeOf(m)
	if mTyp != typ {
		t.Errorf("type does not match %v != %v", mTyp, typ)
	}

	lResult := m.Matches(match)
	if lResult != result {
		t.Errorf("result does not match expected %v != %v", lResult, result)
	}
}

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
