package lab0009

import "testing"

var cases = []struct {
	num    int
	result bool
}{
	{1221, true},
	{123, false},
	{1, true},
	{1234, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, c := range cases {
		r := isPalindrome(c.num)
		if r != c.result {
			t.Errorf("IsPalindrome(%d) = %t,want %t", c.num, r, c.result)
		}
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, c := range cases {
		r := isPalindrome2(c.num)
		if r != c.result {
			t.Errorf("IsPalindrome2(%d) = %t,want %t", c.num, r, c.result)
		}
	}
}
