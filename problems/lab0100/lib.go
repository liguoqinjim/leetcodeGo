package lab0100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {	//全是nil
		return true
	} else {
		if p != nil && q != nil {
			if p.Val != q.Val {
				return false
			}
		} else {
			return false
		}
	}

	if !IsSameTree(p.Left, q.Left) {
		return false
	}

	if !IsSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
