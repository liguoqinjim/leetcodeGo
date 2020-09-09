package lab0003

import "testing"

var cases = []struct {
	Input  string
	Output int
}{
	{Input: "abcabcbb", Output: 3},
	{Input: "bbbbb", Output: 1},
	{Input: "pwwkew", Output: 3},
	{Input: "", Output: 0},
	{Input: "aab", Output: 2},
	{Input: "abba", Output: 2},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, c := range cases {
		result := LengthOfLongestSubstring(c.Input)
		if result != c.Output {
			t.Errorf("LengthOfLongestSubstring(%s) == %d, want %d", c.Input, result, c.Output)
		}
	}
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	for _, c := range cases {
		result := LengthOfLongestSubstring2(c.Input)
		if result != c.Output {
			t.Errorf("LengthOfLongestSubstring2(%s) == %d, want %d", c.Input, result, c.Output)
		}
	}
}
