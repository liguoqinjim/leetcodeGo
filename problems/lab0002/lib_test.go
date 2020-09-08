package lab0002

import (
	"reflect"
	"testing"
)

//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.

var cases = []struct {
	input1 []int
	input2 []int
	output []int
}{
	{
		input1: []int{2, 4, 3},
		input2: []int{5, 6, 4},
		output: []int{7, 0, 8},
	},
	{
		input1: []int{5},
		input2: []int{5},
		output: []int{0, 1},
	},
	{
		input1: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		input2: []int{5, 6, 4},
		output: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, c := range cases {
		result := AddTwoNumbers(NewListNode(c.input1), NewListNode(c.input2))
		//if !result.Equals(NewListNode(c.output)) {
		//	t.Errorf("AddTwoNumbers(%v,%v) == %v, want %v", c.input1, c.input2, result.ToArray(), c.output)
		//}
		if !reflect.DeepEqual(result.ToArray(), c.output) {
			t.Errorf("AddTwoNumbers(%v,%v) == %v, want %v", c.input1, c.input2, result.ToArray(), c.output)
		}
	}
}
