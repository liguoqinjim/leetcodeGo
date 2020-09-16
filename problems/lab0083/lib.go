package lab0083

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[int]bool)
	m[head.Val] = true

	i := head
	for i.Next != nil {
		if _, ok := m[i.Next.Val]; !ok {
			m[i.Next.Val] = true
			i = i.Next
		} else {
			i.Next = i.Next.Next
		}
	}

	return head
}
