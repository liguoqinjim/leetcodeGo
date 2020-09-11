package lab0013

import "testing"

var cases = []struct {
	Input  string
	Output int
}{
	{Input: "III", Output: 3},
	{Input: "IV", Output: 4},
	{Input: "IX", Output: 9},
	{Input: "MCMXCIV", Output: 1994},
}

func TestRomanToInt(t *testing.T) {
	for _, c := range cases {
		result := RomanToInt(c.Input)
		if result != c.Output {
			t.Errorf("RomanToInt(%s)=%d,want %d", c.Input, result, c.Output)
		}
	}
}

func TestRomanToInt2(t *testing.T) {
	for _, c := range cases {
		result := RomanToInt2(c.Input)
		if result != c.Output {
			t.Errorf("RomanToInt2(%s)=%d,want %d", c.Input, result, c.Output)
		}
	}
}
