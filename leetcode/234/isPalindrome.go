package _34

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	p1, p2 := head, head
	var pre *ListNode
	for p2 != nil && p2.Next != nil {
		p2 = p2.Next.Next
		next := p1.Next
		p1.Next = pre
		pre = p1
		p1 = next
	}

	if p2 != nil {
		p1 = p1.Next
	}

	for p1 != nil {
		if pre.Val != p1.Val {
			return false
		}
		pre = pre.Next
		p1 = p1.Next
	}
	return true
}
