package lab0027

import (
	"testing"
)

var cases = []struct {
	nums []int
	val  int
	len  int
}{
	{nums: []int{3, 2, 2, 3}, val: 3, len: 2},
	{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, len: 5},
}

func TestRemoveElement(t *testing.T) {
	for _, c := range cases {
		nums := make([]int, 0)
		nums = append(nums, c.nums...)
		result := RemoveElement(nums, c.val)
		if result != c.len {
			t.Errorf("RemoveElement(%v,%d)=%d,want %d", c.nums, c.val, result, c.len)
		}
	}
}

func TestRemoveElement2(t *testing.T) {
	for _, c := range cases {
		nums := make([]int, 0)
		nums = append(nums, c.nums...)
		result := RemoveElement2(nums, c.val)
		if result != c.len {
			t.Errorf("RemoveElement2(%v,%d)=%d,want %d", c.nums, c.val, result, c.len)
		}
	}
}
