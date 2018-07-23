package lab0014

import "testing"

var cases = []struct {
	in     []string
	result string
}{
	{[]string{"abc", "abdc", "abd", "ab"}, "ab"},
	{[]string{"abc", "abdc", "abd"}, "ab"},
	{[]string{}, ""},
	{[]string{""}, ""},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, c := range cases {
		r := longestCommonPrefix(c.in)
		if r != c.result {
			t.Errorf("longestCommonPrefix(%v) == %s,want %s", c.in, r, c.result)
		}
	}
}

func TestLongestCommonPrefix2(t *testing.T) {
	for _, c := range cases {
		r := longestCommonPrefix2(c.in)
		if r != c.result {
			t.Errorf("longestCommonPrefix(%v) == %s,want %s", c.in, r, c.result)
		}
	}
}
