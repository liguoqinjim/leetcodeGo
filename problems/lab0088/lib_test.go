package lab0088

import (
	"reflect"
	"testing"
)

var cases = []struct {
	nums1  []int
	m      int
	nums2  []int
	n      int
	output []int
}{
	{
		nums1: []int{1, 2, 3, 0, 0, 0}, m: 3,
		nums2: []int{2, 5, 6}, n: 3,
		output: []int{1, 2, 2, 3, 5, 6},
	},
	{
		nums1: []int{1, 0, 0}, m: 1,
		nums2: []int{2, 2}, n: 2,
		output: []int{1, 2, 2},
	},
	{
		nums1: []int{0}, m: 0,
		nums2: []int{1}, n: 1,
		output: []int{1},
	},
}

func TestMerge(t *testing.T) {
	for _, c := range cases {
		nums1Input := append([]int{}, c.nums1...)
		Merge(nums1Input, c.m, c.nums2, c.n)
		if !reflect.DeepEqual(nums1Input, c.output) {
			t.Errorf("Merge(%v,%v)=%v,want %v", c.nums1, c.nums2, nums1Input, c.output)
		}
	}
}
