package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var p *TreeNode
	var q *TreeNode
	if p == q {
		log.Println("q==p")
	}

	p = &TreeNode{}
	if p == q {
		log.Println("q==p2")
	} else {
		log.Println("q!=p")
	}
}
