package lab0070

import "testing"

var cases = []struct {
	input  int
	output int
}{
	{input: 1, output: 1},
	{input: 2, output: 2},
	{input: 3, output: 3},
}

func TestClimbStairs(t *testing.T) {
	for _, c := range cases {
		result := ClimbStairs(c.input)
		if result != c.output {
			t.Errorf("ClimbStairs(%d)=%d,want %d", c.input, result, c.output)
		}
	}
}
