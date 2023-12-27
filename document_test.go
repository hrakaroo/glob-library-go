package glob

import "testing"

func TestExample1(t *testing.T) {
	m, err := Compile("dog*cat\\*goat??")

	if err != nil {
		t.Error(err)
	}

	if !m.Matches("dog horse cat*goat!~") {
		t.Errorf("failed 1")
	}

	if !m.Matches("dogcat*goat..") {
		t.Errorf("failed 2")
	}

	if m.Matches("dog catgoat!/") {
		t.Errorf("failed 3")
	}
}

func TestExample2(t *testing.T) {

	m, err := Compile("dog%cat\\%goat__", WithWildcardChar('%'), WithMatchOneChar('_'), WithHandleEscapes(true))

	if err != nil {
		t.Error(err)
	}

	if !m.Matches("dog horse cat%goat!~") {
		t.Errorf("failed 1")
	}

	if !m.Matches("dogcat%goat..") {
		t.Errorf("failed 2")
	}

	if m.Matches("dog catgoat!/") {
		t.Errorf("failed 3")
	}
}
