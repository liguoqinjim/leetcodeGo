package lab0038

import "testing"

var cases = []struct {
	input  int
	output string
}{
	{input: 1, output: "1"},
	{input: 2, output: "11"},
	{input: 3, output: "21"},
	{input: 4, output: "1211"},
	{input: 5, output: "111221"},
}

func TestCountAndSay(t *testing.T) {
	for _, c := range cases {
		result := CountAndSay(c.input)
		if result != c.output {
			t.Errorf("CountAndSay(%d)=%s,want %s", c.input, result, c.output)
		}
	}
}
