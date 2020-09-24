package lab0068

import "testing"

var cases = []struct {
	input  int
	output int
}{
	{input: 4, output: 2},
	{input: 8, output: 2},
	{input: 3, output: 1},
	{input: 0, output: 0},
	{input: 1, output: 1},
	{input: 2, output: 1},
	{input: 2147483647, output: 46340},
}

func TestMySqrt(t *testing.T) {
	for _, c := range cases {
		result := MySqrt(c.input)
		if result != c.output {
			t.Errorf("MySqrt(%d)=%d,want %d", c.input, result, c.output)
		}
	}
}

func TestMySqrt2(t *testing.T) {
	for _, c := range cases {
		result := MySqrt2(c.input)
		if result != c.output {
			t.Errorf("MySqrt2(%d)=%d,want %d", c.input, result, c.output)
		}
	}
}

func TestMySqrt3(t *testing.T) {
	for _, c := range cases {
		result := MySqrt3(c.input)
		if result != c.output {
			t.Errorf("MySqrt3(%d)=%d,want %d", c.input, result, c.output)
		}
	}
}
