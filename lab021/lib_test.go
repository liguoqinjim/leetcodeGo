package lab021

import (
	"testing"
)

var cases = []struct {
	l1     *ListNode
	l2     *ListNode
	result *ListNode
}{
	{&ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		&ListNode{4, &ListNode{5, &ListNode{6, nil}}},
		&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
	},
}

func TestMergeTwoLists(t *testing.T) {
	for _, c := range cases {
		r := mergeTwoLists(c.l1, c.l2)
		if !r.Equal(c.result) {
			t.Errorf("MergeTwoList([%s],[%s]) = [%s],want [%s]", c.l1.Print(), c.l2.Print(), r.Print(), c.result.Print())
		}
	}
}

func TestMergeTwoLists2(t *testing.T) {
	for _, c := range cases {
		r := mergeTwoLists2(c.l1, c.l2)
		if !r.Equal(c.result) {
			t.Errorf("MergeTwoList2([%s],[%s]) = [%s],want [%s]", c.l1.Print(), c.l2.Print(), r.Print(), c.result.Print())
		}
	}
}
