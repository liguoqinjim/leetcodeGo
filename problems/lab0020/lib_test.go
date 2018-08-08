package lab0020

import "testing"

var cases = []struct {
	s      string
	result bool
}{
	{"{()}[]", true},
	{"{{]]", false},
	{"{}()[]", true},
	{"}{", false},
	{"", true},
	{"{[]}", true},
}

func TestIsValid(t *testing.T) {
	for _, c := range cases {
		r := isValid(c.s)
		if r != c.result {
			t.Errorf("isValid(%t) = %t,want %t", c.result, r, c.result)
		}
	}
}

func TestIsValid2(t *testing.T) {
	for _, c := range cases {
		r := isValid2(c.s)
		if r != c.result {
			t.Errorf("isValid(%t) = %t,want %t", c.result, r, c.result)
		}
	}
}
