package _43

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := midNode(head)
	l2 := reverseList(mid.Next)
	mid.Next = nil
	mergeList(head, l2)
}

func midNode(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return p1
}

func mergeList(n1, n2 *ListNode) {
	for n1 != nil && n2 != nil {
		t1, t2 := n1.Next, n2.Next
		n2.Next = n1.Next
		n1.Next = n2
		n1, n2 = t1, t2
	}
}

func reverseList(head *ListNode) *ListNode {
	var p, c *ListNode = nil, head
	for c != nil {
		c, c.Next, p = c.Next, p, c
	}
	return p
}
