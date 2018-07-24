package lab0008

import "testing"

var cases = []struct {
	s      string
	result int
}{
	{"1234", 1234},
	{"a1234", 0},
	{"01234", 1234},
	{"1234.5", 1234}, //按照 {"  -0012a42", -12}, 这个逻辑得出这个的正确结果应该是1234
	{"+1234", 1234},
	{"-1234", -1234},
	{"", 0},
	{"   010", 10},
	{"+-2", 0},
	{"    -00134", -134},
	{"  -0012a42", -12},
	{"2147483648", 2147483647},
	{"-2147483648", -2147483648},
	{"-2147483647", -2147483647},
	{"123  456", 123},
	{"9223372036854775809", 2147483647},
	{"   -1123u3761867", -1123},
	{"-9223372036854775809", -2147483648},
}

func TestMyAtoi(t *testing.T) {
	for _, c := range cases {
		r := myAtoi(c.s)
		if r != c.result {
			t.Errorf("myAtoi(%s) = %d,want %d", c.s, r, c.result)
		}
	}
}

func TestMyAtoi2(t *testing.T) {
	for _, c := range cases {
		r := myAtoi2(c.s)
		if r != c.result {
			t.Errorf("myAtoi2(%s) = %d,want %d", c.s, r, c.result)
		}
	}
}
