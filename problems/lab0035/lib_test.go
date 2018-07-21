package lab0035

import "testing"

var cases = []struct {
	nums   []int
	target int
	result int
}{
	{[]int{1, 3, 5, 6}, 5, 2},
	{[]int{1, 3, 5, 6}, 2, 1},
	{[]int{1, 3, 5, 6}, 7, 4},
	{[]int{1, 3, 5, 6}, 0, 0},
}

func TestSearchInsert(t *testing.T) {
	for _, c := range cases {
		r := searchInsert(c.nums, c.target)
		if r != c.result {
			t.Errorf("searchInsert(%v,%d) = %d,want %d", c.nums, c.target, r, c.result)
		}
	}
}

func TestSearchInsert2(t *testing.T) {
	for _, c := range cases {
		r := searchInsert2(c.nums, c.target)
		if r != c.result {
			t.Errorf("searchInsert2(%v,%d) = %d,want %d", c.nums, c.target, r, c.result)
		}
	}
}
