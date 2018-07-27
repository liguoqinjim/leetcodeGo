package main

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	l := mergeTwoLists2(l1, l2)
	log.Printf("%s",l.Print())

	//fmt.Println("\n\nnode=")
	//for node := l; node != nil; {
	//	fmt.Println(node)
	//	node = node.Next
	//}

	//s := l.Print()
	//fmt.Println(s)
	//
	//fmt.Println("\n\n\n测试相等")
	//l3 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	//l4 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	//fmt.Println(l3.Equal(l4))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	defer func() {
		if r := recover(); r != nil {
			panic(r)
		}
	}()

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

//递归解法
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		log.Println("l1 is nil")
		return l2
	}
	if l2 == nil {
		log.Println("l2 is nil")
		return l1
	}

	var head = &ListNode{}
	if l1.Val < l2.Val {
		head = l1
		log.Println("merge l1.Next l2", head)
		head.Next = mergeTwoLists2(l1.Next, l2)
		log.Println("merge l1.Next l2 end", head)
	} else {
		head = l2
		log.Println("merge l1 l2.Next", head)
		head.Next = mergeTwoLists2(l1, l2.Next)
		log.Println("merge l1 l2.Next end", head)
	}

	return head
}

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
