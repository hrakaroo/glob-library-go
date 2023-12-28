package glob

import "testing"

func TestGlob_escape_match(t *testing.T) {
	matches(t, globEngineType, "%a%\\%b%", "%af%bd", true)
}

/**
 * This is intended to verify that our stack doubling is working, so we need a big pattern
 */
func TestGlob_big_match(t *testing.T) {
	pattern := "%a%b%c%d%e%f%g%h%i%j%k%l%m%n%o%p%q%r%s%t%u%v"
	str := "ababcabcdabcdeabcdeabcdefabcdefgabcdefghabcdefghiabcdefghijabcdefghijkabcdefghijkl" +
		"abcdefghijklmabcdefghijklmnabcdefghijklmnoabcdefghijklmnopabcdefghijklmnopqabcdefghijklmnopqr" +
		"abcdefghijklmnopqrsabcdefghijklmnopqrstabcdefghijklmnopqrstuabcdefghijklmnopqrstuv"
	matches(t, globEngineType, pattern, str, true)
}

func TestGlob_escape_r(t *testing.T) {
	pattern := "foo\\rbar"
	str := "foo\rbar"
	matches(t, equalToEngineType, pattern, str, true)
}

func TestGlob_escape_n(t *testing.T) {
	pattern := "foo\\nbar"
	str := "foo\nbar"
	matches(t, equalToEngineType, pattern, str, true)
}

func TestGlob_escape_t(t *testing.T) {
	pattern := "foo\\tbar"
	str := "foo\tbar"
	matches(t, equalToEngineType, pattern, str, true)
}

func TestGlob_escape_other(t *testing.T) {
	pattern := "foo\\\\bar"
	str := "foo\\bar"
	matches(t, equalToEngineType, pattern, str, true)
}
