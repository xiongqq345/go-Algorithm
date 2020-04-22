package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	l, r := dummy, dummy
	for n > 0 {
		r = r.Next
		n--
	}

	for r.Next != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return dummy.Next
}
