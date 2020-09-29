package lab0101

import "testing"

var cases = []struct {
	input  *TreeNode
	output bool
}{
	{
		input: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
		output: true,
	},
	{
		input: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		},
		output: false,
	},
}

func TestIsSymmetric(t *testing.T) {
	for _, c := range cases {
		result := IsSymmetric(c.input)
		if result != c.output {
			t.Errorf("IsSymmetric(%v)=%t,want %t", c.input, result, c.output)
		}
	}
}
