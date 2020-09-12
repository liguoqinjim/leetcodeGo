package lab0058

import "testing"

var cases = []struct {
	input  string
	output int
}{
	{input: "Hello World", output: 5},
	{input: "a", output: 1},
	{input: "a ", output: 1},
	{input: "nn a ", output: 1},
}

func TestLengthOfLastWord(t *testing.T) {
	for _, c := range cases {
		result := LengthOfLastWord(c.input)
		if result != c.output {
			t.Errorf("LengthOfLastWord(%s)=%d,want %d", c.input, result, c.output)
		}
	}
}
