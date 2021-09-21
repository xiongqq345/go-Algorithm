package jz_offerII

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func midNode(h *ListNode) *ListNode {
	p1, p2 := h, h.Next
	for p2 != nil && p2.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
	}
	return p1
}

func reverseToNode(h *ListNode, n *ListNode) (p *ListNode) {
	for h != n {
		h, h.Next, p = h.Next, p, h
	}
	return
}
