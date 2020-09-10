package lab0028

import "testing"

var cases = []struct {
	haystack string
	needle   string
	output   int
}{
	{haystack: "hello", needle: "ll", output: 2},
	{haystack: "aaaaa", needle: "bba", output: -1},
	{haystack: "aaa", needle: "", output: 0},
}

func TestStrStr(t *testing.T) {
	for _, c := range cases {
		result := StrStr(c.haystack, c.needle)
		if result != c.output {
			t.Errorf("StrStr(%s,%s)=%d,want %d", c.haystack, c.needle, result, c.output)
		}
	}
}
