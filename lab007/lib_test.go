package lab007

import (
	"testing"
)

var cases = []struct {
	x      int
	result int
}{
	{1, 1},
	{1234, 4321},
	{99891, 19899},
	{1534236469, 0},
	{-132, -231},
	{-2147483648, 0},
	{-214748364811, 0},
}

func TestReverse(t *testing.T) {
	for _, c := range cases {
		r := reverse(c.x)
		if r != c.result {
			t.Errorf("Reverse(%d) = %d,want %d", c.x, r, c.result)
		}
	}
}

func TestReverse2(t *testing.T) {
	for _, c := range cases {
		r := reverse2(c.x)
		if r != c.result {
			t.Errorf("Reverse2(%d) = %d,want %d", c.x, r, c.result)
		}
	}
}
