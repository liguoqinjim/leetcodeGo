package lab0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	l := new(ListNode)
	temp := l
	for n, v := range nums {
		temp.Val = v

		if n != len(nums)-1 {
			temp.Next = &ListNode{}
			temp = temp.Next
		}
	}

	return l
}

func (l1 *ListNode) ToArray() []int {
	nums := make([]int, 0)

	for {
		nums = append(nums, l1.Val)

		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}

	return nums
}

func (l1 *ListNode) Equals(l2 *ListNode) bool {
	for {
		if l1.Val != l2.Val {
			return false
		}

		if (l1.Next == nil && l2.Next != nil) || (l1.Next != nil && l2.Next == nil) {
			return false
		}

		if l1.Next == nil {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	return true
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	temp := l
	first := true

	carry := false
	for {
		val := 0
		if l1 == nil && l2 == nil && !carry {
			break
		}

		if !first {
			temp.Next = &ListNode{}
			temp = temp.Next
		} else {
			first = false
		}

		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		if carry {
			val += 1
			carry = false
		}

		temp.Val = val % 10
		if val >= 10 {
			carry = true
		}
	}

	return l
}
