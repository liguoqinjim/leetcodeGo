package lab0083

import (
	"reflect"
	"testing"
)

var cases = []struct {
	input  *ListNode
	output *ListNode
}{
	{input: NewListNode([]int{1, 1, 2}), output: NewListNode([]int{1, 2})},
	{input: NewListNode([]int{1, 1, 2, 3, 3}), output: NewListNode([]int{1, 2, 3})},
	{input: nil, output: nil},
}

func TestDeleteDuplicates(t *testing.T) {
	for _, c := range cases {
		i := append([]int{}, c.input.ToArray()...)

		result := DeleteDuplicates(NewListNode(i))
		if !reflect.DeepEqual(result.ToArray(), c.output.ToArray()) {
			t.Errorf("DeleteDuplicates(%v)=%v,want %v", c.input.ToArray(), result.ToArray(), c.output.ToArray())
		}
	}
}

func NewListNode(nums []int) *ListNode {
	l := new(ListNode)
	temp := l
	for n, v := range nums {
		temp.Val = v

		if n != len(nums)-1 {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}

	return l
}

func (l1 *ListNode) ToArray() []int {
	nums := make([]int, 0)

	for {
		nums = append(nums, l1.Val)

		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}

	return nums
}
