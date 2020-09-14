package lab0067

import "testing"

var cases = []struct {
	a      string
	b      string
	output string
}{
	//{a: "11", b: "1", output: "100"},
	//{a: "1010", b: "1011", output: "10101"},
	//{a: "0", b: "0", output: "0"},
	{a: "1", b: "111", output: "1000"},
}

func TestAddBinary(t *testing.T) {
	for _, c := range cases {
		result := AddBinary(c.a, c.b)
		if result != c.output {
			t.Errorf("AddBinary(%s,%s)=%s,want %s", c.a, c.b, result, c.output)
		}
	}
}

func TestAddBinary2(t *testing.T) {
	for _, c := range cases {
		result := AddBinary2(c.a, c.b)
		if result != c.output {
			t.Errorf("AddBinary2(%s,%s)=%s,want %s", c.a, c.b, result, c.output)
		}
	}
}