package glob

import (
	"reflect"
	"testing"
)

// Technically this should probably be called compiler_test, but I'm trying to keep the
// 1:1 parity with the glob-library-java library

var containsEngineType = reflect.TypeOf(&containsEngine{})
var everythingEngineType = reflect.TypeOf(&everythingEngine{})
var globEngineType = reflect.TypeOf(&globEngine{})
var endsWithEngineType = reflect.TypeOf(&endsWithEngine{})
var startsWithEngineType = reflect.TypeOf(&startsWithEngine{})
var emptyOnlyEngineType = reflect.TypeOf(&emptyOnlyEngine{})
var equalToEngineType = reflect.TypeOf(&equalToEngine{})

func cmp(bools []bool, mask uint, length int) bool {
	m := uint(1)
	for i := length - 1; i >= 0; i-- {
		if bools[i] {
			if (mask & m) != m {
				return false
			}
		}
		m = m << 1
	}
	return true
}

func compile(t *testing.T, typ reflect.Type, pattern, lowerCase, upperCase string, wildcard, matchOne uint) {
	m, err := Compile(pattern, WithWildcardChar('%'), WithMatchOneChar('_'), WithCaseInsensitive(true), WithHandleEscapes(true))
	if err != nil {
		t.Errorf("error was not nil: %v", err)
	}

	mTyp := reflect.TypeOf(m)
	if mTyp != typ {
		t.Errorf("type does not match %v != %v", mTyp, typ)
	}

	switch v := m.(type) {
	case *globEngine:
		if !cmp(v.wildcard, wildcard, v.length) {
			t.Errorf("wildcard mask not equal: %b != %v", wildcard, v.wildcard)
		}
	}
}

func matches(t *testing.T, typ reflect.Type, pattern string, match string, result bool) {
	m, err := Compile(pattern, WithWildcardChar('%'), WithMatchOneChar('_'), WithCaseInsensitive(true), WithHandleEscapes(true))
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

func TestSqlTest1(t *testing.T) {
	compile(t, everythingEngineType, "%%%", "%", "%", 0b01, 0b00)
}

func TestSqlTest2(t *testing.T) {
	compile(t, globEngineType, "a%b", "A%B", "a%b", 0b010, 0b000)
}

func TestSqlTest3(t *testing.T) {
	compile(t, globEngineType, "a%\\%%b", "A%%%B", "a%%%b", 0b01010, 0b000)
}

func TestSqlTest3a(t *testing.T) {
	compile(t, globEngineType, "a%\\%%b", "A%%%B", "a%%%b", 0b01010, 0b000)
}

func TestSqlTest3b(t *testing.T) {
	compile(t, globEngineType, "a%\\\\%b", "A%\\%B", "a%\\%b", 0b01010, 0b000)
}

func TestSqlTest4(t *testing.T) {
	compile(t, globEngineType, "a%%%b", "A%B", "a%b", 0b010, 0b000)
}

func TestSqlTest5(t *testing.T) {
	compile(t, globEngineType, "a%%_%b", "A%_%B", "a%_%b", 0b01010, 0b00100)
}

func TestSqlTest6(t *testing.T) {
	compile(t, globEngineType, "%%%a%%_%b%a%_", "%A%_%B%A%_", "%a%_%b%a%_", 0b1010101010, 0b0001000001)
}

func TestSqlTest7(t *testing.T) {
	compile(t, globEngineType, "%\\%%\\__\\_", "%%%___", "%%%___", 0b101000, 0b000010)
}

func TestSqlTest1CaseSensitive(t *testing.T) {
	compile(t, everythingEngineType, "%%%", "%", "%", 0b01, 0b00)
}

func TestSqlTest2CaseSensitive(t *testing.T) {
	compile(t, globEngineType, "a%b", "a%b", "a%b", 0b010, 0b000)
}

func TestContainsTest(t *testing.T) {
	compile(t, containsEngineType, "%%f_o%", "%F_O%", "%f_o%", 0b10001, 0b00100)
}

func TestEndsWithTest(t *testing.T) {
	compile(t, endsWithEngineType, "%f_o", "%F_O", "%f_o", 0b1000, 0b0010)
}

func TestStartsWithTest(t *testing.T) {
	compile(t, startsWithEngineType, "f_o%", "F_O%", "f_o%", 0b0001, 0b0100)
}

func TestEqualToTest(t *testing.T) {
	compile(t, equalToEngineType, "f_o", "F_O", "f_o", 0b000, 0b010)
}

func TestEmptyOnlyTest(t *testing.T) {
	compile(t, emptyOnlyEngineType, "", "", "", 0b0, 0b0)
}
