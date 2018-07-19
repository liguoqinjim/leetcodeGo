package lib

import (
	"reflect"
	"testing"
)

var cases = []struct {
	nums   []int
	target int
	result []int
}{
	{[]int{6, 2, 7, 11}, 9, []int{1, 2}},
	{[]int{2, 3, 6}, 8, []int{0, 2}},
}

func TestTwoSum(t *testing.T) {
	for _, c := range cases {
		result := twoSum(c.nums, c.target)
		if !reflect.DeepEqual(result, c.result) { //这里偷懒直接用DeepEqual来判断两个数组是否相同
			t.Errorf("twonums(%v,%v) == %v, want %v", c.nums, c.target, result, c.result)
		}
	}
}

func TestTwoSum2(t *testing.T) {
	for _, c := range cases {
		result := twoSum2(c.nums, c.target)
		if !reflect.DeepEqual(result, c.result) { //这里偷懒直接用DeepEqual来判断两个数组是否相同
			t.Errorf("twonums(%v,%v) == %v, want %v", c.nums, c.target, result, c.result)
		}
	}
}
