package glob

import "testing"

func TestWithoutWildcard(t *testing.T) {
	options := WithoutWildcard()

	co := compileOption{}
	options.apply(&co)

	if co.wildcardChar != NullCharacter {
		t.Error("Wildcard is not null")
	}
}

func TestWithoutMatchOne(t *testing.T) {
	options := WithoutMatchOne()

	co := compileOption{}
	options.apply(&co)

	if co.matchOneChar != NullCharacter {
		t.Error("MatchOne is not null")
	}
}

func TestWitWildcard(t *testing.T) {
	options := WithWildcardChar('d')

	co := compileOption{}
	options.apply(&co)

	if co.wildcardChar != 'd' {
		t.Error("Wildcard was not set")
	}
}

func TestWithMatchOne(t *testing.T) {
	options := WithMatchOneChar('c')

	co := compileOption{}
	options.apply(&co)

	if co.matchOneChar != 'c' {
		t.Error("MatchOne was not set")
	}
}
