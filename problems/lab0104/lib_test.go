package lab0104

import "testing"

var cases = []struct {
	input  *TreeNode
	output int
}{
	{
		input:  &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		output: 3,
	},
}

func TestMaxDepth(t *testing.T) {
	for _, c := range cases {
		result := MaxDepth(c.input)
		if result != c.output {
			t.Errorf("MaxDepth(%v)=%d,want %d", c.input, result, c.output)
		}
	}
}
