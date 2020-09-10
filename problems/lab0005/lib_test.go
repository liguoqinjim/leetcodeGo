package lab0005

import "testing"

var cases = []struct {
	Input  string
	Output string
}{
	//{Input: "babad", Output: "bab"},
	//{Input: "cbbd", Output: "bb"},
	//{Input: "ccc", Output: "ccc"},
	//{Input: "abb", Output: "bb"},
	//{Input: "aaaa", Output: "aaaa"},
	//{Input: "aaaab", Output: "aaaa"},
	{Input: "bananas", Output: "anana"},
}

func TestLongestPalindrome(t *testing.T) {
	for _, c := range cases {
		result := LongestPalindrome(c.Input)
		if result != c.Output {
			t.Errorf("LongestPalindrome(%s)=%s,want %s", c.Input, result, c.Output)
		}
	}
}
