package glob

import "testing"

// public void init() {
//     testUtils = new TestUtils('%', '_', GlobPattern.CASE_INSENSITIVE | GlobPattern.HANDLE_ESCAPES);
// }

func TestEndsWith_noMatch(t *testing.T) {
	matches(t, endsWithEngineType, "%a", "", false)
}

func TestEndsWith_Length1_match(t *testing.T) {
	matches(t, endsWithEngineType, "%a", "a", true)
}

func TestEndsWith_Length2_match(t *testing.T) {
	matches(t, endsWithEngineType, "%a", "ba", true)
}

func TestEndsWith_Length1_noMatch(t *testing.T) {
	matches(t, endsWithEngineType, "%a", "b", false)
}

func TestEndsWith_test1(t *testing.T) {
	matches(t, endsWithEngineType, "%a_c", "badc", true)
}
