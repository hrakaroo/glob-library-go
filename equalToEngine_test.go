package glob

import "testing"

// public void init() {
// testUtils = new TestUtils('%', '_', GlobPattern.CASE_INSENSITIVE | GlobPattern.HANDLE_ESCAPES);
// }

func TestEqualTo_patternHasEscapedGlob_match(t *testing.T) {
	matches(t, equalToEngineType, "ab\\%cd", "ab%cd", true)
}

func TestEqualTo_patternAndInputAreEqualLength1_match(t *testing.T) {
	matches(t, equalToEngineType, "a", "a", true)
}

func TestEqualTo_patternAndInputAreEqualLength2_match(t *testing.T) {
	matches(t, equalToEngineType, "ab", "ab", true)
}

func TestEqualTo_inputStringIsEmpty_noMatch(t *testing.T) {
	matches(t, equalToEngineType, "a", "", false)
}

func TestEqualTo_inputStringIsLongerThanPattern_noMatch(t *testing.T) {
	matches(t, equalToEngineType, "a", "aa", false)
}

func TestEqualTo_patternIsSingleBlank_noMatch(t *testing.T) {
	matches(t, equalToEngineType, "_", "", false)
}

func TestEqualTo_patternIsSingleBlank_match(t *testing.T) {
	matches(t, equalToEngineType, "_", "a", true)
}

func TestEqualTo_patternIsSingleBlank_noMatch2(t *testing.T) {
	matches(t, equalToEngineType, "_", "aa", false)
}

func TestEqualTo_patternStartsWithBlank_match(t *testing.T) {
	matches(t, equalToEngineType, "_b", "ab", true)
}

func TestEqualTo_patternEndsWithBlank_match(t *testing.T) {
	matches(t, equalToEngineType, "a_", "ab", true)
}

func TestEqualTo_patternIsTwoBlanks_match(t *testing.T) {
	matches(t, equalToEngineType, "__", "ab", true)
}

func TestEqualTo_patternHasTwoBlanksAtEnds_match(t *testing.T) {
	matches(t, equalToEngineType, "_b_", "abb", true)
}

/**
 * Test where the lengths are equal, but he strings differ after at the end
 */
func TestEqualTo_patternDiffersFromInputAtEnd_noMatch(t *testing.T) {
	matches(t, equalToEngineType, "abc", "abd", false)
}

/**
 * Test where the lengths are equal, but he strings differ in the middle
 */
func TestEqualTo_patternDiffersFromInputAtMiddle_noMatch(t *testing.T) {
	matches(t, equalToEngineType, "adc", "abc", false)
}

/**
 * Test where the lengths are equal, but he strings differ at the start
 */
func TestEqualTo_patternDiffersFromInputAtStart_noMatch(t *testing.T) {
	matches(t, equalToEngineType, "bbc", "abc", false)
}
