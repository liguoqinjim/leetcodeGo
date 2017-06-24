package lab021

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
