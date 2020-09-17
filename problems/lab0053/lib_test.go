package lab0053

import "testing"

var cases = []struct {
	input  []int
	output int
}{
	{input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, output: 6},
	{input: []int{-2, 1}, output: 1},
	{input: []int{0}, output: 0},
	{input: []int{-1}, output: -1},
}

func TestMaxSubArray(t *testing.T) {
	for _, c := range cases {
		result := MaxSubArray(c.input)
		if result != c.output {
			t.Errorf("MaxSubArray(%v)=%d,want %d", c.input, result, c.output)
		}
	}
}
