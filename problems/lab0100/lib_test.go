package lab0100

import "testing"

var cases = []struct {
	input1 *TreeNode
	input2 *TreeNode
	output bool
}{
	{
		input1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		input2: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
		output: true,
	},
	{
		input1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
		input2: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		output: false,
	},
	{
		input1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
		input2: &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
		output: false,
	},
}

func TestIsSameTree(t *testing.T) {
	for _, c := range cases {
		result := IsSameTree(c.input1, c.input2)
		if result != c.output {
			t.Errorf("IsSameTree(%v,%v)=%t,want %t", c.input1, c.input2, result, c.output)
		}
	}
}
