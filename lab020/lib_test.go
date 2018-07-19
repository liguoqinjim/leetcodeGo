package lab020

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
}

func TestIsValid(t *testing.T) {
	for _, c := range cases {
		r := isValid(c.s)
		if r != c.result {
			t.Errorf("isValid(%t) = %t,want %t", c.result, r, c.result)
		}
	}
}
