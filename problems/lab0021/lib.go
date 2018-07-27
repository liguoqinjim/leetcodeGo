package lab0021

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	l := &ListNode{}
	next := l
	for l1 != nil || l2 != nil {
		if l1 == nil || l2 == nil {
			if l1 == nil {
				next.Val = l2.Val
				l2 = l2.Next
			} else if l2 == nil {
				next.Val = l1.Val
				l1 = l1.Next
			}
		} else {
			if l1.Val <= l2.Val {
				next.Val = l1.Val
				l1 = l1.Next
			} else {
				next.Val = l2.Val
				l2 = l2.Next
			}
		}
		if l1 != nil || l2 != nil {
			next.Next = &ListNode{}
			next = next.Next
		}
	}

	return l
}

//sample
//递归解法
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head = &ListNode{}
	if l1.Val < l2.Val {
		head = l1
		head.Next = mergeTwoLists2(l1.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists2(l1, l2.Next)
	}

	return head
}

//测试里面用到的方法
func (l *ListNode) Print() string {
	s := ""
	for next := l; next != nil; {
		s += fmt.Sprint(next.Val)
		next = next.Next
		if next != nil {
			s += fmt.Sprint(",")
		}
	}
	return s
}

func (l1 *ListNode) Equal(l2 *ListNode) bool {
	n1 := l1
	n2 := l2
	for n1 != nil && n2 != nil {
		if n1 == nil {
			if n2 != nil {
				return false
			}
		}
		if n2 == nil {
			if n1 != nil {
				return false
			}
		}
		if n1.Val != n2.Val {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	return true
}
