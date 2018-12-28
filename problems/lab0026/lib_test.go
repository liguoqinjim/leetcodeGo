package lab0026

import "testing"

var cases = []struct {
	nums   []int
	result int
	nums2  []int
}{
	{nums: []int{1, 1, 2}, result: 2, nums2: []int{1, 2}},
	{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, result: 5, nums2: []int{0, 1, 2, 3, 4}},
	{nums: []int{}, result: 0, nums2: []int{}},
}

func TestRemoveDuplicates(t *testing.T) {
	for _, c := range cases {
		r := removeDuplicates(c.nums)

		if r != c.result {
			t.Errorf("removeDuplicates(%v) = %d ,want %d", c.nums, r, c.result)
		}

		for i := 0; i < len(c.nums2); i++ {
			if c.nums[i] != c.nums2[i] {
				t.Errorf("removeDuplicates nums = %v,want %v", c.nums, c.nums2)
			}
		}

	}
}
