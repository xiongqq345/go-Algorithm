package linked_list

func midNode(head *ListNode) *ListNode {
	p1, p2 := head, head.Next
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
	}
	return p1
}

func reverse(h *ListNode) *ListNode {
	var p *ListNode
	for h != nil {
		h, h.Next, p = h.Next, p, h
	}
	return p
}

func reverseToNode(h, node *ListNode) *ListNode {
	var p *ListNode
	for h != node {
		h, h.Next, p = h.Next, p, h
	}
	return p
}
